package cmds

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	outfile    string
	configType string
)

func CreateConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "createconfig",
		Short: "Creates a stub config file prepopulated with default values",
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, err := os.Stat(outfile); err == nil {
				return err
			}
			viper.SetConfigType(configType)
			fmt.Printf("Creating %s Config: %s\n", configType, outfile)
			return viper.WriteConfigAs(outfile)
		},
	}

	cmd.Flags().StringVarP(&outfile, "output", "o", "config.yaml", "output file name")
	cmd.Flags().StringVarP(&configType, "type", "t", "YAML", "YAML,JSON,TOML")
	return cmd
}
