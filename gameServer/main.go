package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/flatbuffers/go"
	"log"
	"meckz-netcode/UDPserver"
	"net"
	"net/http"
	"runtime"
	"sync"
)

/*
	GameServer waits for WebServer to authenticate the client
	when its done , GameServer receives an update from WebServer
	including :
	{
	 clientId ,
     encryptionKey [base64],
     token [base64]
	}
	an authentication must be implemented later , only WebServer can consume
	this api .

 */
var  wg sync.WaitGroup
var  registrationQueue [] Registration
var  clients [] Client
type Registration struct {
	Token string  `json:"token"`
	EncryptionKey string  `json:"encryptionKey"`
	ClientId int32 `json:"clientId"`
}

type Client struct {

	Token string
	ClientId int32
	Address net.UDPAddr
	SequenceNumber uint32
	Connected bool
}

type myError struct {
	message string
}

func (e *myError) Error() string {
	return e.message
}

func RegisterClient(w http.ResponseWriter, req *http.Request){
	cred := Registration{}
	err := json.NewDecoder(req.Body).Decode(&cred)
	if err != nil {
		log.Fatal(err)
	}
	registrationQueue = append(registrationQueue,cred)
	fmt.Println("A registration request with client id :",cred.ClientId)
}

func StartHttpServer(){
	fmt.Println("Http server started")
	http.HandleFunc("/regclient",RegisterClient )
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	wg.Done()
}

func StartUDPserver(){
	addr := net.UDPAddr{
		Port: 3030,
		IP:  []byte{0,0,0,0},
		Zone:""}
	serverConn, _ := net.ListenUDP("udp",&addr)
	defer serverConn.Close()
	receive,send,_ := UDPserver.Init(serverConn)
	fmt.Println("UDP server started")
	//wait for clients
	waitForClients := true
	// handshake phase
	for waitForClients {
		 packet :=<- receive
		 packetType := UDPserver.GetPacketType(packet.Content)
		 if packetType == 1 {
			 addClient(packet,send)
		 }else if packetType == 0{
		 	 clientIndex,err:=getClient(&packet)
		 	 if err!=nil{
		 	 	log.Println(err)
			 } else {
			 	//later we gonna buffer the input
			 	sequenceNumber,_ :=UDPserver.ReadPacket(packet.Content)
			 	clients[clientIndex].Connected=true
			 	clients[clientIndex].SequenceNumber=sequenceNumber
			 }
		 }
		 if getConnectedClientsCount() == 2 {
		 	waitForClients = false
		 }
	}
	fmt.Println("Clients are ready")
	for {
		packet:=<-receive
		packetType := UDPserver.GetPacketType(packet.Content)
		if packetType == 0 {
			clientIndex,err:=getClient(&packet)
			fmt.Println(clientIndex)
			if err!=nil{
				log.Println(err)
			} else {
				for index,element := range clients {
					if index!=clientIndex {
						fmt.Println("sending to",element.Address)
						send<-UDPserver.Bundle{Address: element.Address, Content: packet.Content}
					}
				}
			}
		}
	}
	wg.Done()
}
//func sendState(clientIndex int ,packet UDPserver.Bundle,send chan UDPserver.Bundle){
//	for index,element := range clients {
//		if index!=clientIndex {
//			send<-UDPserver.Bundle{Address: element.Address, Content: packet.Content}
//		}
//	}
//}
func getConnectedClientsCount()( count int ){
	count=0
	for _,element:= range clients {
		if element.Connected {
			count++
		}
	}
	return count
}
func getClient(packet * UDPserver.Bundle) (clientIndex int, err error){
	clientIndex =-1
	for index,element:= range clients {
		if element.Address.IP.Equal(packet.Address.IP)&&element.Address.Port==packet.Address.Port{
			clientIndex = index
		}
	}
	if clientIndex == -1  {
		return clientIndex,&myError{"client not found"}
	}else {
		return clientIndex,nil
	}
}
func addClient(packet UDPserver.Bundle,send chan UDPserver.Bundle){
		token,sequenceNumber :=UDPserver.ReadConnPacket(packet.Content)
		for index, element := range registrationQueue {
			if element.Token == token {
				newClient :=Client{
					token,
				 	element.ClientId,
					packet.Address,
					0,
					false}
				//remove from registration queue
				registrationQueue= append(registrationQueue[:index], registrationQueue[index+1:]...)
				clients = append(clients,newClient)
				b:=flatbuffers.NewBuilder(200)
				acknowledge :=UDPserver.MakeAck(b,sequenceNumber)
				send<-UDPserver.Bundle{Address: newClient.Address, Content: acknowledge}
			}
		}
		for _,element:= range clients {
			if element.Token == token {
				b:=flatbuffers.NewBuilder(200)
				acknowledge :=UDPserver.MakeAck(b,sequenceNumber)
				send<-UDPserver.Bundle{Address: packet.Address, Content: acknowledge}
			}
		}
}
func init(){
	runtime.GOMAXPROCS(1)
}
func main()  {
	clientA :=&Registration{"tokenA","keyA",1}
	clientB :=&Registration{"tokenB","keyB",2}
	registrationQueue = append(registrationQueue,*clientA,*clientB)
	fmt.Println(registrationQueue)
	wg.Add(1)
	 	go StartUDPserver()
	wg.Wait()

}
