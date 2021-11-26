package logger

import (
	"path"
	"runtime"
)

// getCallerDetails returns the details of the caller of the function.
//
// The return order is: package-name, file-path, line-number.
func getCallerDetails(skip int) (string, string, int) {
	pc, file, line, ok := runtime.Caller(skip)
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
