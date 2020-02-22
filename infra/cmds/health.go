package cmds

import (
	"fmt"

	"github.com/chadgrant/go-http-infra/infra/health"

	"github.com/spf13/cobra"
)

func Health(config ConfigGetter, hc HealthCheckGetter) *cobra.Command {
	return &cobra.Command{
		Use:   "health",
		Short: "runs health checks and displays status",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config(cfgFile)
			if err != nil {
				return err
			}
			hc := hc(cfg)
			rep, err := hc.Report()
			if err != nil {
				return err
			}
			report("Liveness", rep.Liveness)
			report("Readiness", rep.Readiness)
			return nil
		},
	}
}

func report(name string, results []health.Result) {
	fmt.Printf("%s:\n", name)
	for _, t := range results {
		fmt.Printf("\t%s: %s\n", t.Name, t.Status)
	}
}
