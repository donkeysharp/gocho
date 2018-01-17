package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	MULTICAST_ADDRESS     = "239.6.6.6:1337"
	MULTICAST_BUFFER_SIZE = 4096
	NODE_ANNOUNCE_COMMAND = "\x01"
	HEADER                = "\x60\x0D\xF0\x0D"
	MIN_PACKAGE_SIZE      = 6
	EXPIRE_TIMEOUT_SEC    = 60
)

var (
	nodeIdentifier string
	nodeMutex      sync.Mutex
)

type NodeInfo struct {
	Id            string `json:"id"`
	Address       net.IP `json:"-"`
	WebPort       string `json:"port"`
	LastMulticast int64  `json:"-"`
}

func GetAnnouncePacket(n *NodeInfo) (string, error) {
	jsonMessage, err := json.Marshal(n)
	if err != nil {
		return "", err
	}

	message := fmt.Sprintf("%s%s%s", HEADER, NODE_ANNOUNCE_COMMAND, jsonMessage)

	return message, nil
}

// func (n *NodeInfo) String() string {
// 	msg, err := json.Marshal(n)
// 	if err != nil {
// 		return "<err>"
// 	}
// 	return string(msg)
// }

func announceNode(nodeInfo *NodeInfo) {
	address, err := net.ResolveUDPAddr("udp", MULTICAST_ADDRESS)
	if err != nil {
		return
	}

	conn, err := net.DialUDP("udp", nil, address)
	if err != nil {
		return
	}

	for {
		fmt.Println("sending multicast info")

		message, err := GetAnnouncePacket(nodeInfo)
		if err != nil {
			fmt.Println("Could not get announce package")
			fmt.Println(err)
			continue
		}

		conn.Write([]byte(message))
		time.Sleep(10 * time.Second)
	}
}

func listenForNodes(nodeList *list.List) {
	address, err := net.ResolveUDPAddr("udp", MULTICAST_ADDRESS)
	if err != nil {
		return
	}

	conn, err := net.ListenMulticastUDP("udp", nil, address)
	if err != nil {
		return
	}

	conn.SetReadBuffer(MULTICAST_BUFFER_SIZE)

	for {
		packet := make([]byte, MULTICAST_BUFFER_SIZE)
		size, udpAddr, err := conn.ReadFromUDP(packet)
		if err != nil {
			fmt.Println(err)
			continue
		}

		nodeInfo, err := ParseAnnouncePackage(size, udpAddr, packet)
		fmt.Printf("Received multicast packet from %s Id: %s\n", udpAddr.String(), nodeInfo.Id)

		if err != nil {
			fmt.Println(err)
			continue
		}

		go announcedNodeHandler(nodeInfo, nodeList)
	}
}

func announcedNodeHandler(nodeInfo *NodeInfo, nodeList *list.List) {
	addNode(nodeInfo, nodeList)

	fmt.Println("Printing nodes")

	fmt.Print("[")
	for el := nodeList.Front(); el != nil; el = el.Next() {
		fmt.Print(el.Value.(*NodeInfo).Id, " ")
	}
	fmt.Print("]\n\n\n\n")
}

func addNode(nodeInfo *NodeInfo, nodeList *list.List) {
	if isSelfNode(nodeInfo) {
		return
	}

	nodeMutex.Lock()
	nodeExists := false
	for el := nodeList.Front(); el != nil; el = el.Next() {
		tmp := el.Value.(*NodeInfo)

		// Already in list
		if tmp.Id == nodeInfo.Id {
			tmp.LastMulticast = time.Now().Unix()
			nodeExists = true
			continue
		}

		if isNodeExpired(tmp, EXPIRE_TIMEOUT_SEC) {
			fmt.Println("Node expired, removing: ", tmp.Id)
			nodeList.Remove(el)
		}
	}
	if !nodeExists {
		nodeInfo.LastMulticast = time.Now().Unix()
		nodeList.PushBack(nodeInfo)
	}
	nodeMutex.Unlock()
}

func isNodeExpired(nodeInfo *NodeInfo, timeout int) bool {
	diff := time.Now().Unix() - nodeInfo.LastMulticast
	return diff > int64(timeout)
}

func isSelfNode(nodeInfo *NodeInfo) bool {
	return nodeInfo.Id == nodeIdentifier
}

func ParseAnnouncePackage(size int, addr *net.UDPAddr, packet []byte) (*NodeInfo, error) {
	if size <= MIN_PACKAGE_SIZE {
		return nil, fmt.Errorf("foobar")
	}
	// fmt.Println("Packet has the right size")
	if strings.Compare(string(packet[0:len(HEADER)]), HEADER) != 0 {
		return nil, fmt.Errorf("Invalid package header")
	}
	// fmt.Println("Packet header is valid")

	if string(packet[len(HEADER):len(HEADER)+1]) != NODE_ANNOUNCE_COMMAND[0:] {
		return nil, fmt.Errorf("Command different than NODE_ANNOUNCE_COMMAND")
	}
	fmt.Println("Packet command is NODE_ANNOUNCE_COMMAND")

	payload := string(packet[len(HEADER)+1:])
	payload = strings.Trim(payload, "\x00")

	nodeInfo := &NodeInfo{
		Address: addr.IP,
	}

	err := json.Unmarshal([]byte(payload), nodeInfo)
	if err != nil {
		return nil, err
	}

	return nodeInfo, nil
}

func main() {
	fmt.Println("=== Gocho Sharing ===")
	nodeIdentifier = os.Args[1]

	nodeList := list.New()

	nodeInfo := &NodeInfo{
		Id:            nodeIdentifier,
		Address:       nil,
		WebPort:       "1337",
		LastMulticast: 1234,
	}

	go announceNode(nodeInfo)
	go listenForNodes(nodeList)

	for {
		time.Sleep(15 * time.Minute)
	}
}
