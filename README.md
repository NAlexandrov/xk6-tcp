# xk6-tcp

A k6 extension for sending strings to TCP port

## Build

To build a `k6` binary with this plugin, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Install `xk6`:

```sh
go install go.k6.io/xk6/cmd/xk6@latest
```

2. Build the binary:

```sh
xk6 build --with github.com/NAlexandrov/xk6-tcp@latest
```

## Example

```javascript
import tcp from 'k6/x/tcp';
import { check } from 'k6';

const conn = tcp.connect('host:port');

// Or with timeout 5 seconds
// const conn = tcp.connect('host:port', 5000);

export default function () {
  tcp.writeLn(conn, 'Say Hello');

  let res = String.fromCharCode(...tcp.read(conn, 1024))

  // Or with timeout 5 seconds
  // let res = String.fromCharCode(...tcp.read(conn, 1024, 5000))

  check (res, {
    'verify ag tag': (res) => res.includes('Hello')
  });

  tcp.close(conn);
}
```
