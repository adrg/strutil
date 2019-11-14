package strutil

func min(args ...int) int {
	if len(args) == 0 {
		return 0
	}
	if len(args) == 1 {
		return args[0]
	}

	min := args[0]
	for _, arg := range args[1:] {
		if min > arg {
			min = arg
		}
	}

	return min
}

func minf(args ...float64) float64 {
	if len(args) == 0 {
		return 0
	}
	if len(args) == 1 {
		return args[0]
	}

	min := args[0]
	for _, arg := range args[1:] {
		if min > arg {
			min = arg
		}
	}

	return min
}

func max(args ...int) int {
	if len(args) == 0 {
		return 0
	}
	if len(args) == 1 {
		return args[0]
	}

	max := args[0]
	for _, arg := range args[1:] {
		if max < arg {
			max = arg
		}
	}

	return max
}

func maxf(args ...float64) float64 {
	if len(args) == 0 {
		return 0
	}
	if len(args) == 1 {
		return args[0]
	}

	max := args[0]
	for _, arg := range args[1:] {
		if max < arg {
			max = arg
		}
	}

	return max
}
