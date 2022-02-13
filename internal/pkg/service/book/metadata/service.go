package metadata

import (
	"net/http"
	"net/url"
	"time"
)

const (
	httpTimeoutDuration = time.Second * 3
	isbnQueryParamKey   = "q"
)

func NewService(metadataApiURL string) Service {
	return &defaultService{
		httpClient:     &http.Client{Timeout: httpTimeoutDuration},
		metadataApiURL: metadataApiURL,
	}
}

func (service *defaultService) RetrieveBookMetadata(isbn string) (*BookMetadata, error) {
	urlValues := url.Values{}
	isbnQueryParamValue := "isbn:" + isbn
	urlValues.Set(isbnQueryParamKey, isbnQueryParamValue)
	return nil, nil
}
