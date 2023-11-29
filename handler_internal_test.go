package grpcuiredoc

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

var (
	themeCSSPattern       = regexp.MustCompile(`s/theme\.css\?t=\d+`)
	descriptionCSSPattern = regexp.MustCompile(`s/description\.css\?t=\d+`)
)

func TestHandlerViaReflection_ServerHasNoReflection(t *testing.T) {
	t.Parallel()

	_, conn := createGRPCServer(t)

	_, actual := HandlerViaReflection(context.Background(), conn, "")
	expected := status.Error(codes.Unimplemented, "unknown service grpc.reflection.v1alpha.ServerReflection")

	assert.Equal(t, expected, actual)
}

func TestHandlerViaReflection_CouldNotGetFilesViaReflection(t *testing.T) {
	t.Parallel()

	cnt := 0

	_, conn := createGRPCServerWithReflection(t,
		grpc.ChainStreamInterceptor(func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			if cnt > 0 {
				return status.Error(codes.Internal, "reflection error")
			}

			cnt++

			return handler(srv, ss)
		}),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	_, actual := HandlerViaReflection(ctx, conn, "")
	expected := status.Error(codes.Internal, "reflection error")

	assert.Equal(t, expected, actual)
}

func TestHandlerViaReflection(t *testing.T) { //nolint: paralleltest
	testCases := []struct {
		scenario            string
		options             []HandlerOption
		expectedDescription string
	}{
		{
			scenario:            "no extra options",
			expectedDescription: `404 page not found`,
		},
		{
			scenario: "with description",
			options: []HandlerOption{
				AddDescription("test description"),
			},
			expectedDescription: fmt.Sprintf(`.target::after {content: %q;}`, "test description"),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) { //nolint: paralleltest
			_, conn := createGRPCServerWithReflection(t)

			h, err := HandlerViaReflection(context.Background(), conn, "", tc.options...)
			require.NoError(t, err, "could not start http server")

			// Get index.
			index, err := requestContent(h, "/")
			require.NoError(t, err, "could not get index response")

			require.Truef(t, themeCSSPattern.MatchString(index), "could not find %q in index", "theme.css?t=<TIMESTAMP>")
			require.Truef(t, descriptionCSSPattern.MatchString(index), "could not find %q in index", "description.css?t=<TIMESTAMP>")

			// Get theme.css.
			themeFile := themeCSSPattern.FindString(index)

			theme, err := requestContent(h, themeFile)
			require.NoError(t, err, "could not get %q", themeFile)

			assert.Equal(t, removeTrailingNL(cssContent), theme)

			// Get description.css.
			descriptionFile := descriptionCSSPattern.FindString(index)

			desc, err := requestContent(h, descriptionFile)
			require.NoError(t, err, "could not get %q", descriptionFile)

			assert.Equal(t, tc.expectedDescription, desc)

			// Get logo.
			logoCode, logoHeader, _, err := request(h, "grpc-logo.png")
			require.NoError(t, err, "could not get grpc logo: %w", err)

			assert.Equal(t, http.StatusTemporaryRedirect, logoCode)
			assert.Equal(t, svgLogo, logoHeader.Get("location"))
		})
	}
}

func splitGRPCOptions(opts ...interface{}) ([]grpc.ServerOption, []func(s *grpc.Server)) {
	srvOpts := make([]grpc.ServerOption, 0, len(opts))
	configure := make([]func(s *grpc.Server), 0, len(opts))

	for _, o := range opts {
		switch o := o.(type) {
		case grpc.ServerOption:
			srvOpts = append(srvOpts, o)

		case func(s *grpc.Server):
			configure = append(configure, o)
		}
	}

	return srvOpts, configure
}

func createGRPCServer(t *testing.T, opts ...interface{}) (*grpc.Server, *grpc.ClientConn) {
	t.Helper()

	srvOpts, configure := splitGRPCOptions(opts...)

	s := grpc.NewServer(srvOpts...)

	for _, c := range configure {
		c(s)
	}

	l := bufconn.Listen(1024 * 1024)

	go func() {
		defer l.Close() //nolint: errcheck

		err := s.Serve(l)

		require.NoError(t, err, "could not start grpc server")
	}()

	t.Cleanup(func() {
		s.GracefulStop()
	})

	conn, err := grpc.DialContext(context.Background(), "",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return l.Dial()
		}),
	)
	require.NoError(t, err, "could not create grpc connection")

	return s, conn
}

func createGRPCServerWithReflection(t *testing.T, opts ...interface{}) (*grpc.Server, *grpc.ClientConn) {
	t.Helper()

	opts = append(opts, func(s *grpc.Server) {
		reflection.Register(s)
	})

	return createGRPCServer(t, opts...)
}

func request(h http.Handler, path string) (int, http.Header, io.Reader, error) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	u, err := req.URL.Parse(path)
	if err != nil {
		return 0, nil, nil, fmt.Errorf("could not parse path: %w", err)
	}

	req.URL = u

	h.ServeHTTP(rr, req)

	return rr.Code, rr.Header(), rr.Body, err
}

func requestContent(h http.Handler, path string) (string, error) {
	_, _, body, err := request(h, path)
	if err != nil {
		return "", err
	}

	if body == nil {
		return "", nil
	}

	b, err := io.ReadAll(body)
	if err != nil {
		return "", err
	}

	return removeTrailingNL(string(b)), nil
}

func removeTrailingNL(s string) string {
	return strings.TrimSuffix(s, "\n")
}
