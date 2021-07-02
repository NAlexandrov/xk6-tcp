package tcp

import (
	"net"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/tcp", new(TCP))
}

type TCP struct{}

func (TCP) Connect(addr string) (net.Conn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (TCP) Write(conn net.Conn, data string) error {
	_, err := conn.Write([]byte(data + "\n"))
	if err != nil {
		return err
	}

	return nil
}
