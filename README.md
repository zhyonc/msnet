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
	"sync/atomic"

	"github.com/zhyonc/msnet"
)

type server struct {
	addr string
	lis  net.Listener
	idCount atomic.Int32
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
	var idCount int32 = 0
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
		cs := msnet.NewCClientSocket(s, conn, nil, nil)
		go cs.OnRead()
		cs.OnConnect()
		cs.SetID(s.idCount.Add(1))
	}
}

func (s *server) Shutdown() {
	if s.lis != nil {
		s.lis.Close()
		s.lis = nil
	}
}

func main() {
	msnet.New(&msnet.Setting{
		MSRegion:       msnet.GMS,
		MSVersion:      95,
		MSMinorVersion: "1",
	})
	s := NewServer("127.0.0.1:8484")
	s.Run()
}

```
## Setting
- LocaleRegion: Language Regions including `EUCKR(KMS)/ShiftJIS(JMS)/GBK(CMS)/Big5(TMS)`
- MSRegion: MapleStory Regions including `GMSCW(1)/KMS(1)`/`KMST(2)`/`JMS(3)`/`CMS(4)`/`TMS(6)`/`MSEA(7)`/`GMS(8)`/`BMS(9)`
- MSVersion: MapleStory Client Version
- MSMinorVersion: MapleStory Client Minor Version
- RecvCipherType: Defines how `CInPacket` are decrypted
	- AESCipher: Used for the majority of clients (default) 
	- XORCipher: Used in versions about 2004
- SendCipherType: Defines how `COutPacket` are encrypted
	- AESCipher: Used for the majority of clients (default) 
	- XORCipher: Used in versions about 2004
	- LinearCipher: Used for server packet data since 2017 (excluding the login server)
	- NullCipher: Used for connect packet
- DESKey (optional): A 16-byte string used for opcode encryption based on [v193-encryption](https://forum.ragezone.com/threads/v193-encryption.1147967/)
- AESKeyType (optional)
	- DefaultKey: Uses the fixed `AESKeyDefault` for old MSVersion
	- CycleKey: Picks a key from the `CycleAESKeys` using `MSVersion%20`
	- CycleKey13: Same as CycleKey, but adds an offset of +13 before modulo 20
- AESKey (optional): A 32-byte array used for packet data
	- If provided directly, this key is used as‑is
	- If not set (first byte is zero), the key is derived according to the `AESKeyType`

- IsTypeHeader1Byte: Used in versions about 2004~2008
- AESInitType (optional): Compatible with older versions
	- Default: Used in versions after about 2008
	- Duplicate: Used in versions about 2005~2007 (`excluding TMS`)
	- Shuffle: Used in TMS versions about 2005~2007
- RecvBuffXOR (optional): The server must use the same XOR key to recover the original packet buffer
- SendBuffXOR (optional): The client must use the same XOR key to recover the original packet buffer

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