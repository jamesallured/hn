package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jamesallured/hn/internal/hackernews"
)

func main() {
	sort := flag.String("sort", "top", "sort to apply to stories")
	limit := flag.Int("n", 10, "number of stories to return (max 500)")
	flag.Parse()

	stories, err := hackernews.GetStories(*limit, *sort)
	if err != nil {
		log.Fatal(err)
	}

	for i, story := range stories {
		fmt.Printf("%02d - %s (%s)\n", i+1, story.Title, hackernews.GetThreadURL(story.ID))
	}
}
