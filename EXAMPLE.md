# Examples
### Create a game
`curl -X POST http://localhost:3333/games`
```JSON
{
  "Player": {
    "Cards": [
      {
        "Name": "Four",
        "Suit": "Hearts",
        "HighValue": 4,
        "LowValue": 4
      },
      {
        "Name": "Two",
        "Suit": "Clubs",
        "HighValue": 2,
        "LowValue": 2
      }
    ],
    "Score": 6
  },
  "Dealer": {
    "Cards": [
      {
        "Name": "Five",
        "Suit": "Spades",
        "HighValue": 5,
        "LowValue": 5
      }
    ],
    "Score": 5
  },
  "ID": 3,
  "IsFinished": false,
  "IsTie": false,
  "HasPlayerWon": false
}
```

### Hit
`curl -X POST http://localhost:3333/games/3/stand`
```JSON
{
  "Player": {
    "Cards": [
      {
        "Name": "Four",
        "Suit": "Hearts",
        "HighValue": 4,
        "LowValue": 4
      },
      {
        "Name": "Two",
        "Suit": "Clubs",
        "HighValue": 2,
        "LowValue": 2
      },
      {
        "Name": "Three",
        "Suit": "Spades",
        "HighValue": 3,
        "LowValue": 3
      }
    ],
    "Score": 9
  },
  "Dealer": {
    "Cards": [
      {
        "Name": "Five",
        "Suit": "Spades",
        "HighValue": 5,
        "LowValue": 5
      }
    ],
    "Score": 5
  },
  "ID": 3,
  "IsFinished": false,
  "IsTie": false,
  "HasPlayerWon": false
}
```
### Stand
`curl -X POST http://localhost:3333/games/3/stand`
```JSON
{
  "Player": {
    "Cards": [
      {
        "Name": "Four",
        "Suit": "Hearts",
        "HighValue": 4,
        "LowValue": 4
      },
      {
        "Name": "Two",
        "Suit": "Clubs",
        "HighValue": 2,
        "LowValue": 2
      },
      {
        "Name": "Three",
        "Suit": "Spades",
        "HighValue": 3,
        "LowValue": 3
      }
    ],
    "Score": 9
  },
  "Dealer": {
    "Cards": [
      {
        "Name": "Five",
        "Suit": "Spades",
        "HighValue": 5,
        "LowValue": 5
      },
      {
        "Name": "Five",
        "Suit": "Hearts",
        "HighValue": 5,
        "LowValue": 5
      },
      {
        "Name": "Four",
        "Suit": "Spades",
        "HighValue": 4,
        "LowValue": 4
      },
      {
        "Name": "King",
        "Suit": "Diamonds",
        "HighValue": 10,
        "LowValue": 10
      }
    ],
    "Score": 24
  },
  "ID": 3,
  "IsFinished": true,
  "IsTie": false,
  "HasPlayerWon": true
}
```

##Errors
### Wrong game ID
`curl -iX GET http://localhost:3333/games/5`
```
HTTP/1.1 404 Not Found
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Sun, 26 Mar 2017 19:00:34 GMT
Content-Length: 10

Not Found
```
### Hit or Stand on finished game
`curl -iX POST http://localhost:3333/games/3/hit`  
`curl -iX POST http://localhost:3333/games/3/stand`
```
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Sun, 26 Mar 2017 19:01:52 GMT
Content-Length: 17

Game is finished
```