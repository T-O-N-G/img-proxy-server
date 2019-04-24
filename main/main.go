package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// setup a http client
var httpTransport = &http.Transport{}
var httpClient = &http.Client{Transport: httpTransport}

func getImg(c echo.Context) error {
	url := c.QueryParam("url")
	resp, err := http.Get(url)
	if err != nil {
		return c.String(503, "Failed to get image!")

	}
	defer resp.Body.Close()

	body, readRrr := ioutil.ReadAll(resp.Body)
	if readRrr != nil {
		return c.String(503, "Failed to read from body!")

	}
	return c.Blob(http.StatusOK, "image/jpeg", body)
}

func getImgProxy(c echo.Context) error {
	url := c.QueryParam("url")
	resp, err := httpClient.Get(url)
	if err != nil {
		return c.String(503, "Failed to get image!")

	}
	defer resp.Body.Close()

	body, readRrr := ioutil.ReadAll(resp.Body)
	if readRrr != nil {
		return c.String(503, "Failed to read from body!")

	}
	return c.Blob(http.StatusOK, "image/jpeg", body)
}

func getWebProxy(c echo.Context) error {
	url := c.QueryParam("url")
	resp, err := httpClient.Get(url)
	if err != nil {
		return c.String(503, "Failed to get image!")

	}
	defer resp.Body.Close()

	body, readRrr := ioutil.ReadAll(resp.Body)
	if readRrr != nil {
		return c.String(503, "Failed to read from body!")

	}
	return c.Blob(http.StatusOK, "text/html", body)
}

func main() {
	port := flag.String("port", "", "http listen port")
	proxyAddr := flag.String("proxy", "", "sock5 proxy add:port")

	flag.Parse()

	if *proxyAddr == "" || *port == "" {
		flag.Usage()
		fmt.Print("Example:\n")
		fmt.Print("./build_for_mac -port=\":80\"  -proxy=\"0.0.0.0:1086\"\n\n\n")

		fmt.Print("You can not use proxy without address..\n")
		fmt.Print("Using random port..\n")
	}

	var dialer, proxyErr = proxy.SOCKS5("tcp", *proxyAddr, nil, proxy.Direct)
	if proxyErr != nil {
		_, _ = fmt.Fprintln(os.Stderr, "can't connect to the proxy:", proxyErr)
	}
	// set our socks5 as the dialer
	httpTransport.Dial = dialer.Dial
	e := echo.New()
	e.GET("/img", getImg)
	e.GET("/img_proxy", getImgProxy)
	e.GET("/proxy", getWebProxy)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	s := &http.Server{
		Addr:         *port,
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	e.Logger.Fatal(e.StartServer(s))
}
