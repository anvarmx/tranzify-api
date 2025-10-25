package services

import (
	"strings"
	"unicode"
)

var morseMap = map[string]string{
	"A": ".-", "B": "-...", "C": "-.-.", "D": "-..", "E": ".", "F": "..-.",
	"G": "--.", "H": "....", "I": "..", "J": ".---", "K": "-.-", "L": ".-..",
	"M": "--", "N": "-.", "O": "---", "P": ".--.", "Q": "--.-", "R": ".-.",
	"S": "...", "T": "-", "U": "..-", "V": "...-", "W": ".--", "X": "-..-",
	"Y": "-.--", "Z": "--..",
	"0": "-----", "1": ".----", "2": "..---", "3": "...--", "4": "....-",
	"5": ".....", "6": "-....", "7": "--...", "8": "---..", "9": "----.",
	" ": "/", ".": ".-.-.-",
}

var textMap = reverseMap(morseMap)

func reverseMap(m map[string]string) map[string]string {
	newMap := make(map[string]string)

	for k, v := range m {
		newMap[v] = k
	}

	return newMap
}

func TextToMorse(text string) string {
	text = strings.ToUpper(text)

	var result []string

	for _, char := range text {

		if code, ok := morseMap[string(char)]; ok {
			result = append(result, code)
		} else {
			result = append(result, string(char))
		}
	}

	return strings.Join(result, " ")
}

func MorseToText(morse string) string {
	var result []string

	words := strings.Split(morse, " / ")

	for i, word := range words {
		letters := strings.SplitSeq(word, " ")

		for letter := range letters {

			if val, ok := textMap[letter]; ok {
				result = append(result, strings.ToLower(val))
			} else if letter != "" {
				result = append(result, letter)
			}
		}

		if i != len(words)-1 {
			result = append(result, " ")
		}
	}

	text := strings.Join(result, "")

	textRunes := []rune(text)

	if len(textRunes) > 0 {
		textRunes[0] = unicode.ToUpper(textRunes[0])
	}

	for i := 0; i < len(textRunes)-2; i++ {

		if textRunes[i] == '.' || textRunes[i] == '?' || textRunes[i] == '!' {
			j := i + 1

			for j < len(textRunes) && textRunes[j] == ' ' {
				j++
			}

			if j < len(textRunes) {
				textRunes[j] = unicode.ToUpper(textRunes[j])
			}
		}
	}

	return string(textRunes)
}
