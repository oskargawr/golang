package models

type Post struct {
	ID       *int    `json:"id"`
	Date     *string `json:"date"`
	Country  *string `json:"country"`
	Area     *string `json:"area"`
	Activity *string `json:"activity"`
	Injury   *string `json:"injury"`
}
