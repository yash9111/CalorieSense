package models

import "time"

type Food struct {
	Name        string    `json:"name"`
	Calories    string    `json:"calories"`
	Protein     string    `json:"protein"`
	Carbs       string    `json:"carbs"`
	Fat         string    `json:"fat"`
	Fiber       string    `json:"fiber"`
	ImageUrl    string    `json:"image_url"`
	Ingredients []string  `json:"ingredients"`
	TimeStamp   time.Time `json:"timestamp"`
}
