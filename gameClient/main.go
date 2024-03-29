package main

import (
	"fmt"
	"github.com/google/flatbuffers/go"
	"log"
	"meckz-netcode/UDPserver"
	"meckz-netcode/UDPserver/ack"
	"net"
	"sync"
	"time"
)

/* Assumption : client has requested the http server to get credentials
				for sake of simplicity we skip this step.
    clientA :=&Registration{Token :"tokenA",Key:"keyA",ClientId :1}
	clientB :=&Registration{Token : "tokenB",Key : "keyB",ClientId :2}

 */
 var wg sync.WaitGroup
 func  dummyClient(port int ,token string) {
	 addr := net.UDPAddr{
		 Port: port,
		 IP:  []byte{0,0,0,0},
		 Zone:""}
	 connection,err := net.ListenUDP("udp",&addr)
	 if err!=nil {
		 log.Fatal(err)
	 }
	 defer connection.Close()
	 serverAddr := net.UDPAddr{
		 Port: 3030,
		 IP:  []byte{185,94,99,104},
		 Zone:""}
	 b :=flatbuffers.NewBuilder(200)
	 var clientSequenceNumber uint32 = 0
	 data :=UDPserver.MakeConnPacket(b,clientSequenceNumber,token)
	 discChannel :=make(chan bool)
	 go connect(connection, serverAddr,data,discChannel)
	 for {
		 inputBytes := make([]byte, 200)
		 connection.ReadFromUDP(inputBytes)
		 payload := ack.GetRootAsAck(inputBytes,0)
		 payloadType:= payload.Type()
		 if payloadType == 2 {
		 	discChannel<-false
			break
		 }
	 }
	 go getState(connection)
	 for {
		 fmt.Println("state")
		 clientSequenceNumber++
		 data := UDPserver.MakePacket(b,clientSequenceNumber,100,100)
		 connection.WriteToUDP(data,&serverAddr)
		 time.Sleep(time.Millisecond*150)
	 }
 }
func getState(conn * net.UDPConn)  {
	for{
		inputBytes := make([]byte, 200)
		conn.ReadFromUDP(inputBytes)
		payload := ack.GetRootAsAck(inputBytes,0)
		payloadType:= payload.Type()
		fmt.Println("got state",payloadType)
	}
}
func connect(conn * net.UDPConn,addr net.UDPAddr,data [] byte,disChannel chan bool)  {
	for{
		select {

		case<-disChannel:
			fmt.Println("Connected")
			return
		default:
			conn.WriteToUDP(data,&addr)
			fmt.Println("sending")
		}
		time.Sleep(time.Millisecond*250)
	}
}
func main()  {
	//client listens on port 3031
	wg.Add(2)
		go dummyClient(4040,"tokenA")
		go dummyClient(4041,"tokenB")
	wg.Wait()

}