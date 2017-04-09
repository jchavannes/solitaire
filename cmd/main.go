package cmd

import (
	"github.com/spf13/cobra"
	"github.com/jchavannes/solitaire/sol"
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

	verifyCmd = &cobra.Command{
		Use:   "verify",
		Short: "Run solitaire cli app",
		RunE: func (c *cobra.Command, args []string) error {
			game := sol.GetGame3()
			game.OutputMissingCards()
			return nil
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
	solitaireCmd.AddCommand(verifyCmd)
	solitaireCmd.AddCommand(webCmd)
	return solitaireCmd.Execute()
}
