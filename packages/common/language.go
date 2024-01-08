package common

import (
	"fmt"
	"strings"
)

const (
	Beginner     Level = "Beginner"
	Intermediate Level = "Intermediate"
	Advanced     Level = "Advanced"
)

var (
	LangVi = Lang{Code: "vi", Name: "Vietnamese"}
	LangEn = Lang{Code: "en", Name: "English"}
)

type (
	Level string
	Lang  struct {
		Code string `json:"languageCode"` // ISO-[639-1] Code of language based
		Name string `json:"languageName"` // English name of lanague
	}
	Language struct {
		Lang  Lang  `json:"language"`
		Level Level `json:"languageLevel"`
	}
)

func (l Level) String() string {
	return string(l)
}

func (l Lang) String() string {
	return l.Name
}

func (c Language) String() string {
	str := strings.Builder{}
	str.WriteString("[")
	str.WriteString(fmt.Sprintf("language: %s, ", c.Lang))
	str.WriteString(fmt.Sprintf("level: %s", c.Level))
	str.WriteString("]")
	return str.String()
}
