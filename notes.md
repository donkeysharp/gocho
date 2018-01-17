Notes about spider-share
========================

Send a multicast package every 30 seconds
    header
    node_id
    node_ip_address
    web_share_node


Wait for multicast package with information about nodes
    parse package

Information about a node

struct Node {
    nodeId
    nodeIpAddress
    webSharePort
    lastMulticast
}
