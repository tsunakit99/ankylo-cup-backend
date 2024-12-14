package auth

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func InitFirebaseApp(ctx context.Context, credFile string) *firebase.App {
    opt := option.WithCredentialsFile(credFile)
    app, err := firebase.NewApp(ctx, nil, opt)
    if err != nil {
        log.Fatalf("error initializing firebase app: %v", err)
    }
    return app
}
