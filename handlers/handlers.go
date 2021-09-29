package handlers

import "github.com/TrainLCD/CSMan/resolvers"

type Handlers struct {
	resolver *resolvers.Resolver
}

func NewHandlers() (*Handlers, error) {
	resolver, err := resolvers.NewResolver()
	if err != nil {
		return nil, err
	}
	return &Handlers{resolver}, nil
}
