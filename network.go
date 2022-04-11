package main

type MessageStatus int8

const (
	Delivered MessageStatus = iota
	NotDelivered
	Commited
	Delivering
)

type Networker interface {
	Listen()
	Send()
	Bootsrap()
}
