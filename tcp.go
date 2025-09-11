package tcp

import (
	"net"

	"go.k6.io/k6/js/modules"
)

// RootModule is the global module instance that will create module instances for each VU.
type RootModule struct{}

// Ensure the interfaces are implemented correctly.
var _ modules.Module = &RootModule{}

type TCP struct {
	vu modules.VU // provides methods for accessing internal k6 objects
}

// init is called by the Go runtime at application startup.
func init() {
	modules.Register("k6/x/tcp", new(RootModule))
}

// NewModuleInstance implements the modules.Module interface returning a new instance for each VU.
func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &TCP{vu: vu}
}

func (tcp *TCP) Exports() modules.Exports {
	return modules.Exports{Default: tcp}
}

func (tcp *TCP) Connect(addr string) (net.Conn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (tcp *TCP) Write(conn net.Conn, data []byte) error {
	_, err := conn.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (tcp *TCP) Read(conn net.Conn, size int) ([]byte, error) {
	buf := make([]byte, size)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}

func (tcp *TCP) WriteLn(conn net.Conn, data []byte) error {
	return tcp.Write(conn, append(data, []byte("\n")...))
}

func (tcp *TCP) Close(conn net.Conn) error {
	err := conn.Close()
	if err != nil {
		return err
	}
	return nil
}
