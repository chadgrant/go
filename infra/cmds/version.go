package cmds

import (
	"fmt"
	"time"

	"github.com/chadgrant/go-http-infra/infra/metadata"
	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"
)

func Version() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "displays version",
		RunE: func(cmd *cobra.Command, args []string) error {
			md := metadata.Get()
			t, err := time.Parse(time.RFC3339, md.BuildTime)
			if err != nil {
				return fmt.Errorf("parsing build time %v %s", err, md.BuildTime)
			}
			fmt.Printf("BuildNumber: %v\n", md.BuildNumber)
			fmt.Printf("Githash: %s\n", md.GitHash)
			fmt.Printf("Built: %s (%s)\n", md.BuildTime, humanize.Time(t))
			return nil
		},
	}
}
