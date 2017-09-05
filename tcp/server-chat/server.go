package main

import (
	"fmt"
	"golangim/lotteryim/tcp/common"
	configFile "golangim/lotteryim/tcp/file"
	freeUtility "golangim/lotteryim/tcp/utility"
	"io"

	"log"
	"net"
	"os"
	"time"
)

const (
	CONN_HEARTBEAT byte = 3 // s send heartbeat req
	CONN_RES       byte = 6 // cs send ack
	CONN_REQ       byte = 5 // cs send data
)

var (
	connections []net.Conn
)

func main() {

	// Listen for incoming connections.
	l, err := net.Listen(configFile.FlAppInfo.CONN_TYPE, configFile.FlAppInfo.CONN_HOST+":"+configFile.FlAppInfo.CONN_PORT)
	if err != nil {
		freeUtility.Flog.Error("Error listening:" + err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	freeUtility.Flog.Info("Listening on " + configFile.FlAppInfo.CONN_HOST + ":" + configFile.FlAppInfo.CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		//conn.SetDeadline(time.Now().Add(5 * time.Second))
		if err != nil {
			freeUtility.Flog.Info("Error accepting: " + err.Error())
			os.Exit(1)
		}
		// Save connection
		connections = append(connections, conn)
		// Handle connections in a new goroutine.
		go handleRequest(conn)

	}
}

var xmldata = []byte(`<?xml version="1.0" encoding="utf-8"?><PACKAGE><BODY><amount>121</amount></BODY></PACKAGE>`)

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	for {
		// setReadTimeout
		err := conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		if err != nil {
			fmt.Println(err)
		}

		if msg, err := common.ReadMsg(conn); err == nil {
			freeUtility.Flog.Info("Message Received: " + msg)
			cmd := []byte(msg)[0]
			if cmd == CONN_RES {
				freeUtility.Flog.Info("recv client data ack")
				continue
			}
			broadcast(conn, string(xmldata))
			continue
		}
		//server check Heat Beat
		if configFile.FlAppInfo.CONN_HeatBeatEnable {
			heatMsg := string(CONN_HEARTBEAT)
			common.WriteMsg(conn, heatMsg)
			freeUtility.Flog.Info("send ht packet")
			conn.SetReadDeadline(time.Now().Add(3 * time.Second))
			if _, herr := common.ReadMsg(conn); herr == nil {
				freeUtility.Flog.Info("resv client : " + conn.RemoteAddr().String() + " ht packet ack")
			} else {
				removeConn(conn)
				conn.Close()
				freeUtility.Flog.Warn("close client : " + conn.RemoteAddr().String())
				return
			}
		} else {
			if err == io.EOF {
				removeConn(conn)
				conn.Close()
				freeUtility.Flog.Warn("close client : " + conn.RemoteAddr().String())
				return
			}
			//log.Println(err)
		}
	}
}

func removeConn(conn net.Conn) {
	var i int
	for i = range connections {
		if connections[i] == conn {
			break
		}
	}
	connections = append(connections[:i], connections[i+1:]...)
}

func broadcast(conn net.Conn, msg string) {
	for i := range connections {
		if connections[i] != conn {
			err := common.WriteMsg(connections[i], msg)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
