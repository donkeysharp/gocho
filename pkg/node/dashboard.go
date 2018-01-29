package node

import (
	"container/list"
	"encoding/json"
	"fmt"
	"github.com/donkeysharp/gocho/pkg/config"
	"net/http"
)

func configHandler(conf *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(conf)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
	}
}

func nodesHandler(nodeList *list.List) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		nodes := make([]*NodeInfo, 0)
		for el := nodeList.Front(); el != nil; el = el.Next() {
			tmp := el.Value.(*NodeInfo)
			nodes = append(nodes, tmp)
		}

		data, err := json.Marshal(nodes)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
	}
}

func dashboardServe(conf *config.Config, nodeList *list.List) {
	dashboardMux := http.NewServeMux()
	dashboardMux.HandleFunc("/api/config", configHandler(conf))
	dashboardMux.HandleFunc("/api/nodes", nodesHandler(nodeList))

	// We don't want the dashboard to be public
	address := "localhost"
	if conf.Debug {
		address = "0.0.0.0"
	}
	fmt.Printf("Starting dashboard at %s\n", address)
	http.ListenAndServe(fmt.Sprintf("%s:%s", address, conf.LocalPort), dashboardMux)
}
