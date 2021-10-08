package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getPlayerProfile(username string) {

	var url string = "https://api.chess.com/pub/player/" + username
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(body))

}

func getPlayerStats(username string) {

	var url string = " https://api.chess.com/pub/player/" + username + "/stats"
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(body))
}

func isPlayerOnline(username string) {

	var url string = " https://api.chess.com/pub/player/" + username + "/is-online"
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(body))

}

func getCurrentDailyGames(username string) {
	var url string = " https://api.chess.com/pub/player/" + username + "/games"
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(body))
}
