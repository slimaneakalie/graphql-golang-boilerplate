## About
A simple graphql API in golang using gqlgen library to expose the following data about a book using an ISBN:
- Title
- Publishing date
- Number of pages
- Goodreads rating count
- Goodreads reviews count
- Goodreads average rating (on the scale of 5)

## Build
```
make build
```

## Run
```
make run
```

## Generate gqlgen models and resolvers
```
make generate
```
