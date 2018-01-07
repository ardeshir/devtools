package main
// go build -ldflags "-X main.version=0.0.1"

import (
	"net/http"
	// "encoding/base64"
	// "bytes"
   "io/ioutil"
	// "os"
	// "log"
	"fmt"
   u "github.com/ardeshir/version"
)

var version string = "0.0.1"

func main() {

  resp, err := http.Get("https://httpbin.org/get")
  u.ErrNil(err, "Unable to get resp")
  
 defer resp.Body.Close()
 content, err := ioutil.ReadAll(resp.Body)
 u.ErrNil(err, "Unable to read resp")
 fmt.Println(string(content))


  u.V(version)
}
