package models

type ScheduleResponse struct {
	Tz       string         `json:"tz"`
	Schedule []ScheduleItem `json:"schedule"`
}

type ScheduleItem struct {
	Title    string `json:"title"`
	Page     string `json:"page"`
	ImageURL string `json:"image_url"`
	Time     string `json:"time"`
	Aired    bool   `json:"aired"`
}
