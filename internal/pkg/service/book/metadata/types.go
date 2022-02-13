package metadata

import "net/http"

type Service interface {
	RetrieveBookMetadata(isbn string) (*BookMetadata, error)
}

type BookMetadata struct {
	Title          *string
	PublishingDate *string
	NumberOfPages  *int
}

type defaultService struct {
	httpClient     *http.Client
	metadataApiURL string
}
