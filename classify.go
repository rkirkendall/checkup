package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	. "github.com/jbrukh/bayesian"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	Positive Class = "1"
	Negative Class = "0"
)

func main() {
	start := time.Now().Unix()
	train()
	end := time.Now().Unix()
	test()
	fmt.Println("Time: ", end-start, "s")
}

type E struct{}

var stopWords map[string]struct{}

//2% accuracy win!
func loadStopWords() map[string]struct{} {
	if stopWords == nil {
		file, err := os.Open("dataset.csv")
		if err != nil {
			fmt.Println("Error:", err)
			panic(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		stopWords = make(map[string]struct{})
		for scanner.Scan() {
			stopWords[scanner.Text()] = E{}
		}
	}
	return stopWords
}

func cleanTweet(tweet string) []string {
	r := strings.NewReplacer("!", " ",
		".", " ",
		",", " ",
		"&lt;", " ",
		"'", "",
		"&gt;", " ",
		"&amp", " ",
		"?", " ",
		"-", "",
	)

	tweet = r.Replace(tweet)
	tweet = strings.ToLower(tweet)
	words := strings.Fields(tweet)
	interestingWords := []string{}
	s := loadStopWords()
	for _, word := range words {
		_, prs := s["word"]
		if !prs {
			interestingWords = append(interestingWords, word)
		}
	}
	return interestingWords
}

func test() {
	twitSentClassifier, classErr := NewClassifierFromFile("twitSent.classifier")
	if classErr != nil {
		fmt.Println("Error:", classErr)
	}

	right := 0
	wrong := 0

	file, err := os.Open("dataset.csv")
	if err != nil {
		fmt.Println("Error:", err)
		//panic(err)
	}
	defer file.Close()
	i := 0
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if i > 1350000 {
			if err == io.EOF {
				break
			} else if err == nil {
				tweet := record[3]
				words := cleanTweet(tweet)

				_, likely, _ := twitSentClassifier.ProbScores(words)

				expected := record[1]
				if expected == strconv.FormatInt(int64(likely), 10) {
					right++
				} else {
					wrong++
				}
			}
		}
		i++
	}

	ratio := float64(right) / (float64(right) + float64(wrong))
	fmt.Println("Ratio ", ratio)
	fmt.Println("Right ", right)
	fmt.Println("Wrong ", wrong)
}

func train() {

	classifier := NewClassifier(Negative, Positive)

	file, err := os.Open("dataset.csv")
	if err != nil {
		fmt.Println("Error:", err)
		//panic(err)
	}
	defer file.Close()
	i := 0
	reader := csv.NewReader(file)
	for i < 1350000 {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err == nil {
			tweet := record[3]
			words := cleanTweet(tweet)

			if record[1] == "0" {
				classifier.Learn(words, Negative)
			} else if i != 0 {
				classifier.Learn(words, Positive)
			}
			i++
		}
	}

	fmt.Println("Rows: ", i)

	if fileErr := classifier.WriteToFile("twitSent.classifier"); fileErr != nil {
		fmt.Println("Errror: ", fileErr)
		panic(fileErr)
	}
}
