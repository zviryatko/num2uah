package num2uah

import (
	"strings"

	ntw "moul.io/number-to-words"
)

// Convert an amount to a Ukrainian handwriting representation of the amount in hryvnias and kopiyok.
func Convert(amount float64) string {
	integer := int(amount)
	fraction := int(amount*100) - integer*100

	majorWord := toFeminine(ntw.IntegerToUkUa(integer), integer)
	minorWord := toFeminine(ntw.IntegerToUkUa(fraction), fraction)

	currencyMajor := formatPlural(integer, "гривня", "гривні", "гривень")
	currencyMinor := formatPlural(fraction, "копійка", "копійки", "копійок")

	return majorWord + " " + currencyMajor + " та " + minorWord + " " + currencyMinor
}

func formatPlural(number int, singular, plural2to4, pluralOther string) string {
	lastDigit := number % 10
	lastTwoDigits := number % 100

	if lastTwoDigits >= 11 && lastTwoDigits <= 14 {
		return pluralOther
	} else if lastDigit == 1 {
		return singular
	} else if lastDigit >= 2 && lastDigit <= 4 {
		return plural2to4
	}
	return pluralOther
}

func toFeminine(word string, number int) string {
	if number == 1 {
		return "одна"
	}
	if number == 2 {
		return "дві"
	}

	lastTwoDigits := number % 100

	// Convert a final occurrence of masculine forms to feminine
	switch lastTwoDigits {
	case 1:
		// Replace the last "один" with "одна" for numbers ending in 1
		parts := strings.Split(word, " ")
		for i := len(parts) - 1; i >= 0; i-- {
			if parts[i] == "один" {
				parts[i] = "одна"
				break
			}
		}
		return strings.Join(parts, " ")
	case 2, 22:
		// Replace the last "два" with "дві"
		parts := strings.Split(word, " ")
		for i := len(parts) - 1; i >= 0; i-- {
			if parts[i] == "два" {
				parts[i] = "дві"
				break
			}
		}
		return strings.Join(parts, " ")
	case 21:
		return replaceLastOccurrence(word, "двадцять один", "двадцять одна", "мільйон")
	case 32:
		return replaceLastOccurrence(word, "тридцять два", "тридцять дві", "мільйона")
	}

	return word
}

func replaceLastOccurrence(word, old, new, excludeAfter string) string {
	if strings.HasSuffix(word, old) {
		return strings.TrimSuffix(word, old) + new
	}
	lastIdx := strings.LastIndex(word, old)
	if lastIdx != -1 {
		after := word[lastIdx+len(old):]
		if excludeAfter == "" || !strings.Contains(after, excludeAfter) {
			return word[:lastIdx] + new + after
		}
	}
	return word
}
