package node

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

func NewAnnouncePacket(n *NodeInfo) (string, error) {
	jsonMessage, err := json.Marshal(n)
	if err != nil {
		return "", err
	}

	message := fmt.Sprintf("%s%s%s", HEADER, NODE_ANNOUNCE_COMMAND, jsonMessage)

	return message, nil
}

func ParseAnnouncePacket(size int, addr *net.UDPAddr, packet []byte) (*NodeInfo, error) {
	if size <= MIN_PACKET_SIZE {
		return nil, fmt.Errorf("Invalid packet size")
	}
	if strings.Compare(string(packet[0:len(HEADER)]), HEADER) != 0 {
		return nil, fmt.Errorf("Invalid packet header")
	}

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
