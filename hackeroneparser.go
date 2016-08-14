package bbcrawler

import "encoding/json"

type hackerOneBounty struct {
	BugCount      int `json:"bug_count"`
	MinimumBounty int `json:"minimum_bounty"`
}

type hackerOneRecord struct {
	Id                int             `json:"id"`
	Url               string          `json:"url"`
	Name              string          `json:"name"`
	Meta              hackerOneBounty `json:"meta"`
	About             string          `json:"about"`
	StrippedPolicy    string          `json:"stripped_policy"`
	Handle            string          `json:"handle"`
	ProfilePicture    string          `json:"profile_picture"`
	InternetBugBounty bool            `json:"internet_bug_bounty"`
}

type HackerOneResponse struct {
	Limit   int `json:"limit"`
	Total   int `json:"total"`
	Results []hackerOneRecord
}

type hackerOneParser struct{}

func (h hackerOneParser) Read(data []byte) (interface{}, error) {
	var jsonResponse HackerOneResponse
	err := json.Unmarshal(data, &jsonResponse)
	return jsonResponse, err
}

var HackerOneParser = hackerOneParser{}
