package main

/// Problem:

/// Given an array a that contains only numbers in the range from 1 to a.length,
/// find the first duplicate number for which the second occurrence has the minimal index.
/// In other words, if there are more than 1 duplicated numbers,
/// return the number for which the second occurrence has a smaller
/// index than the second occurrence of the other number does.
/// If there are no such elements, return -1.
func FirstDuplicate(a []int) int {
	m := make(map[int]int)

	for _, v := range a {
		if _, ok := m[v]; !ok {
			m[v]++
		} else {
			return v
		}
	}
	return -1
}

/// Problem:

/// Given a string s consisting of small English letters,
/// find and return the first instance of a non-repeating
/// character in it.
/// If there is no such character, return '_'.
func FirstNotRepeatingCharacter(s string) string {

	m := make(map[string]int)
	for _, v := range s {
		m[string(v)]++
	}
	for _, v := range s {
		if m[string(v)] == 1 {
			return string(v)
		}

	}
	return "_"
}

/// Problem:
/// You are given an n x n 2D matrix that represents an image. Rotate the image by 90 degrees (clockwise).
func RotateImage(a [][]int) [][]int {
	b := make([][]int, len(a))
	for r := 0; r < (len(a)); r++ {
		b[r] = make([]int, len(a[r]))
		for c := 0; c < len(a); c++ {
			b[r][c] = a[(len(a)-1)-c][r]
		}
	}
	return b

}
