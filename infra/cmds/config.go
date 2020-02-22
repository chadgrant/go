package cmds

import (
	"github.com/spf13/cobra"
)

var cfgFile string

func Config(loader ConfigLoader) *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "displays config information",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := loader(cfgFile)
			if err != nil {
				return err
			}

			enumerate(0, cfg)
			return nil
		},
	}
}
