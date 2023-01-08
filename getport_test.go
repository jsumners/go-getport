package go_getport

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPort(t *testing.T) {
	t.Run("gets empty tcp port", func(t *testing.T) {
		portResult, err := GetPort(TCP, "127.0.0.1")
		assert.NoError(t, err)
		assert.Contains(t, []string{"127.0.0.1", "::1"}, portResult.IP)
		assert.Greater(t, portResult.Port, 0)
	})

	t.Run("gets empty tcp4 port", func(t *testing.T) {
		portResult, err := GetPort(TCP4, "127.0.0.1")
		assert.NoError(t, err)
		assert.Equal(t, portResult.IP, "127.0.0.1")
		assert.Greater(t, portResult.Port, 0)
	})

	t.Run("gets empty tcp6 port", func(t *testing.T) {
		portResult, err := GetPort(TCP6, "::1")
		assert.NoError(t, err)
		assert.Equal(t, portResult.IP, "::1")
		assert.Greater(t, portResult.Port, 0)
	})

	t.Run("gets empty udp port", func(t *testing.T) {
		portResult, err := GetPort(UDP, "127.0.0.1")
		assert.NoError(t, err)
		assert.Contains(t, []string{"127.0.0.1", "::1"}, portResult.IP)
		assert.Greater(t, portResult.Port, 0)
	})

	t.Run("gets empty udp4 port", func(t *testing.T) {
		portResult, err := GetPort(UDP4, "127.0.0.1")
		assert.NoError(t, err)
		assert.Equal(t, portResult.IP, "127.0.0.1")
		assert.Greater(t, portResult.Port, 0)
	})

	t.Run("gets empty udp6 port", func(t *testing.T) {
		portResult, err := GetPort(UDP6, "::1")
		assert.NoError(t, err)
		assert.Equal(t, portResult.IP, "::1")
		assert.Greater(t, portResult.Port, 0)
	})

	t.Run("errors for bad protocol", func(t *testing.T) {
		portResult, err := GetPort(-1, "127.0.0.1")
		assert.Equal(t, portResult.IP, "")
		assert.Equal(t, portResult.Port, -1)
		assert.Error(t, err, "stack not recognized")
	})
}

func TestGetTcpPort(t *testing.T) {
	portResult, err := GetTcpPort()
	assert.NoError(t, err)
	assert.Greater(t, portResult.Port, 0)
}

func TestGetTcp4Port(t *testing.T) {
	portResult, err := GetTcp4Port()
	assert.NoError(t, err)
	assert.Greater(t, portResult.Port, 0)
}

func TestGetTcp6Port(t *testing.T) {
	portResult, err := GetTcp6Port()
	assert.NoError(t, err)
	assert.Greater(t, portResult.Port, 0)
}

func TestGetUdpPort(t *testing.T) {
	portResult, err := GetUdpPort()
	assert.NoError(t, err)
	assert.Greater(t, portResult.Port, 0)
}

func TestGetUdp4Port(t *testing.T) {
	portResult, err := GetUdp4Port()
	assert.NoError(t, err)
	assert.Greater(t, portResult.Port, 0)
}

func TestGetUdp6Port(t *testing.T) {
	portResult, err := GetUdp6Port()
	assert.NoError(t, err)
	assert.Greater(t, portResult.Port, 0)
}

func TestGetTcpPortForAddress(t *testing.T) {
	portResult, err := GetTcpPortForAddress("127.0.0.1")
	assert.NoError(t, err)
	assert.Greater(t, portResult.Port, 0)
}

func TestGetTcp4PortForAddress(t *testing.T) {
	portResult, err := GetTcp4PortForAddress("127.0.0.1")
	assert.NoError(t, err)
	assert.Greater(t, portResult.Port, 0)
}

func TestGetTcp6PortForAddress(t *testing.T) {
	portResult, err := GetTcp6PortForAddress("::1")
	assert.NoError(t, err)
	assert.Greater(t, portResult.Port, 0)
}

func TestGetUdpPortForAddress(t *testing.T) {
	portResult, err := GetUdpPortForAddress("127.0.0.1")
	assert.NoError(t, err)
	assert.Greater(t, portResult.Port, 0)
}

func TestGetUdp4PortForAddress(t *testing.T) {
	portResult, err := GetUdp4PortForAddress("127.0.0.1")
	assert.NoError(t, err)
	assert.Greater(t, portResult.Port, 0)
}

func TestGetUdp6PortForAddress(t *testing.T) {
	portResult, err := GetUdp6PortForAddress("::1")
	assert.NoError(t, err)
	assert.Greater(t, portResult.Port, 0)
}

func TestListen(t *testing.T) {
	t.Run("errors for bad tcp stack", func(t *testing.T) {
		addr, err := listen("tcp-bad", "127.0.0.1:0")
		assert.Empty(t, addr)
		assert.Error(t, err, "listen tcp-bad: unknown network tcp-bad")
	})

	t.Run("errors for bad udp stack", func(t *testing.T) {
		addr, err := listen("udp-bad", "127.0.0.1:0")
		assert.Empty(t, addr)
		assert.Error(t, err, "listen udp-bad: unknown network udp-bad")
	})

	t.Run("errors for bad stack", func(t *testing.T) {
		addr, err := listen("unix", "127.0.0.1:0")
		assert.Empty(t, addr)
		assert.Error(t, err, "stack not recognized")
	})
}
