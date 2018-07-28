package fighter

type Fighter struct {
	ID              int     `json:"id"`
	FirstName       string  `json:"first_name"`
	LastName        string  `json:"last_name"`
	Nickname        string  `json:"nickname"`
	DOB             string  `json:"dob,omitempty"`
	Weight          int     `json:"Weight"`
	Height          string  `json:"height_ft"`
	Status          string  `json:"fighter_status"`
	WeightClass     string  `json:"weight_class"`
	Titleholder     bool    `json:"title_holder"`
	Rank            int     `json:"rank"`
	CountryResiding string  `json:"country_residing"`
	CityResiding    string  `json:"city_residing"`
	CountryOfOrigin string  `json:"country_of_origin"`
	CityOfOrigin    string  `json:"city_of_origin"`
	KoTkoWins       int     `json:"ko_tko_wins"`
	SubWins         int     `json:"submission_wins"`
	Decisionwins    int     `json:"decision_wins"`
	Fights          []Fight `json:"fights"`
}

type Fight struct {
	ID                    string   `json:"statid"`
	KOOfTheNight          int      `json:"KOOfTheNight"`
	SubmissionOfTheNight  int      `json:"SubmissionOfTheNight"`
	PerformanceOfTheNight int      `json:"PerformanceOfTheNight"`
	WeighIn               string   `json:"WeighIn"`
	Result                Result   `json:"Result"`
	Opponent              Opponent `json:"Opponent"`
}

type Result struct {
	ID              string `json:"OutcomeID"`
	Outcome         string `json:"Outcome"`
	EndingRound     string `json:"EndingRound"`
	EndingTime      string `json:"EndingTime"`
	Method          string `json:"Method"`
	FightOfTheNight string `json:"FightOfTheNight"`
	Submission      string `json:"Submission"`
	EndStrike       string `json:"EndStrike"`
}

type Opponent struct {
	StatID                string `json:"statid"`
	FirstName             string `json:"FirstName"`
	LastName              string `json:"LastName"`
	NickName              string `json:"NickName"`
	Weight                string `json:"Weight"`
	WeighIn               string `json:"WeighIn"`
	PerformanceOfTheNight string `json:"PerformanceOfTheNight"`
	SubmissionOfTheNight  string `json:"SubmissionOfTheNight"`
	KOOfTheNight          string `json:"KOOfTheNight"`
}
