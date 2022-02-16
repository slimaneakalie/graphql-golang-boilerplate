package reviews

type Service interface {
	RetrieveBookReviews(isbn string) (*BookReviews, error)
}

type BookReviews struct {
	NumberOfRatings *int     `json:"work_ratings_count,omitempty"`
	NumberOfReviews *int     `json:"work_text_reviews_count,omitempty"`
	AverageRating   *float64 `json:"average_rating,omitempty"`
}

type apiResponse struct {
	Books []*apiResponseBook `json:"books,omitempty"`
}

type apiResponseBook struct {
	RatingsCount     *int   `json:"work_ratings_count,omitempty"`
	TextReviewsCount *int   `json:"work_text_reviews_count,omitempty"`
	AverageRating    string `json:"average_rating,omitempty"`
}

type defaultService struct {
	reviewsApiURL string
}
