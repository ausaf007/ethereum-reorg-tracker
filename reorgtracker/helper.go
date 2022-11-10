package reorgtracker

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"math/big"
	"reflect"
)

// getLatestHeader returns the last header from the headerArray given (arr []*types.Header)
func getLatestHeader(arr []*types.Header) *types.Header {
	return arr[len(arr)-1]
}

// printHeaderArray prints the []*types.Header in descending order
func printHeaderArray(arr []*types.Header) {
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Println(formatHeader(arr[i]))
	}
}

// debugLogHeaderArray logs the []*types.Header in descending order on debug level
func debugLogHeaderArray(arr []*types.Header) {
	for i := len(arr) - 1; i >= 0; i-- {
		log.Debug(formatHeader(arr[i]))
	}
}

// formatHeader returns a well formatted string with block number & hash of the block header
func formatHeader(header *types.Header) string {
	return "Hash of " + header.Number.String() + " is " + header.Hash().Hex()
}

// uintToBigInt takes in uint64 and converts it to big.Int
func uintToBigInt(num uint64) *big.Int {
	return new(big.Int).SetUint64(num)
}

// blockNum returns block number from the header
func blockNum(header *types.Header) uint64 {
	return header.Number.Uint64()
}

// isEqual checks equality of 2 headers
func isEqual(header1 *types.Header, header2 *types.Header) bool {
	return reflect.DeepEqual(header1, header2)
}
