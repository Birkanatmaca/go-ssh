package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	SSHAddr string
	SSHUser string
	KeyPath string

	LocalListenAddr  string
	RemoteTargetAddr string

	KnownHostsPath        string
	InsecureIgnoreHostKey bool
	KeepAliveSeconds      int
	Verbose               bool

	AskPassphrase bool
}

func expandHome(p string) string {
	if p == "" || p[0] != '~' { // If empty or not starting with ~, return as is
		return p
	}

	home, err := os.UserHomeDir() // Get user's home directory
	if err != nil {               // On error, return path as is
		return p
	}

	if p == "~" { // If path is exactly "~", return home directory
		return home
	}

	if len(p) >= 2 && p[1] == '/' { // If path starts with "~/", join home with the rest of the path
		return filepath.Join(home, p[2:])
	}
	return p
}
