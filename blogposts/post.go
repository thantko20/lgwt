package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	titlePrefix       = "Title: "
	descriptionPrefix = "Description: "
	tagsPrefix        = "Tags: "
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(prefix string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), prefix)
	}

	title := readMetaLine(titlePrefix)
	description := readMetaLine(descriptionPrefix)
	tags := strings.Split(readMetaLine(tagsPrefix), ", ")

	post := Post{Title: title, Description: description, Tags: tags, Body: readBody(scanner)}
	return post, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()

	// var body string
	// for scanner.Scan() {
	// 	body += scanner.Text() + "\n"
	// }
	// body = strings.TrimSuffix(body, "\n")
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	body := strings.TrimSuffix(buf.String(), "\n")
	return body
}
