package main

import (
    "crypto/tls"
    "log"
    //"encoding/json"
    "lib/randip"
    "net"
    "time"
    "fmt"
    //"lib/ipcheck"
)

func main() {
  ripmgr := randip.NewRandIPv4Mgr(true,1249767200)
  for {
    newIP, err := ripmgr.GetNextIP()
    if err != nil {
      log.Println("IP Addr Exhausted")
      return
    } else {
      go func(){
        log.Println(newIP.String())
        config := tls.Config{InsecureSkipVerify: true,ServerName:"google.com"}
        var err error
        var newConn *tls.Conn
        newConn, err = tls.DialWithDialer(&net.Dialer{Timeout:2*time.Second},"tcp",newIP.String() + ":443",&config)
        if err != nil {
          log.Println(err)
        } else {
          conState := newConn.ConnectionState()
          fmt.Println(newConn.RemoteAddr(),conState.PeerCertificates[0].NotBefore,conState.PeerCertificates[0].NotAfter,conState.PeerCertificates[0].SerialNumber)
          //jsonCert,_ := json.MarshalIndent(conState.PeerCertificates[0],""," ")
          //fmt.Println(string(jsonCert))
          newConn.Close()
        }
      }()
    }
  }
}
