package exampleserver

import (
	"context"
	"fmt"
	"io"
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
	_ gorpc.ExampleServer = (*server)(nil)
)

// server is used to implement ExampleServiceServer
type server struct {
	gorpc.UnimplementedExampleServer
}

// Echo implements gorpc.ExampleServiceServer
func (s *server) Echo(ctx context.Context, req *gorpc.EchoRequest) (*gorpc.EchoResponse, error) {
	return &gorpc.EchoResponse{
		Message: req.GetMessage(),
	}, nil
}

func (s *server) EchoStream(stream gorpc.Example_EchoStreamServer) error {
	// see https://github.com/grpc/grpc-go/blob/master/examples/route_guide/server/server.go#L92
	msgCount := 0
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Printf("recv EOF")
				return nil
			}
			log.Printf("recv error %v", err)
			return err
		}
		msgCount++
		log.Printf("recv(%v) %v", msgCount, req.Message)
		if err := stream.Send(&gorpc.EchoResponse{
			Message: req.Message,
		}); err != nil {
			log.Printf("send error %v", err)
			return err
		}
	}
}

// see https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go

func NewCommand() *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use:   "example-server",
		Short: "run example server",
		Run: func(cmd *cobra.Command, _ []string) {
			lis, err := net.Listen("tcp", fmt.Sprintf(":%d", o.port))
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			s := grpc.NewServer()
			gorpc.RegisterExampleServer(s, &server{})
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
