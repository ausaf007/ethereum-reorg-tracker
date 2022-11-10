<h1 align="center">Chain Reorg Tracker</h1>

<h3 align="center">Command Line Tool To Detect Chain Reorganization in Ethereum Blockchain </h3>

<!-- TABLE OF CONTENTS -->
<details open>
  <summary>Table of Contents</summary>
  <ul>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#tech-stack">Tech Stack</a></li>
    <li><a href="#algorithm">Algorithm</a></li>
    <li><a href="#prerequisites">Prerequisites</a></li>
    <li><a href="#how-to-use">How to use?</a></li>
  </ul>
</details>

## About The Project

Command-line application to track and detect chain reorganization in the Ethereum blockchain. In case a chain re-org is detected, the discarded blocks from the ephemetal forks is printed out.

## Tech Stack

[![](https://img.shields.io/badge/Built_with-Go-green?style=for-the-badge&logo=Go)](https://go.dev/)

## Algorithm

It works by querying the Eth-client for the last 9 blocks and storing the blocks in an array(**currentArray**), then a pause of 16 seconds is taken and then the last 9 blocks are retrieved again and stored in an array, **nextArray**. These two arrays are then compared to detect any chain reorgs. And then, the **currentArray** is overwritten with **nexArray** and the process continues until the program is terminated by the user, by hitting Cntrl+C.

Check out `reorgtracker/reorgtracker.go` to learn more about the algorithm.

## Prerequisites

Download and install [Golang 1.19](https://go.dev/doc/install) (or higher).  

## How To Use?

1. Navigate to `ethereum-reorg-tracker/`:
   ``` 
   cd /path/to/folder/ethereum-reorg-tracker/
   ```
2. Open `.env` file and fill in the `CLIENT_LINK` field. This is useful to connect to the Ethereum Node.
3. Get dependencies:
   ``` 
   go mod tidy
   ```
4. Run the app:
   ``` 
   go run main.go 
   # use "--verbose" flag to get additional logs
   go run main.go --verbose 
   ```
5. CD into `reorgtracker/` to run tests:  
   ``` 
   cd reorgtracker/
   go test
   ```
   
Thank you!