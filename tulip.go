package main

import (
    "os"
    "log"
    "runtime"
    "encoding/json"
    cfg "tulip/pkgs/config"
    engine "tulip/pkgs/server"
)

var config = &Configuration{}

type Configuration struct {
    Server engine.Server
}

func init() {
	log.SetFlags(log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
    rootDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
    cfg.Load(rootDir + string(os.PathSeparator) + "config" + string(os.PathSeparator) + "config.json", config)
    engine.Run(config.Server)
}

func (c *Configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}