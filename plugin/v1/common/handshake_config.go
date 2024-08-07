package common

import "github.com/hashicorp/go-plugin"

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "LIMEPIPES_PLUGIN",
	MagicCookieValue: "5ea2ae60-eb31-42a8-9f35-cc4f87dbfc85",
}
