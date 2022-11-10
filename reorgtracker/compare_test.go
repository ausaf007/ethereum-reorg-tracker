package reorgtracker

import (
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"math/big"
	"reflect"
	"testing"
)

func TestCompareHeaderArrCase1(t *testing.T) {
	log.SetLevel(log.InfoLevel)

	header1 := &types.Header{
		Difficulty: big.NewInt(10000000000),
		Number:     big.NewInt(1),
		GasLimit:   8_000_000,
		GasUsed:    8_000_000,
		Time:       555,
		Extra:      make([]byte, 32),
	}
	header2 := &types.Header{
		Difficulty: big.NewInt(10000000000),
		Number:     big.NewInt(2),
		GasLimit:   8_000_000,
		GasUsed:    8_000_000,
		Time:       555,
		Extra:      make([]byte, 32),
	}
	header2a := &types.Header{
		Difficulty: big.NewInt(4864548543),
		Number:     big.NewInt(2),
		GasLimit:   8_000_000,
		GasUsed:    8_000_000,
		Time:       555,
		Extra:      make([]byte, 32),
	}
	header3a := &types.Header{
		Difficulty: big.NewInt(86455456),
		Number:     big.NewInt(3),
		GasLimit:   8_000_000,
		GasUsed:    8_000_000,
		Time:       555,
		Extra:      make([]byte, 32),
	}

	// 1  2
	//    2a 3a
	arr1 := []*types.Header{header1, header2}
	arr2 := []*types.Header{header2a, header3a}

	arr, err := compareHeaderArr(arr1, arr2)
	if err != nil {
		t.Errorf("\nEncountered Error=%s", err)
	}

	if !isEqual(arr[0], header2) {
		t.Errorf("\nTest failed. Headers not equal as expected.")
	}

	arrDiscarded := getDiscardedArr(arr1, arr2, 2, 0)
	if !isEqual(arrDiscarded[0], header2) {
		t.Errorf("\nTest failed. Headers not equal as expected.")
	}

	equalPos, err := findEqualPos(arr1, arr2, 2)
	if err != nil {
		t.Errorf("\nEncountered Error=%s", err)
	}
	if equalPos != 0 {
		t.Errorf("\nTest failed. equalPos expected 0 but found %d", equalPos)
	}

}

func TestCompareHeaderArrCase2(t *testing.T) {
	log.SetLevel(log.InfoLevel)

	header1 := &types.Header{
		Difficulty: big.NewInt(10000000000),
		Number:     big.NewInt(1),
		GasLimit:   8_000_000,
		GasUsed:    8_000_000,
		Time:       555,
		Extra:      make([]byte, 32),
	}
	header2 := &types.Header{
		Difficulty: big.NewInt(10000000000),
		Number:     big.NewInt(2),
		GasLimit:   8_000_000,
		GasUsed:    8_000_000,
		Time:       555,
		Extra:      make([]byte, 32),
	}
	header3 := &types.Header{
		Difficulty: big.NewInt(10000000000),
		Number:     big.NewInt(3),
		GasLimit:   8_000_000,
		GasUsed:    8_000_000,
		Time:       555,
		Extra:      make([]byte, 32),
	}
	header4 := &types.Header{
		Difficulty: big.NewInt(10000000000),
		Number:     big.NewInt(4),
		GasLimit:   8_000_000,
		GasUsed:    8_000_000,
		Time:       555,
		Extra:      make([]byte, 32),
	}
	header4a := &types.Header{
		Difficulty: big.NewInt(5465561),
		Number:     big.NewInt(4),
		GasLimit:   8_000_000,
		GasUsed:    8_000_000,
		Time:       555,
		Extra:      make([]byte, 32),
	}
	header5a := &types.Header{
		Difficulty: big.NewInt(86455456),
		Number:     big.NewInt(5),
		GasLimit:   8_000_000,
		GasUsed:    8_000_000,
		Time:       555,
		Extra:      make([]byte, 32),
	}
	header6a := &types.Header{
		Difficulty: big.NewInt(86455456),
		Number:     big.NewInt(6),
		GasLimit:   8_000_000,
		GasUsed:    8_000_000,
		Time:       555,
		Extra:      make([]byte, 32),
	}
	header7a := &types.Header{
		Difficulty: big.NewInt(86455456),
		Number:     big.NewInt(7),
		GasLimit:   8_000_000,
		GasUsed:    8_000_000,
		Time:       555,
		Extra:      make([]byte, 32),
	}

	// 1  2  3  4
	//			4a 5a 6a 7a
	arr1 := []*types.Header{header1, header2, header3, header4}
	arr2 := []*types.Header{header4a, header5a, header6a, header7a}

	arr, err := compareHeaderArr(arr1, arr2)
	if err != nil {
		t.Errorf("\nEncountered Error=%s", err)
	}

	if !isEqual(arr[0], header4) {
		t.Errorf("\nTest failed. Headers not equal as expected.")
	}

	arrDiscarded := getDiscardedArr(arr1, arr2, 4, 0)
	if !isEqual(arrDiscarded[0], header4) {
		t.Errorf("\nTest failed. Headers not equal as expected.")
	}

	equalPos, err := findEqualPos(arr1, arr2, 4)
	if err != nil {
		t.Errorf("\nEncountered Error=%s", err)
	}
	if equalPos != 0 {
		t.Errorf("\nTest failed. equalPos expected 0 but found %d", equalPos)
	}

}

func TestCompareHeaderArrCase3(t *testing.T) {
	log.SetLevel(log.InfoLevel)

	header1 := &types.Header{
		Difficulty: big.NewInt(10000000000),
		Number:     big.NewInt(1),
		GasLimit:   8_000_000,
		GasUsed:    8_000_000,
		Time:       555,
		Extra:      make([]byte, 32),
	}
	header2 := &types.Header{
		Difficulty: big.NewInt(10000000000),
		Number:     big.NewInt(2),
		GasLimit:   8_000_000,
		GasUsed:    8_000_000,
		Time:       555,
		Extra:      make([]byte, 32),
	}
	header3 := &types.Header{
		Difficulty: big.NewInt(10000000000),
		Number:     big.NewInt(3),
		GasLimit:   8_000_000,
		GasUsed:    8_000_000,
		Time:       555,
		Extra:      make([]byte, 32),
	}
	header4 := &types.Header{
		Difficulty: big.NewInt(10000000000),
		Number:     big.NewInt(4),
		GasLimit:   8_000_000,
		GasUsed:    8_000_000,
		Time:       555,
		Extra:      make([]byte, 32),
	}

	header2a := &types.Header{
		Difficulty: big.NewInt(645416),
		Number:     big.NewInt(2),
		GasLimit:   8_000_000,
		GasUsed:    8_000_000,
		Time:       555,
		Extra:      make([]byte, 32),
	}
	header3a := &types.Header{
		Difficulty: big.NewInt(86455456),
		Number:     big.NewInt(3),
		GasLimit:   8_000_000,
		GasUsed:    8_000_000,
		Time:       555,
		Extra:      make([]byte, 32),
	}
	header4a := &types.Header{
		Difficulty: big.NewInt(86455456),
		Number:     big.NewInt(4),
		GasLimit:   8_000_000,
		GasUsed:    8_000_000,
		Time:       555,
		Extra:      make([]byte, 32),
	}
	header5a := &types.Header{
		Difficulty: big.NewInt(86455456),
		Number:     big.NewInt(5),
		GasLimit:   8_000_000,
		GasUsed:    8_000_000,
		Time:       555,
		Extra:      make([]byte, 32),
	}

	// 1  2  3  4
	//    2a 3a 4a 5a
	arr1 := []*types.Header{header1, header2, header3, header4}
	arr2 := []*types.Header{header2a, header3a, header4a, header5a}

	arr, err := compareHeaderArr(arr1, arr2)
	if err != nil {
		t.Errorf("\nEncountered Error=%s", err)
	}

	if reflect.DeepEqual(arr, []*types.Header{header2, header3, header4}) {
		t.Errorf("\nTest failed. Headers not equal as expected.")
	}

	arrDiscarded := getDiscardedArr(arr1, arr2, 4, 2)
	if reflect.DeepEqual(arrDiscarded, []*types.Header{header2, header3, header4}) {
		t.Errorf("\nTest failed. Headers not equal as expected.")
	}

	equalPos, err := findEqualPos(arr1, arr2, 4)
	if err != nil {
		t.Errorf("\nEncountered Error=%s", err)
	}
	if equalPos != 2 {
		t.Errorf("\nTest failed. equalPos expected 2 but found %d", equalPos)
	}

}
