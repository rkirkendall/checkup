package checkup

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"testing"
)

// testing data link:
// https://docs.google.com/spreadsheets/d/1rY_evVtZO-C64EFUeQpr3jawVOFxVbMnfkBeedWTZGw/pubhtml
// last tweets: http://pastebin.com/QpV9RThH

type testPhrase struct {
	Text string
	Flag bool
}

func TestTweets(t *testing.T) {
	phrases, err := load()
	if err != nil {
		panic(err)
	}

	for _, p := range phrases {
		if p.Flag != Scan(p.Text) {
			t.Errorf("%v should have been flagged: %v", p.Text, p.Flag)
		}
	}
}

func load() ([]*testPhrase, error) {
	pairs := []*testPhrase{}
	file, err := os.Open("phrases.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return []*testPhrase{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil || len(record) != 2 {
			fmt.Println("Error:", err)
			return []*testPhrase{}, err
		}

		if record[0] != "Text" {
			flag := false
			if record[1] == "Yes" {
				flag = true
			}

			pairs = append(pairs, &testPhrase{record[0], flag})
		}
	}
	return pairs, nil
}
