# msnet
msnet is a pure Golang networking package for MapleStory

## Installation
 `$ go get github.com/zhyonc/msnet`

## Quick Start
```golang
package main

import (
	"log/slog"
	"net"

	"github.com/zhyonc/msnet"
)

type server struct {
	addr string
	lis  net.Listener
}

func NewServer(addr string) *server {
	s := &server{
		addr: addr,
	}
	return s
}

func (s *server) Run() {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		slog.Error("Failed to create tcp listener", "err", err)
		return
	}
	slog.Info("TCPListener is starting on " + s.addr)
	s.lis = lis
	for {
		if s.lis == nil {
			slog.Warn("TCPListener is nil")
			break
		}
		conn, err := s.lis.Accept()
		if err != nil {
			slog.Error("Failed to accept conn", "err", err)
			continue
		}
		slog.Info("New client connected", "addr", conn.RemoteAddr())
		cs := msnet.NewCClientSocket(conn, nil, nil, s)
		go cs.OnRead()
		cs.OnConnect()
	}
}

func (s *server) Shutdown() {
	s.lis.Close()
	s.lis = nil
}

func main() {
	msnet.New(&msnet.Setting{
		MSRegion: 8, // MapleStory Regions include KMS(1)/KMST(2)/JMS(3)/CMS(4)/TMS(6)/MSEA(7)/GMS(8)/BMS(9)
		MSVersion: 95, // MapleStory Client Version
		MSMinorVersion: "1", // MapleStory Client Minor Version
	})
	s := NewServer("127.0.0.1:8484")
	s.Run()
}

```

## Packet
|Header|AESOFB|Note|
|:---:|:---:|:---:|
|4 Bytes|Any Bytes|Except for the first packet|
### Decode
PacketLen = (Header[0]+Header[1]*0x100) ^ (Header[2]+Header[3]*0x100)
### Encode
- sVersion = (^clientVer >> 8 & 0xFF) | ((^clientVer << 8) & 0xFF00)
- a = int(sendIV[3])
- a |= int(sendIV[2])<<8
- a ^= sVersion
- b = ((PacketLen << 8) & 0xFF00) | (PacketLen >> 8)
- c = a ^ b
- Header = [a>>8, a, c>>8, b]
### Format
|Opcode|Data|Note|
|:---:|:---:|:---:|
|2 Bytes|Any Bytes|Except for the connect packet|
### Connect Packet
|Name|PacketLen|Version|MinorVersionLen|MinorVersion|RecvIV|SendIV|Region|Note|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
|Connect|2 Bytes|2 Bytes|2 Bytes|1 Byte|4 Bytes|4 Bytes|1 Bytes|The connect packet|