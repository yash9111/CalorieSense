package models

type UserProfileDetails struct {
	Name         string `json:"name"`
	Weight       string `json:"weight"`
	Height       string `json:"height"`
	Age          string `json:"age"`
	TargetWeight string `json:"target_weight" firestore:"target_weight"`
	Activity     string `json:"activity"`
	Goal         string `json:"goal"`
	Gender       string `json:"gender"`
}
