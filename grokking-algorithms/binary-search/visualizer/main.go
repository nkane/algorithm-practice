package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	resultText := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetText("Binary Search Visualizer")

	inputField := tview.NewInputField().
		SetLabel("Enter a sorted list of numbers (comma-separated): ").
		SetFieldWidth(30)

	searchField := tview.NewInputField().
		SetLabel("Enter number to search for: ").
		SetFieldWidth(10)

	grid := tview.NewGrid().
		SetRows(1, 1, 1, 1).
		SetColumns(0).
		AddItem(resultText, 0, 0, 1, 1, 0, 0, false).
		AddItem(inputField, 1, 0, 1, 1, 0, 0, true).
		AddItem(searchField, 2, 0, 1, 1, 0, 0, true)

	inputField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			app.SetFocus(searchField)
		}
	})

	searchField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			listStr := inputField.GetText()
			targetStr := searchField.GetText()
			list, err := parseInput(listStr)
			if err != nil {
				resultText.SetText(fmt.Sprintf("[red]Error: %s", err.Error()))
				return
			}
			target, err := strconv.Atoi(targetStr)
			if err != nil {
				resultText.SetText(fmt.Sprintf("[red]Error: %s", err.Error()))
				return
			}
			result := binarySearchVisualizer(list, target)
			resultText.SetText(result)
			app.Draw()
		}
	})

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}

func parseInput(input string) ([]int, error) {
	parts := strings.Split(input, ",")
	var list []int
	for _, part := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			return nil, err
		}
		list = append(list, num)
	}
	return list, nil
}

func binarySearchVisualizer(list []int, target int) string {
	low, high := 0, len(list)-1
	steps := []string{}

	log.Println("debug out")
	for low <= high {
		mid := (low + high) / 2
		steps = append(steps, fmt.Sprintf("[yellow]Step: low=%d, mid=%d, high=%d", low, mid, high))

		if list[mid] == target {
			steps = append(steps, fmt.Sprintf("[green]Found %d at index %d", target, mid))
			return strings.Join(steps, "\n")
		}
		if list[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	steps = append(steps, fmt.Sprintf("[red]%d not found in list", target))
	return strings.Join(steps, "\n")
}
