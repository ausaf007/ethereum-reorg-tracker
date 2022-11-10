package reorgtracker

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"time"
)

// ArrLength specifies number of block headers stored in the array, as a temporary cache.
// Value chosen is 9, because after 8 blocks, new blocks are created,
// finality is confirmed. See detailed explanation at the end to learn more.
const ArrLength = 9

const (
	ShortPause = 200 * time.Millisecond
	LongPause  = 16 * time.Second
)

var UnreachableStatement = errors.New("this statement should never be reached")

// setVerbosity is a helper function that sets the verbosity level
// true => Debug Level and false => Info Level
func setVerbosity(isVerbose bool) {
	if isVerbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

// printEphemeralFork is a helper function that prints the discarded ephemeral fork
func printEphemeralFork(discardedArr []*types.Header) {
	fmt.Println("\nChain Re-Org Detected. The discarded blocks are {")
	printHeaderArray(discardedArr)
	fmt.Println("\n}")
}

// StartTracking starts tracking for chain-reorgs
// it prints chain reorgs, if it encounters one
func StartTracking(client *ethclient.Client, isVerbose bool) error {

	fmt.Println("Re-Org Tracker Started!")
	setVerbosity(isVerbose)

	currHeaderArr, err := fetchRecentHeaders(client, ArrLength)
	if err != nil {
		return err
	}

	for true {
		time.Sleep(LongPause)

		latestHeader := getLatestHeader(currHeaderArr)
		min := blockNum(latestHeader) + 1

		// fetchRecentHeaders1 makes sure the fetched headers have at least 1 new header
		// that is, blockNum(currHeaderArr[last elem]) < blockNum(nextHeaderArr[last elem])
		nextHeaderArr, err := fetchRecentHeaders1(client, ArrLength, min)
		if err != nil {
			return err
		}

		// discardedArr contains array of discarded headers
		// If len(discardedArr) == 0, it implies that no ephemeral blocks found,
		// that is, no chain re-org detected
		discardedArr, err := compareHeaderArr(currHeaderArr, nextHeaderArr)
		if err != nil {
			return err
		}
		if len(discardedArr) == 0 {
			latestHeader = getLatestHeader(nextHeaderArr)
			fmt.Println("No Ephemeral Fork found. Reached Block Height = ", blockNum(latestHeader))
		} else {
			printEphemeralFork(discardedArr)
		}
		currHeaderArr = nextHeaderArr
	}

	return UnreachableStatement
}

/*
------------------------------------
Working Mechanism of StartTracking()
------------------------------------

Note: In this explanation, by block we mean Block Headers, because Block Headers have all the relevant
information we need.

First StartTracking() fetches 9 (specified by ArrLength) recent blocks, and stores it in currHeaderArr[].
Between each of this fetching, a short pause is taken. After the entire 9 blocks are fetched, then
a long pause is taken.

So the flow goes like -> fetch 9 Blocks with small pauses -> long pause -> fetch 9 Blocks with small pauses
and so on

Note: These pauses are essential to ensure that the remote EthClient is not overloaded with too many requests.

After a long pause is taken, the new 9 Blocks are stored in nextHeaderArr[]. Typically nextHeaderArr[] will have
1 or 2 new blocks, because the long pause is of 16 seconds.

(for eg, if currHeaderArr[]={1,2,3,4,5,6,7,8,9}, then nextHeaderArr[] will be = {2,3,4,5,6,7,8,9,10} )

Now the currHeaderArr[] is compared to nextHeaderArr[] using compareHeaderArr() function
First the overlap between is calculated, for eg, if 1 new block is found, the overlap will be 8 (9-1=8)
ie: overlap will be = arrLength - (no. of new blocks fetched)

(for eg, the overlap in this example would be {2,3,4,5,6,7,8,9})

After that, the equality of the overlapping headers is checked, if they are all equal, then no ephemeral fork
are found.

If some blocks are not equal, then ephemeral forks are found, and the unequal blocks from the currHeaderArr
are the discarded blocks, and they are printed.

Then the currHeaderArr = nextHeaderArr, and the cycle repeats

*/
