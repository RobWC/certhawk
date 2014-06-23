package main

import (
    "crypto/tls"
    "fmt"
    "encoding/json"
)

func main() {
  config := tls.Config{InsecureSkipVerify: true,ServerName:"reddit.com"}
  newConn, err := tls.Dial("tcp","reddit.com:443",&config)
  fmt.Println(err)
  conState := newConn.ConnectionState()
  fmt.Println(newConn.RemoteAddr())
  fmt.Println(conState.PeerCertificates[0].Subject)
  fmt.Println(conState.PeerCertificates[0].NotBefore)
  fmt.Println(conState.PeerCertificates[0].NotAfter)
  fmt.Println(conState.PeerCertificates[0].SerialNumber)
  jsonCert,_ := json.MarshalIndent(conState.PeerCertificates[0],""," ")
  fmt.Println(string(jsonCert))
}
