package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"gitlab.silkrode.com.tw/team_golang/kbc2/sample/cmd/http"
)

var httpCmd = &cobra.Command{
	Use:          "http",
	SilenceUsage: true,
	Short:        "Start http Server",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := http.Initialize(cfgFile)
		if err != nil {
			log.Fatal(err)
		}
		app.Launch()
	},
}
