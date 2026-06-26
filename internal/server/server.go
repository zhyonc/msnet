package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"sync/atomic"

	"github.com/zhyonc/msnet"
	"github.com/zhyonc/msnet/internal/opcode"
)

type Server struct {
	addr    string
	lis     net.Listener
	idCount atomic.Int32
}

func NewServer(addr string) *Server {
	s := &Server{
		addr: addr,
	}
	return s
}

func (s *Server) Run() {
	var lc net.ListenConfig
	lis, err := lc.Listen(context.Background(), "tcp", s.addr)
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
		var conn net.Conn
		conn, err = s.lis.Accept()
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
		cs.SetID(s.idCount.Add(1))
	}
}

func (s *Server) Shutdown() {
	if s.lis != nil {
		err := s.lis.Close()
		if err != nil {
			slog.Error(err.Error())
		}
		s.lis = nil
	}
}

// DebugInPacketLog implements [msnet.CClientSocketDelegate].
func (s *Server) DebugInPacketLog(id int32, iPacket msnet.CInPacket) {
	key := iPacket.GetType()
	_, ok := opcode.NotLogCP[key]
	if !ok {
		tag := opcode.CPMap[key]
		slog.Debug("[CInPacket]",
			"id", id,
			"opcode", fmt.Sprintf("%d/0x%04X", key, key),
			"tag", tag,
			"length", iPacket.GetLength(),
			"data", iPacket.DumpString(-1),
		)
	}
}

// DebugOutPacketLog implements [msnet.CClientSocketDelegate].
func (s *Server) DebugOutPacketLog(id int32, oPacket msnet.COutPacket) {
	key := oPacket.GetType()
	_, ok := opcode.NotLogLP[key]
	if !ok {
		tag := opcode.LPMap[key]
		slog.Debug("[COutPacket]",
			"id", id,
			"opcode", fmt.Sprintf("%d/0x%04X", key, key),
			"tag", tag,
			"length", oPacket.GetLength(),
			"data", oPacket.DumpString(-1),
		)
	}
}

// NewConnectPacket implements [msnet.CClientSocketDelegate].
func (s *Server) NewConnectPacket(
	region msnet.Region,
	version uint16,
	minorVersion string,
	seqRcv []byte,
	seqSnd []byte,
) msnet.COutPacket {
	if region == msnet.GMSCW {
		// Login Server Connect Packet
		oPacket := msnet.NewCOutPacket()
		oPacket.EncodeStr(minorVersion)       // sMinorVersion
		oPacket.EncodeBuffer(seqRcv)          // client uSeqSnd
		oPacket.EncodeBuffer(seqSnd)          // client uSeqRcv
		oPacket.Encode1(int8(region))         // nRegion
		oPacket.Encode1(0)                    // unk
		oPacket.Encode2(int16(version))       // nVersion as temp seq
		oPacket.Encode4(int32(version))       // nVersion
		oPacket.EncodeStr(minorVersion)       // sMinorVersion
		oPacket.EncodeBuffer(seqRcv)          // client uSeqSnd
		oPacket.EncodeBuffer(seqSnd)          // client uSeqRcv
		oPacket.Encode1(int8(region))         // nRegion
		oPacket.Encode4(int32(version * 100)) // nClientVersion_Min
		oPacket.Encode4(int32(version * 100)) // nClientVersion_Max
		oPacket.Encode4(0)                    // nClientVersion_Temp
		oPacket.Encode1(1)                    // nLoginOpt
		oPacket.Encode1(0)                    // unk
		oPacket.Encode1(5)                    // nWvsType
		return oPacket
	}
	return nil
}

// ProcessPacket implements [msnet.CClientSocketDelegate].
func (s *Server) ProcessPacket(cs msnet.CClientSocket, iPacket msnet.CInPacket) {
	op := iPacket.Decode2()
	switch op {
	case opcode.CP_AliveAck:
		cs.OnAliveAck()
	default:
		slog.Warn("Unprocessed CInPacket", "opcode", fmt.Sprintf("%d/0x%04X", op, op))
	}
}

// SocketClose implements [msnet.CClientSocketDelegate].
func (s *Server) SocketClose(id int32) {
	slog.Info("Socket closed", "id", id)
}
