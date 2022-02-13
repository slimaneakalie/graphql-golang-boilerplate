package helper

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
)

func GetStringRequiredArgumentFromContext(ctx context.Context, argumentName string) (arg string) {
	argument, _ := getArgumentFromContext(ctx, argumentName)
	return argument.(string)
}

func getArgumentFromContext(ctx context.Context, argumentName string) (arg interface{}, exists bool) {
	parentContext := graphql.GetFieldContext(ctx)
	exists = false

	for parentContext != nil && !exists {
		arg, exists = parentContext.Args[argumentName]
		parentContext = parentContext.Parent
	}

	return arg, exists
}
