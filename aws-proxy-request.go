package main

import (
  "fmt"
  "log"
  "strings"
  "net/http"
  "io/ioutil"
  "github.com/smartystreets/go-aws-auth"
  "gopkg.in/alecthomas/kingpin.v2"
)

var (
  method = kingpin.Flag("method", "Request method.").Default("POST").String()
  body = kingpin.Flag("body", "JSON body of request.").Default("").String()
  ESUrl = kingpin.Arg("ESUrl", "URL.").Required().String()
)

func main() {
  kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version("1.0").Author("Sergey Zelenov")
  kingpin.CommandLine.Help = "An implementation of AWS signing requests."
  
  kingpin.Parse()

  req, err := http.NewRequest(*method, *ESUrl, strings.NewReader(*body))
  if err != nil {
		log.Fatalln(err)
  }

  awsauth.Sign(req)
  
  resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
  }
  
  defer resp.Body.Close()

  response, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatalln(err)
  }

  fmt.Println(string(response))
}
