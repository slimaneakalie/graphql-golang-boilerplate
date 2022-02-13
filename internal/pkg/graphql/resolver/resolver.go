package resolver

import "github.com/slimaneakalie/graphql-golang-boilerplate/internal/pkg/service/book/metadata"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Services struct {
	Metadata metadata.Service
}

type Resolver struct {
	Services *Services
}
