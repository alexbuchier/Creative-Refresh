package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed with a different value based on current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random number within a range
	result := rand.Intn(100) + 1

	fmt.Println("Your random number is", result)
}

func randomSlice(l int) []int {
	slice := make([]int, l)
	for i := 0; i < l; i++ {
		slice[i] = rand.Intn(100)
	}
	return slice
}

func countEachElement(slice []int) []int {
	// Initialize counts by 0s
	counts := make([]int, 100)
	for _, value := range slice {
		counts[value]++
	}
	return counts
}

func findMostCommon(counts []int) int {
	// Initialize variables to keep track of most common
	commonValue := 0
	commonCount := 0

	// Loop through counts to find most common
	for i, count := range counts {
		if count > commonCount {
			commonValue = i
			commonCount = count
		}
	}
	return commonValue
}

func findLeastCommon(counts []int) int {
	// Initialize variables to keep track of least common
	commonValue := 0
	commonCount := 0

	// Loop through counts to find least common
	for i, count := range counts {
		// Set initial value
		if i == 0 {
			commonValue = i
			commonCount = count
		}
		if count < commonCount {
			commonValue = i
			commonCount = count
		}
	}
	return commonValue
}

func shuffleSlice(slice []int) {
	// Loop through slice and swap two random elements
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func filterOddSlice(slice []int) []int {
	// Initialize a slice to hold the result
	var result []int

	// Loop through slice and add only elements with odd values
	for _, value := range slice {
		if value%2 != 0 {
			result = append(result, value)
		}
	}
	return result
}

func combineSlices(slice1, slice2 []int) []int {
	// Initialize a slice to hold the result
	var result []int

	// Loop through both slices and append values
	for _, value1 := range slice1 {
		result = append(result, value1)
	}
	for _, value2 := range slice2 {
		result = append(result, value2)
	}
	return result
}

func sortSlice(slice []int) {
	// Loop through slice and find smallest value
	for i := 0; i < len(slice); i++ {
		smallestIndex := i
		for j := i; j < len(slice); j++ {
			if slice[j] < slice[smallestIndex] {
				smallestIndex = j
			}
		}
		// Swap current element with smallest element
		slice[i], slice[smallestIndex] = slice[smallestIndex], slice[i]
	}
}

func removeDuplicates(slice []int) []int {
	// Initialize a slice to hold the result
	var result []int

	// Loop through slice and add unique elements
	for _, value := range slice {
		if !contains(result, value) {
			result = append(result, value)
		}
	}
	return result
}

func contains(slice []int, value int) bool {
	// Loop through slice to check if the element is in it
	for _, num := range slice {
		if num == value {
			return true
		}
	}
	return false
}

func countChar(str string) map[string]int {
	// Initialize a map to hold the result
	result := make(map[string]int)

	// Loop through string and count each character
	for _, char := range str {
		charStr := string(char)
		_, exists := result[charStr]
		if exists {
			result[charStr]++
		} else {
			result[charStr] = 1
		}
	}
	return result
}

func containsKV(m map[string]int, key string, value int) bool {
	// Loop through map to check if the key-value pair is in it
	for k, v := range m {
		if k == key && v == value {
			return true
		}
	}
	return false
}

func printMap(m map[string]int) {
	// Loop through map and print every key-value pair
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}

func isPalindrome(str string) bool {
	// Loop through string from both ends and check if they are equal
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func maxOf(nums ...int) int {
	// Initialize variable to hold result
	max := nums[0]

	// Loop through numbers and find max
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func minOf(nums ...int) int {
	// Initialize variable to hold result
	min := nums[0]

	// Loop through numbers and find min
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func sumOf(nums ...int) int {
	// Initialize variable to hold result
	sum := 0

	// Loop through numbers and sum them
	for _, num := range nums {
		sum += num
	}
	return sum
}

func hasPrefix(str, prefix string) bool {
	// Check if string has the given prefix
	return str[0:len(prefix)] == prefix
}

func hasSuffix(str, suffix string) bool {
	// Check if string has the given suffix
	return str[len(str)-len(suffix):] == suffix
}

func splitByWhiteSpace(str string) []string {
	// Initialize a slice to hold the result
	var result []string

	// Split string by white space
	for _, word := range strings.Fields(str) {
		result = append(result, word)
	}
	return result
}

func reverseString(str string) string {
	// Initialize a string to hold the result
	var result string

	// Loop through string in reverse order
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return result
}

func capitalizeString(str string) string {
	// Make sure string is at least two characters
	if len(str) < 2 {
		return str
	}

	// Capitalize first letter and lowercase the rest
	return strings.ToUpper(string(str[0])) + strings.ToLower(str[1:])
}

func findSubstring(str, sub string) int {
	// Loop through string and try to find the substring
	for i := 0; i < len(str); i++ {
		if str[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}

func repeatString(str string, times int) string {
	// Initialize a string to hold the result
	var result string

	// Loop through number of times and concatenate string
	for i := 0; i < times; i++ {
		result += str
	}
	return result
}

func containsOnlyDigits(str string) bool {
	// Loop through string and make sure every character is a digit
	for _, char := range str {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func isValidIdentifier(str string) bool {
	// Make sure string is at least two characters
	if len(str) < 2 {
		return false
	}

	// Make sure first character is a letter
	if !unicode.IsLetter(rune(str[0])) {
		return false
	}

	// Make sure remaining characters are letters or digits
	for _, char := range str[1:] {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func extractData(str string) string {
	// Initialize a string to hold the result
	var result string

	// Loop through string and save data in variable
	for _, char := range str {
		if unicode.IsDigit(char) || unicode.IsLetter(char) {
			result += string(char)
		}
	}
	return result
}

func printChars(str string) {
	// Loop through string and print every character
	for _, char := range str {
		fmt.Printf("%c\n", char)
	}
}

func removeLastChar(str string) string {
	// Make sure string is at least two characters
	if len(str) < 2 {
		return str
	}
	return str[:len(str)-1]
}

func removeFirstChar(str string) string {
	// Make sure string is at least two characters
	if len(str) < 2 {
		return str
	}
	return str[1:]
}

func isAnagram(str1, str2 string) bool {
	// Make sure both strings have the same length
	if len(str1) != len(str2) {
		return false
	}

	// Initialize maps to count each character
	count1 := make(map[rune]int)
	count2 := make(map[rune]int)

	// Loop through strings and count each character
	for _, char1 := range str1 {
		_, exists := count1[char1]
		if exists {
			count1[char1]++
		} else {
			count1[char1] = 1
		}
	}
	for _, char2 := range str2 {
		_, exists := count2[char2]
		if exists {
			count2[char2]++
		} else {
			count2[char2] = 1
		}
	}

	// Compare maps
	if reflect.DeepEqual(count1, count2) {
		return true
	} else {
		return false
	}
}

func isPalindromeNumber(num int) bool {
	// Initialize variables to hold the number and its reversed version
	n := num
	rev := 0

	// Reverse the number
	for n > 0 {
		rev = (rev * 10) + (n % 10)
		n /= 10
	}

	// Compare the reversed number with the original
	if num == rev {
		return true
	}
	return false
}

func validIP(IP string) bool {
	// Make sure IP has four parts
	parts := strings.Split(IP, ".")
	if len(parts) != 4 {
		return false
	}

	// Check each part
	for _, part := range parts {
		// Make sure part has only digits
		if !containsOnlyDigits(part) {
			return false
		}
		// Make sure part is not greater than 255
		num, _ := strconv.Atoi(part)
		if num > 255 {
			return false
		}
	}
	return true
}

func reverseWords(str string) string {
	// Initialize a string to hold the result
	var result string

	// Split string by white space and reverse it
	words := strings.Fields(str)
	for i := len(words) - 1; i >= 0; i-- {
		result += words[i] + " "
	}

	// Remove trailing white space
	return strings.TrimSpace(result)
}

func concatStrings(str1, str2 string) string {
	// Initialize a string to hold the result
	var result string

	// Loop through strings and concatenate them
	for i := 0; i < len(str1) || i < len(str2); i++ {
		if i < len(str1) {
			result += string(str1[i])
		}
		if i < len(str2) {
			result += string(str2[i])
		}
	}
	return result
}

func convertToBinary(num int) string {
	// Initialize a string to hold the result
	var result string

	// Convert number to binary
	for num > 0 {
		remainder := num % 2
		result = fmt.Sprintf("%d", remainder) + result
		num /= 2
	}
	return result
}

func compareSlices(slice1, slice2 []int) bool {
	// Make sure slices have the same length
	if len(slice1) != len(slice2) {
		return false
	}

	// Loop through slices and compare their elements
	for i, num1 := range slice1 {
		if num1 != slice2[i] {
			return false
		}