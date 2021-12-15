package sshelper

import (
	"log"
	"runtime"
)

func Example() {
	// Show OS
	log.Printf("runtime.GOOS: %v", runtime.GOOS)

	// Get ssh global / user-specific config file.
	log.Printf("Global config file: %v", GlobalConfigFile())
	log.Printf("User config file: %v", UserConfigFile())

	// Output:
}
