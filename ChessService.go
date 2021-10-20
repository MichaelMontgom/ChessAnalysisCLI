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

func getOpeningMovePreference(username string) ([]string, []string) {

	t := time.Now()
	games := getMonthlyArchive(username, strconv.Itoa(int(t.Month())), strconv.Itoa(t.Year()))

	whiteRe := regexp.MustCompile("1\\. ...")
	blackRe := regexp.MustCompile("1\\.\\.\\. ...")

	white := "White .." + username
	black := "Black .." + username

	s := strings.Split(games, "url")
	var whiteMoves, blackMoves []string

	for _, element := range s {

		isWhite, err := regexp.Match(white, []byte(element))
		if err != nil {
			log.Fatal(err)
		}

		isBlack, err2 := regexp.Match(black, []byte(element))
		if err2 != nil {
			log.Fatal(err2)
		}

		if isWhite {
			move := fmt.Sprintf("%s\n", whiteRe.Find([]byte(element)))
			whiteMoves = append(whiteMoves, move)
			continue
		}

		if isBlack {
			move := fmt.Sprintf("%s %s\n", whiteRe.Find([]byte(element)), blackRe.Find([]byte(element)))
			blackMoves = append(blackMoves, move)
			continue
		}

	}

	return whiteMoves, blackMoves
}

func getMonths() ([]int, []int) {

	month := int(time.Now().Month())
	year := time.Now().Year()
	var months, years []int

	// month := int(t.Month())

	for i := 0; i < 11; i++ {

		months = append(months, month)
		years = append(years, year)

		if month == 1 {
			month = 12
			year = year - 1
			continue
		}

		month = month - 1

	}

	return months, years
}
