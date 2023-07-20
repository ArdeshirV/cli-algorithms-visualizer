package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func main() {
	arr := generateRandomArray(10)
	fmt.Println("Initial array:", arr)

	// Prompt user to select the sorting algorithm
	prompt := promptui.Select{
		Label: "Select Sorting Algorithm",
		Items: []string{
			"Bubble Sort",
			"Selection Sort",
			"Insertion Sort",
			"Gnome Sort",
			"Cocktail Shaker Sort",
			"Comb Sort",
			"Odd-Even Sort",
		},
	}

	_, algorithm, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	clearConsole()

	switch algorithm {
	case "Bubble Sort":
		bubbleSortVisualizer(arr)
	case "Selection Sort":
		selectionSortVisualizer(arr)
	case "Insertion Sort":
		insertionSortVisualizer(arr)
	case "Gnome Sort":
		gnomeSortVisualizer(arr)
	case "Cocktail Shaker Sort":
		cocktailShakerSortVisualizer(arr)
	case "Comb Sort":
		combSortVisualizer(arr)
	case "Odd-Even Sort":
		oddEvenSortVisualizer(arr)
	default:
		fmt.Println("Invalid selection")
		return
	}

	fmt.Println("Sorted array:", arr)
}
