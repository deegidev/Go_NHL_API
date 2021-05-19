package nhlApi

//copy JSON toGO
//delete copyright
//change type to Team
//delete `json:"teams"` at end of struct and all the Venue structs
type Team struct {
	Teams []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Link  string `json:"link"`
		Venue struct {
			Name     string `json:"name"`
			Link     string `json:"link"`
			City     string `json:"city"`
			Timezone struct {
				ID     string `json:"id"`
				Offset int    `json:"offset"`
				Tz     string `json:"tz"`
			} `json:"timeZone"`
		} `json:"venue,omitempty"`
		Abbreviation    string `json:"abbreviation"`
		Teamname        string `json:"teamName"`
		Locationname    string `json:"locationName"`
		Firstyearofplay string `json:"firstYearOfPlay,omitempty"`
		Division        struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"division,omitempty"`
		Conference struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"conference,omitempty"`
		Franchise struct {
			Franchiseid int    `json:"franchiseId"`
			Teamname    string `json:"teamName"`
			Link        string `json:"link"`
		} `json:"franchise"`
		Shortname       string `json:"shortName,omitempty"`
		Officialsiteurl string `json:"officialSiteUrl"`
		Franchiseid     int    `json:"franchiseId"`
		Active          bool   `json:"active"`
	}
}
