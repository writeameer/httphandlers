package handlers

import (
	"crypto/tls"
	"net/http"
	"net/http/httputil"
)

// ReverseProxyHandler proxies requests to service-now.com
func ReverseProxyHandler(originHost string, next ...http.Handler) http.Handler {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	reverseProxy := &httputil.ReverseProxy{Director: func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", originHost)
		req.Host = originHost
		req.URL.Scheme = "https"
		req.URL.Host = originHost
	}}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reverseProxy.ServeHTTP(w, r)
		if next != nil {
			next[0].ServeHTTP(w, r)
		}
	})
}
