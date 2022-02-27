package main

import (
	"log"

	"github.com/josudoey/gopp/cmd/exampleclient"
	"github.com/josudoey/gopp/cmd/exampleserver"
	"github.com/spf13/cobra"
)

func newCommand() *cobra.Command {
	root := &cobra.Command{
		Use: "gopp",
		CompletionOptions: cobra.CompletionOptions{
			// see https://github.com/spf13/cobra/blob/9054739e08187aab9294b7a773d54c92fabc23d3/completions.go#L599
			DisableDefaultCmd: true,
		},
	}
	root.AddCommand(exampleserver.NewCommand())
	root.AddCommand(exampleclient.NewCommand())
	return root
}

func main() {
	cmd := newCommand()
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
