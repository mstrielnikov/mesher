package main

import (
	"fmt"

	"lukechampine.com/blake3"
)

const (
	MaxMessageSize = 1024
)

var (
	ErrMsgSizeHuge          error = fmt.Errorf("Message content size is greater then allowed")
	ErrMsgSizeZero          error = fmt.Errorf("Message content is zero length or empty")
	ErrMsgSenderNameShort   error = fmt.Errorf("Message sender name is too short or empty")
	ErrMsgRecieverNameShort error = fmt.Errorf("Message receiver name is too short or empty")
)

type Message struct {
	Content []byte `json:"content"`
	From    []byte `json:"from"`
	To      []byte `json:"to"`
}

func NewMessage(content, from, to []byte) (*Message, error) {
	if len(content) > MaxMessageSize {
		fmt.Println(ErrMsgSizeHuge.Error())
		return nil, ErrMsgSizeHuge
	}
	if len(content) == 0 {
		fmt.Println(ErrMsgSizeZero.Error())
		return nil, ErrMsgSizeZero
	}
	if len(from) < 3 {
		fmt.Println(ErrMsgSenderNameShort.Error())
		return nil, ErrMsgSenderNameShort
	}
	if len(to) < 3 {
		fmt.Println(ErrMsgRecieverNameShort.Error())
		return nil, ErrMsgRecieverNameShort
	}
	return &Message{
		Content: content,
		From:    from,
		To:      to,
	}, nil
}

type Chat struct {
	Messages []Message
	Name     []byte
	Hash     [32]byte
}

func NewChat(chatName []byte, message Message) Chat {
	var messages []Message
	messages = append(messages, message)
	return Chat{
		Messages: messages,
		Name:     chatName,
		Hash:     blake3.Sum256(message.Content),
	}
}
