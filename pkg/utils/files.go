package utils

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/RikunjSindhwad/Task-Ninja/v2/pkg/visuals"
)

func CopyMountFiles(mounts []string, hostHiveTaskInputDir, defaultHive string) error {
	for _, mount := range mounts {
		name, err := GetHostHiveTaskDirectory(mount, defaultHive)
		if err != nil {
			return err
		}
		_, dockermountTaskOutputDir := GetInputOutputDocker(name)

		// copy files from input to task hive
		err = CopyDir(dockermountTaskOutputDir, hostHiveTaskInputDir, SanitizeTaskName(mount))
		if err != nil {
			return err
		}
		// Get all folder names in and log in specified folder recuresively
		folders, err := GetDirectorywithkeyword(hostHiveTaskInputDir)
		if err != nil {
			return err
		}
		for _, folder := range folders {
			// delete folders
			err = os.RemoveAll(folder)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func GetDirectorywithkeyword(input string) ([]string, error) {
	var folders []string
	err := filepath.Walk(input, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // return the error to stop the walk
		}
		if info.IsDir() && (strings.HasSuffix(path, "/logs") || strings.HasSuffix(path, "/in")) && input != path {
			folders = append(folders, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return folders, nil
}

func CopyInputFiles(inputs []string, hostHiveTaskInputDir string) error {

	for _, input := range inputs {
		filename := ""
		if strings.Contains(input, ",") {
			parts := strings.Split(input, ",")
			input = parts[0]
			filename = parts[1]
		}
		var err error
		if IsURL(input) {
			if filename == "" {
				err = DownloadFile(input, filepath.Join(hostHiveTaskInputDir, filepath.Base(input)))
			} else {
				err = DownloadFile(input, filepath.Join(hostHiveTaskInputDir, filename))
			}
			if err != nil {
				return err
			}
			continue
		}
		// if input is folder
		if strings.HasSuffix(input, "/") || IsFolder(input) {
			err = CopyDir(input, hostHiveTaskInputDir, filepath.Base(input))
			if err != nil {
				return err
			}
			continue
		}
		// if input is file
		if filename == "" {
			err = CopyFile(input, filepath.Join(hostHiveTaskInputDir, filepath.Base(input)))
		} else {
			err = CopyFile(input, filepath.Join(hostHiveTaskInputDir, filename))
		}
		if err != nil {
			return err
		}

	}
	return nil
}

func ParseDockerLogs(reader io.Reader) (stdoutLogs, stderrLogs []byte, err error) {
	const headerSize = 8 // Docker API log header is 8 bytes

	for {
		header := make([]byte, headerSize)
		_, err := io.ReadFull(reader, header)
		if err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				break
			}
			return nil, nil, err
		}

		size := binary.BigEndian.Uint32(header[4:])
		frame := make([]byte, size)
		_, err = io.ReadFull(reader, frame)
		if err != nil {
			return nil, nil, err
		}

		if header[0] == 1 { // Stdout
			stdoutLogs = append(stdoutLogs, frame...)
		} else if header[0] == 2 { // Stderr
			stderrLogs = append(stderrLogs, frame...)
		}
	}

	return stdoutLogs, stderrLogs, nil
}

func WriteLogsToFile(logs []byte, filePath string) error {
	if filePath != "" {
		if err := os.WriteFile(filePath, logs, 0o644); err != nil {
			return err
		}
	}
	return nil
}

func GetInputOutput(source string) (in, out string) {
	in = filepath.Join(source, "in")
	out = filepath.Join(source, "out")
	err := EnsurePathExists(in)
	if err != nil {
		visuals.PrintState("fatal", "", "Error creating input directory: "+err.Error())
	}
	err = EnsurePathExists(out)
	if err != nil {
		visuals.PrintState("fatal", "", "Error creating output directory: "+err.Error())
	}
	return in, out
}

func IsFolderExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			// The folder does not exist
			return false, nil
		}
		// Some other error occurred
		return false, err
	}
	return info.IsDir(), nil
}

func GetInputOutputDocker(source string) (in, out string) {
	check, _ := IsFolderExists(source + "out")
	if check {
		return filepath.Join(source, "in"), filepath.Join(source, "out")
	}
	return source, source
}

func ReadDynamicRangeFromFile(dynamicRange string) string {
	lines, err := ReadLinesFromFile(dynamicRange)
	if err != nil {
		return ""
	}
	if len(lines) == 1 {
		return lines[0]
	}
	if len(lines) > 2 {
		return ""
	}
	return strings.Join(lines, ",")
}

func ReplaceTaskPlaceholders(command, hive, mode string) string {
	// Regex to find placeholders like {{{TaskName:Type}}}
	// re := regexp.MustCompile(`\{\{\{([^:}]+):([^}]+)\}\}\}`)
	re := regexp.MustCompile(`\{{3}([^:}]+):([^}]+)\}{3}`)

	return re.ReplaceAllStringFunc(command, func(placeholder string) string {
		matches := re.FindStringSubmatch(placeholder)

		if len(matches) != 3 {
			// Not a valid placeholder, return as is.
			return placeholder
		}

		taskName := matches[1]
		taskType := strings.ToLower(matches[2])

		if taskType != "file" && taskType != "folder" {
			// Invalid type, return placeholder as is.
			return placeholder
		}

		// Get the path for the task.
		path := getPathForTask(taskName, taskType, hive, mode)

		return path
	})
}

func getPathForTask(taskName, taskType, hive, mode string) string {
	sanitizedTaskName := SanitizeTaskName(taskName)
	// Get the path for the task.
	if mode == "dynamic" {
		switch taskType {
		case "file":
			return filepath.Join(hive, sanitizedTaskName, "out", "output.txt")
		case "folder":
			return filepath.Join(hive, sanitizedTaskName)
		default:
			return ""
		}
	} else {
		switch taskType {
		case "file":
			return filepath.Join(hive, "in", sanitizedTaskName, "out", "output.txt")
		case "folder":
			return filepath.Join(hive, "in", sanitizedTaskName)
		default:
			return ""
		}
	}

}

func DownloadFile(source, destination string) error {
	// Create the destination file
	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(source)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func GeterrorLogPath(source string) (stdout, stderr string) {
	stdout = filepath.Join(source, "logs", "stdout.log")
	stderr = filepath.Join(source, "logs", "stderr.log")
	err := EnsurePathExists(filepath.Join(source, "logs"))
	if err != nil {
		visuals.PrintState("fetal", "", "Error creating logs directory: "+err.Error())
	}
	return stdout, stderr
}

func GetLogPaths(stdoutDir, stderrDir, taskName string) (stdoutFile, stderrFile string) {
	if stdoutDir != "" {
		stdouterr := EnsurePathExists(stdoutDir)
		if stdouterr != nil {
			visuals.PrintState("fetal", taskName, "Error creating stdout directory: "+stdouterr.Error())
		}
		stdoutFile = stdoutDir + "/" + strings.ReplaceAll(taskName, " ", "-") + ".stdout"
	}
	if stderrDir != "" {
		stderrerr := EnsurePathExists(stderrDir)
		if stderrerr != nil {
			visuals.PrintState("fetal", taskName, "Error creating stderr directory: "+stderrerr.Error())
		}
		stderrFile = stderrDir + "/" + strings.ReplaceAll(taskName, " ", "-") + ".stderr"
	}
	return stdoutFile, stderrFile
}

func EnsurePathExists(path string) error {
	// Check if the path already exists.
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Path does not exist, create it.
		err := os.MkdirAll(path, 0o755)
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

func GetHostHiveTaskDirectory(taskName, defaultHive string) (string, error) {
	sanitizedTaskName := SanitizeTaskName(taskName)

	hostHiveDir := defaultHive
	if !filepath.IsAbs(defaultHive) {
		absPath, err := filepath.Abs(filepath.Join(".", defaultHive))
		if err != nil {
			return "", fmt.Errorf("failed to get absolute path for Hive directory: %v", err)
		}
		hostHiveDir = absPath
	}

	hostHiveTaskDir := filepath.Join(hostHiveDir, sanitizedTaskName)
	if err := os.MkdirAll(hostHiveTaskDir, 0o755); err != nil {
		return "", fmt.Errorf("failed to create host hive directory for task '%s': %v", sanitizedTaskName, err)
	}

	return hostHiveTaskDir, nil
}

func WriteLogs(stdoutFile, stderrFile string, stdoutLogs, stderrLogs []byte) error {
	if err := WriteLogsToFile(stdoutLogs, stdoutFile); err != nil {
		return err
	}
	if err := WriteLogsToFile(stderrLogs, stderrFile); err != nil {
		return err
	}
	return nil
}

func CopyDir(src, dst, taskName string) error {
	// Append taskName to dst only for the top-level call
	if taskName != "" {
		dst = filepath.Join(dst, taskName)
	}

	// Create the destination directory
	err := os.MkdirAll(dst, 0o755)
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		fileInfo, err := entry.Info()
		if err != nil {
			return err
		}

		if fileInfo.IsDir() {
			// Recursive call for subdirectories without taskName
			err = CopyDir(srcPath, dstPath, "")
			if err != nil {
				return err
			}
		} else {
			// Copy file
			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func IsFolder(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func CopyFile(src, dst string) error {
	var (
		err  error
		sfp  *os.File
		dfp  *os.File
		size int64
	)

	if sfp, err = os.Open(src); err != nil {
		return err
	}
	defer sfp.Close()

	if dfp, err = os.Create(dst); err != nil {
		return err
	}
	defer dfp.Close()

	if size, err = io.Copy(dfp, sfp); err != nil {
		return err
	}

	si, err := os.Stat(src)
	if err != nil {
		return err
	}

	if size != si.Size() {
		return fmt.Errorf("size of copied file %s is %d, expected %d", dst, size, si.Size())
	}

	return nil
}
