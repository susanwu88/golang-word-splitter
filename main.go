package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Cloud Run Jobs automatically sets these environment variables.
	// taskIndex is the index of the task in the job execution, from 0 to task-count-1.
	// taskAttempt is the number of times this task has been retried.
	taskIndex := os.Getenv("CLOUD_RUN_TASK_INDEX")
	taskAttempt := os.Getenv("CLOUD_RUN_TASK_ATTEMPT")

	// Default text to split. You can also pass this in as an environment variable
	// or another configuration method.
	textToSplit := "Hello, World from a Cloud Run Job!"

	// You can optionally get the text from an environment variable.
	if textFromEnv := os.Getenv("INPUT_TEXT"); textFromEnv != "" {
		textToSplit = textFromEnv
	}


	log.Printf("Starting Task #%s, Attempt #%s", taskIndex, taskAttempt)
	log.Printf("Input text: '%s'", textToSplit)

	// Split the string by spaces.
	// The strings.Fields function is a good choice as it handles
	// multiple spaces gracefully.
	words := strings.Fields(textToSplit)

	fmt.Println("--- Words ---")
	for _, word := range words {
		// Print each word on a new line.
		fmt.Println(word)
	}
	fmt.Println("-------------")
}
