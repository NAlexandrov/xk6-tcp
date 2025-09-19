import tcp from 'k6/x/tcp';
import { check } from 'k6';

const conn = tcp.connect('host:port');

// Or with timeout 5 seconds
// const conn = tcp.connect('host:port', 5000);

export default function () {
  tcp.writeLn(conn, 'Say Hello');

  // Or with timeout 5 seconds
  // tcp.writeLn(conn, 'Say Hello', 5000);

  let res = String.fromCharCode(...tcp.read(conn, 1024))

  // Or with timeout 5 seconds
  // let res = String.fromCharCode(...tcp.read(conn, 1024, 5000))

  check (res, {
    'verify ag tag': (res) => res.includes('Hello')
  });

  tcp.close(conn);
}