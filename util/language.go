package util

type Language string

const LANG_ID Language = "id"
const LANG_EN Language = "en"

type Languages struct {
	ID string `json:"id"`
	EN string `json:"en"`
}
