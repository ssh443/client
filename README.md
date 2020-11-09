# SSH443 Client
The https proxy client for [https://www.ssh443.com/]()

### Setup
- Install golang [https://golang.org/doc/install]()
- Build the binary: `go build ssh443client.go`
- Move the `ssh443client` into your `$PATH`

### Basic Usage
```bash
# Inline example:
ssh example.com -o 'ProxyCommand=ssh443client -auth TOKEN:SECRET -address %h:%p'
```
```bash
# ssh config example: ~/.ssh/config
Host *
  ProxyCommand ssh443client -auth TOKEN:SECRET -address %h:%p
```


### Custom Usage
```bash
Usage of ssh443client:
  -address string
        <host>:<port>
  -auth string
        <token>:<secret>
  -http
        use http
  -proxy string
        <host>:<port>
```
