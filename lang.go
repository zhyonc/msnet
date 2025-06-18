package msnet

import (
	"bytes"
	"io"
	"log/slog"
	"strings"

	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

var (
	langEncoder *encoding.Encoder
	langDecoder *encoding.Decoder
)

func GetLangBuf(s string) []byte {
	reader := strings.NewReader(s)
	transformer := transform.NewReader(reader, langEncoder)
	buf, err := io.ReadAll(transformer)
	if err != nil {
		slog.Error("Failed to get local buf", "str", s)
		return nil
	}
	return buf
}

func GetLangStr(buf []byte) string {
	reader := bytes.NewReader(buf)
	transformer := transform.NewReader(reader, langDecoder)
	decodedBytes, err := io.ReadAll(transformer)
	if err != nil {
		slog.Error("Failed to get local str", "buf", buf)
		return ""
	}
	return string(decodedBytes)
}
