package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func UpdateReadme(readme, pattern string) error {
	content, err := readFile(readme)
	if err != nil {
		return err
	}

	updatedContent := updateContent(readme, content, pattern)

	return writeFile(readme, updatedContent)
}

func readFile(filename string) (string, error) {
	result, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("reading %s: %s", filename, err)
	}
	return string(result), nil
}

func writeFile(filename, content string) error {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("writing to %s: %s", filename, err)
	}
	return nil
}

func updateContent(filename, content, pattern string) string {
	re := regexp.MustCompile(pattern)

	return re.ReplaceAllStringFunc(
		content,
		func(match string) string {
			dirName := re.FindStringSubmatch(match)[1]
			subReadme := fmt.Sprintf("./%s/%s", dirName, filename)
			subContent, err := readFile(subReadme)
			if err != nil {
				return fmt.Sprintf("<WARNING: %s not found>", subReadme)
			}
			lines := strings.Split(subContent, "\n")
			if len(lines) > 1 {
				subContent = strings.Join(lines[1:], "\n")
			}
			updatedContent := updateDirLinks(subContent, filename)
			return updateFileLinks(updatedContent, dirName)
		},
	)
}

func updateDirLinks(content, pattern string) string {
	re := regexp.MustCompile(fmt.Sprintf(`\[(.*?)\]\(\.\./%s#(\w+)\)`, pattern))

	return re.ReplaceAllStringFunc(
		content,
		func(match string) string {
			header := re.FindStringSubmatch(match)[1]
			headerLink := re.FindStringSubmatch(match)[2]
			return fmt.Sprintf("[%s](#%s)", header, headerLink)
		},
	)
}

func updateFileLinks(content, dir string) string {
	re := regexp.MustCompile(`\[(.*?)\]\((\w+)\)`)

	return re.ReplaceAllStringFunc(
		content,
		func(match string) string {
			header := re.FindStringSubmatch(match)[1]
			headerLink := re.FindStringSubmatch(match)[2]
			return fmt.Sprintf("[%s](./%s/%s)", header, dir, headerLink)
		},
	)
}
