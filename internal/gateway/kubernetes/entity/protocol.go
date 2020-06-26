package entity

type Protocol string

const (
	ProtocolTCP  Protocol = "TCP"
	ProtocolUDP  Protocol = "UDP"
	ProtocolSCTP Protocol = "SCTP"
)

func (p Protocol) String() string {
	return string(p)
}
