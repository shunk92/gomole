package main

import (
	"errors"
	"io/ioutil"
	"strings"

	"github.com/BurntSushi/toml"
)

type protocol struct {
	port map[string]interface{}
}

func (p *protocol) new() {
	bytes, err := ioutil.ReadFile("./protocol.toml")
	if err != nil {
		panic(err)
	}
	toml.Unmarshal(bytes, &p.port)
}

func (p *protocol) getPortNo(protocolName interface{}) (int, error) {
	switch protocolName.(type) {
	case int64:
		return int(protocolName.(int64)), nil
	case string:
		if port := p.port[strings.ToLower(protocolName.(string))]; port != nil {
			return int(port.(int64)), nil
		} else {
			return 0, errors.New("Protocol doesn't exist in list, you have to assign port number.")
		}
	default:
		return 0, errors.New("Protocol doesn't assigned, you have to assign protocol name or port number.")
	}

}
