# Redoc Theme for gRPC UI

[![GitHub Releases](https://img.shields.io/github/v/release/nhatthm/grpcui-redoc)](https://github.com/nhatthm/grpcui-redoc/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/nhatthm/grpcui-redoc)](https://goreportcard.com/report/github.com/nhatthm/grpcui-redoc)
[![GoDevDoc](https://img.shields.io/badge/dev-doc-00ADD8?logo=go)](https://pkg.go.dev/github.com/nhatthm/grpcui-redoc)
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/donate/?hosted_button_id=PJZSGJN57TDJY)

Redoc Theme for [gRPC UI](https://github.com/fullstorydev/grpcui/)

[Don't gRPC without it!](https://www.fullstory.com/blog/grpcui-dont-grpc-without-it/)

## Prerequisites

- `Go >= 1.17`

## Install

```bash
go get github.com/nhatthm/grpcui-redoc
```

## Usage

Use `grpcuiredoc.HandlerViaReflection()` or `grpcuiredoc.Handler()` the same way as `fullstorydev/grpcui`.

For example:

```go
package exmaple

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	grpcuiredoc "github.com/nhatthm/grpcui-redoc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// ----------------------------
	// Start up our gRPC server
	// ----------------------------
	svr := grpc.NewServer()

	// we need the reflection service, to power the UI
	reflection.Register(svr)

	// ...
	// register our gRPC services and then start up the server in a goroutine
	// ...

	// ----------------------------
	// Create a client to local server
	// ----------------------------
	cc, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", grpcPort))
	if err != nil {
		panic(fmt.Errorf("failed to create client to local server: %w", err))
	}

	// ----------------------------
	// Create gRPCui handler
	// ----------------------------
	target := fmt.Sprintf("%s:%d", filepath.Base(os.Args[0]), grpcPort)
	// This one line of code is all that is needed to create the UI handler!
	h, err := grpcuiredoc.HandlerViaReflection(ctx, cc, target)
	if err != nil {
		panic(fmt.Errorf("failed to create client to local server: %w", err))
	}

	// ----------------------------
	// Now wire it up to an HTTP server
	// ----------------------------
	serveMux := http.NewServeMux()

	serveMux.Handle("/grpcui/", http.StripPrefix("/grpcui", h))
}
```

![image](https://user-images.githubusercontent.com/1154587/153202990-13515fc3-9a9e-433c-909e-101d167381ae.png)

## Donation

If this project help you reduce time to develop, you can give me a cup of coffee :)

### Paypal donation

[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/donate/?hosted_button_id=PJZSGJN57TDJY)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;or scan this

<img src="https://user-images.githubusercontent.com/1154587/113494222-ad8cb200-94e6-11eb-9ef3-eb883ada222a.png" width="147px" />
