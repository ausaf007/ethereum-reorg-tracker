package main

import (
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	tracker "reorg-tracker/reorgtracker"
	"time"
)

const MediumPause = 5 * time.Second

// loadFlags is a helper function that sets the verbosity level of the program through user specified flags
func loadFlags() bool {
	isVerbose := flag.Bool("verbose", false, "Specifies verbosity of logs. True means Debug Level. "+
		"False means Info Level")
	flag.Parse()
	if *isVerbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	return *isVerbose
}

// getClientLink is a helper function, it loads CLIENT_LINK from .env file
// If the program fails to load client link from .env file,
// then the ClientLink is taken as an input from user
func getClientLink() string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Warn("Error: Unable to read ClientLink from .env file."+
			"Error: ", err)
		fmt.Println("Please enter Eth ClientLink:")
		return getStringInput()
	}
	return viper.GetString("CLIENT_LINK")
}

// getStringInput is a helper function that takes string input from the user
func getStringInput() string {
	var input string
	for true {
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Invalid Input. Please try again.")
			continue
		}
		break
	}
	return input
}

// establishConnection is a helper function which keeps on running until a
// successful connection to EthClient is made
func establishConnection(ClientLink string) *ethclient.Client {
	var client *ethclient.Client

	// This loop keeps running until a successful connection to EthClient is established
	for true {
		var err error
		client, err = ethclient.Dial(ClientLink)
		if err != nil {
			log.Warn("Encountered Error Connecting to ETH Client. Error: ", err)
			log.Warn("Trying again...")
			ClientLink = getClientLink()

			// This pause is taken before trying again to connect with the ETH Client,
			// A pause is taken to prevent sending too many requests to the remote eth client
			time.Sleep(MediumPause)
			continue
		}
		err = nil
		log.Debug("Successfully connected to EthClient.")
		break
	}
	return client
}

// gracefulShutdown is a helper function that stops the program
// When user presses Ctrl+C, the connection is closed and the program closes
func gracefulShutdown(client *ethclient.Client) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	go func() {
		<-sigs
		client.Close()
		fmt.Println("Chain-Reorg Tracker stopped successfully.")
		os.Exit(0)
	}()
}

// startTacking is a helper function which calls tracker.StartTracking()
func startTacking(client *ethclient.Client, isVerbose bool) {
	for true {
		err := tracker.StartTracking(client, isVerbose)
		if err != nil {
			log.Warn("Encountered Error Tracking Chain ReOrg. Error: ", err)
			log.Warn("Restarting the tracker in 5 seconds.")

			// This pause is taken before trying again to start the tracker,
			// A pause is taken to prevent sending too many requests to the remote eth client
			// through StartTracking method
			time.Sleep(MediumPause)
			continue
		}
	}
}

// driver code
// This shows usage of reorg-tracker package
func main() {
	isVerbose := loadFlags()
	ClientLink := getClientLink()
	client := establishConnection(ClientLink)
	gracefulShutdown(client)
	startTacking(client, isVerbose)
}
