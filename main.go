package main

import (
	config "./internal/config"
	"flag"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var confPtr *string
var cfg config.ServerConfig
func init() {
	confPtr = flag.String("conf", "server-config.json", "Path to the config file")
	cfg = config.ReadConfig(*confPtr)

	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func oauth2Callback(resp http.ResponseWriter, req * http.Request) {
	log.Debugln(req)
	resp.Write([]byte("Great!"))
}

func main() {
	http.HandleFunc("/oauth-callback", oauth2Callback)
	err := http.ListenAndServe(cfg.Listen, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
