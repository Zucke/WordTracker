package data

//ScrapData has the data to make the scraping
type ScrapData struct {
	ToFind           string   `json:"to_find"`
	UseURLsWhiteList bool     `json:"use_urls_white_list"`
	URLsWhiteList    []string `json:"urls_white_list,omitempty"`
	UseRegex         bool     `json:"use_regex"`
	CurrendURL       string   `json:"currend_url"`
}
