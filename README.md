# xk6-tcp

A k6 extension for sending strings to TCP port

## Build

To build a `k6` binary with this plugin, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Install `xk6`:

  ```shell
  go install github.com/k6io/xk6/cmd/xk6@latest
  ```

2. Build the binary:

  ```shell
  xk6 build master \
    --with github.com/NAlexandrov/xk6-tcp
  ```

## Example

```javascript
import tcp from 'k6/x/tcp';
import { check } from 'k6';

const conn = tcp.connect('host:port');

export default function () {
  tcp.writeLn(conn, 'Say Hello');
  let res = String.fromCharCode(...tcp.read(conn, 1024))
  check (res, {
    'verify ag tag': (res) => res.includes('Hello')
  });
  tcp.close(conn);
}
```
