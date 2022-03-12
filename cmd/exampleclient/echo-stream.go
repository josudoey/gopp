package exampleclient

import (
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/josudoey/gopp/gorpc"
	"github.com/spf13/cobra"
)

// see https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go
func NewEchoStreamCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "echo-stream <message>",
		Short: "rpc echo-stream",
		RunE: func(cmd *cobra.Command, argv []string) error {
			internalDuration, err := cmd.Flags().GetDuration("interval")
			if err != nil {
				return err
			}

			timeoutDuration, err := cmd.Flags().GetDuration("timeout")
			if err != nil {
				return err
			}

			client, err := dialExampleClient(cmd)
			if err != nil {
				return err
			}

			ctx := cmd.Context()
			stream, err := client.EchoStream(ctx)
			if err != nil {
				return err
			}
			req := &gorpc.EchoRequest{
				Message: strings.Join(argv, " "),
			}
			wg := sync.WaitGroup{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					log.Printf("begin stream recv")
					res, err := stream.Recv()
					log.Printf("end stream recv")
					if err == io.EOF {
						log.Printf("recv EOF")
						return
					}
					if err != nil {
						log.Printf("recv error %v", err)
						return
					}
					fmt.Printf("%s\n", res.GetMessage())
				}
			}()

			wg.Add(1)
			go func() {
				defer wg.Done()
				timeout := time.After(timeoutDuration)
				for {
					select {
					case <-timeout:
						if err := stream.CloseSend(); err != nil {
							log.Printf("close send error %v", err)
							return
						}
						log.Printf("close sent")
						return
					case <-time.After(internalDuration):
						if err := stream.Send(req); err != nil {
							if err == io.EOF {
								log.Printf("send EOF")
								return
							}
							log.Printf("send error %v", err)
							return
						}
						log.Printf("sent msg")
					}

				}
			}()

			<-stream.Context().Done()
			log.Printf("stream done")
			wg.Wait()
			return nil
		},
	}
	cmd.Flags().Duration("interval", time.Second, "time interval")
	cmd.Flags().Duration("timeout", 3*time.Second, "timeout")
	return cmd
}
