// Given an array of integers, find the sum of its elements.

// Function Description

// Complete the  function with the following parameter(s):

// : an array of integers

// Returns
// : the sum of the array elements

// STDIN           Function
// -----           --------
// 6               ar[] size n = 6
// 1 2 3 4 10 11   ar = [1, 2, 3, 4, 10, 11]

// output 
// 31
// Print the sum of the array's elements: 1 + 2 + 3 + 4 + 10 + 11 = 31

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
 * Complete the 'simpleArraySum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY ar as parameter.
 */

func simpleArraySum(ar []int32) int32 {
    var result int32
    for i := range ar{
        result += ar[i]
    }
return result
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    arCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    arTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var ar []int32

    for i := 0; i < int(arCount); i++ {
        arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
        checkError(err)
        arItem := int32(arItemTemp)
        ar = append(ar, arItem)
    }

    result := simpleArraySum(ar)

    fmt.Fprintf(writer, "%d\n", result)

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



