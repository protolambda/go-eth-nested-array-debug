//go:generate go run "$GOPATH/src/github.com/ethereum/go-ethereum/cmd/abigen/main.go" --sol DeeplyNestedArray.sol --pkg bind_nested_array --out bind_nested_array/deeply_nested_array.go
package main

import "log"

func main() {
	log.Println("Test!")
}
