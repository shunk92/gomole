package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/BurntSushi/toml"
)

func main() {
	configFile := flag.String("config", "config.toml", "config file")
	flag.Parse()

	var cfg config
	if _, err := toml.DecodeFile(*configFile, &cfg); err != nil {
		log.Fatal(err)
	}

	var p protocol
	p.new()
	bst := cfg.Bastion
	wg := new(sync.WaitGroup)
	for _, s := range cfg.Server {
		wg.Add(1)
		go makeTunnel(wg, s, bst, p)
	}
	wg.Wait()
}

func makeTunnel(wg *sync.WaitGroup, srv server, bst bastion, p protocol) {
	port, err := p.getPortNo(srv.DistinationPort)
	if err != nil {
		log.Fatal(err)
		return
	}
	distination := fmt.Sprintf("%v:%v", srv.Distination, port)

	auth := PrivateKeyFile(bst.AuthFile)
	tunnel := NewSSHTunnel(bst.Server, auth, distination, srv.LocalPort)
	tunnel.Log = log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds)
	tunnel.logf("Start connection [%s] -> port %d", srv.getDistinationString(), srv.LocalPort)
	tunnel.Start()
}

type bastion struct {
	Server   string
	AuthFile string
}

type server struct {
	ServerName      string
	LocalPort       int
	DistinationPort interface{}
	Distination     string
}

type config struct {
	Bastion bastion
	Server  []server
}

func (srv *server) getDistinationString() string {
	if srv.ServerName != "" {
		return srv.ServerName
	}
	return srv.Distination
}
