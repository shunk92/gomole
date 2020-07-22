package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

func main() {
	var cfg config
	if _, err := toml.DecodeFile("config.toml", &cfg); err != nil {
		log.Fatal(err)
	}
	// for _, s := range cfg.Server {
	s := cfg.Server[0]
	auth := PrivateKeyFile(cfg.Bastion.AuthFile)
	tunnel := NewSSHTunnel(cfg.Bastion.Server, auth, s.Distination, s.LocalPort)
	tunnel.Log = log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds)
	tunnel.Start()
	// }

}

func makeTunnel(cfg config) {

}

type bastion struct {
	Server   string
	AuthFile string
}

type server struct {
	LocalPort   int
	Distination string
}

type config struct {
	Bastion bastion
	Server  []server
}
