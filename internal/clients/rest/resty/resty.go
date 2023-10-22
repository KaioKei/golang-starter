package resty

import (
	"github.com/go-resty/resty/v2"
	"log"
	"net/url"
)

// This tutorial shows the usage of a rest client using the module Resty.
// Use this client with the tutorial Gin :
// [Gin tutorial](../../../api/rest/gin)

var client *resty.Client

var myUrl *url.URL

func Run() {
	// init connexion
	client = resty.New()
	myUrl, _ = url.Parse("http://localhost:8080")

	// GET
	log.Println("GET method")
	restyGet(myUrl)

}

func restyGet(u *url.URL) {
	// use this method to add path dynamically
	u = u.JoinPath("/albums")
	// get response
	resp, err := client.R().EnableTrace().Get(u.String())
	if err != nil {
		log.Fatal("GET error:", err)
	}

	// Explore response object
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()

	// Explore trace info
	log.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	log.Println("  DNSLookup     :", ti.DNSLookup)
	log.Println("  ConnTime      :", ti.ConnTime)
	log.Println("  TCPConnTime   :", ti.TCPConnTime)
	log.Println("  TLSHandshake  :", ti.TLSHandshake)
	log.Println("  ServerTime    :", ti.ServerTime)
	log.Println("  ResponseTime  :", ti.ResponseTime)
	log.Println("  TotalTime     :", ti.TotalTime)
	log.Println("  IsConnReused  :", ti.IsConnReused)
	log.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	log.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	log.Println("  RequestAttempt:", ti.RequestAttempt)
	log.Println("  RemoteAddr    :", ti.RemoteAddr.String())
}
