package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	graphql1 "github.com/slimaneakalie/graphql-golang-boilerplate/internal/pkg/graphql"
)

func (r *bookInfoResolver) Metadata(ctx context.Context, obj *graphql1.BookInfo) (*graphql1.BookMetadata, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *bookInfoResolver) Reviews(ctx context.Context, obj *graphql1.BookInfo) (*graphql1.BookReviews, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) RetrieveBookInfo(ctx context.Context, isbn string) (*graphql1.BookInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

// BookInfo returns graphql1.BookInfoResolver implementation.
func (r *Resolver) BookInfo() graphql1.BookInfoResolver { return &bookInfoResolver{r} }

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

type bookInfoResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
