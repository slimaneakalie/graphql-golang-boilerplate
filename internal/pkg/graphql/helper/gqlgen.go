package helper

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
)

func RetrieveStringRequiredArgumentFromContext(ctx context.Context, argumentName string) (arg string) {
	argument, _ := retrieveArgumentFromContext(ctx, argumentName)
	return argument.(string)
}

func retrieveArgumentFromContext(ctx context.Context, argumentName string) (arg interface{}, exists bool) {
	parentContext := graphql.GetFieldContext(ctx)
	exists = false

	for parentContext != nil && !exists {
		arg, exists = parentContext.Args[argumentName]
		parentContext = parentContext.Parent
	}

	return arg, exists
}
