// What is the output of this program?
package main

import "fmt"

func appendSlice() {
	a := make([]int, 3, 4)
	a[0], a[1], a[2] = 0, 1, 2

	b := append(a, 66)
	b[0] = 6
	c := append(a, 77)
	c[0] = 7
	d := append(a, 88, 99)
	d[0] = 9

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

// Line number 10 appends 66 to slice a. The slice a has a capacity of 4, so it can accommodate the new element. The slice b is a new slice that points to a new underlying array. The slice b has the same length as a, but it has a different capacity. The slice b is a copy of a, so the value of a is not changed. The value of b is changed to [0 1 2 66].
// Line numer 11 will change the value of the first element of slice b to 6. As b and a share the same underlying array, the value of the first element of a will also change to 6.
// Line number 12 appends 77 to slice a. The slice a has a capacity of 4, so it can accommodate the new element. The slice c is a new slice that points to a new underlying array. The slice c has the same length as a, but it has a different capacity. The slice c is a copy of a, so the value of a is not changed. The value of c is changed to [6 1 2 77].
// Line number 13 will change the value of the first element of slice c to 7. As c and a share the same underlying array, the value of the first element of a will also change to 7.
// Line number 14 appends 88 and 99 to slice a. The slice a has a capacity of 4, and as we are adding multiple elements at the same time it will automatically update the capacity. The slice d is a new slice that points to a new underlying array. The slice d has the same length as a, but it has a different capacity. The slice d is a copy of a, so the value of a is not changed. The value of d is changed to [7 1 2 88 99].
// Line number 15 will change the value to 9. As d is a new array only the value of d will change to [9 1 2 88 99].
// The output of this program will be:
// [7 1 2]
// [7 1 2 77]
// [7 1 2 77]
// [9 1 2 88 99]
