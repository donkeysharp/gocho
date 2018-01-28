package node

import (
	"fmt"
	"github.com/donkeysharp/gocho/pkg/config"
	"io"
	"net/http"
	"time"
)

func fileServe(conf *config.Config) {
	fileMux := http.NewServeMux()
	fileMux.Handle("/", http.FileServer(http.Dir(conf.ShareDirectory)))
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", conf.WebPort), fileMux)
}

func startAnnouncer(conf *config.Config) {
	announcer := &Announcer{
		config: conf,
	}
	announcer.Start()
}

func dashboardServe(conf *config.Config) {
	dashboardMux := http.NewServeMux()

	dashboardMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "foobarken")
	})
	http.ListenAndServe(fmt.Sprintf("localhost:%s", conf.LocalPort), dashboardMux)
}

func Serve(conf *config.Config) {
	startAnnouncer(conf)

	go fileServe(conf)
	go dashboardServe(conf)

	for {
		time.Sleep(time.Minute * 15)
	}
}
