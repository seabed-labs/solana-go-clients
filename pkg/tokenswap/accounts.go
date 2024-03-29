// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package tokenswap

import ag_solanago "github.com/gagliardetto/solana-go"

type TokenSwap struct {
	Version                     uint8
	IsInitialized               uint8
	BumpSeed                    uint8
	TokenProgramId              ag_solanago.PublicKey
	TokenAccountA               ag_solanago.PublicKey
	TokenAccountB               ag_solanago.PublicKey
	TokenPool                   ag_solanago.PublicKey
	MintA                       ag_solanago.PublicKey
	MintB                       ag_solanago.PublicKey
	FeeAccount                  ag_solanago.PublicKey
	TradeFeeNumerator           uint64
	TradeFeeDenominator         uint64
	OwnerTradeFeeNumerator      uint64
	OwnerTradeFeeDenominator    uint64
	OwnerWithdrawFeeNumerator   uint64
	OwnerWithdrawFeeDenominator uint64
	HostFeeNumerator            uint64
	HostFeeDenominator          uint64
}
