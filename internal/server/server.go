package server

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/zhyonc/msnet"
	"github.com/zhyonc/msnet/internal/opcode"
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
		cs.LoopAliveAck(0)
		cs.LoopAliveReq(0, opcode.LP_AliveReq)
		cs.SetID(idCount)
		idCount++
	}
}

func (s *server) Shutdown() {
	if s.lis != nil {
		s.lis.Close()
		s.lis = nil
	}
}

// DebugInPacketLog implements [msnet.CClientSocketDelegate].
func (s *server) DebugInPacketLog(id int32, iPacket msnet.CInPacket) {
	key := iPacket.GetType()
	_, ok := opcode.NotLogCP[key]
	if !ok {
		tag, _ := opcode.CPMap[key]
		slog.Debug("[CInPacket]", "id", id, "opcode", fmt.Sprintf("%d/0x%04X", key, key), "tag", tag, "length", iPacket.GetLength(), "data", iPacket.DumpString(-1))
	}
}

// DebugOutPacketLog implements [msnet.CClientSocketDelegate].
func (s *server) DebugOutPacketLog(id int32, oPacket msnet.COutPacket) {
	key := oPacket.GetType()
	_, ok := opcode.NotLogLP[key]
	if !ok {
		tag, _ := opcode.LPMap[key]
		slog.Debug("[COutPacket]", "id", id, "opcode", fmt.Sprintf("%d/0x%04X", key, key), "tag", tag, "length", oPacket.GetLength(), "data", oPacket.DumpString(-1))
	}
}

// NewConnectPacket implements [msnet.CClientSocketDelegate].
func (s *server) NewConnectPacket(region msnet.Region, version uint16, minorVersion string, seqRcv [4]byte, seqSnd [4]byte) msnet.COutPacket {
	return nil
}

// ProcessPacket implements [msnet.CClientSocketDelegate].
func (s *server) ProcessPacket(cs msnet.CClientSocket, iPacket msnet.CInPacket) {
	op := iPacket.Decode2()
	switch op {
	case opcode.CP_AliveAck:
		cs.OnAliveAck()
	default:
		slog.Warn("Unprocessed CInPacket", "opcode", fmt.Sprintf("%d/0x%04X", op, op))
	}
}

// SocketClose implements [msnet.CClientSocketDelegate].
func (s *server) SocketClose(id int32) {
	slog.Info("Socket closed", "id", id)
}
