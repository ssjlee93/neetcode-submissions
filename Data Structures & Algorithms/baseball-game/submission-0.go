func calPoints(operations []string) int {
	// stack of records
	stack := make([]int, 0, len(operations))
	sum := 0

	var add2 func()
	add2 = func() {
		// edge case : not enough elements
		if len(stack) < 2 {
			return
		}

		n := len(stack)
		added := stack[n-1] + stack[n-2]
		stack = append(stack, added)
		sum += added
	}


	for _, s := range operations {
		// edge case : empty
		if len(s) == 0 {
			continue
		}
		// see if it's a number
		i, err := strconv.Atoi(s)
		// if number -> add
		if err == nil {
			stack = append(stack, i)
			sum += i
		// if + -> look at 2 previous => requires a stack
		} else if s == "+" {
			add2()
		// if D -> double last
		} else if s == "D" {
			// edge case : no element
			if len(stack) == 0 {
				continue
			}
			doubled := 2 * stack[len(stack)-1]
			sum += doubled
			stack = append(stack, doubled)
		// if C -> pop()
		} else if s == "C" {
			if len(stack) == 0 {
				continue
			}
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			sum -= popped
		} else {
			fmt.Println("operation not supported")
			return 99999
		}

	}
	return sum
}
