/** * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *
 * net stream
 * generate by DavidYang 2017.9.13
 *
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package common

import (
	protocol "lotteryim/tcp/net"
	"net"
)

type binaryProtocol struct {
	Conn net.Conn
}

func NewNetStream(conn net.Conn) protocol.NetStreamProtocol {

	return &binaryProtocol{
		Conn: conn,
	}
}

func (binaryPro *binaryProtocol) Receive() ([]byte, error) {
	receiveMsg, err := ReadMsg(binaryPro.Conn)
	return []byte(receiveMsg), err
}

func (binaryPro *binaryProtocol) Send(msg interface{}) error {

	msgStr, ok := msg.(string)
	if ok {
		WriteMsg(binaryPro.Conn, msgStr)
	}
	return nil
}

func (binaryPro *binaryProtocol) Close() error {
	if binaryPro.Conn != nil {
		binaryPro.Conn.Close()
	}
	return nil
}
