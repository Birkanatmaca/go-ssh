package config

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
}
