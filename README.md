# msnet
msnet is a pure Golang networking package for MapleStory

## Installation
 `$ go get github.com/zhyonc/msnet@latest`

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
		MSRegion: 8,
		MSVersion: 95,
		MSMinorVersion: "1",
	})
	s := NewServer("127.0.0.1:8484")
	s.Run()
}

```
## Setting
- MSRegion: MapleStory Regions including `KMS(1)`/`KMST(2)`/`JMS(3)`/`CMS(4)`/`TMS(6)`/`MSEA(7)`/`GMS(8)`/`BMS(9)`
- MSVersion: MapleStory Client Version
- MSMinorVersion: MapleStory Client Minor Version
- RecvXOR: The server must use the same XOR key to recover the original packet
- SendXOR: The client must use the same XOR key to recover the original packet
- IsCycleAESKey: 
	- Default is false, `old AES key` will be used, which is compatible with most earlier versions
	- If set true, `cycle AES key` will be used, which is compatible with newer versions
- CustomAESKey(optional): It's used to instead of `old AES key` and `cycle AES key`
	- Decrypt: A 32-byte array used for decrypting data in CInPacket::DecryptData
	- Encrypt: A 32-byte array used for encrypting data in COutPacket::MakeBufferList
## Packet
|Header|AESOFB|Note|
|:---:|:---:|:---:|
|4 Bytes|Any Bytes|Except for the first packet|
### Decode
Decode packet length using XOR of two little-endian uint16 values from header  
`PacketLen = (Header[0]+Header[1]*0x100) ^ (Header[2]+Header[3]*0x100)`
### Encode
Encode packet length using little-endian XOR with sVersion and sendIV
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