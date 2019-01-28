package news

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const categoryApiUrlPattern = "https://newsapi.org/v1/sources?language=en&category=%s"
const topicApiUrlPattern = "https://newsapi.org/v1/articles?source=%s&apiKey=2f6050f30a744e9ca308171a52e7df69"

type source struct {
	ID string `json:"id"`
}

type sourcesAPI struct {
	Sources []source `json:"sources"`
}

type TopicsAPI struct {
	Articles []Topic `json:"articles"`
}

func getSources(category string) []source {
	body := getData(sourceURL(category))

	var sourcesAPI sourcesAPI
	_ = json.Unmarshal(body, &sourcesAPI)

	return sourcesAPI.Sources
}

func getTopics(sources []source) []Topic {
	var topics []Topic

	for _, source := range sources {
		body := getData(topicURl(source.ID))

		var topicsAPI TopicsAPI
		_ = json.Unmarshal(body, &topicsAPI)

		topics = append(topics, topicsAPI.Articles...)
	}

	return topics
}

func topicURl(id string) string {
	return fmt.Sprintf(topicApiUrlPattern, id)
}

func sourceURL(category string) string {
	return fmt.Sprintf(categoryApiUrlPattern, category)
}

func getData(url string) []byte {
	result, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}

	return body
}
