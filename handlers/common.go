package handlers

var apiKey string
var baseURL string

func Init(key string, url string) {
	apiKey = key
	baseURL = url
}
