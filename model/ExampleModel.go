package model

import (
	"Golang-Echo-MVC-Pattern/settings"
)

type ExamplePost struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ExampleModel struct {
	db settings.DatabaseConfig
}

type ExampleModelInterface interface {
	GetPosts() []ExamplePost
}

// Get Example Post
func (ExampleModel ExampleModel) GetPosts() []ExamplePost {
	posts := []ExamplePost{
		{
			Title:       "NewsOne",
			Description: "NewsOneDescription",
		},
		{
			Title:       "NewsTwo",
			Description: "NewsTwoDescription",
		},
	}

	return posts
}
