// +build darwin freebsd linux netbsd openbsd

package main

import (
	"net"
	"syscall"
)

func isErrConnReset(err error) bool {
	if ne, ok := err.(*net.OpError); ok {
		return ne.Err == syscall.ECONNRESET
	}
	return false
}

func isDNSError(err error) bool {
	if _, ok := err.(*net.DNSError); ok {
		return true
	}
	return false
}
