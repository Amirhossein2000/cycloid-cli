package creds

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "delete a credential",
		Example: `
	# delete a credential with ID 123
	cy --org my-org creds delete --id 123
`,
		RunE:    del,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	common.RequiredFlag(common.WithFlagID, cmd)
	return cmd
}

func del(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	id, err := cmd.Flags().GetUint32("id")
	if err != nil {
		return err
	}

	if err := m.DeleteCredential(org, id); err != nil {
		return errors.Wrap(err, "unable to delete credential")
	}
	return nil
}
