package entity

type Event struct {
	Id          string `json:"id"`
	IdUser      string `json:"id_user"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DateTime    string `json:"dateTime"`
	Duration    string `json:"duration"`
	Notes       string `json:"notes"`
}
