package contract

import (
	"github.com/ethereum/go-ethereum/signer/core"
)

var withDrawType = core.Types{
	"EIP712Domain": {
		{
			Name: "name",
			Type: "string",
		},
		{
			Name: "version",
			Type: "string",
		},
		{
			Name: "chainId",
			Type: "uint256",
		},
		{
			Name: "verifyingContract",
			Type: "address",
		},
	},
	"Withdraw": {
		{
			Name: "account",
			Type: "address",
		},
		{
			Name: "amount",
			Type: "uint256",
		},
		{
			Name: "nonce",
			Type: "uint256",
		},
		{
			Name: "deadline",
			Type: "uint256",
		},
	},
}

const withdrawPrimaryType = "Withdraw"
const claimPrimaryType = "Claim"

var claimType = core.Types{
	"EIP712Domain": {
		{
			Name: "name",
			Type: "string",
		},
		{
			Name: "version",
			Type: "string",
		},
		{
			Name: "chainId",
			Type: "uint256",
		},
		{
			Name: "verifyingContract",
			Type: "address",
		},
	},
	"Claim": {
		{
			Name: "account",
			Type: "address",
		},
		{
			Name: "currency",
			Type: "uint256",
		},
		{
			Name: "amount",
			Type: "uint256",
		},
		{
			Name: "nonce",
			Type: "uint256",
		},
		{
			Name: "deadline",
			Type: "uint256",
		},
	},
}
