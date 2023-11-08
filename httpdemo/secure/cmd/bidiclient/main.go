package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func main() {
	cert, err := tls.LoadX509KeyPair("./cert/client.crt", "./cert/client.key")
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
		RootCAs:            clientCertPool,
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: false,
		ServerName:         "rabbit.com", // server cert common name
	}

	conn, err := tls.Dial("tcp", "localhost:9973", cfg)
	if err != nil {
		logrus.Fatal(err)
	}
	defer conn.Close()

	logrus.Info("dial succeeded")

	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)

	n, err := w.WriteString("Hello from client.\n")
	if err != nil {
		logrus.WithField("n", n).Fatal(err)
	}
	err = w.Flush()
	if err != nil {
		logrus.Fatal(err)
	}

	msg, err := r.ReadString('\n')
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.WithField("msg", msg).Info("message from server")
}
