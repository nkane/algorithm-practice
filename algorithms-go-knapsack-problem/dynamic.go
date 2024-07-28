package knapsack

func DynamicProgramming(items []Item, allowedWeight int) ([]Item, int, int) {
	numItems := len(items)
	// Allocate the arrays.
	solutionValue := make([][]int, numItems)
	prevWeight := make([][]int, numItems)
	for i := 0; i < numItems; i++ {
		solutionValue[i] = make([]int, allowedWeight+1)
		prevWeight[i] = make([]int, allowedWeight+1)
	}

	// Initialize the row item 0.
	for w := 0; w <= allowedWeight; w++ {
		if items[0].Weight <= w {
			// items[0] fits.
			solutionValue[0][w] = items[0].Value
			prevWeight[0][w] = -1
		} else {
			// items[0] does not fit.
			solutionValue[0][w] = 0
			prevWeight[0][w] = w
		}
	}

	// Fill in the remaining table rows.
	for i := 1; i < numItems; i++ {
		for w := 0; w <= allowedWeight; w++ {
			// Calculate the value if we do not use the new item i.
			valueWithoutI := solutionValue[i-1][w]

			// Calculate the value if we do use the new item i.
			valueWithI := 0
			if items[i].Weight <= w { // Make sure it fits.
				valueWithI = solutionValue[i-1][w-items[i].Weight] + items[i].Value
			}

			// See which is better.
			if valueWithoutI >= valueWithI {
				// We're better off omitting item i.
				solutionValue[i][w] = valueWithoutI
				prevWeight[i][w] = w
			} else {
				// We're better off including item i.
				solutionValue[i][w] = valueWithI
				prevWeight[i][w] = w - items[i].Weight
			}
		}
	}

	// Set isSelected to false for all items.
	// (It should already be false, but let's play it safe.)
	for i := range items {
		items[i].IsSelected = false
	}

	// Reconstruct the solution.
	// Get the row and column for the final solution.
	i := numItems - 1
	w := allowedWeight

	// Work backwards until we reach an initial solution.
	for i >= 0 {
		// Check prevWeight for the current solution.
		prevW := prevWeight[i][w]
		if w == prevW {
			// We skipped item i.
			// Leave w unchanged.
		} else {
			// We added item i.
			items[i].IsSelected = true // Select this item in the solution.
			w = prevW                  // Move to the previous solution's weight.
		}
		i -= 1 // Move to the previous row.
	}
	return items, solutionValue[numItems-1][allowedWeight], 1
}
