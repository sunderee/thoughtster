package main

import "fmt"

func main() {
	envFile, envFileLoadingError := loadDotEnv()
	if envFileLoadingError != nil {
		panic(envFileLoadingError)
	}
	fmt.Println(envFile)
}
