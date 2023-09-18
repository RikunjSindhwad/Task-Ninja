package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func DetermineOutputPaths(stdoutDir, stderrDir, taskName string) (string, string) {
	stdoutFile := ""
	stderrFile := ""
	if stdoutDir != "" || stderrDir != "" {
		stdoutFile = stdoutDir + "/" + strings.ReplaceAll(taskName, " ", "-") + ".stdout"
		stderrFile = stderrDir + "/" + strings.ReplaceAll(taskName, " ", "-") + ".stderr"
	}
	return stdoutFile, stderrFile
}

func GetOutputFileNames(taskName, stdoutDir, stderrDir string) (string, string) {
	stdoutFile := ""
	stderrFile := ""
	if stdoutDir != "" || stderrDir != "" {
		formatName := strings.ReplaceAll(taskName, " ", "-")
		stdoutFile = fmt.Sprintf("%s/%s.stdout", stdoutDir, formatName)
		stderrFile = fmt.Sprintf("%s/%s.stderr", stderrDir, formatName)
	}
	return stdoutFile, stderrFile
}

func FindLineNumber(input, pattern string) int {
	lines, err := ReadLinesFromFile(input)
	if err != nil {
		return -1
	}
	for i, line := range lines {
		if strings.Contains(line, pattern) {
			// Add 1 to convert from zero-based index to one-based line number
			return i + 1
		}
	}
	return -1 // Return -1 if the pattern is not found
}

func ReadLinesFromFile(filePath string) ([]string, error) {
	// Check if the file exists
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	// Read lines from the file
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}