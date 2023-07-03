package clienthellod

import (
	"bytes"
	"reflect"
	"testing"

	"golang.org/x/exp/slices"
)

func TestParseQUICClientHello(t *testing.T) {
	t.Run("Google Chrome", testParseQUICClientHelloGoogleChrome)
}

func testParseQUICClientHelloGoogleChrome(t *testing.T) {
	var rawQCH []byte = []byte{
		0x01, 0x00, 0x01, 0x22, 0x03, 0x03, 0xe3, 0x1b,
		0x6b, 0x88, 0xce, 0x0e, 0xff, 0x48, 0x08, 0x52,
		0xa6, 0x21, 0x03, 0x90, 0x84, 0x92, 0x5d, 0xf6,
		0x8a, 0xcb, 0xad, 0x66, 0xdb, 0x9f, 0x3c, 0x94,
		0x3f, 0x0e, 0xba, 0xf2, 0x4a, 0x3c, 0x00, 0x00,
		0x06, 0x13, 0x01, 0x13, 0x02, 0x13, 0x03, 0x01,
		0x00, 0x00, 0xf3, 0x44, 0x69, 0x00, 0x05, 0x00,
		0x03, 0x02, 0x68, 0x33, 0x00, 0x39, 0x00, 0x5d,
		0x09, 0x02, 0x40, 0x67, 0x0f, 0x00, 0x01, 0x04,
		0x80, 0x00, 0x75, 0x30, 0x05, 0x04, 0x80, 0x60,
		0x00, 0x00, 0xe2, 0xd0, 0x11, 0x38, 0x87, 0x0c,
		0x6f, 0x9f, 0x01, 0x96, 0x07, 0x04, 0x80, 0x60,
		0x00, 0x00, 0x71, 0x28, 0x04, 0x52, 0x56, 0x43,
		0x4d, 0x03, 0x02, 0x45, 0xc0, 0x20, 0x04, 0x80,
		0x01, 0x00, 0x00, 0x08, 0x02, 0x40, 0x64, 0x80,
		0xff, 0x73, 0xdb, 0x0c, 0x00, 0x00, 0x00, 0x01,
		0xba, 0xca, 0x5a, 0x5a, 0x00, 0x00, 0x00, 0x01,
		0x80, 0x00, 0x47, 0x52, 0x04, 0x00, 0x00, 0x00,
		0x01, 0x06, 0x04, 0x80, 0x60, 0x00, 0x00, 0x04,
		0x04, 0x80, 0xf0, 0x00, 0x00, 0x00, 0x33, 0x00,
		0x26, 0x00, 0x24, 0x00, 0x1d, 0x00, 0x20, 0xf8,
		0x82, 0xf6, 0x48, 0x2b, 0x20, 0x0c, 0xa0, 0x60,
		0x79, 0x1c, 0x45, 0xa5, 0xb8, 0x43, 0x58, 0x11,
		0x26, 0x64, 0xec, 0x4f, 0xf7, 0xd6, 0xea, 0x10,
		0x30, 0xf6, 0x9f, 0x36, 0x80, 0x49, 0x43, 0x00,
		0x0d, 0x00, 0x14, 0x00, 0x12, 0x04, 0x03, 0x08,
		0x04, 0x04, 0x01, 0x05, 0x03, 0x08, 0x05, 0x05,
		0x01, 0x08, 0x06, 0x06, 0x01, 0x02, 0x01, 0x00,
		0x10, 0x00, 0x05, 0x00, 0x03, 0x02, 0x68, 0x33,
		0x00, 0x0a, 0x00, 0x08, 0x00, 0x06, 0x00, 0x1d,
		0x00, 0x17, 0x00, 0x18, 0x00, 0x00, 0x00, 0x1a,
		0x00, 0x18, 0x00, 0x00, 0x15, 0x71, 0x2e, 0x63,
		0x6c, 0x69, 0x65, 0x6e, 0x74, 0x68, 0x65, 0x6c,
		0x6c, 0x6f, 0x2e, 0x67, 0x61, 0x75, 0x6b, 0x2e,
		0x61, 0x73, 0x00, 0x2b, 0x00, 0x03, 0x02, 0x03,
		0x04, 0x00, 0x2d, 0x00, 0x02, 0x01, 0x01, 0x00,
		0x1b, 0x00, 0x03, 0x02, 0x00, 0x02,
	}

	qch, err := ParseQUICClientHello(rawQCH)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(qch.Raw(), rawQCH) {
		t.Fatal("marshaled QUIC Client Hello does not match original")
	}

	if qch.TLSHandshakeVersion != 0x0303 {
		t.Fatalf("TLS handshake version does not match, expecting 0x0303, got %.4x", qch.TLSHandshakeVersion)
	}

	if !slices.Equal(
		qch.CipherSuites,
		[]uint16{0x1301, 0x1302, 0x1303}) {
		t.Fatalf("cipher suites do not match, expecting [0x1301, 0x1302, 0x1303], got %v", qch.CipherSuites)
	}

	if qch.ServerName != "q.clienthello.gauk.as" {
		t.Fatal("server name does not match")
	}

	if !slices.Equal(
		qch.NamedGroupList,
		[]uint16{0x001d, 0x0017, 0x0018}) {
		t.Fatalf("named group list does not match, expecting [0x001d, 0x0017, 0x0018], got %v", qch.NamedGroupList)
	}

	if len(qch.ECPointFormatList) != 0 {
		t.Fatalf("EC point format list does not match, expecting [], got %v", qch.ECPointFormatList)
	}

	if !slices.Equal(
		qch.SignatureSchemeList,
		[]uint16{
			0x0403,
			0x0804,
			0x0401,
			0x0503,
			0x0805,
			0x0501,
			0x0806,
			0x0601,
			0x0201,
		}) {
		t.Fatalf("signature scheme list does not match, expecting [0x0403, 0x0804, 0x0401, 0x0503, 0x0805, 0x0501, 0x0806, 0x0601, 0x0201], got %v", qch.SignatureSchemeList)
	}

	if !slices.Equal(
		qch.ALPN,
		[]string{"h3"}) {
		t.Fatalf("ALPN does not match, expecting [\"h3\"], got %v", qch.ALPN)
	}

	if !slices.Equal(
		qch.CertCompressAlgo,
		[]uint16{0x0002}) {
		t.Fatalf("cert compression algorithm does not match, expecting [0x0002], got %v", qch.CertCompressAlgo)
	}

	if len(qch.RecordSizeLimit) != 0 {
		t.Fatalf("record size limit does not match, expecting [], got %v", qch.RecordSizeLimit)
	}

	if !slices.Equal(
		qch.SupportedVersions,
		[]uint16{
			0x0304,
		}) {
		t.Fatalf("supported versions does not match, expecting [0x0304], got %v", qch.SupportedVersions)
	}

	if !slices.Equal(
		qch.PSKKeyExchangeModes,
		[]uint8{
			0x01,
		}) {
		t.Fatalf("PSK key exchange modes does not match, expecting [0x01], got %v", qch.PSKKeyExchangeModes)
	}

	if !slices.Equal(
		qch.KeyShare,
		[]uint16{0x001d}) {
		t.Fatalf("key share does not match, expecting [0x001d], got %v", qch.KeyShare)
	}

	if !slices.Equal(
		qch.ApplicationSettings,
		[]string{"h3"}) {
		t.Fatalf("application settings does not match, expecting [\"h3\"], got %v", qch.ApplicationSettings)
	}

	if !reflect.DeepEqual(qch.qtp, expectedQTPGoogleChrome) {
		t.Errorf("ParseQUICTransportParameters failed: expected %v, got %v", expectedQTPGoogleChrome, qch.qtp)
	}

}
