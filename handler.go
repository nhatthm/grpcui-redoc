package grpcuiredoc

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/fullstorydev/grpcui"
	"github.com/fullstorydev/grpcui/standalone"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"
	"google.golang.org/grpc"
)

//go:generate make generate

// HandlerOption is an alias of github.com/fullstorydev/grpcui/standalone.HandlerOption to avoid multiple imports.
type HandlerOption = standalone.HandlerOption

// HandlerViaReflection is a wrapper around github.com/fullstorydev/grpcui/standalone.HandlerViaReflection to provide
// the theme.
func HandlerViaReflection(ctx context.Context, cc grpc.ClientConnInterface, target string, opts ...HandlerOption) (http.Handler, error) {
	m, err := grpcui.AllMethodsViaReflection(ctx, cc)
	if err != nil {
		return nil, err
	}

	f, err := grpcui.AllFilesViaReflection(ctx, cc)
	if err != nil {
		return nil, err
	}

	return Handler(cc, target, m, f, opts...), nil
}

// Handler is a wrapper around github.com/fullstorydev/grpcui/standalone.Handler to provide the theme.
func Handler(ch grpcdynamic.Channel, target string, methods []*desc.MethodDescriptor, files []*desc.FileDescriptor, opts ...HandlerOption) http.Handler {
	handlerOpts := make([]HandlerOption, 3, len(opts)+3)
	handlerOpts[0] = standalone.ServeAsset("theme.css", []byte(cssContent))
	handlerOpts[1] = standalone.AddCSS(fmt.Sprintf("theme.css?t=%d", time.Now().UnixMilli()), nil)
	handlerOpts[2] = standalone.AddCSS(fmt.Sprintf("description.css?t=%d", time.Now().UnixMilli()), nil)

	handlerOpts = append(handlerOpts, opts...)
	handler := standalone.Handler(ch, target, methods, files, handlerOpts...)

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Original logo is blurry on mac/edge, so render the original svg.
		if req.URL.Path == "/grpc-logo.png" {
			http.Redirect(w, req, "https://cncf-branding.netlify.app/img/projects/grpc/horizontal/white/grpc-horizontal-white.svg", http.StatusTemporaryRedirect)

			return
		}

		handler.ServeHTTP(w, req)
	})
}

// AddDescription adds description after the target section in the header.
func AddDescription(desc string) HandlerOption {
	const css = `.target::after {content: %q;}`

	return standalone.ServeAsset("description.css", []byte(fmt.Sprintf(css, desc)))
}
