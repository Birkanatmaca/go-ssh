package config

import "fmt"

func (c Config) Validate() error {
	if c.SSHAddr == "" {
		return fmt.Errorf("missing -ssh")
	}
	if c.SSHUser == "" {
		return fmt.Errorf("missing -user")
	}
	if c.KeyPath == "" {
		return fmt.Errorf("missing -key")
	}
	if c.RemoteTargetAddr == "" {
		return fmt.Errorf("missing -remote")
	}
	if c.KnownHostsPath == "" && !c.InsecureIgnoreHostKey {
		return fmt.Errorf("missing -known-hosts or set -insecure-ignore-host-key")
	}
	return nil
}
