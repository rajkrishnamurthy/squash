package cli

import (
	"context"

	"github.com/solo-io/squash/pkg/utils"
	"github.com/spf13/cobra"
)

func App(version string) (*cobra.Command, error) {
	app := &cobra.Command{
		Use:   "squash",
		Short: "debug microservices with squash",
		Long: `debug microservices with squash
	Find more information at https://solo.io`,
		Version: version,
	}

	opts := Options{}
	if err := initializeOptions(&opts); err != nil {
		return &cobra.Command{}, err
	}

	// pFlags := app.PersistentFlags()
	// pFlags.BoolVarP(&opts.Top.Static, "static", "s", false, "disable interactive mode")
	// pFlags.StringVarP(&opts.Top.File, "filename", "f", "", "file input")

	app.SuggestionsMinimumDistance = 1
	app.AddCommand(
		DebugContainerCmd(&opts),
		DebugRequestCmd(&opts),
		ListCmd(&opts),
		WaitAttCmd(&opts),
	)

	app.PersistentFlags().BoolVar(&opts.Json, "json", false, "output json format")

	// // Fail fast if we cannot connect to kubernetes
	// err := setup.CheckConnection()
	// if err != nil {
	// 	fmt.Println(errors.Wrap(err, "Failed to connect to Kubernetes. Please check whether the current-context "+
	// 		"in your kubeconfig file points to a running cluster"))
	// 	os.Exit(1)
	// }

	// err = setup.InitCache(&opts)
	// if err != nil {
	// 	fmt.Println(errors.Wrap(err, "Error during initialization!"))
	// 	os.Exit(1)
	// }

	// err = setup.InitSupergloo(&opts)
	// if err != nil {
	// 	fmt.Println(errors.Wrap(err, "Error during initialization!"))
	// 	os.Exit(1)
	// }

	return app, nil
}

func initializeOptions(o *Options) error {
	ctx := context.Background()
	daClient, err := utils.GetDebugAttachmentClient(ctx)
	if err != nil {
		return err
	}
	o.ctx = ctx
	o.daClient = daClient
	return nil
}