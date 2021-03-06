package models

import "github.com/go-bongo/bongo"

// Recipe model
type Recipe struct {
	bongo.DocumentBase `bson:",inline"`
	Title              string
	CategoryID         string
	Description        string
	CookingTime        int
}
