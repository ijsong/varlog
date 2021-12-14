package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v2"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/kakao/varlog/internal/varlogctl"
	"github.com/kakao/varlog/pkg/varlog"
)

func main() {
	os.Exit(run())
}

func run() int {
	app := newVarlogControllerApp()
	if err := app.Run(os.Args); err != nil {
		return -1
	}
	return 0
}

func execute(c *cli.Context, f varlogctl.ExecuteFunc) (err error) {
	var logger *zap.Logger
	if c.Bool(flagVerbose.name) {
		logger, err = zap.NewProduction()
		if err != nil {
			return err
		}
	} else {
		logger = zap.NewNop()
	}
	defer func() {
		_ = logger.Sync()
	}()

	timeout := c.Duration(flagTimeout.name)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	adminAddr := c.String(flagAdminAddress.name)
	admin, err := varlog.NewAdmin(ctx, adminAddr)
	if err != nil {
		return err
	}

	opts := []varlogctl.Option{
		varlogctl.WithAdmin(admin),
		varlogctl.WithExecuteFunc(f),
		varlogctl.WithLogger(logger),
	}
	if c.Bool(flagPrettyPrint.name) {
		opts = append(opts, varlogctl.WithPrettyPrint())
	}

	ctl, err := varlogctl.New(opts...)
	if err != nil {
		return err
	}
	res := ctl.Execute(ctx)
	return multierr.Append(res.Err(), ctl.Print(res, os.Stdout))
}
