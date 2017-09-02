package main

import (
	"bufio"
	"fmt"
	"golangim/lotteryim/tcp/common"
	configFile "golangim/lotteryim/tcp/file"
	freeUtility "golangim/lotteryim/tcp/utility"
	"io"
	"log"
	"net"
	"os"
)

const (
	CONN_HEARTBEAT byte = 3 // s send heartbeat req
	CONN_RES       byte = 6 // cs send ack
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr(configFile.FlAppInfo.CONN_TYPE, configFile.FlAppInfo.CONN_HOST+":"+configFile.FlAppInfo.CONN_PORT)
	if err != nil {
		freeUtility.Flog.Error(err.Error())
	}
	// Connect to server through tcp.
	conn, err := net.DialTCP(configFile.FlAppInfo.CONN_TYPE, nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go printOutput(conn)
	writeInput(conn)
}

func writeInput(conn *net.TCPConn) {
	freeUtility.Flog.Info("Enter username: ")
	// Read from stdin.
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	username = username[:len(username)-1]
	if err != nil {
		freeUtility.Flog.Fatal(err.Error())
	}
	freeUtility.Flog.Info("Enter text: ")
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			freeUtility.Flog.Fatal(err.Error())
		}
		err = common.WriteMsg(conn, username+": "+text)
		if err != nil {
			freeUtility.Flog.Error(err.Error())
		}
	}
}

func printOutput(conn *net.TCPConn) {
	for {
		// heat beat and receiv ack  CONN_HEARTBEAT
		msg, err := common.ReadMsg(conn)
		cmd := []byte(msg)[0]
		if cmd == CONN_HEARTBEAT {
			freeUtility.Flog.Info("client recv server ht pack ")
			heatMsg := string(CONN_RES)
			common.WriteMsg(conn, heatMsg)
			freeUtility.Flog.Info("send ht pack ack")
		}
		// Receiving EOF means that the connection has been closed
		if err == io.EOF {
			conn.Close()
			freeUtility.Flog.Warn("Connection Closed. Bye bye.")
			os.Exit(0)
		}
		if err != nil {
			freeUtility.Flog.Fatal(err.Error())
		}
		fmt.Println(msg)
		/*
			newMap, err := mxj.NewMapXml([]byte(msg))
			if err == nil {
				newJson, _ := newMap.Json()
				fmt.Printf(string(newJson)) //marshal
			} else {
				fmt.Println(msg)
			}
		*/

	}
}
