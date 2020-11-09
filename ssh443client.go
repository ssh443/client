package main
import "bufio"
import "crypto/tls"
import "encoding/base64"
import "errors"
import "flag"
import "fmt"
import "io"
import "net"
import "os"
import "strings"

type args struct {
  address string
  auth string
  http bool
  proxy string
}

func cliArgs() args {
  address := flag.String("address", "", "<host>:<port>")
  auth := flag.String("auth", "", "<token>:<secret>")
  proxy := flag.String("proxy", "", "<host>:<port>")
  http := flag.Bool("http", false, "use http")

  defaultProxy := "proxy.ssh443.com:443"

  flag.Parse()

  proxyAddress := *proxy

  if (proxyAddress == "") {
    proxyAddress = defaultProxy
  }

  if (*address == "" || *auth == "" ) {
    flag.Usage()
    os.Exit(2)
  }

  return args{
    address: *address,
    auth: *auth,
    proxy: proxyAddress,
    http: *http,
  }
}

func getConn(proxy string, ssl bool) (net.Conn) {
  var conn net.Conn
  var err error
  if (ssl) {
    conn, err = tls.Dial("tcp", proxy, &tls.Config{})
  } else {
    conn, err = net.Dial("tcp", proxy)
  }

  if (err != nil) {
    panic(err)
  }

  return conn
}

func sendHeaders(conn net.Conn, auth string, address string, proxy string) {
  fmt.Fprintf(conn, "CONNECT " + address + " HTTP/1.1\r\n")
  fmt.Fprintf(conn, "Proxy-Authorization: Basic " + auth + "\r\n")
  fmt.Fprintf(conn, "Host: " + proxy + "\r\n")
  fmt.Fprintf(conn, "\r\n")
}

func awaitConnectionReady(conn net.Conn) {
  status, err := bufio.NewReader(conn).ReadString('\n')

  if err == io.EOF {
    panic(errors.New("ERROR: connection closed by remote host"))
  }
  if err != nil {
    panic(err)
  }

  fields := strings.Fields(status)
  if (len(fields) < 3 || fields[1] != "200") {
    panic(errors.New("ERROR: proxy connection failure"))
  }
}

func pipeStdIO(conn net.Conn) {
  defer conn.Close()
  go io.Copy(conn, os.Stdin)
  io.Copy(os.Stdout, conn)
}

func main() {
  args := cliArgs()
  auth := base64.StdEncoding.EncodeToString([]byte(args.auth))
  conn := getConn(args.proxy, !args.http)
  sendHeaders(conn, auth, args.address, args.proxy)
  awaitConnectionReady(conn)
  pipeStdIO(conn)
}
