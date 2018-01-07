package main
// go build -ldflags "-X main.version=0.0.1"

import (
	"net/http"
	// "encoding/json"
	// "encoding/base64"
	// "bytes"
   "io/ioutil"
   "os"
   "log"
	"fmt"
   u "github.com/ardeshir/version"
)

var version string = "0.0.1"
var router  = NewRouter()

func init() {
    router.Register(200, func(resp *http.Response) {
            defer resp.Body.Close()
            content, err := ioutil.ReadAll(resp.Body)
            u.ErrNil(err, "Unable to read content")
    
            fmt.Println(string(content))  
    })
    router.Register(404, func(resp *http.Response) {
         log.Fatalln("Not Found (404):", resp.Request.URL.String())
    })
    router.Register(403, func(resp *http.Response) {
         log.Fatalln("Not Found (403):", resp.Request.URL.String())
    })
}

func main() {
    
    resp, err := http.Get(os.Args[1])
    u.ErrNil(err, "Unable to read response")
    router.Process(resp)
 // -----------------  footer ----------- //    
 u.V(version)
}

