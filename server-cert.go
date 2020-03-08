package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//func main() {
//	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
//		io.WriteString(w, "hello, world!\n")
//	})
//	if e := http.ListenAndServeTLS(":443", "/home/zzy/test-tls/server-cert.pem", "/home/zzy/test-tls/server-key.pem", nil); e != nil {
//		log.Fatal("ListenAndServe: ", e)
//	}
//
//}

func main() {
	s := &http.Server{
		Addr: ":443",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello World!\n")
		}),
		TLSConfig: &tls.Config{
			ClientCAs:  loadCA("/home/zzy/test-tls/ca-key.pem"),
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	e := s.ListenAndServeTLS("/home/zzy/test-tls/server2-cert.pem", "/home/zzy/test-tls/server-key.pem")
	if e != nil {
		log.Fatal("ListenAndServeTLS: ", e)
	}
}

func loadCA(caFile string) *x509.CertPool {
	pool := x509.NewCertPool()

	if ca, e := ioutil.ReadFile(caFile); e != nil {
		log.Fatal("ReadFile: ", e)
	} else {
		pool.AppendCertsFromPEM(ca)
	}
	return pool
}
