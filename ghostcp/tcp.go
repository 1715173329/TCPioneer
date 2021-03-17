package ghostcp

type ConnInfo struct {
	Option uint32
	SeqNum uint32
	TTL    byte
	MAXTTL byte
}

var ConnInfo4 [65536]*ConnInfo
var ConnInfo6 [65536]*ConnInfo
var CookiesMap map[string][]byte
var SynOption []byte

const (
	TCP_FIN = byte(0x01)
	TCP_SYN = byte(0x02)
	TCP_RST = byte(0x04)
	TCP_PSH = byte(0x08)
	TCP_ACK = byte(0x10)
	TCP_URG = byte(0x20)
	TCP_ECE = byte(0x40)
	TCP_CWR = byte(0x80)
)

const domainBytes = "abcdefghijklmnopqrstuvwxyz0123456789-"
