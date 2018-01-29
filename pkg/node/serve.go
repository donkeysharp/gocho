package node

import (
	"container/list"
	"fmt"
	"github.com/donkeysharp/gocho/pkg/config"
	"net/http"
	"time"
)

func fileServe(conf *config.Config) {
	fileMux := http.NewServeMux()
	fileMux.Handle("/", http.FileServer(http.Dir(conf.ShareDirectory)))
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", conf.WebPort), fileMux)
}

func startAnnouncer(conf *config.Config, nodeList *list.List) {
	announcer := &Announcer{
		config: conf,
	}
	announcer.Start(nodeList)
}

func Serve(conf *config.Config) {
	nodeList := list.New()

	go startAnnouncer(conf, nodeList)
	go fileServe(conf)
	go dashboardServe(conf, nodeList)

	for {
		time.Sleep(time.Minute * 15)
	}
}
