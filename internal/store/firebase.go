package store

import (
	"context"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/storage"
	"github.com/faridEmilio/api_go_viajate_corporativo/internal/logs"
	"google.golang.org/api/option"
)

// FirebaseClient contiene la instancia del servicio Firebase
type FirebaseClient struct {
	StorageClient *storage.Client
	BucketName    string
}

// NewFirebaseService servicio cliente Firebase
func NewFirebaseClient() *FirebaseClient {
	ctx := context.Background()
	sa := option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENCIALS_FILE"))
	app, err := firebase.NewApp(ctx, nil, sa)

	if err != nil {
		logs.Error("error initializing Firebase app: " + err.Error())
		panic(err)
	}

	storageClient, err := app.Storage(ctx)
	if err != nil {
		logs.Error("error getting Firebase Storage client: " + err.Error())
		panic(err)
	}

	return &FirebaseClient{
		StorageClient: storageClient,
		BucketName:    os.Getenv("FIREBASE_BUCKET_NAME"),
	}
}
