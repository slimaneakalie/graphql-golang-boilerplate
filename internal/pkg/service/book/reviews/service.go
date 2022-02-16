package reviews

import (
	"context"
	"github.com/carlmjohnson/requests"
	"net/http"
	"strconv"
)

const (
	isbnQueryParamKey = "isbns"
)

func NewService(reviewsApiURL string) Service {
	return &defaultService{
		reviewsApiURL: reviewsApiURL,
	}
}

func (service *defaultService) RetrieveBookReviews(isbn string) (*BookReviews, error) {
	response, err := service.retrieveBookReviewsFromExternalAPI(isbn)
	if err != nil {
		return nil, err
	}

	return mapAPIResponseToBookReviews(response), nil
}

func (service *defaultService) retrieveBookReviewsFromExternalAPI(isbn string) (*apiResponse, error) {
	var response apiResponse

	err := requests.
		URL(service.reviewsApiURL).
		Method(http.MethodGet).
		Param(isbnQueryParamKey, isbn).
		ToJSON(&response).
		Fetch(context.Background())

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func mapAPIResponseToBookReviews(response *apiResponse) *BookReviews {
	if len(response.Books) == 0 {
		return &BookReviews{}
	}

	book := response.Books[0]
	averageRating, _ := strconv.ParseFloat(book.AverageRating, 64)

	return &BookReviews{
		NumberOfRatings: book.RatingsCount,
		NumberOfReviews: book.TextReviewsCount,
		AverageRating:   &averageRating,
	}
}
