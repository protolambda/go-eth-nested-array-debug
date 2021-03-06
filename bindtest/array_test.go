package bindtest

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"math/big"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
	"fmt"
)

func TestNestedArray(t *testing.T) {
	// Generate a new random account and a funded simulator
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	sim := backends.NewSimulatedBackend(core.GenesisAlloc{auth.From: {Balance: big.NewInt(10000000000)}})

	//deploy the test contract
	_, _, testContract, err := DeployDeeplyNestedArray(auth, sim)
	if err != nil {
		t.Fatalf("Failed to deploy test contract: %v", err)
	}

	// Finish deploy.
	sim.Commit()

	//Create coordinate-filled array, for testing purposes.
	testArr := [5][4][3]uint64{}
	for i := 0; i < 5; i++ {
		testArr[i] = [4][3]uint64{}
		for j := 0; j < 4; j++ {
			testArr[i][j] = [3]uint64{}
			for k := 0; k < 3; k++ {
				//pack the coordinates, each array value will be unique, and can be validated easily.
				testArr[i][j][k] = uint64(i) << 16 | uint64(j) << 8 | uint64(k)
			}
		}
	}

	if _, err := testContract.StoreDeepUintArray(&bind.TransactOpts{
		From: auth.From,
		Signer: auth.Signer,
	}, testArr); err != nil {
		t.Fatalf("Failed to store nested array in test contract: %v", err)
	}

	sim.Commit()

	retrievedArr, err := testContract.RetrieveDeepArray(&bind.CallOpts{
		From: auth.From,
		Pending: false,
	})
	if err != nil {
		t.Fatalf("Failed to retrieve nested array from test contract: %v", err)
	}

	//quick check to see if contents were copied
	// (See accounts/abi/unpack_test.go for more extensive testing)
	if retrievedArr[4][3][2] != testArr[4][3][2] {
		t.Fatalf("Retrieved value does not match expected value! got: %d, expected: %d. %v", retrievedArr[4][3][2], testArr[4][3][2], err)
	}

}

func printContentComparison(testArr [5][4][3]uint64, retrievedArr [5][4][3]uint64) {

	//comparing every single value:
	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 3; k++ {
				expected := testArr[i][j][k]
				got := retrievedArr[i][j][k]

				e_x := (expected >> 16) & 0xff
				e_y := (expected >> 8) & 0xff
				e_z := (expected) & 0xff

				g_x := (got >> 16) & 0xff
				g_y := (got >> 8) & 0xff
				g_z := (got) & 0xff

				fmt.Printf("(%d; %d; %d) expected: %d %d %d got: %d %d %d\n", i, j, k, e_x, e_y, e_z, g_x, g_y, g_z)

				//don't panic or t.Fatalf, print complete comparison.
			}
		}
	}
}