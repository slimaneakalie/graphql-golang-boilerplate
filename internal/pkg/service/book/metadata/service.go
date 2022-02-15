package metadata

import (
	"context"
	"github.com/carlmjohnson/requests"
	"net/http"
)

const (
	isbnQueryParamKey         = "q"
	isbnQueryParamValuePrefix = "isbn:"
)

func NewService(metadataApiURL string) Service {
	return &defaultService{
		metadataApiURL: metadataApiURL,
	}
}

func (service *defaultService) RetrieveBookMetadata(isbn string) (*BookMetadata, error) {
	apiResponse, err := service.fetchBookMetadataFromExternalAPI(isbn)
	if err != nil {
		return nil, err
	}

	return mapAPIResponseToBookMetadata(apiResponse), nil
}

func (service *defaultService) fetchBookMetadataFromExternalAPI(isbn string) (*APIResponse, error) {
	isbnQueryParamValue := isbnQueryParamValuePrefix + isbn
	var apiResponse APIResponse

	err := requests.
		URL(service.metadataApiURL).
		Method(http.MethodGet).
		Param(isbnQueryParamKey, isbnQueryParamValue).
		ToJSON(&apiResponse).
		Fetch(context.Background())

	if err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

func mapAPIResponseToBookMetadata(apiResponse *APIResponse) *BookMetadata {
	if len(apiResponse.Items) == 0 {
		return &BookMetadata{}
	}

	volumeInfo := apiResponse.Items[0].VolumeInfo
	return &BookMetadata{
		Title:          volumeInfo.Title,
		PublishingDate: volumeInfo.PublishedDate,
		NumberOfPages:  volumeInfo.PageCount,
	}
}
