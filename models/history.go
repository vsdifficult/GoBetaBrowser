package models

var History []string

func AddToHistory(url string) {
    History = append(History, url)
}

func GetHistory() []string {
    return History
}