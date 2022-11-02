package leetcode

import "fmt"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	mid := n/2 + 1
	res := isBadVersion(mid)
	if res {
		return findNext(0, mid)
	} else {
		return findNext(mid+1, n)
	}
}

func findNext(f int, t int) int {
	if f == t {
		return f
	}
	c := (t-f)/2 + f
	a := isBadVersion(c)
	if a {
		return findNext(f, c)
	} else {
		return findNext(c+1, t)
	}
}

func isBadVersion(n int) bool {
	fmt.Println("Call to ", n, "result", n >= 4)
	return n >= 4
}
