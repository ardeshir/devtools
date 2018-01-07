package main
// go build -ldflags "-X main.version=0.0.1"

import (
	"net/http"
	// "encoding/base64"
	// "bytes"
	"io/ioutil"
	// "os"
	"log"
	"fmt"
)

func main(){

	fmt.Println("Headers.go")
	req, err := http.NewRequest("GET", "https://httpbin.org/basic-auth/user/passw0rd", nil)
	errNil(err, "Unable to create request")
	
	req.SetBasicAuth("user","passw0rd")
	
	/*
	buffer := &bytes.Buffer{}
	
	enc := base64.NewEncoder(base64.URLEncoding, buffer)
	enc.Write([]byte("user:passw0rd"))
	enc.Close()
	
	encodedCreds, err := buffer.ReadString('\n')
	
	if err != nil && err.Error() != "EOF" { 
		os.Exit(-1)
	}
    req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedCreds))
	*/
	
	resp, err := http.DefaultClient.Do(req)
	errNil(err, "Failed to do request")
	defer resp.Body.Close()
	
	content, err := ioutil.ReadAll(resp.Body)
	errNil(err, "Unable to read body")
	

    fmt.Println( string(content) )
    fmt.Println( resp.StatusCode )
}

func errNil ( err error, mesg string) {
		if err != nil {
		log.Fatalln(mesg)
	}
}