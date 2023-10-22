package tutorial

import (
	"log"
	"net"
	"net/url"
)

func Schemes() {
	log.Println("Welcome to schemes tutorial !")
	urlScheme()
}

func urlScheme() {
	url1 := "http://myuser:pass@myhost.com:8080/mypath?mykey=myvalue#f"

	parsedUrl1, err := url.Parse(url1)
	if err != nil {
		panic(err)
	}

	// scheme, host, port and path
	log.Println(parsedUrl1.Scheme)
	log.Println(parsedUrl1.Host)
	host, port, _ := net.SplitHostPort(parsedUrl1.Host)
	log.Println(host)
	log.Println(port)
	log.Println(parsedUrl1.Path)

	// auth details
	log.Println(parsedUrl1.User.Username())
	pass, _ := parsedUrl1.User.Password()
	log.Println(pass)

	// query details
	query, _ := url.ParseQuery(parsedUrl1.RawQuery)
	log.Println(query)
	log.Println("Query key value:", query["mykey"])

}

func fileScheme() {
	log.Println("TODO !")
}
