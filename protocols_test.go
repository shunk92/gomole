package main

import (
	"testing"
)

func TestProtocolNew(t *testing.T) {
	var p protocol
	p.new()

	if p.port["ssh"].(int64) != 22 {
		t.Errorf("ssh = %v; expected 22", p.port["ssh"])
	}
}

func TestProtocolGetPortNo(t *testing.T) {
	var p protocol
	p.new()

	port, err := p.getPortNo("SSH")
	if port != 22 {
		t.Error("cannot parse ssh protocl")
	}

	port, err = p.getPortNo("XXX")
	if err == nil {
		t.Error("something error occured", port, err)
	}
}
