package go_util

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func HostReverseProxy(w http.ResponseWriter, req *http.Request, targetHost *TargetHost) {
	host := ""
	if targetHost.IsHttps {
		host = host + "https://"
	} else {
		host = host + "http://"
	}
	remote, err := url.Parse(host + targetHost.Host)
	if err != nil {
		Errorf("err:%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	if targetHost.IsHttps {
		tls, err := GetVerTLSConfig(targetHost.CAPath)
		if err != nil {
			Errorf("https crt error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var pTransport http.RoundTripper = &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*time.Duration(100000))
				if err != nil {
					return nil, err
				}
				return c, nil
			},
			ResponseHeaderTimeout: time.Second * time.Duration(10000),
			TLSClientConfig:       tls,
		}
		proxy.Transport = pTransport
	}
	proxy.ServeHTTP(w, req)
}

type TargetHost struct {
	Host    string
	IsHttps bool
	CAPath  string
}

var TlsConfig *tls.Config
func GetVerTLSConfig(CAPath string) (*tls.Config, error) {
	caData, err := ioutil.ReadFile(CAPath)
	if err != nil {
		Errorf("read wechat ca fail", err)
		return nil, err
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caData)

	TlsConfig = &tls.Config{
		RootCAs: pool,
	}
	return TlsConfig, nil
}