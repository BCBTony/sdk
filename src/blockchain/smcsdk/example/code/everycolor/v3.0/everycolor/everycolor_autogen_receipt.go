package everycolor

import (
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
)

// This file is auto generated by BCB-goland-plugin.
// Don't modified it

func (e *Everycolor) emitSetSecretSigner(newSecretSigner types.PubKey) {
	type setSecretSigner struct {
		NewSecretSigner types.PubKey `json:"newSecretSigner"`
	}

	e.sdk.Helper().ReceiptHelper().Emit(setSecretSigner{
		NewSecretSigner: newSecretSigner,
	})
}

func (e *Everycolor) emitSetSettings(settings map[string]Setting, betExpirationBlocks int64) {
	type setSettings struct {
		Settings            map[string]Setting `json:"settings"`
		BetExpirationBlocks int64              `json:"betExpirationBlocks"`
	}

	e.sdk.Helper().ReceiptHelper().Emit(setSettings{
		Settings:            settings,
		BetExpirationBlocks: betExpirationBlocks,
	})
}

func (e *Everycolor) emitSetRecvFeeInfos(infos RecvFeeInfo) {
	type setRecvFeeInfos struct {
		Infos RecvFeeInfo `json:"infos"`
	}

	e.sdk.Helper().ReceiptHelper().Emit(setRecvFeeInfos{
		Infos: infos,
	})
}

func (e *Everycolor) emitWithdrawFunds(tokenName string, beneficiary types.Address, withdrawAmount bn.Number) {
	type withdrawFunds struct {
		TokenName      string        `json:"tokenName"`
		Beneficiary    types.Address `json:"beneficiary"`
		WithdrawAmount bn.Number     `json:"withdrawAmount"`
	}

	e.sdk.Helper().ReceiptHelper().Emit(withdrawFunds{
		TokenName:      tokenName,
		Beneficiary:    beneficiary,
		WithdrawAmount: withdrawAmount,
	})
}

func (e *Everycolor) emitPlaceBet(tokenName string, amount bn.Number, betData []BetData, possibleWinAmount bn.Number, commitLastBlock, betCount int64, commit, signData []byte, refAddress types.Address) {
	type placeBet struct {
		TokenName         string        `json:"tokenName"`
		Amount            bn.Number     `json:"amount"`
		BetData           []BetData     `json:"betData"`
		PossibleWinAmount bn.Number     `json:"possibleWinAmount"`
		CommitLastBlock   int64         `json:"commitLastBlock"`
		BetCount          int64         `json:"betCount"`
		Commit            []byte        `json:"commit"`
		SignData          []byte        `json:"signData"`
		RefAddress        types.Address `json:"refAddress"`
	}

	e.sdk.Helper().ReceiptHelper().Emit(placeBet{
		TokenName:         tokenName,
		Amount:            amount,
		BetData:           betData,
		PossibleWinAmount: possibleWinAmount,
		CommitLastBlock:   commitLastBlock,
		BetCount:          betCount,
		Commit:            commit,
		SignData:          signData,
		RefAddress:        refAddress,
	})
}

func (e *Everycolor) emitSettleBet(reveal, commit []byte, winNumber string, startIndex, endIndex int64, amountOfWin, amountOfUnLock map[string]bn.Number, finished bool) {
	type settleBet struct {
		Reveal         []byte               `json:"reveal"`
		Commit         []byte               `json:"commit"`
		WinNumber      string               `json:"winNumber"`
		StartIndex     int64                `json:"startIndex"`
		EndIndex       int64                `json:"endIndex"`
		AmountOfWin    map[string]bn.Number `json:"amountOfWin"`
		AmountOfUnLock map[string]bn.Number `json:"amountOfUnLock"`
		Finished       bool                 `json:"finished"`
	}

	e.sdk.Helper().ReceiptHelper().Emit(settleBet{
		Reveal:         reveal,
		Commit:         commit,
		WinNumber:      winNumber,
		StartIndex:     startIndex,
		EndIndex:       endIndex,
		AmountOfWin:    amountOfWin,
		AmountOfUnLock: amountOfUnLock,
		Finished:       finished,
	})
}

func (e *Everycolor) emitWithdrawWin(commit []byte, amountOfWin, unLockAmount map[string]bn.Number) {
	type withdrawWin struct {
		Commit       []byte               `json:"commit"`
		AmountOfWin  map[string]bn.Number `json:"amountOfWin"`
		UnLockAmount map[string]bn.Number `json:"unLockAmount"`
	}

	e.sdk.Helper().ReceiptHelper().Emit(withdrawWin{
		Commit:       commit,
		AmountOfWin:  amountOfWin,
		UnLockAmount: unLockAmount,
	})
}

func (e *Everycolor) emitRefundBet(commit []byte, refundCount int64, unlockAmount map[string]bn.Number, finished bool) {
	type refundBet struct {
		Commit       []byte               `json:"commit"`
		RefundCount  int64                `json:"refundCount"`
		UnlockAmount map[string]bn.Number `json:"unlockAmount"`
		Finished     bool                 `json:"finished"`
	}

	e.sdk.Helper().ReceiptHelper().Emit(refundBet{
		Commit:       commit,
		RefundCount:  refundCount,
		UnlockAmount: unlockAmount,
		Finished:     finished,
	})
}