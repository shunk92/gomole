package main

import (
	"errors"
	"strings"
)

// go:embed protocol.toml
// var protocolsFile []byte

type protocol struct {
	port map[string]interface{}
}

func (p *protocol) new() {
	p.port["ssh"] = 22
	p.port["http"] = 80
	p.port["https"] = 443
	p.port["rdp"] = 3389
	p.port["kibana"] = 5601
	p.port["ne4j"] = 7474
	p.port["bolt"] = 7687
	p.port["elasticsearch"] = 9200
	p.port["scaler"] = 30600
	// toml.Unmarshal(protocolsFile, &p.port)
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
