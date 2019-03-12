package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync/atomic"
)

/*
	HTTPS encryption is "RSA" ≥ 2048-bit
	Key considerations for algorithm "ECDSA" (X25519 || ≥ secp384r1)
	https://safecurves.cr.yp.to/
    List ECDSA the supported curves (openssl ecparam -list_curves)
	cert files = [ server.crt , server.key ]
 */
const (
	tokenSize = 64
	encryptionKeySize = 32
)
var clientId int32

type ClientResponse struct {
	 Token string  `json:"token"`
	 EncryptionKey string  `json:"encryptionKey"`
}
type Registration struct {

	Token string  `json:"token"`
	EncryptionKey string  `json:"encryptionKey"`
	ClientId int32 `json:"clientId"`
}

func GenerateToken (w http.ResponseWriter, req *http.Request) {
	 //Authentication goes her

	 //Check gameServer if there is room fro client

	 //Generate random byte as token
	 tokenByte,err := RandomBytes(tokenSize)
	 if  err!= nil {
	 	log.Fatal(err)
	 }
	 //Generate random byte as encryption key
	keyByte,err := RandomBytes(encryptionKeySize)
	if  err!= nil {
		log.Fatal(err)
	}
	 base64Token := base64.StdEncoding.EncodeToString(tokenByte)
	 base64Key := base64.StdEncoding.EncodeToString(keyByte)
	 atomic.AddInt32(&clientId,1 )
	 //Send Registration to GameServer
	sendCredsToGameServer(base64Token,base64Key,clientId)
	resp := ClientResponse{base64Token,base64Key}
	json.NewEncoder(w).Encode(resp)
	//log
	fmt.Println("Client Joined the gameServer with ID:",clientId)
}
func sendCredsToGameServer(token string , encryptionKey string,clientId int32  ){

	body := &Registration{
		Token: token,
		EncryptionKey: encryptionKey,
		ClientId: clientId,
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)
	client := &http.Client{}
	req, _ := http.NewRequest("POST","http://localhost:5000/regclient", buf)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
func RandomBytes(bytes int) ([]byte, error) {
	b := make([]byte, bytes)
	_, err := rand.Read(b)
	return b, err
}
func main()  {
	http.HandleFunc("/token", GenerateToken)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
