package main
// go build -ldflags "-X main.version=0.0.1"

import (
	"net/http"
	"encoding/json"
	// "encoding/base64"
	// "bytes"
   "io/ioutil"
	// "os"
	// "log"
	"fmt"
   u "github.com/ardeshir/version"
)

var version string = "0.0.1"

type GetResponse struct {
  Origin string `json:"origin`
  URL string    `json:"url"`
  Headers map[string]string  `json:"headers"`
}

func (r *GetResponse) ToString() string { 
    s := fmt.Sprintf("GET Response\n Origin IP: %s\nRequest URL: %s\n",
         r.Origin, r.URL)
         
         for k, v := range r.Headers {
             s += fmt.Sprintf("Header: %s = %s \n", k, v)
         }
    return s
}

func main() {

  resp, err := http.Get("https://httpbin.org/get")
  u.ErrNil(err, "Unable to get resp")
  
 defer resp.Body.Close()
 content, err := ioutil.ReadAll(resp.Body)
 u.ErrNil(err, "Unable to read resp")
 
  respContent := GetResponse{}
  
 //fmt.Println(string(content))
 json.Unmarshal(content, &respContent )
  
 fmt.Println(respContent.ToString())

}
