package metadata

type Service interface {
	RetrieveBookMetadata(isbn string) (*BookMetadata, error)
}

type BookMetadata struct {
	Title          *string
	PublishingDate *string
	NumberOfPages  *int
}

type apiResponse struct {
	Items []*apiResponseItem `json:"items,omitempty"`
}

type apiResponseItem struct {
	VolumeInfo struct {
		Title         *string `json:"title,omitempty"`
		PublishedDate *string `json:"publishedDate,omitempty"`
		PageCount     *int    `json:"pageCount,omitempty"`
	} `json:"volumeInfo,omitempty"`
}

type defaultService struct {
	metadataApiURL string
}
