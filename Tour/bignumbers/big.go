package bignumbers

import (
	"fmt"
	"math/big"
)

func Biggy() {
	fmt.Println("Result is ", new(big.Int).Lsh(big.NewInt(1), 100))
}
