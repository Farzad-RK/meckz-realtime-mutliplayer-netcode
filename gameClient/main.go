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
 func  dummyClient(port int ) {
	 addr := net.UDPAddr{
		 Port: port,
		 IP:  []byte{127,0,0,1},
		 Zone:""}
	 connection,err := net.ListenUDP("udp",&addr)
	 if err!=nil {
		 log.Fatal(err)
	 }
	 defer connection.Close()
	 serverAddr := net.UDPAddr{
		 Port: 3030,
		 IP:  []byte{127,0,0,1},
		 Zone:""}
	 b :=flatbuffers.NewBuilder(200)
	 var clientSequenceNumber uint32 = 0
	 data :=UDPserver.MakeConnPacket(b,clientSequenceNumber,"tokenA")
	 //stop and wait protocol with 250ms delay
	 for {
		 connection.WriteToUDP(data,&serverAddr)
		 inputBytes := make([]byte, 200)
		 connection.ReadFromUDP(inputBytes)
		 payload := ack.GetRootAsAck(inputBytes,0)
		 payloadType:= payload.Type()
		 if payloadType == 2 {
			 fmt.Println("Done")
			 break
		 } else {
			 connection.WriteToUDP(data,&serverAddr)
			 fmt.Println("resend")
		 }
		 time.Sleep(time.Millisecond*250)
	 }
	 for {
	 	clientSequenceNumber++
	 	data = UDPserver.MakePacket(b,clientSequenceNumber,100,100)
		connection.WriteToUDP(data,&serverAddr)
	 }
	 wg.Done()
 }
func main()  {
	//client listens on port 3031
	wg.Add(2)
		go dummyClient(4040)
		go dummyClient(4041)
	wg.Wait()

}