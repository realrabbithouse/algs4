package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"net"

	"algs4/config"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load server certificate & private key.
	cert, err := tls.LoadX509KeyPair("./cert/server.crt", "./cert/server.key")
	if err != nil {
		logrus.Fatal(err)
	}

	certBytes, err := ioutil.ReadFile("./cert/ca.crt")
	if err != nil {
		logrus.Fatal(err)
	}
	clientCertPool := x509.NewCertPool()
	ok := clientCertPool.AppendCertsFromPEM(certBytes)
	if !ok {
		logrus.Fatal("failed to parse root certificate")
	}

	cfg := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		ClientAuth:         tls.RequireAndVerifyClientCert,
		ClientCAs:          clientCertPool,
		InsecureSkipVerify: false,
		ServerName:         "example", // client cert common name
	}
	ln, err := tls.Listen("tcp", config.DefaultAddr, cfg)
	if err != nil {
		logrus.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			if err != io.EOF {
				logrus.Error(err)
			}
			continue
		}

		logrus.Info("connection accepted")

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)

	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				logrus.Error(err)
			}
			return
		}
		logrus.WithField("msg", msg).Info("message from client")

		n, err := w.WriteString("Hello from server.\n")
		if err != nil {
			logrus.WithField("n", n).Error(err)
			return
		}
		if err = w.Flush(); err != nil {
			logrus.Error(err)
			return
		}
	}
}
