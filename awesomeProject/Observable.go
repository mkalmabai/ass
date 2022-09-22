package main

type Observable interface {
	Subscribe(observer Observer)
	Unsubscribe(observer Observer)
	sendAll()
}
