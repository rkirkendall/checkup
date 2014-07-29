package checkup_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/r1cky1337/checkup"
)

var _ = Describe("Checkup", func() {
	It("Should classify the sentiment of a tweet", func() {
		negativeTweet := "I am mad. I hate everything."
		positiveTweet := "love life"
		negativeTweetScore := checkup.ClassifySentiment(negativeTweet, nil)
		positiveTweetScore := checkup.ClassifySentiment(positiveTweet, nil)
		Expect(positiveTweetScore).To(Equal(int64(1)))
		Expect(negativeTweetScore).To(Equal(int64(0)))
	})

	It("Should scan a tweet for flag phrases", func() {
		testTweet := "I want to kill myself"
		testTweet2 := "What a great day!"
		ret := checkup.CheckForPhrases(testTweet)
		ret2 := checkup.CheckForPhrases(testTweet2)
		Expect(ret).To(Equal(true))
		Expect(ret2).To(Equal(false))
	})

	It("Should assess a short history of tweets to have a predominately negative sentiment", func() {
		pos := []string{"beautiful day",
			"love life",
			"everything is perfect"}
		negs := []string{"Today has gone terribly wrong",
			"I never thought it could get this bad",
			"Where did I go wrong"}

		retNeg := checkup.CheckPreviousTweetSentiments(negs, nil)
		Expect(retNeg).To(Equal(true))

		retPos := checkup.CheckPreviousTweetSentiments(pos, nil)
		Expect(retPos).To(Equal(false))
	})
})
