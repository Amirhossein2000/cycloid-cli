package projects

import (
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "project",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
	}
	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	cmd.AddCommand(NewDeleteCommand(),
		NewCreateCommand(),
		NewListCommand(),
		NewDeleteEnvCommand(),
		NewCreateEnvCommand(),
		NewGetCommand())

	return cmd
}