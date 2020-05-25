package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/jefridev/animenotifier/pkg/notifierlib"
	"google.golang.org/api/option"
)

const (
	googleApplicationCredentials = "[PATH/YOUR_FILE_NAME.json]"
	projectID = "[YOUR_PROJECT_ID]"
	databaseURL = "[YOUR_FIREBASE_URL]"
)

func main(){

	l := log.New(os.Stdout, "notifier: ", log.Ldate)
	ctx := context.Background()
	opt := option.WithCredentialsFile(googleApplicationCredentials)
	config := &firebase.Config{
		ProjectID: projectID,
		DatabaseURL: databaseURL,
	}

	app, err := firebase.NewApp(ctx, config, opt)
	dbclient, err := app.Database(ctx)
	if err != nil{
		l.Fatalf("There was an error connecting to remote database. %v", err)
	}

	var animeRepository notifierlib.AnimeRepository = notifierlib.NewAnimeRepository(ctx, dbclient)
	var animeService notifierlib.AnimeService = notifierlib.NewAnimeService(l, animeRepository)
	err = animeService.LoadAnimeToFirebase() 
	if err != nil{
		l.Fatalf("It was not possible to save extracted animes %v", err)
	}

	l.Println("Process completed")
}