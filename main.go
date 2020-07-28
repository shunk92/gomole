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

	wg := new(sync.WaitGroup)
	for _, s := range cfg.Server {
		wg.Add(1)
		go makeTunnel(wg, s, cfg, p)
	}
	wg.Wait()

}

func makeTunnel(wg *sync.WaitGroup, srv server, cfg config, p protocol) {
	port, err := p.getPortNo(srv.DistinationPort)
	if err != nil {
		log.Println(err)
		return
	}
	distination := fmt.Sprintf("%v:%v", srv.Distination, port)

	auth := PrivateKeyFile(cfg.Bastion.AuthFile)
	tunnel := NewSSHTunnel(cfg.Bastion.Server, auth, distination, srv.LocalPort)
	tunnel.Log = log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds)
	tunnel.Start()
}

type bastion struct {
	Server   string
	AuthFile string
}

type server struct {
	LocalPort       int
	DistinationPort interface{}
	Distination     string
}

type config struct {
	Bastion bastion
	Server  []server
}
