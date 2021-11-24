package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

// getLogFilePointer returns the pointer to the file present at the provided path.
//
// If the file is not present, it is created. All the steps are logged.
func getLogFilePointer(filePath string) (*os.File, error) {
	funcName := "logger.getLogFilePointer"
	fmt.Printf("%s: Getting pointer for file: %s\n", funcName, filePath)

	info, err := os.Stat(filePath)
	if err == nil {
		if info.IsDir() {
			return nil, fmt.Errorf("%s: file path is a directory", funcName)
		}
		fmt.Printf("%s: File already present.\n", funcName)
		return os.OpenFile(filePath, os.O_WRONLY, os.ModeAppend)
	}

	if !os.IsNotExist(err) {
		return nil, fmt.Errorf("%s: failed to get information on file: %w", funcName, err)
	}

	fmt.Printf("%s: File absent. Creating...\n", funcName)
	err = os.MkdirAll(path.Dir(filePath), os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return nil, fmt.Errorf("%s: error while creating file: %w", funcName, err)
	}

	return os.Create(filePath)
}

// getCallerDetails returns the details of the caller of the function.
//
// The return order is: package-name, file-path, line-number.
func getCallerDetails() (string, string, int) {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return "unknown", "unknown", 0
	}

	details := runtime.FuncForPC(pc)
	return details.Name(), trimPath(file, 2), line
}

// trimPath trims a path to the given number of segments.
//
// Example: trimPath("/some/random/path", 2) will give "/random/path".
func trimPath(p string, count int) string {
	var elements []string

	for i := 0; i < count; i++ {
		elements = append([]string{path.Base(p)}, elements...)
		p = path.Dir(p)
	}

	return path.Join(elements...)
}
