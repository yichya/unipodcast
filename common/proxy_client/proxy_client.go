package proxy_client

import (
	"log"
	"net/http"
	"net/url"
)

func ProxyClient(proxyUrl string) *http.Client {
	if proxyUrl == "" {
		return http.DefaultClient
	}
	proxyUrlParsed, err := url.Parse(proxyUrl)
	if err != nil {
		log.Printf("parse proxy url %v failed %v", proxyUrl, err)
		return http.DefaultClient
	}
	return &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrlParsed)}}
}
