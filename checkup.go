package checkup

import "fmt"
import s "strings"

type Phrase struct {
	Include    []string
	Exclude    []string
	RiskFactor string
}

//Define Risk Factor
const RF_DEPRESSIVE_FEELINGS string = "Depressive Feelings"
const RF_DEPRESSION_SYMPTOMS string = "Depression Symptoms"
const RF_DRUG_ABUSE string = "Drug Abuse"
const RF_PRIOR_SUICIDE_ATTEMPTS string = "Prior Suicide Attempts"
const RF_SUICIDE_AROUND_INDIVIDUAL string = "Suicide Around Individual"
const RF_SUICIDE_IDEATION string = "Suicide Ideation"
const RF_SELF_HARM string = "Self-Harm"
const RF_BULLYING string = "Bullying"
const RF_GUN_OWNERSHIP string = "Gun Ownership"
const RF_PSYCHOLOGICAL_DISORDERS string = "Psychological Disorders"
const RF_FAMILY_VIOLENCE_DISCORD string = "Family Violence/Discord"
const RF_IMPULSIVITY string = "Impulsivity"

var phrases []Phrase
var globalExcludes []string

func BuildSelfHarmPhrases() {
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

func TestContains() string {
	if contains("aa", "aaa") {
		return "compare works"
	} else {
		return "compare is broken"
	}
}

func TestPhrases() int {
	return len(phrases)
}

func Scan(tweet string) bool {
	for _, phrase := range phrases {
		//Make everything upper case
		tweet = s.ToUpper(tweet)
		var includeMatch bool = true
		for _, ph := range phrase.Include {
			ph = s.ToUpper(ph)
			//c.Infof(fmt.Sprintln(tweet, ":", ph))
			if contains(ph, tweet) == false {
				includeMatch = false
				break
			}
		}

		var excludeMatch bool = true
		for _, ph := range phrase.Exclude {
			ph = s.ToUpper(ph)
			if contains(ph, tweet) {
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

func contains(phrase string, text string) bool {
	return s.Contains(text, phrase)
}

func add(include []string, exclude []string) {
	phrases = append(phrases, Phrase{Include: include, Exclude: exclude})
}
