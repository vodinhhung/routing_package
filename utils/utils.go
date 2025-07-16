package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func StringToUint64Slice(input string) ([]uint64, error) {
	if strings.TrimSpace(input) == "" {
		return nil, errors.New("input string is empty")
	}

	parts := strings.Split(input, ",")
	var result []uint64

	for _, part := range parts {
		p := strings.TrimSpace(part)
		if p == "" {
			continue
		}

		num, err := strconv.ParseUint(p, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid uint64 value '%s': %w", p, err)
		}
		result = append(result, num)
	}

	return result, nil
}

func Uint64SliceToString(input []uint64) string {
	if len(input) == 0 {
		return ""
	}

	parts := make([]string, len(input))
	for i, num := range input {
		parts[i] = strconv.FormatUint(num, 10)
	}
	return strings.Join(parts, ",")
}

func Float64ToString(f float64) string {
	// Use 'f' format, with -1 precision to remove trailing zeroes
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func StringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
