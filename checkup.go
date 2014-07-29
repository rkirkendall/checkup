package checkup

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//Define Risk Factor
const (
	RFDepressiveFeelings      string = "Depressive Feelings"
	RFDepressionSymptoms      string = "Depression Symptoms"
	RFDrugAbuse               string = "Drug Abuse"
	RFPriorSuicideAttempts    string = "Prior Suicide Attempts"
	RFSuicideAroundIndividual string = "Suicide Around Individual"
	RFSuicideIdeation         string = "Suicide Ideation"
	RFSelfHarm                string = "Self-Harm"
	RFBullying                string = "Bullying"
	RFGunOwnership            string = "Gun Ownership"
	RFPsychologicalDisorders  string = "Psychological Disorders"
	RFFamilyViolenceDiscord   string = "Family Violence/Discord"
	RFImpulsivity             string = "Impulsivity"
)

type Phrase struct {
	Include    []string
	Exclude    []string
	RiskFactor string
}

var (
	phrases        []Phrase = []Phrase{}
	globalExcludes []string = []string{}
)

func Scan(tweet anaconda.Tweet, httpClient *http.Client) bool {
	var favorited, retweeted, isReply, verified, protected bool
	if tweet.FavoriteCount > 0 {
		favorited = true
	}
	if tweet.InReplyToStatusID > 0 {
		isReply = true
	}
	if tweet.RetweetCount > 0 {
		retweeted = true
	}
	if tweet.User.Verified {
		verified = true
	}
	if tweet.User.Protected {
		protected = true
	}

	//1. If the tweet has recieved interaction or if the author is verified, don't flag.
	if favorited || retweeted || isReply || verified || protected {
		return false
	}

	//2. Check for flag phrases
	containsFlagPhrase := CheckForPhrases(tweet.Text)
	return containsFlagPhrase
}
func CheckPreviousTweetSentiments(tweets []string, httpClient *http.Client) bool {
	if len(tweets) == 0 {
		return true
	}
	var sum int64 = 0
	var negativeHistory bool
	for _, pt := range tweets {
		//Returns 1 for positive, 0 for negative
		s := ClassifySentiment(pt, httpClient)
		sum += s

	}
	sentScore := float64(sum) / float64(len(tweets))
	if sentScore < 0.5 {
		negativeHistory = true
	} else {
		negativeHistory = false
	}

	return negativeHistory
}

func CheckForPhrases(tweetText string) bool {
	if len(phrases) == 0 {
		buildSelfHarmPhrases()
	}
	for _, phrase := range phrases {
		tweetText = strings.Join(cleanTweet(tweetText), " ")
		var includeMatch bool = true
		for _, ph := range phrase.Include {
			ph = strings.ToLower(ph)
			if strings.Contains(tweetText, ph) == false {
				includeMatch = false
				break
			}
		}

		var excludeMatch bool = true
		for _, ph := range phrase.Exclude {
			ph = strings.ToUpper(ph)
			if strings.Contains(tweetText, ph) {
				excludeMatch = false
				break
			}
		}

		if includeMatch && excludeMatch {
			return true
		}
	}
	return false
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
	return words
}
func ClassifySentiment(text string, httpClient *http.Client) int64 {
	urlEncoded := url.QueryEscape(text)
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	resp, reqErr := httpClient.Get("http://sent-classifier.herokuapp.com/classify/" + urlEncoded)
	if reqErr != nil {
		fmt.Println(reqErr)
	}
	sentString, bodyEr := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if bodyEr != nil {
		fmt.Println(bodyEr)
	}

	if string(sentString) == "0" {
		return int64(0)
	} else if string(sentString) == "1" {
		return int64(1)
	} else {
		fmt.Println("Wait what")
		fmt.Println(sentString)
		return int64(1)
	}
}

func buildSelfHarmPhrases() {
	phrases = []Phrase{}
	//globalExcludes := []string{"lol"}
	add([]string{"feel", "alone", "depressed"}, []string{})
	add([]string{"i", "feel", "helpless"}, []string{"girl", "without", "when"})
	add([]string{"i", "feel", "sad"}, []string{"episode", "lakers", "game", "sorry", "you", "when"})
	add([]string{"i", "feel", "empty"}, []string{"stomach", "phone", "hungry", "food"})
	add([]string{"sleeping", "a lot", "lately"}, []string{"haven't been"})
	add([]string{"i", "feel", "irritable"}, []string{"was"})
	add([]string{"depressed", "alchol", "irritable"}, []string{"Ronan"})
	add([]string{"sertaline"}, []string{"special class", "viagra", "study", "clinical", "http"})
	add([]string{"zoloft"}, []string{"toma", "para", "necesito",
		"gracioso", "desde", "palabra", "vida", "sabor", "aborto", "gusta"})
	add([]string{"prozac"}, []string{"toma", "para", "necesito",
		"gracioso", "desde", "palabra", "vida", "sabor", "aborto", "gusta"})
	add([]string{"pills", "depressed"}, []string{"http"})
	add([]string{"suicide", "once", "more"}, []string{"will", "live", "by"})
	add([]string{"suicide", "tried", "commit"}, []string{"dog", "fish", "cat", "who"})
	add([]string{"himself", "tried", "kill"}, []string{"dog", "fish", "cat"})
	add([]string{"herself", "tried", "kill"}, []string{"dog", "fish", "cat"})
	add([]string{"suicide", "attempted"}, []string{"dog", "fish", "cat"})
	add([]string{"killing", "myself"}, []string{})
	add([]string{"kill", "myself"}, []string{})
	add([]string{"cutting", "myself"}, []string{"shaving", "hair", "shave", "accidentally", "off"})
	add([]string{"being", "bullied"}, []string{""})
	add([]string{"feel", "bullied"}, []string{""})
	add([]string{"stop", "bullying", "me"}, []string{"#stop"})
	add([]string{"gun", "suicide"}, []string{""})
	add([]string{"diagnosed", "anorexia"}, []string{""})
	add([]string{"diagnosed", "ocd"}, []string{"http"})
	add([]string{"dad", "fight", "again"}, []string{"food"})
	add([]string{"parents", "fight", "again"}, []string{"food"})
}

func add(include []string, exclude []string) {
	phrases = append(phrases, Phrase{Include: include, Exclude: exclude})
}
