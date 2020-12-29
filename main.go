package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	base            = "https://hacker-news.firebaseio.com/v0"
	storiesEndpoint = base + "/topstories.json"
	storyEndpoint   = base + "/item/%d.json"
	hackerNewsURL   = "https://news.ycombinator.com/item?id=%d"
)

type Story struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func main() {
	stories, err := GetStories(10)
	if err != nil {
		log.Fatal(err)
	}

	for i, story := range stories {
		URL := fmt.Sprintf(hackerNewsURL, story.ID)
		fmt.Printf("%02d - %s (%s)\n", i+1, story.Title, URL)
	}
}

// GetStories returns a slice of stories of specified length and an error.
func GetStories(n int) ([]Story, error) {
	resp, err := http.Get(storiesEndpoint)
	if err != nil {
		return nil, fmt.Errorf("GetStories: %v", err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("GetStories: %v", err)
	}

	IDs := []int{}
	err = json.Unmarshal(data, &IDs)
	if err != nil {
		return nil, fmt.Errorf("GetStories: %v", err)
	}

	stories := []Story{}
	for _, ID := range IDs[:n] {
		story, err := GetStory(ID)
		if err != nil {
			return nil, fmt.Errorf("GetStories: %v", err)
		}
		stories = append(stories, story)
	}

	return stories, nil
}

// GetStory returns a Story and an error given a story ID.
func GetStory(id int) (Story, error) {
	resp, err := http.Get(fmt.Sprintf(storyEndpoint, id))
	if err != nil {
		return Story{}, fmt.Errorf("GetStory: %v", err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Story{}, fmt.Errorf("GetStory: %v", err)
	}

	story := Story{}
	err = json.Unmarshal(data, &story)
	if err != nil {
		return Story{}, fmt.Errorf("GetStory: %v", err)
	}

	return story, nil
}
