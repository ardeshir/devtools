package main
// go build -ldflags "-X main.version=0.0.1"

import (
	"net/http"
	// "encoding/json"
	// "encoding/base64"
	// "bytes"
   "io/ioutil"
   "os"
	// "log"
	"fmt"
   u "github.com/ardeshir/version"
)

var version string = "0.0.1"

func main() {
    
    resp, err := http.Get(os.Args[1])
    u.ErrNil(err, "Unable to read response")
    defer resp.Body.Close()
    content, err := ioutil.ReadAll(resp.Body)
    u.ErrNil(err, "Unable to read content")
    
    fmt.Println(string(content))
 // -----------------  footer ----------- //    
 u.V(version)
}
