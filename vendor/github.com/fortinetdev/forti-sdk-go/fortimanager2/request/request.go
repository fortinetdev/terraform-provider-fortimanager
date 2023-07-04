package request

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/fortinetdev/forti-sdk-go/fortimanager2/config"
)

// Request describes the request to FMG service
type Request struct {
	Config       config.Config
	HTTPRequest  *http.Request
	HTTPResponse *http.Response
	Path         string
	Params       interface{}
	Data         *bytes.Buffer
}

// New creates reqeust object with http method, path, params and data,
// It will save the http request, path, etc. for the next operations
// such as sending data, getting response, etc.
// It returns the created request object to the gobal plugin client.
func New(c config.Config, method string, path string, params interface{}, data *bytes.Buffer) *Request {
	var h *http.Request

	if data == nil {
		h, _ = http.NewRequest(method, "", nil)
	} else {
		h, _ = http.NewRequest(method, "", data)
	}

	r := &Request{
		Config:      c,
		Path:        path,
		HTTPRequest: h,
		Params:      params,
		Data:        data,
	}
	return r
}

// func New(c config.Config, method string, path string, params interface{}, data *bytes.Buffer) *Request {
// 	var h *http.Request

// 	if data == nil {
// 		h, _ = http.NewRequest(method, "", nil)
// 	} else {
// 		h, _ = http.NewRequest(method, "", data)
// 	}

// 	r := &Request{
// 		Config:      c,
// 		Path:        path,
// 		HTTPRequest: h,
// 		Params:      params,
// 		Data:        data,
// 	}
// 	return r
// }

// Build Request header

// Build Request Sign/Login Info

// Send request data to FMG.
// If errors are encountered, it returns the error.
func (r *Request) Send() error {


	//Build FMG
	//build Sign/Login INfo

	//httpReq.URL, err = url.Parse(clientInfo.Endpoint + operation.HTTPPath)

	r.HTTPRequest.Header.Set("Content-Type", "application/json")
	u := buildURL(r)

	var err error
	r.HTTPRequest.URL, err = url.Parse(u)
	if err != nil {
		log.Fatal(err)
		return err
	}

	retry := 0
	for {
		//Send
		rsp, errdo := r.Config.HTTPCon.Do(r.HTTPRequest)
		r.HTTPResponse = rsp
		if errdo != nil {
			if strings.Contains(errdo.Error(), "x509: ") {
				err = fmt.Errorf("Error found: %v", errdo)
				break
			}

			if retry > 15 {
				err = fmt.Errorf("lost connection to firewall with error: %v", errdo)
				break
			}
			time.Sleep(time.Second)
			log.Printf("Error found: %v, will resend again %s, %d", errdo, u, retry)

			retry++

		} else {
			break
		}
	}

	return err
}

func buildURL(r *Request) string {
	u := "https://"
	u += r.Config.FwTarget
	u += r.Path
	u += "?"

	return u
}

// func (r *Request) Send() error {
// 	//Build FMG
// 	//build Sign/Login INfo

// 	//httpReq.URL, err = url.Parse(clientInfo.Endpoint + operation.HTTPPath)

// 	r.HTTPRequest.Header.Set("Content-Type", "application/json")
// 	u := buildURL(r)

// 	var err error
// 	r.HTTPRequest.URL, err = url.Parse(u)
// 	if err != nil {
// 		log.Fatal(err)
// 		return err
// 	}

// 	retry := 0
// 	for {
// 		//Send
// 		rsp, errdo := r.Config.HTTPCon.Do(r.HTTPRequest)
// 		r.HTTPResponse = rsp
// 		if errdo != nil {
// 			if strings.Contains(errdo.Error(), "x509: ") {
// 				err = fmt.Errorf("Error found: %v", errdo)
// 				break
// 			}

// 			if retry > 15 {
// 				err = fmt.Errorf("lost connection to firewall with error: %v", errdo)
// 				break
// 			}
// 			time.Sleep(time.Second)
// 			log.Printf("Error found: %v, will resend again %s, %d", errdo, u, retry)

// 			retry++

// 		} else {
// 			break
// 		}
// 	}

// 	return err
// }

// func buildURL(r *Request) string {
// 	u := "https://"
// 	u += r.Config.FwTarget
// 	u += r.Path
// 	u += "?"

// 	return u
// }
