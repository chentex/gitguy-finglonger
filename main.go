package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/SUSE/gitguy-finglonger/pkg/config"
	"github.com/SUSE/gitguy-finglonger/server"
)

var (
	configFile string
)

func main() {
	conf, err := initConfig()
	if err != nil {
		return
	}

	s := &http.Server{
		Addr:         conf.Server.Address,
		ReadTimeout:  time.Duration(conf.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(conf.Server.WriteTimeout) * time.Second,
	}

	ser := server.NewServer(s, conf)

	log.Fatal(ser.ListenAndServe())
}

func initConfig() (*config.Config, error) {
	flag.StringVar(&configFile, "config", "config.devel.yml", "Set the path to the file with the configuration for the project.")
	flag.Parse()

	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	conf := new(config.Config)
	err = yaml.Unmarshal(b, conf)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return conf, nil
}
