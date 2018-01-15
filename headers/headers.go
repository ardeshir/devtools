package main
// go build -ldflags "-X main.version=0.0.1"
import (
	"net/http"
	"encoding/base64"
	"bytes"
	"io/ioutil"
	"os"
	// "log"
	"fmt"
  u "github.com/ardeshir/version"
)

var version string = "0.0.1"

func main(){

	fmt.Println("Headers.go")
	req, err := http.NewRequest("GET", "https://httpbin.org/basic-auth/user/passw0rd", nil)
	u.ErrNil(err, "Unable to create request")


	buffer := &bytes.Buffer{}

	enc := base64.NewEncoder(base64.URLEncoding, buffer)
	enc.Write([]byte("user:passw0rd"))
	enc.Close()

	encodedCreds, err := buffer.ReadString('\n')

	if err != nil && err.Error() != "EOF" {
		os.Exit(-1)
	}

    req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedCreds))

	resp, err := http.DefaultClient.Do(req)
	u.ErrNil(err, "Failed to do request")
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	u.ErrNil(err, "Unable to read body")


    fmt.Println( string(content) )
    fmt.Println( resp.StatusCode )

    u.V(version)
}

/* func errNil ( err error, mesg string) {
		if err != nil {
		log.Fatalln(mesg)
	}
} */
