package reorgtracker

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"time"
)

var MinusOne = big.NewInt(-1)
var NumTooShort = errors.New("number of recent headers required too short in fetchRecentHeaders()")

// fetchRecentHeaders fetches given number of recent block headers,
// and returns it as header array
func fetchRecentHeaders(client *ethclient.Client, num int) ([]*types.Header, error) {

	if num < 1 {
		return nil, NumTooShort
	}
	// create HeaderArray of length num
	headerArr := make([]*types.Header, num)

	currBlockNumUint, err := client.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}
	currBlockNum := uintToBigInt(currBlockNumUint)

	for i := num - 1; i >= 0; i-- {
		headerArr[i], err = client.HeaderByNumber(context.Background(), currBlockNum)
		if err != nil {
			return nil, err
		}
		// decrementing bigint by 1
		// currBlockNum--
		currBlockNum.Add(currBlockNum, MinusOne)

		time.Sleep(ShortPause)
	}

	// if program reaches here, it implies no errors encountered
	debugLogHeaderArray(headerArr)
	log.Debug("Printing Header Array Completed")
	return headerArr, nil
}

// fetchRecentHeaders1 fetches given number of recent block headers,
// but it keeps fetching until it finds a header with block number >= min
// and returns it as header array
func fetchRecentHeaders1(client *ethclient.Client, num int, min uint64) ([]*types.Header, error) {
	for true {
		headerArr, err := fetchRecentHeaders(client, num)

		if err != nil {
			return nil, err
		}
		if blockNum(headerArr[num-1]) >= min {
			return headerArr, err
		}

		log.Warn("Couldn't find a new block. Trying again.")
		time.Sleep(ShortPause)
	}

	// unreachable statement
	return nil, UnreachableStatement
}
