package http

import (
	"log"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:          "http",
	SilenceUsage: true,
	Short:        "Start http Server",
	Run: func(cmd *cobra.Command, args []string) {
		cfgFile, err := cmd.InheritedFlags().GetString("config")
		if err != nil {
			log.Fatal(err)
		}
		app, err := Initialize(cfgFile)
		if err != nil {
			log.Fatal(err)
		}
		app.Launch()
	},
}
