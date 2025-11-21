/*
  (25 points)

  While [][]int is the type of two dimensional integer slices,
  [n][m]int is the type of two dimensional arrays. We may
  represent an n*m matrix as a two dimensional array using
  the following declaration.

    var matrix [n][m]int = [n][m]int{}

  Now implement the matrix transpose operation by following
  the instructions below.

  1. Declare the constants `n` and `m` that determine the dimension.
     Set the values of `n` and `m` to 3 and 5, respectively.
  2. Declare a two dimensional array variable named `mat`, which
     stores the matrix to be transposed.
     Initialize `mat` with the following value.
       [n][m]int{ {2, 5, 7, 9, 1}, {2, 1, 6, 8, 4}, {7, 3, 4, 6, 2} }
     So `mat` represents the following matrix.
       [ 2 5 7 9 1
         2 1 6 8 4
         7 3 4 6 2 ]
  3. Declare another two dimensional array variable named `trans`,
     which is used to store the transposed matrix. The dimension
     of `trans` is m*n.
  4. Write nested for loops to assign values from `mat` to `trans`
     such that for indices i and j, trans[j][i] = mat[i][j]. After
     the nested for loops, `trans` should represent the following
     matrix.
       [ 2 2 7
         5 1 3
         7 6 4
         9 8 6
         1 4 2 ]
  5. At the end of the program, print the matrices `mat` and
     `trans`. (Just use `fmt.Println` to print the variables.)
*/

package main

import "fmt"

func main() {
	// --- <code begin> ---
	const n = 3
	const m = 5
	mat := [n][m]int{{2, 5, 7, 9, 1}, {2, 1, 6, 8, 4}, {7, 3, 4, 6, 2}}
	var trans [m][n]int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			trans[j][i] = mat[i][j]
		}
	}

	fmt.Println(mat)
	fmt.Println(trans)
	// --- <code end> ---
}
