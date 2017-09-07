/** * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *
 * net stream
 * generate by DavidYang 2017.9.7
 *
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package net

import (
	"io"
	"net"
	"time"
)

type NetStreamProtocol interface {
	Receive() ([]byte, error)
	Send(interface{}) error
	Close() error
}

type NetStream interface {
	NewNetStream(rw io.ReadWriter) NetStreamProtocol
}

func StartServe(network, address string, netStream NetStream, sendChanSize int) (*Server, error) {
	listener, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}
	return NewServer(listener, netStream, sendChanSize), nil
}

func Connect(network, address string, netStream NetStream, sendChanSize int) (*Client, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return NewClient(netStream.NewNetStream(conn), sendChanSize), nil
}

func ConnectTimeout(network, address string, timeout time.Duration, netStream NetStream, sendChanSize int) (*Client, error) {
	conn, err := net.DialTimeout(network, address, timeout)
	if err != nil {
		return nil, err
	}
	return NewClient(netStream.NewNetStream(conn), sendChanSize), nil
}
