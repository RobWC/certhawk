package main

import (
  "encoding/binary"
  "net"
  "log"
  "math"
)

type RandIPv4Mgr struct {
  Seq bool
  ip uint32
}

func NewRandIPv4Mgr (seq bool,seed uint32) *RandIPv4Mgr {
  return &RandIPv4Mgr{Seq:seq,ip:seed}
}

func (r *RandIPv4Mgr) GetNextIP() net.IP {
  if r.ip == uint32(math.Pow(2,32) - 1) {
    //max IP reached
    ipBytes := make([]byte,4,4)
    binary.BigEndian.PutUint32(ipBytes,r.ip)
    return net.IP(ipBytes)
  } else {
    //increment and return next
    r.ip = r.ip + 1
    ipBytes := make([]byte,4,4)
    binary.BigEndian.PutUint32(ipBytes,r.ip)
    return net.IP(ipBytes)
  }
}

func (r *RandIPv4Mgr) GetNextIPStr() string {
  nextIP := r.GetNextIP()
  return nextIP.String()
  //get current ip
  //increment
  //return
}

func main () {
  ipmgr := NewRandIPv4Mgr(true,167772417)
  for i := 0; i < 1200 ; i ++ {
    //log.Println(ipmgr.GetNextIPStr())
    ipmgr.GetNextIPStr()
  }
  log.Println(ipmgr.GetNextIPStr())
  log.Println("Done!")
}
