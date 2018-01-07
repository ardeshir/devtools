package main

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

func main() {

  resp, err := http.Get("https://httpbin.org/get")
  u.ErrNil(err, "Unable to get resp")
  
 defer resp.Body.Close()
 content, err := ioutil.ReadAll(resp.Body)
 u.ErrNil(err, "Unable to read resp")
 fmt.Println(string(content))

}
