package sshelper

import (
	"os"
	"path"
	"runtime"
)

func HashKnownHosts() bool {
	return false
}

// GlobalConfigFile returns the global ssh client config file name.
// See https://man.openbsd.org/ssh#FILES
func GlobalConfigFile() string {
	if runtime.GOOS == "windows" {
		return path.Join(os.Getenv("ALLUSERSPROFILE"), "/ssh/ssh_config")
	} else {
		return "/etc/ssh/ssh_config"
	}
}

// UserConfigFile returns the per-user ssh client config file name.
// See https://man.openbsd.org/ssh#FILES
func UserConfigFile() string {
	d, _ := os.UserHomeDir()
	return path.Join(d, ".ssh/config")
}
