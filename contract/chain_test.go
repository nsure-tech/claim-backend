package contract

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"
)

func TestChainId(t *testing.T) {
	id := GetChainId()
	fmt.Println("aaaaaaa")
	fmt.Println(id)
}

func TestClaim(t *testing.T) {
	fmt.Println("11111111111")
	sigHash, _, err := SignClaim("0x2e9475c282069675fFAc22a8cd5038E4DAC01634", "0x2e9475c282069675fFAc22a8cd5038E4DAC0", "123", "1", "1333333")
	fmt.Println(err)
	fmt.Println(sigHash.String())
}

func TestKKK(t *testing.T) {
	buyClaim, err := GetBuyClaim(big.NewInt(1000))
	fmt.Println("aaaaaaa")
	fmt.Println(err)
	fmt.Println(buyClaim)
	see, err := json.Marshal(buyClaim)
	fmt.Println(string(see))
}
