package gumble

import (
	"crypto/tls"
	"io"
	"net"

	"github.com/bontibon/gumble/gumble/MumbleProto"
)

type Config struct {
	// User name used when authenticating with the server.
	Username string
	// Password used when authenticating with the server. A password is not
	// usually required to connect to a server.
	Password string
	// Server address, including port (e.g. localhost:64738).
	Address string
	Tokens  AccessTokens

	Listener      EventListener
	AudioListener AudioListener

	TlsConfig tls.Config
	Dialer    net.Dialer
}

type AccessTokens []string

func (at AccessTokens) writeTo(client *Client, w io.Writer) (int64, error) {
	packet := MumbleProto.Authenticate{
		Tokens: at,
	}
	proto := protoMessage{&packet}
	return proto.writeTo(client, w)
}
