package main

import (
	"flag"
	"io"
	"log"
	"net"
	"unsafe"

	"github.com/gorilla/websocket"
)

type Packet struct {
	ID        int16
	DeviceID  string
	SrcMAC    string
	DstMAC    string
	SrcIP     string
	DstIP     string
	SrcPort   string
	DstPort   string
	SYN       bool
	ACK       bool
	Sequence  int64
	Protocol  string
	Length    int64
	DataChank []byte
}

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options
const SOCK_ADDRESS = "/tmp/Vortoj-Packet.sock"

func main() {
	log.SetFlags(log.Lshortfile)

	conn, err := net.Dial("unix", SOCK_ADDRESS)
	if err != nil {
		panic(err)
	}
	data := make([]byte, 0)
	for {
		println("starting packetserver")
		packetbuff := Packet{}
		println(unsafe.Sizeof(packetbuff))
		buf := make([]byte, 500)
		nr, err := conn.Read(buf)
		println("starting redding")

		if err != nil {
			if err != io.EOF {
				log.Printf("error: %v", err)
			}
			break
		}

		buf = buf[:nr]
		log.Printf("receive: %s\n", buf)
		data = append(data, buf...)
	}
	log.Printf("send: %s\n", string(append(data, data...)))

	// flag.Parse()
	// log.SetFlags(0)

	//http.HandleFunc("/statuses/sample", ServerDataAll)
	//http.HandleFunc("/", home)
	//log.Fatal(http.ListenAndServe(*addr, nil))
}
