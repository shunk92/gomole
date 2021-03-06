package main

import (
	"testing"
)

func TestNewEndpoint(t *testing.T) {
	var endpointStr string
	endpointStr = "user@10.0.0.1:22"
	endpoint := NewEndpoint(endpointStr)

	if endpoint.User != "user" {
		t.Errorf("Endpoint.User = %s; expected user", endpoint.User)
	}

	if endpoint.Host != "10.0.0.1" {
		t.Errorf("Endpoint.Host = %s; expected 10.0.0.1", endpoint.Host)
	}

	if endpoint.Port != 22 {
		t.Errorf("Endpoint.Host = %d; expected 22", endpoint.Port)
	}

}
