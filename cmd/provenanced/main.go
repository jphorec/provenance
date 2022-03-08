package main

import (
	//nolint:typecheck // this import is used
	"os"

	//nolint:typecheck // this import is used
	"github.com/cosmos/cosmos-sdk/server"

	"github.com/provenance-io/provenance/cmd/provenanced/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)
		default:
			os.Exit(1)
		}
	}
}
