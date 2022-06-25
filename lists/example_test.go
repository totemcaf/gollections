package lists

import (
	"fmt"
	"strings"
)

// This an example of using a List
func Example_using_lists() {
	emptyList := Of[string]()

	wordList := Of("hello", "world")

	newList := emptyList.Append("a value")

	completeList := newList.Concat(wordList)

	fmt.Printf("total count of elements: %d", completeList.Count())

	isLongWord := func(s string) bool { return len(s) > 4 }

	longWords := completeList.FilterBy(isLongWord)

	longWordCount1 := completeList.FilterBy(isLongWord).Count()
	longWordCount2 := completeList.CountBy(isLongWord)

	fmt.Printf("Long words: %v or %d or %d", longWords, longWordCount1, longWordCount2)

	longWordsInUppercaseJoined := completeList.FilterBy(isLongWord).Map(strings.ToUpper).Join(",")

	fmt.Println("Long words in uppercase and joined", longWordsInUppercaseJoined)
}

// Shows how to convert list to a different element type
func Example_using_map_to_different_type() {

	words := Of("one", "ring", "to", "rule", "them", "all")

	wordSize := func(word string) int { return len(word) }

	sizes := Map(words, wordSize)

	addInts := func(sum, e int) int { return sum + e }

	sumAllSizes := sizes.Reduce(addInts)

	fmt.Println("Sum is", sumAllSizes)

	oddNumber := func(n int) bool { return n%2 != 0 }

	sumWordLengthHavingOddLetterCount := Map(words, wordSize).FilterBy(oddNumber).Reduce(addInts)

	fmt.Println("Sum is", sumWordLengthHavingOddLetterCount)
}
