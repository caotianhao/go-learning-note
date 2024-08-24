package timelock_puzzle

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestGeneratePrime(t *testing.T) {
	x, _ := GeneratePrime(DefaultBits)
	log.Println("x =", x.Text(10))
}

func TestGenerateRandomBigInt(t *testing.T) {
	x, _ := GenerateRandomBigInt(DefaultBits)
	log.Println("x =", x.Text(10))
}

func TestGenerateTimeLockPuzzle(t *testing.T) {
	puzzleT, _ := GenerateRandomBigInt(100)
	log.Println("T =", puzzleT.Text(10))
	x, _ := GenerateRandomBigInt(100)
	log.Println("x =", x.Text(10))
	puzzle, _ := GenerateTimeLockPuzzle(DefaultBits, puzzleT, x)
	puzzleBytes, _ := json.Marshal(puzzle)
	log.Println("puzzle =", string(puzzleBytes))
}

func TestSolveByPrivate(t *testing.T) {
	puzzleT, _ := GenerateRandomBigInt(10)
	log.Println("T =", puzzleT.Text(10))
	x, _ := GenerateRandomBigInt(10)
	log.Println("x =", x.Text(10))
	puzzle, _ := GenerateTimeLockPuzzle(10, puzzleT, x)
	puzzleBytes, _ := json.Marshal(puzzle)
	log.Println("puzzle =", string(puzzleBytes))
	y, _ := SolveByPrivate(puzzle)
	fmt.Println("y =", y.Text(10))
}

func TestSolveByPublic(t *testing.T) {
	puzzleT, _ := GenerateRandomBigInt(3)
	log.Println("T =", puzzleT.Text(10))
	x, _ := GenerateRandomBigInt(6)
	log.Println("x =", x.Text(10))
	puzzle, _ := GenerateTimeLockPuzzle(6, puzzleT, x)
	puzzleBytes, _ := json.Marshal(puzzle)
	log.Println("puzzle =", string(puzzleBytes))
	y, _ := SolveByPrivate(puzzle)
	fmt.Println("y by private =", y.Text(10))
	yByPublic, _ := SolveByPublic(puzzle)
	fmt.Println("y by public =", yByPublic.Text(10))
}
