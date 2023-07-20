package main

import "time"

func cocktailShakerSortVisualizer(arr []int, delay time.Duration) {
	left := 0
	right := len(arr) - 1

	for left <= right {
		for i := left; i < right; i++ {
			visualizeIteration(White, arr, i, i+1, delay)
			if arr[i] > arr[i+1] {
				visualizeIteration(LightYellow, arr, i, i+1, delay)
				arr[i], arr[i+1] = arr[i+1], arr[i]
				visualizeIteration(LightBlue, arr, i, i+1, delay)
			}
		}
		right--

		for i := right; i > left; i-- {
			visualizeIteration(White, arr, i-1, i, delay)
			if arr[i-1] > arr[i] {
				visualizeIteration(LightYellow, arr, i-1, i, delay)
				arr[i], arr[i-1] = arr[i-1], arr[i]
				visualizeIteration(LightBlue, arr, i-1, i, delay)
			}
		}
		left++
	}
}
