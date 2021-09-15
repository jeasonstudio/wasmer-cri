package utils

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"strings"
	"syscall"
)

// GetListener get a listener for an address.
func GetListener(addr string, tlsConfig *tls.Config) (net.Listener, error) {
	addrParts := strings.SplitN(addr, "://", 2)
	if len(addrParts) != 2 {
		return nil, fmt.Errorf("invalid listening address %s: must be in format [protocol]://[address]", addr)
	}

	switch addrParts[0] {
	case "tcp":
		l, err := net.Listen("tcp", addrParts[1])
		if err != nil {
			return l, err
		}
		if tlsConfig != nil {
			l = tls.NewListener(l, tlsConfig)
		}
		return l, err
	case "unix":
		return newUnixSocket(addrParts[1])

	default:
		return nil, fmt.Errorf("only unix socket or tcp address is support")
	}
}

func newUnixSocket(path string) (net.Listener, error) {
	if err := syscall.Unlink(path); err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	oldmask := syscall.Umask(0777)
	defer syscall.Umask(oldmask)
	l, err := net.Listen("unix", path)
	if err != nil {
		return nil, err
	}

	// chmod unix socket, make other group writable
	if err := os.Chmod(path, 0660); err != nil {
		l.Close()
		return nil, fmt.Errorf("failed to chmod %s: %s", path, err)
	}

	gid, err := ParseID(GroupFile, "wasmer", func(line, str string, idInt int, idErr error) (uint32, bool) {
		var (
			name, placeholder string
			id                int
		)

		ParseString(line, &name, &placeholder, &id)
		if str == name {
			return uint32(id), true
		}
		return 0, false
	})
	if err != nil {
		// ignore error when group wasmer not exist, group wasmer should to be
		// created before wasmerd started, it means code not create wasmer group
		fmt.Printf("failed to find group wasmer, cannot change unix socket %s to wasmer group", path)
		return l, nil
	}

	// chown unix socket with group wasmer
	if err := os.Chown(path, 0, int(gid)); err != nil {
		l.Close()
		return nil, fmt.Errorf("failed to chown %s: %s", path, err)
	}
	return l, nil
}
