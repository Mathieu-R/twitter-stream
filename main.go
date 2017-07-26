package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// Package clientcredentials implements the OAuth2.0 "client credentials" token flow,
// also known as the "two-legged OAuth 2.0".
// This should be used when the client is acting
// on its own behalf or when the client is the resource owner

var (
	track = flag.String("track", "neymar", "Tweets subject to track")
)

func main() {
	loadEnvFile()
	flag.Parse()

	config := &clientcredentials.Config{
		ClientID:     os.Getenv("CONSUMER_KEY"),
		ClientSecret: os.Getenv("CONSUMER_SECRET"),
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	params := &twitter.StreamFilterParams{
		Track:         []string{*track},
		StallWarnings: twitter.Bool(true),
	}

	httpClient := config.Client(oauth2.NoContext)

	client := twitter.NewClient(httpClient)
	trackTweets(client, params)

}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("error loading .env file, cannot track tweets...")
	}
}

func trackTweets(client *twitter.Client, params *twitter.StreamFilterParams) {
	stream, err := client.Streams.Filter(params)
	if err != nil {
		fmt.Printf("error with stream: %v", err)
	}

	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		fmt.Println(tweet.Text)
	}

	for message := range stream.Messages {
		demux.Handle(message)
	}

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	stream.Stop()
}
