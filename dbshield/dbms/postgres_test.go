package dbms_test

import (
	"errors"
	"io"
	"testing"

	"github.com/nim4/DBShield/dbshield/dbms"
	"github.com/nim4/mock"
)

var samplePostgres = [...][]byte{
	{0x00, 0x00, 0x00, 0x08, 0x04, 0xd2, 0x16, 0x2f}, //Client
	{0x4e},
	{
		0x00, 0x00, 0x00, 0x25, 0x00, 0x03, 0x00, 0x00, 0x75, 0x73, 0x65, 0x72,
		0x00, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x00, 0x64, 0x61,
		0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x00, 0x74, 0x65, 0x73, 0x74, 0x00,
		0x00,
	},
	{
		0x52, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x05, 0x15, 0x81, 0xd2,
		0xfb,
	},
	{
		0x70, 0x00, 0x00, 0x00, 0x28, 0x6d, 0x64, 0x35, 0x61, 0x65, 0x66, 0x34,
		0x36, 0x64, 0x61, 0x32, 0x30, 0x31, 0x31, 0x37, 0x63, 0x61, 0x35, 0x30,
		0x65, 0x37, 0x34, 0x64, 0x33, 0x64, 0x66, 0x61, 0x65, 0x31, 0x65, 0x62,
		0x33, 0x37, 0x35, 0x32, 0x00,
	},
	{0x52, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00},
	{
		0x51, 0x00, 0x00, 0x00, 0x19, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x20,
		0x2a, 0x20, 0x46, 0x52, 0x4f, 0x4d, 0x20, 0x73, 0x74, 0x6f, 0x63, 0x6b,
		0x73, 0x00,
	},
	{
		0x54, 0x00, 0x00, 0x00, 0x4e, 0x00, 0x03, 0x69, 0x64, 0x00, 0x00, 0x00,
		0x40, 0x18, 0x00, 0x01, 0x00, 0x00, 0x00, 0x17, 0x00, 0x04, 0xff, 0xff,
		0xff, 0xff, 0x00, 0x00, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x00, 0x00,
		0x00, 0x40, 0x18, 0x00, 0x02, 0x00, 0x00, 0x04, 0x13, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x0e, 0x00, 0x00, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
		0x00, 0x00, 0x00, 0x40, 0x18, 0x00, 0x03, 0x00, 0x00, 0x04, 0x13, 0xff,
		0xff, 0x00, 0x00, 0x01, 0x03, 0x00, 0x00, 0x44, 0x00, 0x00, 0x00, 0x19,
		0x00, 0x03, 0x00, 0x00, 0x00, 0x01, 0x31, 0x00, 0x00, 0x00, 0x01, 0x50,
		0x00, 0x00, 0x00, 0x05, 0x50, 0x61, 0x72, 0x74, 0x61, 0x44, 0x00, 0x00,
		0x00, 0x19, 0x00, 0x03, 0x00, 0x00, 0x00, 0x01, 0x32, 0x00, 0x00, 0x00,
		0x01, 0x58, 0x00, 0x00, 0x00, 0x05, 0x58, 0x61, 0x72, 0x74, 0x61, 0x44,
		0x00, 0x00, 0x00, 0x19, 0x00, 0x03, 0x00, 0x00, 0x00, 0x01, 0x33, 0x00,
		0x00, 0x00, 0x01, 0x5a, 0x00, 0x00, 0x00, 0x05, 0x5a, 0x61, 0x72, 0x74,
		0x61, 0x43, 0x00, 0x00, 0x00, 0x0d, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54,
		0x20, 0x33, 0x00, 0x5a, 0x00, 0x00, 0x00, 0x05, 0x49,
	},
	{0x58, 0x00, 0x00, 0x00, 0x04},
}

var postgresCount int

func postgresDummyReader(c io.Reader) (buf []byte, err error) {
	if postgresCount < len(samplePostgres) {
		buf = samplePostgres[postgresCount]
		postgresCount++
	} else {
		err = errors.New("EOF")
	}
	return
}

func TestPostgres(t *testing.T) {
	p := new(dbms.Postgres)
	port := p.DefaultPort()
	if p.DefaultPort() != 5432 {
		t.Error("Expected 5432, got ", port)
	}
	err := p.SetCertificate("", "")
	if err == nil {
		t.Error("Expected error")
	}
	p.SetReader(postgresDummyReader)
	var s mock.ConnMock
	p.SetSockets(s, s)
	err = p.Handler()
	if err != nil {
		t.Error("Got error", err)
	}
	p.Close()
}

func BenchmarkPostgres(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	var s mock.ConnMock
	var m = dbms.Postgres{}
	m.SetReader(postgresDummyReader)
	m.SetSockets(s, s)
	for i := 0; i < b.N; i++ {
		postgresCount = 0
		err := m.Handler()
		if err != nil {
			b.Fatal(err)
		}
		m.Close()
	}
}
