package msnet

import (
	"bytes"
	"log/slog"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

var (
	localeEncoder *encoding.Encoder
	localeDecoder *encoding.Decoder
)

func SetLocale(region Region) {
	switch region {
	case KMS, KMST:
		localeEncoder = korean.EUCKR.NewEncoder()
		localeDecoder = korean.EUCKR.NewDecoder()
	case JMS:
		localeEncoder = japanese.ShiftJIS.NewEncoder()
		localeDecoder = japanese.ShiftJIS.NewDecoder()
	case CMS:
		localeEncoder = simplifiedchinese.GBK.NewEncoder()
		localeDecoder = simplifiedchinese.GBK.NewDecoder()
	case TMS:
		localeEncoder = traditionalchinese.Big5.NewEncoder()
		localeDecoder = traditionalchinese.Big5.NewDecoder()
	default:
		localeEncoder = encoding.Nop.NewEncoder()
		localeDecoder = encoding.Nop.NewDecoder()
	}
}

func GetLocaleBuf(s string) []byte {
	if s == "" {
		return nil
	}
	result, _, err := transform.String(localeEncoder, s)
	if err != nil {
		slog.Error(err.Error(), "str", s)
		return []byte(s)
	}
	return []byte(result)
}

func GetLocaleStr(rawBuf []byte) string {
	if len(rawBuf) == 0 {
		return ""
	}

	if before, _, ok := bytes.Cut(rawBuf, []byte{0}); ok {
		rawBuf = before
	}

	if len(rawBuf) == 0 {
		return ""
	}
	result, _, err := transform.Bytes(localeDecoder, rawBuf)
	if err != nil {
		slog.Error(err.Error(), "rawBuf", len(rawBuf))
		return string(rawBuf)
	}
	return string(result)
}
