package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"room/commom/message"
)

//关联到结构体
type Transfer struct {
	//分析传输需要哪些字段
	Conn net.Conn
	Buf  [8096]byte //这是缓冲
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		fmt.Printf("conn.Read  err = %v ", err)
		return
	}
	fmt.Printf("读buf = %v\n", this.Buf[:4])

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])

	b, err := this.Conn.Read(this.Buf[0:pkgLen])
	if b != int(pkgLen) || err != nil {
		err = errors.New("Read pkg err")
		return
	}
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		return
	}

	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	//发送对方一个长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Printf("conn.Write err %v\n", err)
		return
	}

	//发送消息本身

	n, err = this.Conn.Write(data)
	if n != 4 || err != nil {
		fmt.Printf("conn.Write err %v\n", err)
		return
	}

	return
}
