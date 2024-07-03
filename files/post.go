package blogposts

import (
	"bufio"
	"io"
	"strings"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	titleLine := readLine(titleSeparator)
	descriptionLine := readLine(descriptionSeparator)
	tagsLine := readLine(tagsSeparator)

	tags := strings.Split(tagsLine, ", ")

	post := Post{
		Title:       titleLine,
		Description: descriptionLine,
		Tags:        tags,
	}

	return post, nil
}
