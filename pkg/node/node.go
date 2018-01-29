package node

import (
	"container/list"
	"github.com/donkeysharp/gocho/pkg/config"
)

const (
	MULTICAST_ADDRESS     = "239.6.6.6:1337"
	MULTICAST_BUFFER_SIZE = 4096

	NODE_ANNOUNCE_COMMAND = "\x01"
	HEADER                = "\x60\x0D\xF0\x0D"
	MIN_PACKET_SIZE       = 6

	EXPIRE_TIMEOUT_SEC    = 50
	ANNOUNCE_INTERVAL_SEC = 10
)

type NodeInfo struct {
	Id            string `json:"nodeId"`
	Address       string `json:"ipAddress"`
	WebPort       string `json:"webPort"`
	LastMulticast int64  `json:"-"`
}

type Announcer struct {
	config *config.Config
}

func (a *Announcer) Start(nodeList *list.List) {
	nodeInfo := &NodeInfo{
		Id:            a.config.NodeId,
		Address:       "",
		WebPort:       a.config.WebPort,
		LastMulticast: 0,
	}

	go announceNode(nodeInfo)
	go listenForNodes(nodeList)

}
