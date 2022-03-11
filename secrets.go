package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func loadDotEnv() (map[string]string, error) {
	dotEnvFile, fileOpeningError := os.Open(".env")
	if fileOpeningError != nil {
		return nil, &envParsingExceptionModel{message: fileOpeningError.Error()}
	}
	defer dotEnvFile.Close()

	var envFilesMap map[string]string = make(map[string]string)
	var scanner *bufio.Scanner = bufio.NewScanner(dotEnvFile)
	for scanner.Scan() {
		line := scanner.Text()
		separatedVariables := strings.Split(line, "=")
		if len(separatedVariables) != 2 {
			return nil, &envParsingExceptionModel{message: ""}
		}

		envFilesMap[separatedVariables[0]] = separatedVariables[len(separatedVariables)-1]
	}

	return envFilesMap, nil
}

type envParsingExceptionModel struct {
	message string
}

func (exception *envParsingExceptionModel) Error() string {
	return fmt.Sprintf(
		"ENV parsing error %s",
		exception.message,
	)
}
