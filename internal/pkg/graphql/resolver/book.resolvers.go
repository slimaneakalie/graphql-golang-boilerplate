package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/slimaneakalie/graphql-golang-boilerplate/internal/pkg/graphql/helper"

	graphql1 "github.com/slimaneakalie/graphql-golang-boilerplate/internal/pkg/graphql"
)

const (
	isbnArgumentName = "isbn"
)

func (r *bookInfoResolver) Metadata(ctx context.Context, obj *graphql1.BookInfo) (*graphql1.BookMetadata, error) {
	isbn := helper.RetrieveStringRequiredArgumentFromContext(ctx, isbnArgumentName)
	metadata, err := r.Services.Metadata.RetrieveBookMetadata(isbn)
	if err != nil {
		return nil, err
	}

	graphqlMetadata := &graphql1.BookMetadata{
		Title:          metadata.Title,
		PublishingDate: metadata.PublishingDate,
		NumberOfPages:  metadata.NumberOfPages,
	}

	return graphqlMetadata, nil
}

func (r *bookInfoResolver) Reviews(ctx context.Context, obj *graphql1.BookInfo) (*graphql1.BookReviews, error) {
	isbn := helper.RetrieveStringRequiredArgumentFromContext(ctx, isbnArgumentName)
	reviews, err := r.Services.Reviews.RetrieveBookReviews(isbn)
	if err != nil {
		return nil, err
	}

	graphqlReviews := &graphql1.BookReviews{
		NumberOfRatings: reviews.NumberOfRatings,
		NumberOfReviews: reviews.NumberOfReviews,
		AverageRating:   reviews.AverageRating,
	}

	return graphqlReviews, nil
}

func (r *queryResolver) RetrieveBookInfo(ctx context.Context, isbn string) (*graphql1.BookInfo, error) {
	return &graphql1.BookInfo{}, nil
}

// BookInfo returns graphql1.BookInfoResolver implementation.
func (r *Resolver) BookInfo() graphql1.BookInfoResolver { return &bookInfoResolver{r} }

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

type bookInfoResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
