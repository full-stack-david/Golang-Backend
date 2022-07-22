package firestore

import (
	"cloud.google.com/go/firestore"
	"firebase.google.com/go"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
)

func GetFirestoreClient() (*firestore.Client, context.Context, error) {

	ctx := context.Background()
	// Fetch the service account key JSON file contents
	log.Printf("## Using service account file %s", viper.GetString("FirestoreConfig.ServiceAccountKey"))
	opt := option.WithCredentialsFile(viper.GetString("FirestoreConfig.ServiceAccountKey"))

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Println("Error initializing app:", err)
		return nil, nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Println("Error initializing database client:", err)
		return nil, nil, err
	}

	return client, ctx, nil
}

func GetDropshipDllFirestoreClient() (*firestore.Client, context.Context, error) {

	ctx := context.Background()
	// Fetch the service account key JSON file contents
	log.Printf("## Using service account file %s", viper.GetString("FirestoreConfig.DropshipDllServiceAccountKey"))
	opt := option.WithCredentialsFile(viper.GetString("FirestoreConfig.DropshipDllServiceAccountKey"))

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Println("Error initializing app:", err)
		return nil, nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Println("Error initializing database client:", err)
		return nil, nil, err
	}

	return client, ctx, nil
}