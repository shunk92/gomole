package main

import (
	_ "embed"
	"errors"
	"strings"

	"github.com/BurntSushi/toml"
)

//go:embed protocol.toml
var protocolFile []byte

type protocol struct {
	port map[string]interface{}
}

func (p *protocol) new() {
	toml.Unmarshal(protocolFile, &p.port)
}

func (p *protocol) getPortNo(protocolName interface{}) (int, error) {
	switch protocolName.(type) {
	case int64:
		return int(protocolName.(int64)), nil
	case string:
		if port := p.port[strings.ToLower(protocolName.(string))]; port != nil {
			return int(port.(int64)), nil
		}
		return 0, errors.New("protocol doesn't exist in list, you have to assign port number")
	default:
		return 0, errors.New("protocol doesn't assigned, you have to assign protocol name or port number")
	}

}
