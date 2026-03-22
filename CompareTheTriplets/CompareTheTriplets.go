// Alice and Bob each created one problem for HackerRank. A reviewer rates the two challenges, awarding points on a scale from 1 to 100 for three categories: problem clarity, originality, and difficulty.

// The rating for Alice's challenge is the triplet a = (a[0], a[1], a[2]), and the rating for Bob's challenge is the triplet b = (b[0], b[1], b[2]).

// The task is to calculate their comparison points by comparing each category:

// If a[i] > b[i], then Alice is awarded 1 point.
// If a[i] < b[i], then Bob is awarded 1 point.
// If a[i] = b[i], then neither person receives a point.

// Example

// a = [1, 2, 3]
// b = [3, 2, 1]

// For elements *0*, Bob is awarded a point because a[0] < b[0].
// For the equal elements a[1] and b[1], no points are earned.
// Finally, for elements 2, a[2] > b[2] so Alice receives a point.
// The return array is [1, 1] with Alice's score first and Bob's second.

// Function Description

// Complete the function compareTriplets with the following parameter(s):

// int a[3]: Alice's challenge rating
// int b[3]: Bob's challenge rating
// Returns

// int[2]: the first element is Alice's score and the second is Bob's score
// Input Format

// The first line contains 3 space-separated integers, a[0], a[1], and a[2], the respective values in triplet a.
// The second line contains 3 space-separated integers, b[0], b[1], and b[2], the respective values in triplet b.

// Constraints

// 1 ≤ a[i] ≤ 100
// 1 ≤ b[i] ≤ 100
// Sample Input 0

// 5 6 7
// 3 6 10
// Sample Output 0

// 1 1
// Explanation 0

// Now, let's compare each individual score:

// , so Alice receives  point.
// , so nobody receives a point.
// , so Bob receives  point.
// Alice's comparison score is , and Bob's comparison score is . Thus, we return the array

// Sample Input 1

// 17 28 30
// 99 16 8
// Sample Output 1

// 2 1

// Comparing the  elements,  so Bob receives a point.
// Comparing the  and  elements,  and  so Alice receives two points.
// The return array is

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'compareTriplets' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func compareTriplets(a []int32, b []int32) []int32 {
	if len(a) != len(b) {
		return []int32{0}
	}
	var AlicePoint, BobPoint int32
	for i := range a {
		alice := a[i]
		bob := b[i]
		switch {
		case alice > bob:
			AlicePoint += 1
		case alice < bob:
			BobPoint += 1
		}
	}
	return []int32{AlicePoint, BobPoint}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < 3; i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	bTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var b []int32

	for i := 0; i < 3; i++ {
		bItemTemp, err := strconv.ParseInt(bTemp[i], 10, 64)
		checkError(err)
		bItem := int32(bItemTemp)
		b = append(b, bItem)
	}

	result := compareTriplets(a, b)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
