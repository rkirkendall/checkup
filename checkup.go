package checkup

import "strings"

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

var (
	phrases        []Phrase = []Phrase{}
	globalExcludes []string = []string{}
)

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
			if strings.Contains(tweet, ph) == false {
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
}

func add(include []string, exclude []string) {
	phrases = append(phrases, Phrase{Include: include, Exclude: exclude})
}
