package main

import "strconv"

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

/* Problem:

A cryptarithm is a mathematical puzzle for which the goal is to find
the correspondence between letters and digits, such that the given
arithmetic equation consisting of letters holds true
when the letters are converted to digits.
You have an array of strings crypt, the cryptarithm,
and an an array containing the mapping of letters and digits, solution.
The array crypt will contain three non-empty strings
that follow the structure: [word1, word2, word3], which should
be interpreted as the word1 + word2 = word3 cryptarithm.
If crypt, when it is decoded by replacing all of the letters in
the cryptarithm with digits using the mapping in solution,
becomes a valid arithmetic equation containing no numbers with leading zeroes,
the answer is true. If it does not become a valid arithmetic solution, the answer is false.
Note that number 0 doesn't contain leading zeroes (while for example 00 or 0123 do).
*/

func isCryptSolution(crypt []string, solution [][]string) bool {
	m := make(map[string]string)

	for i, v := range solution {
		m[v[0]] = solution[i][1]
	}
	list := [3]int{}

	for i, v := range crypt {
		var s string
		if i == 0 && m[string(v[0])] == "0" && len(v) > 1 {
			return false
		}
		for c, l := range v {
			if c == 0 && m[string(l)] == "0" && len(v) > 1 {
				return false
			}
			s += m[string(l)]
		}
		tmp, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		list[i] = tmp

	}

	return list[0]+list[1] == list[2]
}
