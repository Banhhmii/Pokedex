package main

import ("strings")

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	trimmed := strings.TrimSpace(lowered)
	splitText := strings.Fields(trimmed)

	return splitText
}