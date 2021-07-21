package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"room/commom/message"
)

type Transfer struct {
	Conn net.Conn
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	buf := make([]byte, 8096)
	_, err = this.Conn.Read(buf[:4])
	if err != nil {
		fmt.Printf("conn.Read  err = %v ", err)
		return
	}
	fmt.Printf("读buf = %v\n", buf[:4])

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])

	b, err := this.Conn.Read(buf[0:pkgLen])
	if b != int(pkgLen) || err != nil {
		err = errors.New("Read pkg err")
		return
	}
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		return
	}

	return
}

func (this *Transfer) WritePkg(conn net.Conn, data []byte) (err error) {
	//发送对方一个长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Printf("conn.Write err %v\n", err)
		return
	}
	//发送消息本身
	n, err = conn.Write(data)
	if n != 4 || err != nil {
		fmt.Printf("conn.Write err %v\n", err)
		return
	}

	return
}
