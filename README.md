This is simple code breaking game with backend in golang. I made this project mainly to learn web development in go.

### Rules

The rules of the game are simple. There is a four digit code that you have to guess. Each digit of the code is unique and first digit is not zero.
After you make a guess, you will be replied with number of bulls and cows.
A bull is a digit which is in the same position in both the secret code and your guess, whereas a cow is a digit which is present both in the guess and in the secret number, but in a different position.

### Setup

Clone the repository and run

`go run *.go`

Then open 127.0.0.1:9000 in your browser
