package sshelper_test

import (
	"github.com/northbright/sshelper"
	"log"
	"runtime"
)

func Example() {
	// Show OS
	log.Printf("runtime.GOOS: %v", runtime.GOOS)

	// Get ssh global / user-specific config file.
	log.Printf("Global config file: %v", sshelper.GlobalConfigFile())
	log.Printf("User config file: %v", sshelper.UserConfigFile())

	// Output:
}
