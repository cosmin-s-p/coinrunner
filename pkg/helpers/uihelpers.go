package helpers

func CursorIncrease(cursor int, length int) int {
	if cursor >= length-1 {
		return length - 1
	}

	return cursor + 1
}

func CursorDecrease(cursor int) int {
	if cursor <= 0 {
		return 0
	}

	return cursor - 1
}
