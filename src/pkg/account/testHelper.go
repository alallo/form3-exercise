package account

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func ReadMockedResponseFromFile(t *testing.T, fileName string) string {
	jsonFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	var buf strings.Builder
	written, err := io.Copy(&buf, jsonFile)
	if err != nil || written < 1 {
		t.Errorf("Something went wrong while reading file: %v", err.Error())
	}

	// body string in JSON format used for the mock response
	body := buf.String()

	return body
}
