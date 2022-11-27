package main

import (
	"crypto/tls"
	"flag"
	"io/ioutil"
	"net/http"
	"time"

	"algs4/config"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
)

var httpVersion = flag.Int("version", 1, "The http version HTTP/1 or HTTP/2")

func main() {
	flag.Parse()

	server := &http.Server{
		Addr:         config.DefaultAddr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		TLSConfig:    tlsConfig(),
	}

	if *httpVersion == 2 {
		// Having this does not change anything but just showing.
		if err := http2.ConfigureServer(server, nil); err != nil {
			logrus.Fatal(err)
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, I am your server for your service."))
	})

	if err := server.ListenAndServeTLS("", ""); err != nil {
		logrus.Fatal(err)
	}
}

func tlsConfig() *tls.Config {
	crt, err := ioutil.ReadFile("./cert/server.crt")
	if err != nil {
		logrus.Fatal(err)
	}

	key, err := ioutil.ReadFile("./cert/server.key")
	if err != nil {
		logrus.Fatal(err)
	}

	cert, err := tls.X509KeyPair(crt, key)
	if err != nil {
		logrus.Fatal(err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
}
