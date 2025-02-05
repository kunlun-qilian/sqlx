package database

import (
	"github.com/kunlun-qilian/sqlx/v3"
)

var DBTest = sqlx.NewDatabase("test")

type Gender int

const (
	GenderMale Gender = iota + 1
	GenderFemale
)

func (g Gender) String() string {
	switch g {
	case GenderMale:
		return "male"
	case GenderFemale:
		return "female"
	}
	return ""
}
