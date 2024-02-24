package firebase

import (
    "context"
    "firebase.google.com/go/v4"
    "google.golang.org/api/option"
	"SC2024/pkg/config"
    "log"
	"os"
)

var app *firebase.App

func InitializeFirebaseApp() {
	currentDirectory, err := os.Getwd()

	log.Println(currentDirectory + config.Get("FIREBASE_CREDENTIALS_FILE"))
    opt := option.WithCredentialsFile(currentDirectory + config.Get("FIREBASE_CREDENTIALS_FILE"))
	
    firebaseApp, err := firebase.NewApp(context.Background(), nil, opt)
    if err != nil {
        log.Fatalf("error initializing Firebase app: %v\n", err)
    }
	
    app = firebaseApp
}

func GetFirebaseApp() *firebase.App {
    return app
}