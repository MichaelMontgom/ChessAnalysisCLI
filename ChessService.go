package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"strings"
)

func getPlayerProfile(username string) map[string]interface{} {

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

	var profile map[string]interface{}

	err2 := json.Unmarshal([]byte(body), &profile)

	if err2 != nil {
		log.Fatal(err2)
	}

	log.Println(profile)

	return profile

}

func getPlayerStats(username string) map[string]interface{} {

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

	var games map[string]interface{}

	err2 := json.Unmarshal([]byte(body), &games)

	if err2 != nil {
		log.Fatal(err2)
	}

	log.Println(games)

	return games
}

func isPlayerOnline(username string) map[string]interface{} {

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

	var games map[string]interface{}

	err2 := json.Unmarshal([]byte(body), &games)

	if err2 != nil {
		log.Fatal(err2)
	}

	log.Println(games)

	return games

}

func getCurrentDailyGames(username string) map[string]interface{} {

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

	var games map[string]interface{}

	err2 := json.Unmarshal([]byte(body), &games)

	if err2 != nil {
		log.Fatal(err2)
	}

	log.Println(games)

	return games
}

func getMonthlyArchive(username string, month string, year string) map[string]interface{} {

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

	var games map[string]interface{}

	err2 := json.Unmarshal([]byte(body), &games)

	if err2 != nil {
		log.Fatal(err2)
	}

	log.Println(games)

	return games

}

func getOpeningMovePreference(username string) string {

	t := time.Now()
	games := getMonthlyArchive(username, strconv.Itoa(int(t.Month())), strconv.Itoa(t.Year()))

	for _, game := range games {
		log.Print(game)
	}

	return ""
}
