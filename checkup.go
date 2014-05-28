package checkup

import (
	"fmt"
	"strings"
)

type Phrase struct {
	Include    []string
	Exclude    []string
	RiskFactor string
}

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

var (
	phrases        []Phrase
	globalExcludes []string
)

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
	add([]string{"zoloft"}, []string{"toma", "para", "necesito", "gracioso", "desde", "palabra", "vida", "sabor", "aborto", "gusta"})
	add([]string{"prozac"}, []string{"toma", "para", "necesito", "gracioso", "desde", "palabra", "vida", "sabor", "aborto", "gusta"})
	add([]string{"pills", "depressed"}, []string{"http"})
	add([]string{"suicide", "once", "more"}, []string{"will", "live", "by"})
	add([]string{"suicide", "tried", "commit"}, []string{"dog", "fish", "cat", "who"})
	add([]string{"himself", "tried", "kill"}, []string{"dog", "fish", "cat"})
	add([]string{"herself", "tried", "kill"}, []string{"dog", "fish", "cat"})
	add([]string{"suicide", "attempted"}, []string{"dog", "fish", "cat"})
	add([]string{"killing", "myself"}, []string{})
	add([]string{"cutting", "myself"}, []string{"shaving", "hair", "shave", "accidentally", "off"})
	add([]string{"being", "bullied"}, []string{""})
	add([]string{"feel", "bullied"}, []string{""})
	add([]string{"stop", "bullying", "me"}, []string{"#stop"})
	add([]string{"gun", "suicide"}, []string{""})
	add([]string{"diagnosed", "anorexia"}, []string{""})
	add([]string{"diagnosed", "ocd"}, []string{"http"})
	add([]string{"dad", "fight", "again"}, []string{"food"})
	add([]string{"parents", "fight", "again"}, []string{"food"})
	fmt.Println(phrases)
}

func Scan(tweet string) bool {
	if len(phrases) == 0 {
		buildSelfHarmPhrases()
	}

	for _, phrase := range phrases {
		//Make everything upper case
		tweet = strings.ToUpper(tweet)
		var includeMatch bool = true
		for _, ph := range phrase.Include {
			ph = strings.ToUpper(ph)
			//c.Infof(fmt.Sprintln(tweet, ":", ph))
			if strings.Contains(tweet, ph) {
				includeMatch = false
				break
			}
		}

		var excludeMatch bool = true
		for _, ph := range phrase.Exclude {
			ph = strings.ToUpper(ph)
			if strings.Contains(tweet, ph) {
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

func add(include []string, exclude []string) {
	phrases = append(phrases, Phrase{Include: include, Exclude: exclude})
}
