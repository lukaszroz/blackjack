# blackjack
Simple Blackjack HTTP server with RESTful JSON API written in Go

## Install Dependencies
`go get -u github.com/pressly/chi`

## Build Project
`go build`

## Run tests
`go test`

## Run
`./blackjack`

## API
[Examples](EXAMPLE.md)
### Game JSON
```JSON
{
  "Player": {
    "Cards": [
      {
        "Name": "Ten",
        "Suit": "Spades",
        "HighValue": 10,
        "LowValue": 10
      },
      {
        "Name": "Ten",
        "Suit": "Diamonds",
        "HighValue": 10,
        "LowValue": 10
      }
    ],
    "Score": 20
  },
  "Dealer": {
    "Cards": [
      {
        "Name": "Jack",
        "Suit": "Clubs",
        "HighValue": 10,
        "LowValue": 10
      }
    ],
    "Score": 10
  },
  "ID": 0,
  "IsFinished": false,
  "IsTie": false,
  "HasPlayerWon": false
}
```
#### List games
**GET** /games  
**response** array of games
#### Create a game
**POST** /games  
**response** a game
#### Retrieve a game
**GET** /games/:ID  
**response** a game  
If there is no game with provided ID  
**error** 404 Not found
#### Ask for another card (hit)
**POST** /games/:ID/hit  
**response** a game  
If game is finished  
**error** 400 Game is finished
#### Hold your cards and end the game (stand)
**POST** /games/:ID/stand  
**response** a game  
If game is finished  
**error** 400 Game is finished 
