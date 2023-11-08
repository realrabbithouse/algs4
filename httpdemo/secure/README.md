# TLS双向认证

## 证书

### 概念

首先要有一个 CA 根证书，然后用 CA 根证书来签发用户证书。
用户进行证书申请：一般先生成一个私钥，然后用私钥生成证书请求（证书请求里应含有公钥信息），再利用 CA 根证书来签发证书。

1. 自签名证书（一般用于顶级证书、根证书）：证书的名称和认证机构的名称相同。
2. 根证书：根证书是 CA 认证中心给自己颁发的证书，是信任链的起始点。
3. 数字证书：由证书认证机构（CA）对证书申请者真实身份验证之后，用 CA 的根证书对申请人的一些基本信息以及申请人的公钥进行签名后形成的一个数字文件。数字证书包含证书中所标识的实体的公钥，数字证书将公钥与特定的个人匹配，并且证书的真实性由颁发机构保证。

### x509证书

[X.509](https://zh.wikipedia.org/wiki/X.509) 一般会用到三类文件，`key`，`csr`，`crt`

- `key`私钥，`openssl`格式，通常由`rsa`算法生成
- `csr (Certificate Signing Request)`是证书请求文件，用于申请证书，包含公钥信息。在申请的时候，必须使用自己的私钥来签署申请，可以设定一个密钥
- `crt`证书文件，是签署人用自己的 key 给申请人签署的凭证，**可以自签署**
- `pem`是以上三种文件的编码方式之一，另外一种是`DER`，`Base64`编码后的文本格式，可以单独存放证书或密钥
- `crl (Certificate Revocation List)`证书吊销列表

### 证书生成

> 当然，如果我们要弄 `ssl`，又没钱请 CA 给我们签署的时候（他们的验证费都好贵的），可以考虑自己给自己签署。生成一份 `key`，生成一个 `req`，然后用自己的 `key` 签署自己的 `req`. 当你使用这份证书做 `ssl` 的时候，就会产生不受信任的证书警告。你可以在客户那里安装这份证书为根证书，就不会弹出这个警告了。当然，考究的人，签署用证书和服务器身份认证证书是分开的。先做一个自签署证书作为根证书，然后用这个根证书签署一个服务器证书。这样在客户那里安装根证书后，服务器证书就会自动信任。==一本证书只能颁发给一个特定域名。==如果你有多个域名，就要反复在客户这里安装证书。然而如果采用根证书签名，那么只要安装一张根证书，其他都是自动的。

两种处理方式：

1. 不需要根证书，直接生成用户证书
2. 首先需要自签署根证书，然后再签发服务器身份认证证书

### CA 根证书生成

```sh
# Generate CA private key
openssl genrsa -out ca.key 2048
# Generate csr (Certificate Signing Request)
# CN is not imperative?
openssl req -new -key ca.key -out ca.csr \
-subj "/C=CN/ST=Sichuan/L=Chengdu/O=RabbitHole, Inc./CN=RabbitHole Root CA"
# Generate self-signed certificate as CA
openssl x509 -req -days 365 -in ca.csr -signkey ca.key -out ca.crt
```

或者

```sh
# Generate CA private key
openssl genrsa -out ca.key 2048
# Generate self-signed certificate
openssl req -new -x509 -key ca.key -out ca.crt -days 365
```

- `genrsa` - generate an RSA private key
- `req` - PKCS#10 certificate request and certificate generating utility
- `x509` - Certificate display and sign utility
- `-days` 证书有效期

接下来继续使用`ca.crt`来签署服务器证书。

### 用户证书

**根证书只需生成一次，每个用户都需要使用根证书签署自己的证书。**

1. 生成私钥：`openssl genrsa -out mydomain.com.key 2048`

2. 检查私钥：`openssl rsa -in mydomain.com.key -noout -text`

3. 生成`csr`，注意`Common Name`指向服务器地址或域名，有互动、命令行和配置文件三种生成方式

   ```sh
   # 互动式
   openssl req -new -key mydomain.com.key -out mydomain.com.csr
   # 命令行
   openssl req -new -sha256 -key mydomain.com.key -out mydomain.com.csr \
   -subj "/C=CN/ST=Sichuan/L=Chengdu/O=Rabbit Inc./OU=Tech/CN=mydomain.com" \
   -addext "subjectAltName=DNS:mydomain.com"
   # 配置文件
   openssl req -new -out mydomain.com.csr -config oats.conf
   ```

4. 检验`csr`：`openssl req -in mydomain.com.csr -noout -text`

5. 由根证书签发用户证书

   ```sh
   openssl x509 -req -in mydomain.com.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out mydomain.com.crt -days 3650 -sha256
   ```

6. 检验用户证书：`openssl x509 -in mydomain.com.crt -text -noout`

Since Go version 1.15, the deprecated, legacy behavior of treating the CommonName field on X.509 certificates as a hostname when no **Subject Alternative Names** (SAN) are present is now disabled by default. 

This may affect existing SSL certificates and any secure connection relying on these certificates, including databases, clients (such as Docker), and applications.

A new valid certificate needs to be created to include the subjectAltName property, and should be added directly when creating an SSL self-signed certificate using openssl command, by specifying an `-addext` flag.

As a workaround, the behavior in which the CommonName field is being treated can be temporarily re-enabled by adding the value `x509ignoreCN=0` to the `GODEBUG` environment variable. For instance, by running the following command in the terminal:

```sh
export GODEBUG="x509ignoreCN=0"
```

**i.e.**

```sh
# 生成自签名根证书
openssl genrsa -out cert/ca.key 2048

# openssl req -new -key cert/ca.key -out cert/ca.csr \
# -subj "/C=CN/ST=Sichuan/L=Chengdu/O=RabbitHole, Inc./CN=RabbitHole Root CA"
#
# openssl x509 -req -days 365 -in cert/ca.csr -signkey cert/ca.key -out cert/ca.crt
openssl req -new -x509 -days 365 -key cert/ca.key -subj "/C=CN/ST=Sichuan/L=Chengdu/O=RabbitHole, Inc./CN=RabbitHole Root CA" -out cert/ca.crt

# 生成证书（由根证书进行签名）
# openssl genrsa -out cert/server.key 2048
#
# openssl req -new -sha256 -key cert/server.key \
# -subj "/C=CN/ST=Sichuan/L=Chengdu/O=Tesla, Inc./CN=rabbit.com" \
# -addext "subjectAltName=DNS:rabbit.com" -out cert/server.csr
#
# openssl x509 -req -in cert/server.csr -CA cert/ca.crt -CAkey cert/ca.key -CAcreateserial \
#  -days 365 -sha256 -out cert/server.crt

openssl req -newkey rsa:2048 -nodes -keyout cert/server.key -subj "/C=CN/ST=Sichuan/L=Chengdu/O=Tesla, Inc./CN=rabbit.com" -out cert/server.csr
# ServerName in client's tls config is rabbit.com
openssl x509 -req -extfile <(printf "subjectAltName=DNS:rabbit.com") -days 365 -in cert/server.csr -CA cert/ca.crt -CAkey cert/ca.key -CAcreateserial -out cert/server.crt

# client example certificate, also signed by root CA
openssl req -newkey rsa:2048 -nodes -keyout cert/client.key -subj "/C=CN/ST=Sichuan/L=Chengdu/O=Tesla, Inc./CN=example" -out cert/client.csr

openssl x509 -req -extfile <(printf "subjectAltName=DNS:example") -days 365 -in cert/client.csr -CA cert/ca.crt -CAkey cert/ca.key -CAcreateserial -out cert/client.crt

# PS: ca.sr 由 -CAcreateserial flag 生成，用于记录 CA 当前签名证书的序列号
```

### 流程

1. 服务端需要使用的文件是：server 私钥 (key)， server 证书
2. 客户端需要使用的文件是：CA file (根证书，用以验证 server 下发的证书)

## Demo in http

### server.go

```go
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

```

### client.go

```go
package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"algs4/config"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
)

var httpVersion = flag.Int("version", 1, "The http version HTTP/1 or HTTP/2")

func main() {
	flag.Parse()

	var t http.RoundTripper
	switch *httpVersion {
	case 1:
		t = transport1()
	case 2:
		t = transport2()
	default:
		logrus.Fatal("unknown http version")
	}

	client := &http.Client{Transport: t}

	res, err := client.Get("https://" + config.DefaultAddr)
	if err != nil {
		logrus.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.WithFields(logrus.Fields{
		"code": res.StatusCode,
		"body": string(body),
	}).Info("server respond")
}

func transport1() *http.Transport {
	return &http.Transport{
		// Original configurations from `http.DefaultTransport` variable.
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		// Set it false to enforce HTTP/1.
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       60 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,

		// Our custom configurations.
		ResponseHeaderTimeout: 10 * time.Second,
		DisableCompression:    true,
		// Set DisableKeepAlives to true when using HTTP/1 otherwise it will
		// cause error: "dial tcp [::1]:8090: socket: too many open files".
		DisableKeepAlives: false,
		TLSClientConfig:   tlsConfig(),
	}
}

func transport2() *http2.Transport {
	return &http2.Transport{
		TLSClientConfig:    tlsConfig(),
		DisableCompression: true,
		AllowHTTP:          false,
	}
}

func tlsConfig() *tls.Config {
	crt, err := ioutil.ReadFile("./cert/ca.crt")
	if err != nil {
		logrus.Fatal(err)
	}

	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM(crt)

	return &tls.Config{
		RootCAs:            rootCAs,
		InsecureSkipVerify: false,
		ServerName:         "localhost",
	}
}
```


