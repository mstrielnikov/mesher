package msg

const (
	MsgMaxLen = 65535
)

type Message struct {
	text    string
	address string
}
