checkup
=======
[![Build Status](https://drone.io/github.com/r1cky1337/checkup/status.png)](https://drone.io/github.com/r1cky1337/checkup/latest)

checkup is a package for actively monitoring Twitter feeds to detect suicidal language.

Example
-------
```Go
for _, tweet := range tweets {
  //Scan tweet for flag phrases
	
	flag := checkup.Scan(tweet, nil)

	//If there is a flag, check history of tweets for negative sentiments
	
	if flag {
		//Get last tweets from that user
		prevTweetStrings := []string{}
		for _, pt := range previousTweets {
			prevTweetStrings = append(prevTweetStrings, pt.Text)
		}
		negativeHistory := checkup.CheckPreviousTweetSentiments(prevTweetStrings, httpClient)
		if negativeHistory {
			fmt.Println("Flagged Tweet!")
			//Send email alert
		}
	}
}
```

Dependencies
------------
checkup uses [the anaconda package](https://github.com/ChimeraCoder/anaconda) as a Twitter client
Research
--------
checkup was inspired by [BYU's research](http://news.byu.edu/archive13-oct-suicide.aspx)
