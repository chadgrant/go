package cmds

import (
	"github.com/chadgrant/go-http-infra/infra/metadata"
	"github.com/fatih/structs"

	"github.com/spf13/cobra"
)

func Metadata() *cobra.Command {
	return &cobra.Command{
		Use:   "metadata",
		Short: "displays metadata",
		Run: func(cmd *cobra.Command, args []string) {
			enumerate(0, structs.Map(metadata.Get()))
		},
	}
}
