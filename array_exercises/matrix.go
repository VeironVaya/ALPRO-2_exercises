package main

import "fmt"

func main() {
	var n, m, p, q int
	var i, j, k int
	var a [20][20]int
	var b [20][20]int
	var result [20][20]int

	// Input Matrix A
	fmt.Print("Matrix A Size (n) :")
	fmt.Scan(&n)
	fmt.Print("Matrix A Size (m) :")
	fmt.Scan(&m)

	for i = 0; i < n; i++ {
		for j = 0; j < m; j++ {
			fmt.Printf("Matrix A[%d][%d] = ", i, j)
			fmt.Scan(&a[i][j])
		}
	}
	fmt.Println()

	// Input Matrix B
	fmt.Print("Matrix B Size (p) :")
	fmt.Scan(&p)
	fmt.Print("Matrix B Size (q) :")
	fmt.Scan(&q)

	for i = 0; i < p; i++ {
		for j = 0; j < q; j++ {
			fmt.Printf("Matrix B[%d][%d] = ", i, j)
			fmt.Scan(&b[i][j])
		}
	}
	fmt.Println()

	// Matrix Addition & Subtraction
	if n == p && m == q {
		for i = 0; i < n; i++ {
			for j = 0; j < m; j++ {
				result[i][j] = a[i][j] + b[i][j]
			}
		}
		fmt.Println("Matrix A + B:")
		printMatrix(result, n, m)

		for i = 0; i < n; i++ {
			for j = 0; j < m; j++ {
				result[i][j] = a[i][j] - b[i][j]
			}
		}
		fmt.Println("Matrix A - B:")
		printMatrix(result, n, m)
	} else {
		fmt.Println("Addition/Subtraction not possible. Matrices must be the same size.")
	}

	// Matrix Multiplication (n x m * p x q -> result n x q)
	if m == p {
		for i = 0; i < n; i++ {
			for j = 0; j < q; j++ {
				result[i][j] = 0
				for k = 0; k < m; k++ {
					result[i][j] += a[i][k] * b[k][j]
				}
			}
		}
		fmt.Println("Matrix A * B:")
		printMatrix(result, n, q)
	} else {
		fmt.Println("Matrix multiplication not possible. Columns of A must equal rows of B.")
	}

	// Transpose of Matrix A
	for i = 0; i < n; i++ {
		for j = 0; j < m; j++ {
			result[j][i] = a[i][j]
		}
	}
	fmt.Println("Transpose of Matrix A:")
	printMatrix(result, m, n)

	// Transpose of Matrix B
	for i = 0; i < p; i++ {
		for j = 0; j < q; j++ {
			result[j][i] = b[i][j]
		}
	}
	fmt.Println("Transpose of Matrix B:")
	printMatrix(result, q, p)
}

// Helper function to print a matrix
func printMatrix(matrix [20][20]int, rows int, cols int) {
	var i, j int
	for i = 0; i < rows; i++ {
		for j = 0; j < cols; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}
