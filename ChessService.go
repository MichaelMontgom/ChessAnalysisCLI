package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
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

func getOpeningMovePreference(username string) string {

	t := time.Now()
	games := getMonthlyArchive(username, strconv.Itoa(int(t.Month())), strconv.Itoa(t.Year()))

	// re := regexp.MustCompile(`1\. ...`)

	whiteRe := regexp.MustCompile("1\\. ...")
	blackRe := regexp.MustCompile("1\\.\\.\\. ...")

	white := "White .." + username
	black := "Black .." + username

	s := strings.Split(games, "url")
	// var moves []string

	for _, element := range s {

		isWhite, err := regexp.Match(white, []byte(element))
		if err != nil {
			log.Fatal(err)
		}

		isBlack, err2 := regexp.Match(black, []byte(element))
		if err2 != nil {
			log.Fatal(err2)
		}

		fmt.Println(isBlack, isWhite)

		if isWhite {
			fmt.Println("User is playing white")
			fmt.Printf("%q\n", whiteRe.Find([]byte(element)))
			// fmt.Println(element)
			continue
		}

		if isBlack {
			fmt.Println("User is playing black")
			fmt.Printf("%q\n", blackRe.Find([]byte(element)))
			// fmt.Println(element)
			continue
		}

	}

	return ""
}
