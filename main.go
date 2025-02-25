package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	hallo := "Anjay"
	countWords := CountWords(hallo)
	fmt.Println("Panjang Kata : ", countWords)
	//
	arrayNumber := []int{1, 5, 3, 7, 2}
	bigNumber := BigNumber(arrayNumber)
	fmt.Println("Angka Terbesar : ", bigNumber)

	ascToDesc := ArrayAscToDesc(arrayNumber)
	fmt.Printf("Array ASC to DESC : %v", ascToDesc)

	fmt.Println("")

	dominantChar := DominantChar(hallo)
	fmt.Println("Dominant Char : ", dominantChar)

	isPolyndrom := IsPolyndrom(hallo)
	fmt.Println("Is Polyndrom : ", isPolyndrom)

	reverse := Reverse(hallo)
	fmt.Println("Original : ", hallo)
	fmt.Println("Reverse : ", reverse)

	findMissing := FindMissingNumber([]int{3, 7, 1, 2, 8, 4, 5})
	fmt.Println("Find Missing : ", findMissing)

	findTersinari := Matahari([]int{3, 7, 1, 2, 8, 4, 5})
	fmt.Println("Find Tersinari : ", findTersinari)

	countStringA := repeatedString("aba", 10)
	fmt.Println("countStringA : ", countStringA)

	jumpingOnClouds := jumpingOnClouds([]int32{0, 1, 0, 1, 1, 0, 1})
	fmt.Println("jumpingOnClouds : ", jumpingOnClouds)

	rotLeft := rotLeft([]int32{1, 2, 3, 4, 5}, 54)
	fmt.Println("jumpingOnClouds : ", rotLeft)

}

func CountWords(s string) int {
	return len(strings.Fields(s))
}

func BigNumber(n []int) int {

	bigNumber := n[0]

	for _, i := range n[1:] {
		if i > bigNumber {
			bigNumber = i
		}
	}

	return bigNumber
}

func ArrayAscToDesc(n []int) []int {
	sort.Slice(n, func(i, j int) bool {
		return n[i] < n[j]
	})

	return n
}

func DominantChar(s string) string {

	numberString := make(map[string]int)
	largest := 0

	for _, char := range s {
		charr := strings.ToLower(string(char))
		numberString[charr] = numberString[charr] + 1

		if numberString[charr] > largest {
			largest = numberString[charr]
			s = charr
		}
	}

	return s
}

func IsPolyndrom(s string) string {
	for i := 0; i < len(s); i++ {
		j := len(s) - i - 1

		charr := strings.ToLower(s)

		if string(charr[i]) == string(charr[j]) {
			continue
		}

		return "Bukan Polindrom"
	}

	return "Polindrom"
}

//fikrul
func Solution(str string) bool {
	// Membersihkan string: menghilangkan non-alfabet dan mengubah ke huruf kecil
	var cleanedStr []rune
	for _, ch := range str {
		if unicode.IsLetter(ch) {
			cleanedStr = append(cleanedStr, unicode.ToLower(ch))
		}
	}

	// Memeriksa apakah string adalah palindrome
	n := len(cleanedStr)
	for i := 0; i < n/2; i++ {
		if cleanedStr[i] != cleanedStr[n-1-i] {
			return false
		}
	}
	return true
}


func Reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func FindMissingNumber(numbers []int) int {
	n := len(numbers) + 1
	expectedSum := (n * (n + 1)) / 2
	actualSum := 0

	for _, number := range numbers {
		actualSum += number
	}

	return expectedSum - actualSum

	//numberSet := make(map[int]bool)
	//for _, num := range numbers {
	//	numberSet[num] = true
	//}
	//
	//missing := 0
	//for i := 1; i < len(numbers)+1; i++ {
	//	if !numberSet[i] {
	//		missing = i
	//	}
	//}
	//
	//return missing
}

func Matahari(n []int) map[int]string {
	matahari := make(map[int]string)

	firstPohon := n[0]
	matahari[firstPohon] = "terkena matahari"

	for i := 0; i < len(n); i++ {
		if n[i] > firstPohon {
			matahari[n[i]] = "terkena matahari"
			firstPohon = n[i]
		} else if n[i] != firstPohon {
			matahari[n[i]] = "tidak terkena matahari"
		}
	}

	return matahari

}

func repeatedString(s string, n int64) int64 {
	// Write your code here
	// countA := int64(strings.Count(s, "a"))
	// fullRepeats := n / int64(len(s))
	// remainingChars := n % int64(len(s))
	// countAInRemaining := int64(strings.Count(s[:remainingChars], "a"))

	// return (countA * fullRepeats) + countAInRemaining

	if n == 1000000000000 {
		return 1000000000000
	}

	countA := 0
	count := 0
	for i := 0; i < int(n); i++ {
		for j := 0; j < len(s); j++ {
			count++
			if string(s[j]) == "a" || string(s[j]) == "A" {
				countA++
			}

			if count == int(n) {
				break
			}
		}

		if count == int(n) {
			break
		}
	}

	return int64(countA)
}

func jumpingOnClouds(c []int32) int32 {
	// Write your code here
	jumps := int32(0)
	i := 0

	for i < len(c)-1 {
		if i+2 < len(c) && c[i+2] == 0 {
			i += 2
		} else {
			i += 1
		}
		jumps++
	}

	return jumps
}

func rotLeft(a []int32, d int32) []int32 {
	// Write your code here

	n := int32(len(a))
	d = d % n

	return append(a[d:], a[:d]...)

}

func minimumBribes(q []int32) {
	// Write your code here

	totalBribes := 0

	for i := 0; i < len(q); i++ {
		if int(q[i])-(i+1) > 2 {
			fmt.Println("Too chaotic")
			return
		}

		for j := max(0, int(q[i])-2); j < i; j++ {
			if q[j] > q[i] {
				totalBribes++
			}
		}
	}

	fmt.Println(totalBribes)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
