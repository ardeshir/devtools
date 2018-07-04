package main
// Help for developers:
// go build -ldflags "-X main.version=0.0.1"
// godoc -http :9090
// root@ardeshir ~/httpdoc (master) $ graphpkg -stdout net > graph
// ardeshir.org/graph to view the SVG of all dependencies 

import (
	"net/http"
	"flag"
	// "encoding/json"
	// "encoding/base64"
	// "bytes"
   "io/ioutil"
   "os"
	// "log"
	"fmt"
	"errors"
   u "github.com/ardeshir/version"
)

var mymap = map[string]string{
       "first": "Ardeshir",
       "last": "Sepahsalar",
       "age": "47",
       "role": "Developer",
       }

type Person struct {
    First string
    Last string
    Age  int 
    Role string
    Email string 
}

func (p Person) Salary() (int, error) {
     if p.Role == "" {
         return 0, errors.New("Unable to handle empty value for Role!")
     }
      switch p.Role {
          
          case "Developer" : 
             return 100, nil
          case "Architect" : 
             return 200, nil
       }
       return 0, errors.New( fmt.Sprintf("Unable to parse role: '%s' ", p.Role), )
}

func (p *Person) UpdateEmail(email string) {
    p.Email = email;
}


func main() {
    
   ardeshir := Person{   First: "Ardeshir",  Last: "Sepahsalar", Age: 47, Role: "Developer" }
   john     := Person{   First: "John",  Last: "Jupyter", Age: 70, Role: "Architect" }
   jack     := Person{   First: "Jack",  Last: "Johnson", Age: 30, Role: "Designer" }
   jean     := Person{   First: "Jean",  Last: "Julin", Age: 20, Role: "" }
   
   if aSal, erra :=  ardeshir.Salary(); erra != nil {
       fmt.Println(erra)
   } else {
         fmt.Println(ardeshir.First, aSal)
   }
   
   if jSal, errj :=  john.Salary(); errj != nil {
       fmt.Println(errj)
   } else {
       fmt.Println(john.First,jSal )
   }
   
   if kSal, errk := jack.Salary(); errk != nil {
       fmt.Println(errk)
   } else {
      fmt.Println(jack.First, kSal) 
   }
   
   if nSal, errn := jean.Salary();  errn != nil {
       fmt.Println(errn)
   } else {
       fmt.Println(jean.First, nSal) 
   }
 
   
   
   
   ardeshir.UpdateEmail("ardeshir.org@gmail.com")
   fmt.Println(ardeshir.Email)
   
     first, ok := mymap["first"]

     if ok == true {
         fmt.Println("First:", first)
     } else {
         fmt.Println("First is not defined")
     }

    for key, value := range mymap {
        fmt.Println("key: ", key, " value:", value)
    }


    url2get  := flag.String("url", defaultURL() , "Link to download")
    flag.Parse()

    // resp, err := http.Get(os.Args[1])

    resp, err := http.Get(*url2get)
    u.ErrNil(err, "Unable to read response")
    defer resp.Body.Close()
    content, err := ioutil.ReadAll(resp.Body)
    u.ErrNil(err, "Unable to read content")

    fmt.Println(string(content))


 // --------- basic housekeeping ----------- // 
 // if we're int debug mode, print out info 

   if debugTrue() {
    u.V( defaultVersion() )
   }
}

// Function to check env variable DEFAULT_DEBUG bool
func debugTrue() bool {

     if os.Getenv("DEFAULT_DEBUG") != "" {
        return true
     }
     return false
}

// Function to check env variable DEFAULT_VERSION string
func defaultVersion() string {

 if os.Getenv("DEFAULT_VERSION") != "" {
     return os.Getenv("DEFAULT_VERSION")
  }

    var version string = "0.0.1"
 return version
}

// Function to check env variable DEFAULT_URL to http get
func defaultURL() string {
    if os.Getenv("DEFAULT_URL") != "" {
        return os.Getenv("DEFAULT_URL")
    }
    return "https://httpbin.org/get"
}
