package main

import "github.com/tken2039/montyhall/internal/cmd"

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		cmd.HandleCmdErr(err)
	}
}
