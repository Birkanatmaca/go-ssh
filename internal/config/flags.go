package config

import "flag"

func Parse() Config {
	var cfg Config

	flag.StringVar(&cfg.SSHAddr, "ssh", "", "SSH sunucu adresi (host:port) örn: 203.0.113.10:22")
	flag.StringVar(&cfg.SSHUser, "user", "", "SSH kullanıcı adı")
	flag.StringVar(&cfg.KeyPath, "key", "", "Private key yolu (örn: ~/.ssh/id_ed25519)")

	flag.StringVar(&cfg.LocalListenAddr, "local", "127.0.0.1:8080", "Local dinleme adresi (host:port)")
	flag.StringVar(&cfg.RemoteTargetAddr, "remote", "", "SSH üzerinden gidilecek hedef (host:port) örn: 127.0.0.1:80")

	flag.StringVar(&cfg.KnownHostsPath, "known-hosts", "", "Known hosts dosya yolu (örn: ~/.ssh/known_hosts)")
	flag.BoolVar(&cfg.InsecureIgnoreHostKey, "insecure-ignore-host-key", false, "Host anahtar doğrulamasını atla (güvenli değil!)")
	flag.IntVar(&cfg.KeepAliveSeconds, "keep-alive", 30, "SSH bağlantısı için keep-alive aralığı (saniye cinsinden)")
	flag.BoolVar(&cfg.Verbose, "verbose", false, "Ayrıntılı çıktı")
	flag.Parse()

	return cfg
}
