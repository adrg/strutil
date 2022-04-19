package mathutil

// Min returns the value of the smallest argument,
// or 0 if no arguments are provided.
func Min(args ...int) int {
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

// Max returns the value of the largest argument,
// or 0 if no arguments are provided.
func Max(args ...int) int {
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

// Minf returns the value of the smallest argument,
// or 0 if no arguments are provided.
func Minf(args ...float64) float64 {
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

// Maxf returns the value of the largest argument,
// or 0 if no arguments are provided.
func Maxf(args ...float64) float64 {
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
