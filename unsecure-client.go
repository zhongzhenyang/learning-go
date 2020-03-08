package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//func main() {
//	c := &http.Client{
//		Transport: &http.Transport{
//			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
//		}}
//
//	if resp, e := c.Get("https://localhost"); e != nil {
//		log.Fatal("http.Client.Get: ", e)
//	} else {
//		defer resp.Body.Close()
//		io.Copy(os.Stdout, resp.Body)
//	}
//}

func main() {
	pair, e := tls.LoadX509KeyPair("/home/zzy/test-tls/client/client.crt", "/home/zzy/test-tls/client/client.key")
	if e != nil {
		log.Fatal("LoadX509KeyPair:", e)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      loadClientCA("/home/zzy/test-tls/ca-key.pem"),
				Certificates: []tls.Certificate{pair},
			},
		}}

	resp, e := client.Get("https://localhost")
	if e != nil {
		log.Fatal("http.Client.Get: ", e)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

func loadClientCA(caFile string) *x509.CertPool {
	pool := x509.NewCertPool()

	if ca, e := ioutil.ReadFile(caFile); e != nil {
		log.Fatal("ReadFile: ", e)
	} else {
		pool.AppendCertsFromPEM(ca)
	}
	return pool
}
