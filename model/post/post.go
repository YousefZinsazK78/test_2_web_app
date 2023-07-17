package post

import "fmt"

type Post struct {
	Title       string
	Description string
}

func (p Post) String() string {
	return fmt.Sprintf("title: %s, description: %s\n", p.Title, p.Description)
}
