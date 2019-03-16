package UDPserver

import (
	"fmt"
	"net"
)

type Bundle struct {
	Address net.UDPAddr
	Content  []byte
}

func listen(receive chan Bundle,serverConn *net.UDPConn,control chan bool) {

	for  {
		select {
		case <-control:
			fmt.Println("UDPserver listening is finished")
			return
		default:
			inputBytes := make([]byte, 200)
			_,source,_ := serverConn.ReadFromUDP(inputBytes)
			receive<-Bundle{*source,inputBytes}
		}
	}
}
func broadcast (send chan Bundle,serverConn *net.UDPConn,control chan bool){

	for  {
		select {
		case <-control:
			fmt.Println("UDPserver writing is finished")
			return
		default:
			payload:=<-send
			serverConn.WriteToUDP(payload.Content ,&payload.Address)
		}
	}
}
func Init(serverConn *net.UDPConn) (chan Bundle, chan Bundle,chan bool){
	receive := make(chan Bundle, 10)
	send := make(chan Bundle, 10)
	control := make(chan bool)
	go listen(receive,serverConn,control)
	go broadcast(send,serverConn,control)
	return receive, send,control
}