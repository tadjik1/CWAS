package main

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type HTTPHeaderValue http.Header

var (
	region  = kingpin.Flag("region", "AWS region").Required().String()
	service = kingpin.Flag("service", "AWS Service short name").Required().String()
	url     = kingpin.Arg("url", "Request url").Required().String()
	method  = kingpin.Flag("method", "Request method.").Default("GET").Short('X').String()
	body    = kingpin.Flag("body", "Request body.").Default("").Short('d').String()
	headers = HTTPHeader(kingpin.Flag("header", "Request headers (could be an array).").Short('H'))
)

func (h *HTTPHeaderValue) Set(value string) error {
	parts := strings.SplitN(value, ":", 2)
	if len(parts) != 2 {
		return fmt.Errorf("Expected Header:Value got '%s'", value)
	}
	(*http.Header)(h).Add(parts[0], parts[1])
	return nil
}

func (i *HTTPHeaderValue) IsCumulative() bool {
	return true
}

func (h *HTTPHeaderValue) String() string {
	return ""
}

func HTTPHeader(s kingpin.Settings) (target *http.Header) {
	target = &http.Header{}
	s.SetValue((*HTTPHeaderValue)(target))
	return
}

func replaceBody(req *http.Request) []byte {
	if req.Body == nil {
		return []byte{}
	}
	payload, _ := ioutil.ReadAll(req.Body)
	req.Body = ioutil.NopCloser(bytes.NewReader(payload))
	return payload
}

func main() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version("1.0").Author("Sergey Zelenov")
	kingpin.CommandLine.Help = "[K]WAS - Curl With AWS Signing."

	kingpin.Parse()

	sess, err := session.NewSession()
	if err != nil {
		log.Fatalln(err)
	}
	signer := v4.NewSigner(sess.Config.Credentials)

	req, err := http.NewRequest(*method, *url, strings.NewReader(*body))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header = *headers

	payload := bytes.NewReader(replaceBody(req))
	_, err = signer.Sign(req, payload, *service, *region, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

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
