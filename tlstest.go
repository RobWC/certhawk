package main

import (
    "crypto/tls"
    "fmt"
    "encoding/json"
)

func main() {
  config := tls.Config{InsecureSkipVerify: true,ServerName:"74.125.239.128"}
  newConn, err := tls.Dial("tcp","74.125.239.128:443",&config)
  if err != nil {
    fmt.Println(err)
  } else {
    conState := newConn.ConnectionState()
    fmt.Println(newConn.RemoteAddr())
    fmt.Println(conState.PeerCertificates[0].Subject)
    fmt.Println(conState.PeerCertificates[0].NotBefore)
    fmt.Println(conState.PeerCertificates[0].NotAfter)
    fmt.Println(conState.PeerCertificates[0].SerialNumber)
    jsonCert,_ := json.MarshalIndent(conState.PeerCertificates[0],""," ")
    fmt.Println(string(jsonCert))
  }
}
