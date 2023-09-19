package utils

import (
	"bufio"
	"os"
	"strings"

	"github.com/RikunjSindhwad/Task-Ninja/pkg/visuals"
)

func GetLogPaths(stdoutDir, stderrDir, taskName string) (string, string) {
	stdoutFile := ""
	stderrFile := ""
	if stdoutDir != "" || stderrDir != "" {
		stdouterr := ensurePathExists(stdoutDir)
		if stdouterr != nil {
			visuals.PrintState("fetal", taskName, "Error creating stdout directory: "+stdouterr.Error())
		}
		stderrerr := ensurePathExists(stderrDir)
		if stderrerr != nil {
			visuals.PrintState("fetal", taskName, "Error creating stderr directory: "+stderrerr.Error())
		}
		stdoutFile = stdoutDir + "/" + strings.ReplaceAll(taskName, " ", "-") + ".stdout"
		stderrFile = stderrDir + "/" + strings.ReplaceAll(taskName, " ", "-") + ".stderr"
	}
	return stdoutFile, stderrFile
}

func ensurePathExists(path string) error {
	// Check if the path already exists.
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Path does not exist, create it.
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
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
