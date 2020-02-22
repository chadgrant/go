package cmds

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/chadgrant/go-http-infra/infra/health"
	"github.com/spf13/cobra"
)

type ConfigGetter func(string) (interface{}, error)
type HealthCheckGetter func(interface{}) health.HealthChecker

func Register(rootCmd *cobra.Command, configLoader ConfigGetter, healthchecks HealthCheckGetter) {
	rootCmd.AddCommand(Config(configLoader))
	rootCmd.AddCommand(CreateConfig())
	rootCmd.AddCommand(Metadata())
	rootCmd.AddCommand(Version())
	rootCmd.AddCommand(Health(configLoader, healthchecks))
}

func enumerate(depth int, m map[string]interface{}) {
	for k, v := range m {
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Map {
			fmt.Printf("%s%s: \n", strings.Repeat("  ", depth), k)
			if im, ok := v.(map[string]interface{}); ok {
				enumerate(depth+1, im)
			}
		} else {
			fmt.Printf("%s%s: %v\n", strings.Repeat("  ", depth), k, v)
		}
	}
}
