package net

import (
	"container/list"
	"errors"
	"sync"
	"sync/atomic"
)

var ClientClosedError = errors.New("Client Closed")
var ClientBlockedError = errors.New("Client Blocked")

var globalSessionId uint64

type Client struct {
	id             uint64
	netStream      NetStreamProtocol // net stream send receive msg
	manager        *Manager
	sendChan       chan interface{}
	closeFlag      int32
	closeChan      chan int
	closeMutex     sync.Mutex
	closeCallbacks *list.List
	State          interface{}
}

type closeCallback struct {
	Handler interface{}
	Func    func()
}

func NewClient(netStream NetStreamProtocol, sendChanSize int) *Client {
	return newClient(nil, netStream, sendChanSize)
}
func newClient(manager *Manager, netStream NetStreamProtocol, sendChanSize int) *Client {
	client := &Client{
		netStream: netStream,
		manager:   manager,
		closeChan: make(chan int),
		id:        atomic.AddUint64(&globalSessionId, 1),
	}

	if sendChanSize > 0 {
		client.sendChan = make(chan interface{}, sendChanSize)
		go client.sendLoop()
	}
	return client
}

func (client *Client) NetStream() NetStreamProtocol {
	return client.netStream
}

func (client *Client) Receive() ([]byte, error) {
	msg, err := client.netStream.Receive()
	if err != nil {
		client.Close()
	}
	return msg, err
}

func (client *Client) sendLoop() {
	defer client.Close()
	for {
		select {
		case msg := <-client.sendChan:
			if client.netStream.Send(msg) != nil {
				return
			}
		case <-client.closeChan:
			return
		}
	}
}

func (client *Client) ID() uint64 {
	return client.id
}

func (client *Client) IsClosed() bool {

	return atomic.LoadInt32(&client.closeFlag) == 1

}

func (client *Client) Send(msg interface{}) (err error) {

	if client.IsClosed() {
		return ClientClosedError
	}
	if client.sendChan == nil {
		return client.netStream.Send(msg)
	}
	select {
	case client.sendChan <- msg:
		return nil
	default:
		return ClientBlockedError
	}
}

func (client *Client) Close() error {
	if atomic.CompareAndSwapInt32(&client.closeFlag, 0, 1) {
		err := client.netStream.Close()
		close(client.closeChan)
		if client.manager != nil {
			client.manager.delClient(client)
		}
		client.invokeCloseCallbacks()
		return err
	}
	return ClientClosedError
}

func (client *Client) addCloseCallback(handler interface{}, callback func()) {
	if client.IsClosed() {
		return
	}
	client.closeMutex.Lock()
	defer client.closeMutex.Unlock()

	if client.closeCallbacks == nil {
		client.closeCallbacks = list.New()
	}
	client.closeCallbacks.PushBack(closeCallback{handler, callback})
}

func (client *Client) removeCloseCallback(handler interface{}) {
	if client.IsClosed() {
		return
	}
	client.closeMutex.Lock()
	defer client.closeMutex.Unlock()
	for elm := client.closeCallbacks.Front(); elm != nil; elm = elm.Next() {
		if elm.Value.(closeCallback).Handler == handler {
			client.closeCallbacks.Remove(elm)
			return
		}
	}
}

func (client *Client) invokeCloseCallbacks() {
	client.closeMutex.Lock()
	defer client.closeMutex.Unlock()
	if client.closeCallbacks == nil {
		return
	}
	for elm := client.closeCallbacks.Front(); elm != nil; elm = elm.Next() {
		callback := elm.Value.(closeCallback)
		callback.Func()
	}
}
