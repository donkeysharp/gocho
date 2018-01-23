package node

import (
	"fmt"
	"github.com/donkeysharp/gocho/pkg/config"
	"net/http"
	"time"
)

func fileServe(conf *config.Config) {
	http.Handle("/", http.FileServer(http.Dir(conf.ShareDirectory)))
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", conf.WebPort), nil)
}

func startAnnouncer(conf *config.Config) {
	announcer := &Announcer{
		config: conf,
	}
	announcer.Start()
}

func Serve(conf *config.Config) {
	startAnnouncer(conf)

	go fileServe(conf)

	for {
		time.Sleep(time.Minute * 15)
	}
}
