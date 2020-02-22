package cmds

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/chadgrant/go-http-infra/infra/health"
	"github.com/spf13/cobra"
)

type ConfigLoader func(string) (interface{}, error)

func RegisterDefault(rootCmd *cobra.Command, loader ConfigLoader, hc health.HealthChecker) {
	rootCmd.AddCommand(Config(loader))
	rootCmd.AddCommand(CreateConfig())
	rootCmd.AddCommand(Metadata())
	rootCmd.AddCommand(Version())
	rootCmd.AddCommand(Health(hc))
}

func enumerate(depth int, obj interface{}) {
	otype := reflect.TypeOf(obj)
	oval := reflect.ValueOf(obj)
	for i := 0; i < otype.NumField(); i++ {
		f := otype.Field(i)
		v := oval.FieldByName(f.Name)
		if f.Type.Kind() == reflect.Struct {
			fmt.Printf("%s%s: \n", strings.Repeat("  ", depth), f.Name)
			enumerate(depth+1, v.Interface())
		} else {
			fmt.Printf("%s%s: %v\n", strings.Repeat("  ", depth), f.Name, v)
		}
	}
}
