package exampleserver

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/josudoey/gopp/gorpc"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

const (
	defaultPort    = 8080
	defaultAddress = "127.0.0.1"
)

type Options struct {
	port    int
	address string
}

func NewOptions() *Options {
	return &Options{
		port:    defaultPort,
		address: defaultAddress,
	}
}

var (
	_ gorpc.ExampleServiceServer = (*server)(nil)
)

// server is used to implement ExampleServiceServer
type server struct {
	gorpc.UnimplementedExampleServiceServer
}

// SayHello implements gorpc.ExampleServiceServer
func (s *server) Echo(ctx context.Context, req *gorpc.EchoRequest) (*gorpc.EchoResponse, error) {
	return &gorpc.EchoResponse{
		Message: req.GetMessage(),
	}, nil
}

// see https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go

func NewCommand() *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use:   "server",
		Short: "run example server",
		Run: func(cmd *cobra.Command, _ []string) {
			lis, err := net.Listen("tcp", fmt.Sprintf(":%d", o.port))
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			s := grpc.NewServer()
			gorpc.RegisterExampleServiceServer(s, &server{})
			log.Printf("server listening at %v", lis.Addr())
			if err := s.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
		},
	}
	cmd.Flags().IntVarP(&o.port, "port", "p", o.port, "The port on which to run the proxy. Set to 0 to pick a random port.")
	cmd.Flags().StringVar(&o.address, "address", o.address, "The IP address on which to serve on.")

	return cmd
}
