package resolvers

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

type Resolver struct {
	app       *firebase.App
	firestore *firestore.Client
	ctx       context.Context
}

func NewResolver() (*Resolver, error) {
	ctx := context.Background()

	opt := option.WithCredentialsFile("service-account-file.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	fs, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return &Resolver{
		app,
		fs,
		ctx,
	}, nil
}
