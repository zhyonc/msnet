package msnet

import (
	"bytes"
	"log/slog"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

//nolint:gochecknoglobals // globals used to reduce parameter passing across functions
var (
	localeDecoder *encoding.Decoder
	localeEncoder *encoding.Encoder
)

func SetLocale(region Region) {
	switch region {
	case KMS, KMST: // GMSCW
		localeDecoder = korean.EUCKR.NewDecoder()
		localeEncoder = korean.EUCKR.NewEncoder()
	case JMS:
		localeDecoder = japanese.ShiftJIS.NewDecoder()
		localeEncoder = japanese.ShiftJIS.NewEncoder()
	case CMS, CMST, MSEA:
		localeDecoder = simplifiedchinese.GBK.NewDecoder()
		localeEncoder = simplifiedchinese.GBK.NewEncoder()
	case TMS:
		localeDecoder = traditionalchinese.Big5.NewDecoder()
		localeEncoder = traditionalchinese.Big5.NewEncoder()
	case GMS, BMS:
		localeDecoder = charmap.Windows1252.NewDecoder()
		localeEncoder = charmap.Windows1252.NewEncoder()
	default:
		localeDecoder = unicode.UTF8.NewDecoder()
		localeEncoder = unicode.UTF8.NewEncoder()
	}
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
