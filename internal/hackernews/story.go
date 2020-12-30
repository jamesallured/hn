package hackernews

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	serviceBaseURL    = "https://hacker-news.firebaseio.com/v0"
	hackerNewsBaseURL = "https://news.ycombinator.com"
)

var sorts = map[string]string{
	"best": "/beststories.json",
	"new":  "/newstories.json",
	"top":  "/topstories.json",
}

func getItemURL(id int) string {
	return fmt.Sprintf(serviceBaseURL+"/item/%d.json", id)
}

// GetThreadURL returns a URL for the Hacker News thread of a story.
func GetThreadURL(id int) string {
	return fmt.Sprintf(hackerNewsBaseURL+"/item?id=%d", id)
}

func getStoriesURL(sort string) string {
	URI, ok := sorts[sort]
	if !ok {
		fmt.Println("[!] Specified sort not supported, reverting to default (top)")
		return "top"
	}

	return fmt.Sprint(serviceBaseURL + URI)
}

// Story represents a Hacker News story with an ID and a title.
type Story struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// GetStories returns a slice of stories of specified length and an error.
func GetStories(n int, sort string) ([]Story, error) {
	resp, err := http.Get(getStoriesURL(sort))
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	IDs := []int{}
	err = json.Unmarshal(data, &IDs)
	if err != nil {
		return nil, err
	}

	stories := []Story{}
	for _, ID := range IDs[:n] {
		story, err := GetStory(ID)
		if err != nil {
			return nil, err
		}
		stories = append(stories, story)
	}

	return stories, nil
}

// GetStory returns a Story and an error given a story ID.
func GetStory(id int) (Story, error) {
	resp, err := http.Get(getItemURL(id))
	if err != nil {
		return Story{}, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Story{}, err
	}

	story := Story{}
	err = json.Unmarshal(data, &story)
	if err != nil {
		return Story{}, err
	}

	return story, nil
}
