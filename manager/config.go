package manager

import (
	"os"

	"github.com/Ripolak/reddit-snapshots/catcher"
)

var (
	clientID            = os.Getenv("CLIENT_ID")
	clientSecret        = os.Getenv("CLIENT_SECRET")
	username            = os.Getenv("USERNAME")
	password            = os.Getenv("PASSWORD")
	dbUrl               = os.Getenv("DB_URL")
	dbName              = os.Getenv("DB_NAME")
	snapshotsCollection = os.Getenv("SNAPSHOTS_COLLECTION")
	configCollection    = os.Getenv("CONFIG_COLLECTION")
	reddit              = catcher.RedditClient{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Username:     username,
		Password:     password,
	}
)
