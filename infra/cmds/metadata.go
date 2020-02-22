package cmds

import (
	"github.com/chadgrant/go-http-infra/infra/metadata"

	"github.com/spf13/cobra"
)

func Metadata() *cobra.Command {
	return &cobra.Command{
		Use:   "metadata",
		Short: "displays metadata",
		Run: func(cmd *cobra.Command, args []string) {
			md := metadata.Get()
			enumerate(0, *md)
		},
	}
}
