package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
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

func getMonthlyArchive(username string, year string, month string) string {

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

func getOpeningMovePreference(username string, year int, month int) [][]string {

	var games string

	if month < 10 {
		games = getMonthlyArchive(username, strconv.Itoa(year), "0"+strconv.Itoa(month))
	} else {
		games = getMonthlyArchive(username, strconv.Itoa(year), strconv.Itoa(month))
	}

	// fmt.Println(games)

	whiteRe := regexp.MustCompile("1\\. ...")
	blackRe := regexp.MustCompile("1\\.\\.\\. ...")

	white := "White .." + username
	black := "Black .." + username

	s := strings.Split(games, "url")
	var whiteMoves, blackMoves []string
	var moves [][]string

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
	moves = append(moves, whiteMoves)
	moves = append(moves, blackMoves)

	return moves
}

func getMonths() ([]int, []int) {

	month := int(time.Now().Month())
	year := time.Now().Year()

	var months, years []int

	for i := 0; i < 12; i++ {

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

func processOpeningMove(username string) [][][]string {

	months, years := getMonths()
	var moves [][][]string

	var wg sync.WaitGroup

	// Adding 12 to the wait group for 12 months woth of threads
	wg.Add(12)

	// Doing research on ways to make this cleaner
	out1 := make(chan [][]string)
	out2 := make(chan [][]string)
	out3 := make(chan [][]string)
	out4 := make(chan [][]string)
	out5 := make(chan [][]string)
	out6 := make(chan [][]string)
	out7 := make(chan [][]string)
	out8 := make(chan [][]string)
	out9 := make(chan [][]string)
	out10 := make(chan [][]string)
	out11 := make(chan [][]string)
	out12 := make(chan [][]string)

	go func() {
		defer wg.Done()

		out1 <- getOpeningMovePreference(username, years[0], months[0])

	}()
	go func() {
		defer wg.Done()

		out2 <- getOpeningMovePreference(username, years[1], months[1])

	}()
	go func() {
		defer wg.Done()

		out3 <- getOpeningMovePreference(username, years[2], months[2])

	}()
	go func() {
		defer wg.Done()

		out4 <- getOpeningMovePreference(username, years[3], months[3])

	}()
	go func() {
		defer wg.Done()

		out5 <- getOpeningMovePreference(username, years[4], months[4])

	}()
	go func() {
		defer wg.Done()

		out6 <- getOpeningMovePreference(username, years[5], months[5])

	}()
	go func() {
		defer wg.Done()

		out7 <- getOpeningMovePreference(username, years[6], months[6])

	}()
	go func() {
		defer wg.Done()

		out8 <- getOpeningMovePreference(username, years[7], months[7])

	}()
	go func() {
		defer wg.Done()

		out9 <- getOpeningMovePreference(username, years[8], months[8])

	}()
	go func() {
		defer wg.Done()

		out10 <- getOpeningMovePreference(username, years[9], months[9])

	}()
	go func() {
		defer wg.Done()

		out11 <- getOpeningMovePreference(username, years[10], months[10])

	}()
	go func() {
		defer wg.Done()

		out12 <- getOpeningMovePreference(username, years[11], months[11])

	}()

	counter := 0

	for {

		select {
		case msg := <-out1:
			// fmt.Println(msg, "func 1")
			moves = append(moves, msg)
			counter++

		case msg := <-out2:
			// fmt.Println(msg, "func 2")
			moves = append(moves, msg)
			counter++
		case msg := <-out3:
			// fmt.Println(msg, "func 3")
			moves = append(moves, msg)
			counter++
		case msg := <-out4:
			// fmt.Println(msg, "func 4")
			moves = append(moves, msg)
			counter++
		case msg := <-out5:
			// fmt.Println(msg, "func 5")
			moves = append(moves, msg)
			counter++
		case msg := <-out6:
			// fmt.Println(msg, "func 6")
			moves = append(moves, msg)
			counter++
		case msg := <-out7:
			// fmt.Println(msg, "func 7")
			moves = append(moves, msg)
			counter++
		case msg := <-out8:
			// fmt.Println(msg, "func 8")
			moves = append(moves, msg)
			counter++
		case msg := <-out9:
			// fmt.Println(msg, "func 9")
			moves = append(moves, msg)
			counter++
		case msg := <-out10:
			// fmt.Println(msg, "func 10")
			moves = append(moves, msg)
			counter++
		case msg := <-out11:
			// fmt.Println(msg, "func 11")
			moves = append(moves, msg)
			counter++
		case msg := <-out12:
			// fmt.Println(msg, "func 12")
			moves = append(moves, msg)
			counter++
		}

		if counter >= 11 {
			break
		}

	}
	return moves

}
