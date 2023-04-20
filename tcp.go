package tcp

import (
	"go.k6.io/k6/js/modules"
	"net"
	"time"
)

func init() {
	modules.Register("k6/x/tcp", new(TCP))
}

type TCP struct{}

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

func (tcp *TCP) Read(conn net.Conn, size int, timeout_opt ...int) ([]byte, error) {
	timeout_ms := 0
	if len(timeout_opt) > 0 {
		timeout_ms = timeout_opt[0]
	}
	err := conn.SetReadDeadline(time.Now().Add(time.Millisecond * time.Duration(timeout_ms)))
	if err != nil {
		return nil, err
	}
	buf := make([]byte, size)
	_, err = conn.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
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
