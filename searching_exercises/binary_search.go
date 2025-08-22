package main

import "fmt"


func binarySearch(arr arrOfNum, target int) int {
    left := 0
    right := len(arr) - 1

    for left <= right {
        mid := (left + right) / 2
        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}