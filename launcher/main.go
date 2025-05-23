package main

import (
	"github.com/luskaner/ageLANServer/common"
	"github.com/luskaner/ageLANServer/launcher/internal/cmd"
	"github.com/spf13/cobra"
)

const version = "development"

func main() {
	cobra.MousetrapHelpText = ""
	cmd.Version = version
	common.ChdirToExe()
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
