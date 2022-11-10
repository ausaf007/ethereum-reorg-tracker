package reorgtracker

import (
	"errors"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	ArrNotEqual      = errors.New("array not equal error in compareHeaderArr()")
	NoCommonElements = errors.New("no common elements in compareHeaderArr()")
)

// findEqualPos finds the extent of overlap between 2 header arrays
// It basically finds, index of header in arr2 that is equal to the last element of arr1.
/*
For eg:
Input:
arr1 =	2 3 4 5 6 7
arr2 =	  3 4 5 6 7 8

Note: the headers are represented by block numbers in this example

In this the 4th index of arr2 = arr1[last Index]
Therefore, output will return 4
*/
func findEqualPos(arr1 []*types.Header, arr2 []*types.Header, n int) (int, error) {
	equalPos := 0
	for i := n - 2; i >= 0; i-- {
		if blockNum(arr2[i]) == blockNum(arr1[n-1]) {
			equalPos = i
			break
		}
		if i == 0 {
			return 0, NoCommonElements
		}
	}

	return equalPos, nil
}

// getDiscardedArr returns all the (Discarded) blocks in Ephemeral Fork, if it exists.
/*
getDiscardedArr takes in equalPos from equalPos(), ie, the extent of overlap, and then compares the overlapping
block headers, and returns overlapping block headers, if found unequal.
For eg:
Input:
	arr1=	2  3  4  5  6  7
	arr2=	   3  4  5  6' 7' 8'

Note: the headers are represented by block numbers in this example. Note that 6' is not same as 6,
because a block may have same block number, but different block hash(ie block contents)

Output:
	6 7

*/
func getDiscardedArr(arr1 []*types.Header, arr2 []*types.Header, n int, equalPos int) []*types.Header {
	var discardedArray []*types.Header
	j := n - 1
	for i := equalPos; i >= 0; i-- {
		if isEqual(arr1[j], arr2[i]) {
			// Normal behaviour, no forks found.
		} else {
			// Found Ephemeral Fork!
			discardedArray = append(discardedArray, arr1[j])
		}
		j--
	}
	return discardedArray
}

// compareHeaderArr takes in arr1 (currHeaderArr) and arr2 (nextHeaderArr) and
// returns (Discarded) blocks in Ephemeral Fork, if it exists.
func compareHeaderArr(arr1 []*types.Header, arr2 []*types.Header) ([]*types.Header, error) {
	n := len(arr1)
	if n != len(arr2) {
		return nil, ArrNotEqual
	}
	// n = len of arr1 = len of arr2

	equalPos, err := findEqualPos(arr1, arr2, n)
	if err != nil {
		return nil, err
	}
	discardedArray := getDiscardedArr(arr1, arr2, n, equalPos)
	// If program reaches here, it means no error encountered
	return discardedArray, nil
}
