package main

import "fmt"

func areSentencesSimilar(sentence1, sentence2 []string, similarPairs [][]string) bool {
	sim := map[[2]string]struct{}{}
	for _, v := range similarPairs {
		sim[[2]string{v[0], v[1]}] = struct{}{}
	}
	if len(sentence1) != len(sentence2) {
		return false
	}
	for i := range sentence1 {
		_, ok1 := sim[[2]string{sentence1[i], sentence2[i]}]
		_, ok2 := sim[[2]string{sentence2[i], sentence1[i]}]
		if !(sentence1[i] == sentence2[i] || ok1 || ok2) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(areSentencesSimilar([]string{"this", "summer", "thomas", "get",
		"actually", "actually", "rich", "and", "possess", "the", "actually", "great",
		"and", "fine", "vehicle", "every", "morning", "he", "drives", "one", "nice",
		"car", "around", "one", "great", "city", "to", "have", "single", "super",
		"excellent", "super", "as", "his", "brunch", "but", "he", "only", "eat",
		"single", "few", "fine", "food", "as", "some", "fruits", "he", "wants", "to",
		"eat", "an", "unique", "and", "actually", "healthy", "life"},
		[]string{"this", "summer", "thomas", "get", "very", "very", "rich", "and",
			"possess", "the", "very", "fine", "and", "well", "car", "every",
			"morning", "he", "drives", "a", "fine", "car", "around", "unique",
			"great", "city", "to", "take", "any", "really", "wonderful", "fruits",
			"as", "his", "breakfast", "but", "he", "only", "drink", "an", "few",
			"excellent", "breakfast", "as", "a", "super", "he", "wants", "to",
			"drink", "the", "some", "and", "extremely", "healthy", "life"},
		[][]string{{"good", "nice"}, {"good", "excellent"}, {"good", "well"},
			{"good", "great"}, {"fine", "nice"}, {"fine", "excellent"},
			{"fine", "well"}, {"fine", "great"}, {"wonderful", "nice"},
			{"wonderful", "excellent"}, {"wonderful", "well"}, {"wonderful", "great"},
			{"extraordinary", "nice"}, {"extraordinary", "excellent"},
			{"extraordinary", "well"}, {"extraordinary", "great"}, {"one", "a"},
			{"one", "an"}, {"one", "unique"}, {"one", "any"}, {"single", "a"},
			{"single", "an"}, {"single", "unique"}, {"single", "any"}, {"the", "a"},
			{"the", "an"}, {"the", "unique"}, {"the", "any"}, {"some", "a"}, {"some", "an"},
			{"some", "unique"}, {"some", "any"}, {"car", "vehicle"}, {"car", "automobile"},
			{"car", "truck"}, {"auto", "vehicle"}, {"auto", "automobile"},
			{"auto", "truck"}, {"wagon", "vehicle"}, {"wagon", "automobile"},
			{"wagon", "truck"}, {"have", "take"}, {"have", "drink"}, {"eat", "take"},
			{"eat", "drink"}, {"entertain", "take"}, {"entertain", "drink"},
			{"meal", "lunch"}, {"meal", "dinner"}, {"meal", "breakfast"}, {"meal", "fruits"},
			{"super", "lunch"}, {"super", "dinner"}, {"super", "breakfast"}, {"super", "fruits"},
			{"food", "lunch"}, {"food", "dinner"}, {"food", "breakfast"}, {"food", "fruits"},
			{"brunch", "lunch"}, {"brunch", "dinner"}, {"brunch", "breakfast"},
			{"brunch", "fruits"}, {"own", "have"}, {"own", "possess"}, {"keep", "have"},
			{"keep", "possess"}, {"very", "super"}, {"very", "actually"}, {"really", "super"},
			{"really", "actually"}, {"extremely", "super"}, {"extremely", "actually"}}))
}
