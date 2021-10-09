package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func getPlayerProfile(username string) string {

	var url string = "https://api.chess.com/pub/player/" + username
	resp, err := http.Get(strings.TrimSpace(url))

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(body)

}

func getPlayerStats(username string) string {

	var url string = "https://api.chess.com/pub/player/" + username + "/stats"
	resp, err := http.Get(strings.TrimSpace(url))

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

func isPlayerOnline(username string) string {

	var url string = "https://api.chess.com/pub/player/" + username + "/is-online"
	resp, err := http.Get(strings.TrimSpace(url))

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(body)

}

func getCurrentDailyGames(username string) string {

	var url string = "https://api.chess.com/pub/player/" + username + "/games"
	resp, err := http.Get(strings.TrimSpace(url))

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

func getMonthlyArchive(username string, month string, year string) string {

	var url string = "https://api.chess.com/pub/player/" + username + "/games/" + year + "/" + month
	resp, err := http.Get(strings.TrimSpace(url))

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(body)

}
