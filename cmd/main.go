package cmd

import (
	"github.com/spf13/cobra"
)

var (
	solitaireCmd = &cobra.Command{
		Use:   "solitaire",
		Short: "Run solitaire app",
	}

	cliCmd = &cobra.Command{
		Use:   "cli",
		Short: "Run solitaire cli app",
		RunE: func (c *cobra.Command, args []string) error {
			return runCli()
		},
	}

	webCmd = &cobra.Command{
		Use:   "web",
		Short: "Run solitaire web app",
		RunE: func (c *cobra.Command, args []string) error {
			return runWeb()
		},
	}
)

func Run() error {
	solitaireCmd.AddCommand(cliCmd)
	solitaireCmd.AddCommand(webCmd)
	return solitaireCmd.Execute()
}
