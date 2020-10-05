package main
import "bufio"
import "encoding/base64"
import "fmt"
import "io"
import "net"
import "os"

func main() {
  args := os.Args

  if (len(args) != 4) {
    fmt.Println("ERROR: Usage `./ssh443client <auth> <proxyHost>:<proxyPort> <host>:<port>`")
    os.Exit(1)
  }

  auth := base64.StdEncoding.EncodeToString([]byte(args[1]))
  proxyAddress := args[2]
  address := args[3]

  // TODO: add crypto/tls
  conn, err := net.Dial("tcp", proxyAddress)

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  fmt.Fprintf(conn, "CONNECT " + address + " HTTP/1.1\r\n")
  fmt.Fprintf(conn, "Proxy-Authorization: Basic " + auth + "\r\n")
  fmt.Fprintf(conn, "\r\n")
  status, err := bufio.NewReader(conn).ReadString('\n')

  if err == io.EOF {
    fmt.Println("ERROR: connection closed by remote host")
    os.Exit(0)
  }
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  fmt.Println(status)

  defer conn.Close()
  go io.Copy(conn, os.Stdin)
  io.Copy(os.Stdout, conn)
}
