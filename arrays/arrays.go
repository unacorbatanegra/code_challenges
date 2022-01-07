package main

/* Problem:
Given an array a that contains only numbers in the range from 1 to a.length,
find the first duplicate number for which the second occurrence has the minimal index.
In other words, if there are more than 1 duplicated numbers,
return the number for which the second occurrence has a smaller
index than the second occurrence of the other number does.
If there are no such elements, return -1.
*/
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

/* Problem:
Given a string s consisting of small English letters,
find and return the first instance of a non-repeating
character in it.
If there is no such character, return '_'.
*/
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

/* Problem:
You are given an n x n 2D matrix that represents an image. Rotate the image by 90 degrees (clockwise).
*/
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

/* Problem:
Sudoku is a number-placement puzzle. The objective is to fill a 9 × 9 grid
with numbers in such a way that each column,
each row, and each of the nine 3 × 3 sub-grids that
compose the grid all contain all of the numbers
from 1 to 9 one time.
Implement an algorithm that will check whether the given grid of numbers
represents a valid Sudoku puzzle according to the layout rules described above.
Note that the puzzle represented by grid does not have to be solvable.
*/
func sudoku(grid [][]string) bool {

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			current := grid[row][col]
			sR := row + 3
			sC := col + 3
			if sR%3 == 0 && sC%3 == 0 {
				m := make(map[string]int)

				for iR := row; iR < sR; iR++ {
					for iC := col; iC < sC; iC++ {
						v := grid[iR][iC]
						if v == "." {
							continue
						}
						if _, ok := m[v]; !ok {
							m[v]++
						} else {
							return false
						}
					}
				}
			}
			if current == "." {
				continue
			}
			for r := row + 1; r < len(grid); r++ {
				m := make(map[string]int)
				cc := grid[r][col]

				if cc == "." {
					continue
				}
				if cc == current {
					return false
				}
				if _, ok := m[cc]; !ok {
					m[cc]++
				} else {
					return false
				}

			}
			for c := col + 1; c < len(grid); c++ {
				m := make(map[string]int)
				cc := grid[row][c]
				if cc == "." {
					continue
				}
				if cc == current {
					return false
				}
				if _, ok := m[cc]; !ok {
					m[cc]++
				} else {
					return false
				}

			}
		}
	}

	return false

}
