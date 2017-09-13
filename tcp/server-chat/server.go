package main

import (
	"fmt"
	"io"
	"log"
	"lotteryim/tcp/common"
	configFile "lotteryim/tcp/file"
	jsonMsg "lotteryim/tcp/net"
	freeUtility "lotteryim/tcp/utility"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	CONN_HEARTBEAT byte = 3 // s send heartbeat req
	CONN_RES       byte = 6 // cs send ack
	CONN_REQ       byte = 5 // cs send data
	MAXAttempt          = 5
)

var (
	connections       []net.Conn
	connectionsTimout map[string]int
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(configFile.FlAppInfo.CONN_TYPE, configFile.FlAppInfo.CONN_HOST+":"+configFile.FlAppInfo.CONN_PORT)
	if err != nil {
		freeUtility.Flog.Error("Error listening:" + err.Error())
		os.Exit(1)
	}
	defer l.Close()
	freeUtility.Flog.Info("Listening on " + configFile.FlAppInfo.CONN_HOST + ":" + configFile.FlAppInfo.CONN_PORT)

	go handeServer(l)

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)
}

func handeServer(l net.Listener) {

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

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	for {
		// setReadTimeout
		errRead := conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		if errRead != nil {
			fmt.Println(errRead)
		}
		msg, err := common.ReadMsg(conn)
		if err == nil {
			freeUtility.Flog.Info("Message Received: " + msg)
			cmd := []byte(msg)[0]
			if cmd == CONN_RES {
				freeUtility.Flog.Info("recv client data ack")
				continue
			}
			broadcast(conn, msg)
			//testJsonMsg(conn, msg)
			continue
		} else {
			//server check Heat Beat
			if configFile.FlAppInfo.CONN_HeatBeatEnable {
				heatMsg := string(CONN_HEARTBEAT)
				common.WriteMsg(conn, heatMsg)
				freeUtility.Flog.Info("send ht packet")
				conn.SetReadDeadline(time.Now().Add(3 * time.Second))
				if _, herr := common.ReadMsg(conn); herr == nil {
					freeUtility.Flog.Info("resv client : " + conn.RemoteAddr().String() + " ht packet ack")
				} else {
					cleseClient(conn)
					return
				}
			} else {
				if err == io.EOF {
					cleseClient(conn)
					return
				} else if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
					//time out attempt try
					if connectionsTimout[conn.RemoteAddr().String()] >= MAXAttempt {
						cleseClient(conn)
						return
					}
					if connectionsTimout == nil {
						connectionsTimout = make(map[string]int)
					}
					connectionsTimout[conn.RemoteAddr().String()] = connectionsTimout[conn.RemoteAddr().String()] + 1
				} else if err != nil {
					log.Println("run to the end " + err.Error())
				}
			}
		}
	}
}

func cleseClient(conn net.Conn) {
	removeConn(conn)
	delete(connectionsTimout, conn.RemoteAddr().String())
	conn.Close()
	freeUtility.Flog.Warn("close client : " + conn.RemoteAddr().String())
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

func testJsonMsg(conn net.Conn, msg string) {
	msgPack := jsonMsg.NewNetMsgPack()
	jsonByte := msgPack.BuildMsgPack(msg, jsonMsg.NetGroupMsgCMD)
	broadcast(conn, jsonByte)
}
