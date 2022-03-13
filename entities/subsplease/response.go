package subsplease

type Response struct {
	Tz       string `json:"tz"`
	Schedule []Item `json:"schedule"`
}

type Item struct {
	Title    string `json:"title"`
	Page     string `json:"page"`
	ImageURL string `json:"image_url"`
	Time     string `json:"time"`
	Aired    bool   `json:"aired"`
}
