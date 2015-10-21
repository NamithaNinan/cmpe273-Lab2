package main

import (
    "encoding/json"
    "encoding/binary"
    "log"
    "net/http"
    "io/ioutil"
    "fmt"
    "bytes"
)

type input struct {
   // Tag string
    Name string
}
type output struct {
    Greeting string
    
}
func prettyprint(b []byte) ([]byte, error) {
    var out bytes.Buffer
    err := json.Indent(&out, b, "", "  ")
    return out.Bytes(), err
}

func test(rw http.ResponseWriter, req *http.Request) {
    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        panic("run")
    }
    log.Println(string(body))
    var t input
    
    err = json.Unmarshal(body, &t)
    if err != nil {
        panic("run")
    }
    Responseoutput :=" Hello, "+t.Name+" !"
    x:=output{Responseoutput}
    b,err:=json.Marshal(x)   
    b, _ = prettyprint(b)
    //n := bytes.Index(b, []byte{0})
    n:=binary.Size(b)
    s := string(b[:n])
    fmt.Fprintf(rw,s)
    
}

func main() {
    http.HandleFunc("/hello", test)
    log.Fatal(http.ListenAndServe(":8082", nil))
}
/* 
OUTPUT:
request:
{
    "Name" : "Namitha Ninan"
}
respone:
{
  "Greeting": " Hello, Namitha Ninan !"
}
*/