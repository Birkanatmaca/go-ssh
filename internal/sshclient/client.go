package sshclient

import (
	"context"
	"net"
	"os"
	"time"

	"github.com/birkan-is/go-ssh/internal/config"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

func Dial(ctx context.Context, cfg config.Config) (*ssh.Client, error) {
	signer, err := readSigner(cfg.KeyPath) // Read private key who am I?
	if err != nil {
		return nil, err
	}

	hk, err := hostKeyCallback(cfg) // The other side is true?
	if err != nil {
		return nil, err
	}

	sshCfg := &ssh.ClientConfig{
		User:            cfg.SSHUser,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: hk,
		Timeout:         10 * time.Second,
	}

	dialer := net.Dialer{Timeout: 10 * time.Second}
	conn, err := dialer.DialContext(ctx, "tcp", cfg.SSHAddr) // TCP connection create
	if err != nil {
		return nil, err
	}

	cc, chans, reqs, err := ssh.NewClientConn(conn, cfg.SSHAddr, sshCfg) // SSH client create
	if err != nil {
		_ = conn.Close() // Close connection on error
		return nil, err
	}
	return ssh.NewClient(cc, chans, reqs), nil // Return SSH client
}

func readSigner(keyPath string) (ssh.Signer, error) {
	keyBytes, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(keyBytes)
	if err != nil {
		return nil, err
	}

	return signer, nil
}

func hostKeyCallback(cfg config.Config) (ssh.HostKeyCallback, error) {
	if cfg.InsecureIgnoreHostKey {
		return ssh.InsecureIgnoreHostKey(), nil
	}
	return knownhosts.New(cfg.KnownHostsPath)
}
