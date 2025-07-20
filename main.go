package main

import (
	"github.com/beyondEllie/elliecode/cmd"
	"github.com/beyondEllie/elliecode/internal/logging"
)

func main() {
	defer logging.RecoverPanic("main", func() {
		logging.ErrorPersist("Application terminated due to unhandled panic")
	})

	cmd.Execute()
}
