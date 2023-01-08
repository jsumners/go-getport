package go_getport

import (
	"errors"
	"github.com/spf13/cast"
	"net"
	"strings"
)

// Protocol indicates the communication protocol (tcp or udp) and network
// stack (IPv4, IPv6, or OS choice) to target when finding an available port.
type Protocol int

const (
	// TCP indicates to let the OS decide between IPv4 and IPv6 when finding
	// an open TCP based port.
	TCP Protocol = iota

	// TCP4 indicates to find an open IPv4 port.
	TCP4

	// TCP6 indicates to find an open IPv6 port.
	TCP6

	// UDP indicates to let the OS decide between IPv4 and IPv6 when finding
	// an open UDP based port.
	UDP

	// UDP4 indicates to find an open IPv4 port.
	UDP4

	// UDP6 indicates to find an open IPv6 port.
	UDP6
)

// PortResult represents the result of [GetPort]. It indicates the IP address
// and port number combination that resulted in finding an open port.
type PortResult struct {
	// IP is either an IPv4 or IPv6 string as returned by [net.SplitHostPort].
	IP string

	// Port is the determined available port number.
	Port int
}

// GetPort finds an open port for a given [Protocol] and address and returns
// that port number. If the [Protocol] is not recognized, or some problem is
// encountered while verifying the port, then the returned [PortResult.Port]
// number will be `-1` along with an error. The address parameter should be a
// simple IP address string, e.g. `127.0.0.1` or `::1`. The [PortResult.IP] will
// be set to the IP address that was actually used to find the open port. If
// address is the empty string (`""`), then the returned IP address will be the
// one determined by the OS when finding the port.
//
// Note: it is not guaranteed the port will remain open long enough to actually
// be used. Errors should still be checked when attempting to use the found
// port.
func GetPort(protocol Protocol, address string) (PortResult, error) {
	stack := resolveProtocol(protocol)

	result := PortResult{
		IP:   "",
		Port: -1,
	}
	resolvedAddress, listenError := listen(
		stack,
		net.JoinHostPort(address, "0"),
	)
	if listenError != nil {
		return result, listenError
	}

	// I do not see how it's possible to get an error from [net.SplitHostPort]
	// here given how we have already validated the stack and successfully
	// issued a [net.Listen].
	addr, port, _ := net.SplitHostPort(resolvedAddress.String())
	result.IP = addr
	result.Port = cast.ToInt(port)

	return result, nil
}

// GetTcpPort gets a port for some random available address using either
// TCP4 or TCP6. See [GetPort] for more detail
func GetTcpPort() (PortResult, error) {
	return GetPort(TCP, "")
}

// GetTcp4Port gets a port for some random available address using TCP4.
// See [GetPort] for more detail
func GetTcp4Port() (PortResult, error) {
	return GetPort(TCP4, "")
}

// GetTcp6Port gets a port for some random available address using TCP6.
// See [GetPort] for more detail
func GetTcp6Port() (PortResult, error) {
	return GetPort(TCP6, "")
}

// GetUdpPort gets a port for some random available address using either
// UDP4 or UDP6. See [GetPort] for more detail
func GetUdpPort() (PortResult, error) {
	return GetPort(UDP, "")
}

// GetUdp4Port gets a port for some random available address using UDP4.
// See [GetPort] for more detail
func GetUdp4Port() (PortResult, error) {
	return GetPort(UDP4, "")
}

// GetUdp6Port gets a port for some random available address using UDP6.
// See [GetPort] for more detail
func GetUdp6Port() (PortResult, error) {
	return GetPort(UDP6, "")
}

// GetTcpPortForAddress gets either a TCP4 or TCP6 port for the given address.
// See [GetPort] for more detail.
func GetTcpPortForAddress(address string) (PortResult, error) {
	return GetPort(TCP, address)
}

// GetTcp4PortForAddress gets a TCP4 port for the given address.
// See [GetPort] for more detail.
func GetTcp4PortForAddress(address string) (PortResult, error) {
	return GetPort(TCP4, address)
}

// GetTcp6PortForAddress gets a TCP6 port for the given address.
// See [GetPort] for more detail.
func GetTcp6PortForAddress(address string) (PortResult, error) {
	return GetPort(TCP6, address)
}

// GetUdpPortForAddress gets either a UDP4 or UDP6 port for the given address.
// See [GetPort] for more detail.
func GetUdpPortForAddress(address string) (PortResult, error) {
	return GetPort(UDP, address)
}

// GetUdp4PortForAddress gets a UDP4 port for the given address.
// See [GetPort] for more detail.
func GetUdp4PortForAddress(address string) (PortResult, error) {
	return GetPort(UDP4, address)
}

// GetUdp6PortForAddress gets a UDP6 port for the given address.
// See [GetPort] for more detail.
func GetUdp6PortForAddress(address string) (PortResult, error) {
	return GetPort(UDP6, address)
}

// listen is an internal wrapper for [net.Listen] and [net.ListenPacket].
func listen(stack string, addrWithPort string) (net.Addr, error) {
	if strings.HasPrefix(stack, "tcp") {
		l, err := net.Listen(stack, addrWithPort)
		if err != nil {
			return nil, err
		}
		defer l.Close()
		return l.Addr(), nil
	}

	if strings.HasPrefix(stack, "udp") {
		l, err := net.ListenPacket(stack, addrWithPort)
		if err != nil {
			return nil, err
		}
		defer l.Close()
		return l.LocalAddr(), nil
	}

	return nil, errors.New("stack not recognized")
}

// resolveProtocol maps the [Protocol] value to a network stack string
// that is supported by [net.Listen] and [net.ListenPacket].
func resolveProtocol(protocol Protocol) string {
	var stack string

	switch protocol {
	case TCP:
		stack = "tcp"
	case TCP4:
		stack = "tcp4"
	case TCP6:
		stack = "tcp6"
	case UDP:
		stack = "udp"
	case UDP4:
		stack = "udp4"
	case UDP6:
		stack = "udp6"
	}

	return stack
}
