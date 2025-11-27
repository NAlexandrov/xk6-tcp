package tcp

import (
	"net"
	"time"

	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/metrics"
)

// RootModule is the global module instance that will create module instances for each VU.
type RootModule struct{}

// Ensure the interfaces are implemented correctly.
var _ modules.Module = &RootModule{}

type TCP struct {
	vu    modules.VU // provides methods for accessing internal k6 objects
	dsCtr *metrics.Metric
	drCtr *metrics.Metric
}

// init is called by the Go runtime at application startup.
func init() {
	modules.Register("k6/x/tcp", new(RootModule))
}

// NewModuleInstance implements the modules.Module interface returning a new instance for each VU.
func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	e := vu.InitEnv()
	if e == nil {
		panic("vu env is nil")
	}
	dsCtr := e.Registry.Get("data_sent")
	drCtr := e.Registry.Get("data_received")
	return &TCP{vu: vu, dsCtr: dsCtr, drCtr: drCtr}
}

func (tcp *TCP) Exports() modules.Exports {
	return modules.Exports{Default: tcp}
}

func (tcp *TCP) Connect(addr string, timeoutMs ...int) (net.Conn, error) {
	timeout := 60 * time.Second // default timeout

	if len(timeoutMs) > 0 {
		if timeoutMs[0] > 0 {
			timeout = time.Duration(timeoutMs[0]) * time.Millisecond
		}
	}

	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (tcp *TCP) Write(conn net.Conn, data []byte, timeoutMs ...int) error {
	timeout := 60 * time.Second // default timeout

	if len(timeoutMs) > 0 {
		if timeoutMs[0] > 0 {
			timeout = time.Duration(timeoutMs[0]) * time.Millisecond
		}
	}

	err := conn.SetReadDeadline(time.Now().Add(timeout))
	if err != nil {
		return err
	}

	n, err := conn.Write(data)
	if err != nil {
		return err
	}

	state := tcp.vu.State()
	tags := state.Tags.GetCurrentValues().Tags
	metrics.PushIfNotDone(tcp.vu.Context(), state.Samples, metrics.Sample{
		TimeSeries: metrics.TimeSeries{Metric: tcp.dsCtr, Tags: tags},
		Value:      float64(n),
		Time:       time.Now(),
	})

	return nil
}

func (tcp *TCP) Read(conn net.Conn, size int, timeoutMs ...int) ([]byte, error) {
	timeout := 60 * time.Second // default timeout

	if len(timeoutMs) > 0 {
		if timeoutMs[0] > 0 {
			timeout = time.Duration(timeoutMs[0]) * time.Millisecond
		}
	}

	err := conn.SetReadDeadline(time.Now().Add(timeout))
	if err != nil {
		return nil, err
	}

	buf := make([]byte, size)

	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}

	state := tcp.vu.State()
	tags := state.Tags.GetCurrentValues().Tags
	metrics.PushIfNotDone(tcp.vu.Context(), state.Samples, metrics.Sample{
		TimeSeries: metrics.TimeSeries{Metric: tcp.drCtr, Tags: tags},
		Value:      float64(n),
		Time:       time.Now(),
	})

	return buf[:n], nil
}

func (tcp *TCP) WriteLn(conn net.Conn, data []byte, timeoutMs ...int) error {
	return tcp.Write(conn, append(data, []byte("\n")...), timeoutMs...)
}

func (tcp *TCP) Close(conn net.Conn) error {
	err := conn.Close()
	if err != nil {
		return err
	}
	return nil
}
