package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/gookit/color"
)

type Color int

const (
	White Color = iota
	LightYellow
	LightBlue
)

func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func clearConsole() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Unsupported platform. Cannot clear console.")
	}
}

func printColoredArray(c Color, array []int, idx1, idx2 int) {
	var firstSlice string
	var secondSlice string
	var thirdSlice string

	if len(array[:idx1]) > 0 {
		firstSlice = strings.Replace(fmt.Sprintf("%v ", array[:idx1]), "[", "", -1)
		firstSlice = strings.Replace(firstSlice, "]", "", -1)
	}
	if idx1 != idx2 && len(array[idx1+1:idx2]) > 0 {
		secondSlice = strings.Replace(fmt.Sprintf("%v ", array[idx1+1:idx2]), "[", "", -1)
		secondSlice = strings.Replace(secondSlice, "]", "", -1)
	}
	if len(array[idx2+1:]) > 0 {
		thirdSlice = strings.Replace(fmt.Sprintf(" %v", array[idx2+1:]), "[", "", -1)
		thirdSlice = strings.Replace(thirdSlice, "]", "", -1)
	}

	var str string
	switch c {
	case LightYellow:
		str = firstSlice +
			color.OpUnderscore.Sprint(color.LightYellow.Sprint(array[idx1])) +
			" " +
			secondSlice +
			color.OpUnderscore.Sprint(color.LightYellow.Sprint(array[idx2])) +
			thirdSlice
	case LightBlue:
		if idx1 != idx2 {
			str = firstSlice +
				color.OpUnderscore.Sprint(color.LightBlue.Sprint(array[idx1])) +
				" " +
				secondSlice +
				color.OpUnderscore.Sprint(color.LightBlue.Sprint(array[idx2])) +
				thirdSlice
		} else {
			str = firstSlice +
				color.OpUnderscore.Sprint(color.LightBlue.Sprint(array[idx1])) +
				" " +
				thirdSlice
		}
	case White:
		str = firstSlice +
			color.OpUnderscore.Sprint(array[idx1]) +
			" " +
			secondSlice +
			color.OpUnderscore.Sprint(array[idx2]) +
			thirdSlice
	}

	fmt.Printf("[ %v ]", str)
}

func visualizeIteration(c Color, array []int, idx1, idx2 int, delay time.Duration) {
	printColoredArray(c, array, idx1, idx2)
	time.Sleep(delay)
	clearConsole()
}

func printAlgorithmDescription(algorithm int) {
	switch algorithm {
	case BubbleSort:
		color.Bold.Println(color.Yellow.Sprint("Bubble Sort"))
		fmt.Println()
		fmt.Println(color.Cyan.Sprint("Bubble Sort is a simple comparison-based sorting algorithm. It works by repeatedly stepping through the list to be sorted,"))
		fmt.Println(color.Cyan.Sprint("compares adjacent elements, and swaps them if they are in the wrong order. The pass through the list is repeated until the list"))
		fmt.Println(color.Cyan.Sprint("is sorted. The algorithm gets its name from the way smaller elements 'bubble' to the top of the list."))
	case SelectionSort:
		color.Bold.Println(color.Yellow.Sprint("Selection Sort"))
		fmt.Println()
		fmt.Println(color.Cyan.Sprint("Selection Sort is an in-place comparison-based sorting algorithm. It works by dividing the input list into two parts:"))
		fmt.Println(color.Cyan.Sprint("the sorted part and the unsorted part. The algorithm repeatedly selects the minimum element from the unsorted part"))
		fmt.Println(color.Cyan.Sprint("and moves it to the end of the sorted part. This process continues until the entire list is sorted."))
	case InsertionSort:
		color.Bold.Println(color.Yellow.Sprint("Insertion Sort"))
		fmt.Println()
		fmt.Println(color.Cyan.Sprint("Insertion Sort is an in-place comparison-based sorting algorithm. It works by dividing the input list into two parts:"))
		fmt.Println(color.Cyan.Sprint("the sorted part and the unsorted part. The algorithm repeatedly picks an element from the unsorted part and inserts"))
		fmt.Println(color.Cyan.Sprint("it into the correct position in the sorted part. This process continues until the entire list is sorted."))
	case GnomeSort:
		color.Bold.Println(color.Yellow.Sprint("Gnome Sort"))
		fmt.Println()
		fmt.Println(color.Cyan.Sprint("Gnome Sort is an in-place comparison-based sorting algorithm. It works by repeatedly comparing adjacent elements."))
		fmt.Println(color.Cyan.Sprint("If the two elements are in the wrong order, it swaps them and moves one step backward. If the two elements"))
		fmt.Println(color.Cyan.Sprint("are in the correct order, it moves one step forward. This process continues until the entire list is sorted."))
	case CocktailShakerSort:
		color.Bold.Println(color.Yellow.Sprint("Cocktail Shaker Sort"))
		fmt.Println()
		fmt.Println(color.Cyan.Sprint("Cocktail Shaker Sort, also known as Bidirectional Bubble Sort, is a variation of Bubble Sort. It works by repeatedly"))
		fmt.Println(color.Cyan.Sprint("sorting the list in both directions, first from the beginning to the end (like Bubble Sort), and then from the end"))
		fmt.Println(color.Cyan.Sprint("to the beginning. The algorithm stops when the list becomes sorted."))
	case CombSort:
		color.Bold.Println(color.Yellow.Sprint("Comb Sort"))
		fmt.Println()
		fmt.Println(color.Cyan.Sprint("Comb Sort is an in-place comparison-based sorting algorithm. It works by dividing the input list into a series of"))
		fmt.Println(color.Cyan.Sprint("gaps and repeatedly sorting the list with a specific shrink factor. The shrink factor reduces the gaps until it becomes 1."))
		fmt.Println(color.Cyan.Sprint("At this point, the algorithm behaves similar to Bubble Sort. Comb Sort is an improvement over Bubble Sort for large lists."))
	case OddEvenSort:
		color.Bold.Println(color.Yellow.Sprint("Odd-Even Sort"))
		fmt.Println()
		fmt.Println(color.Cyan.Sprint("Odd-Even Sort is an in-place comparison-based sorting algorithm. It works by repeatedly comparing and swapping"))
		fmt.Println(color.Cyan.Sprint("adjacent elements at even and odd indices. The process continues until the list is sorted. Odd-Even Sort is known"))
		fmt.Println(color.Cyan.Sprint("for its simplicity but is not very efficient for large lists."))
	default:
		fmt.Println("Invalid selection")
	}
}
