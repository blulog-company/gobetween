package udpproxyprotocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
)

const MagicNumber uint16 = 0x56EC

func addrToIPAndPort(addr net.Addr) (ip net.IP, port uint16, err error) {
	ipString, portString, err := net.SplitHostPort(addr.String())
	if err != nil {
		return
	}

	ip = net.ParseIP(ipString)
	if ip == nil {
		err = fmt.Errorf("Could not parse IP")
		return
	}

	p, err := strconv.ParseInt(portString, 10, 64)
	if err != nil {
		return
	}
	port = uint16(p)
	return
}

func NewByteHeader(clientAddr *net.UDPAddr) (frame []byte, size int, err error) {
	clientIP, clientPort, err := addrToIPAndPort(clientAddr)
	if err != nil {
		return nil, 0, err
	}

	frame = make([]byte, 20)

	buf := bytes.NewBuffer(frame)
	buf.Reset()
	err = binary.Write(buf, binary.BigEndian, MagicNumber)
	if err != nil {
		return nil, 0, err
	}

	_, err = buf.Write(clientIP.To16())
	if err != nil {
		return nil, 0, err
	}

	err = binary.Write(buf, binary.BigEndian, clientPort)
	if err != nil {
		return nil, 0, err
	}

	byteHeader := buf.Bytes()
	return byteHeader, buf.Len(), nil
}
