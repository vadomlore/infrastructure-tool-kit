package opshelper

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func File2lines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return LinesFromReader(f)
}

func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

/**
 * Insert sting to n-th line of file.
 * If you want to insert a line, append newline '\n' to the end of the string.
 */
func InsertStringToFile(path, str string, index int) error {
	lines, err := File2lines(path)
	if err != nil {
		return err
	}

	fileContent := ""
	for i, line := range lines {
		if i == index {
			fileContent += str
		}
		fileContent += line
		fileContent += "\n"
	}

	return ioutil.WriteFile(path, []byte(fileContent), 0644)
}

/**
 * Insert sting to lines end with.
 * If you want to insert a line, append newline '\n' to the end of the string.
 */
func InsertStringToFileEndWith(src, dest, str string, ends string, num int) error {
	lines, err := File2lines(src)
	if err != nil {
		return err
	}

	count := 0
	cursor := len(lines)
	for i := len(lines) - 1; i > 0; i-- {
		if strings.HasSuffix(lines[i], ends) {
			count++
		} else {
			count = 0
		}
		if count == num {
			cursor = i
			break
		}
	}

	if cursor < 0 {
		cursor = 0
	}

	fileContent := ""
	for i, line := range lines {
		if i == cursor {
			fileContent += str
		}
		fileContent += line
		fileContent += "\n"
	}

	return ioutil.WriteFile(dest, []byte(fileContent), 0644)
}
