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
	response, err := service.retrieveBookMetadataFromExternalAPI(isbn)
	if err != nil {
		return nil, err
	}

	return mapAPIResponseToBookMetadata(response), nil
}

func (service *defaultService) retrieveBookMetadataFromExternalAPI(isbn string) (*apiResponse, error) {
	isbnQueryParamValue := isbnQueryParamValuePrefix + isbn
	var response apiResponse

	err := requests.
		URL(service.metadataApiURL).
		Method(http.MethodGet).
		Param(isbnQueryParamKey, isbnQueryParamValue).
		ToJSON(&response).
		Fetch(context.Background())

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func mapAPIResponseToBookMetadata(response *apiResponse) *BookMetadata {
	if len(response.Items) == 0 {
		return &BookMetadata{}
	}

	volumeInfo := response.Items[0].VolumeInfo
	return &BookMetadata{
		Title:          volumeInfo.Title,
		PublishingDate: volumeInfo.PublishedDate,
		NumberOfPages:  volumeInfo.PageCount,
	}
}
