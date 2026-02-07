package web

import _ "embed"

//go:embed dist.zip
var staticFile []byte

func Static() []byte {
	return staticFile
}
