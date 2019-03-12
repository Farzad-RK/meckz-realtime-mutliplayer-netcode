package UDPserver

import (
	"fmt"
	"net"
)

type Bundle struct {
	Address net.UDPAddr
	Content  []byte
}

func listen(receive chan Bundle, port int,control chan bool) {
	addr := net.UDPAddr{
		Port: port,
		IP:  []byte{0,0,0,0},
		Zone:""}
	ServerConn, _ := net.ListenUDP("udp",&addr)
	defer ServerConn.Close()
	for  {
		select {
		case <-control:
			fmt.Println("UDPserver listening is finished")
			return
		default:
			inputBytes := make([]byte, 200)
			_,source,_ := ServerConn.ReadFromUDP(inputBytes)
			receive<-Bundle{*source,inputBytes}
		}
	}
}
func broadcast (send chan Bundle,port int ,control chan bool){
	addr := net.UDPAddr{
		Port: port,
		IP:  []byte{0,0,0,0},
		Zone:""}
	ServerConn, _ := net.ListenUDP("udp",&addr)
	for  {
		select {
		case <-control:
			fmt.Println("UDPserver writing is finished")
			return
		default:
			payload:=<-send
			ServerConn.WriteToUDP(payload.Content ,&payload.Address)
		}
	}
}
func Init(readPort int,writePort int) (chan Bundle, chan Bundle,chan bool){
	receive := make(chan Bundle, 10)
	send := make(chan Bundle, 10)
	control := make(chan bool)
	go listen(receive, readPort,control)
	go broadcast(send ,writePort,control)
	return receive, send,control
}