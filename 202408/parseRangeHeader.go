package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseRangeHeader(rangeHeader string, fileSize int64) (int64, int64, error) {
	if !strings.HasPrefix(rangeHeader, "bytes=") {
		return 0, 0, fmt.Errorf("invalid range header: must start with 'bytes='")
	}

	rangeSpec := strings.TrimPrefix(rangeHeader, "bytes=")
	rangeParts := strings.Split(rangeSpec, "-")

	if len(rangeParts) != 2 {
		return 0, 0, fmt.Errorf("invalid range header: must contain two parts separated by '-'")
	}

	var start, end int64
	var err error

	switch {
	case rangeParts[0] == "":
		// If start is empty, it means read from the end of the file.
		// Format is `bytes=-end`.
		end, err = strconv.ParseInt(rangeParts[1], 10, 64)
		if err != nil || end < 0 {
			return 0, 0, fmt.Errorf("invalid range end: %v", err)
		}
		start = fileSize - end
		end = fileSize - 1
	case rangeParts[1] == "":
		// If end is empty, it means read until the end of the file.
		// Format is `bytes=start-`.
		start, err = strconv.ParseInt(rangeParts[0], 10, 64)
		if err != nil || start < 0 {
			return 0, 0, fmt.Errorf("invalid range start: %v", err)
		}
		end = fileSize - 1
	default:
		// Both start and end are present.
		start, err = strconv.ParseInt(rangeParts[0], 10, 64)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid range start: %v", err)
		}
		end, err = strconv.ParseInt(rangeParts[1], 10, 64)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid range end: %v", err)
		}
	}

	// Validate the range against the file size.
	if start > end || start < 0 || end >= fileSize {
		return 0, 0, fmt.Errorf("invalid range: start=%d, end=%d, fileSize=%d", start, end, fileSize)
	}

	return start, end, nil
}

func main() {
	fmt.Println(parseRangeHeader("bytes=0-499", 1000))
	fmt.Println(parseRangeHeader("bytes=499", 1000))
}
