//go:generate go run "$GOPATH/src/github.com/ethereum/go-ethereum/cmd/abigen/main.go" --sol TestContract.sol --pkg bind_nested_array --out bind_nested_array/test_contract.go
package main

import "log"

func main() {
	log.Println("Test!")
}
