package main

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	storev1 "ticketbooth/proto/store/v1"
	"ticketbooth/proto/store/v1/storev1connect"
)

type GreetServer struct{}

func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[storev1.GreetRequest],
) (*connect.Response[storev1.GreetResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&storev1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func main() {
	mux := http.NewServeMux()

	greeter := &GreetServer{}
	path, greetSvcHandler := storev1connect.NewGreetServiceHandler(greeter)
	mux.Handle(path, greetSvcHandler)

	handler := cors.AllowAll().Handler(mux)

	http.ListenAndServe(
		"0.0.0.0:9100",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(handler, &http2.Server{}),
	)
}
