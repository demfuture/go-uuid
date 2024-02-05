package uuid

import (
	"encoding/base64"
	"strings"
)

var escaper = strings.NewReplacer("9", "99", "-", "90", "_", "91")
var unescaper = strings.NewReplacer("99", "9", "90", "-", "91", "_")

func (u UUID) Base64() string {
	return base64.RawStdEncoding.EncodeToString(u.Bytes())
}

func (u UUID) Base64URL() string {
	return base64.RawURLEncoding.EncodeToString(u.Bytes())
}

func (u UUID) ContinuousBase64() string {
	return escaper.Replace(base64.RawURLEncoding.EncodeToString(u.Bytes()))
}

func FromBase64(s string) (UUID, error) {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return UUID{}, err
	}
	return FromBytes(b)
}

func FromBase64URL(s string) (UUID, error) {
	b, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return UUID{}, err
	}
	return FromBytes(b)
}

func FromContinuousBase64(s string) (UUID, error) {
	b, err := base64.RawURLEncoding.DecodeString(unescaper.Replace(s))
	if err != nil {
		return UUID{}, err
	}
	return FromBytes(b)
}
