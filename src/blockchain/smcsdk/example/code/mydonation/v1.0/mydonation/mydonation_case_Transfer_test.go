package mydonation

import (
	"blockchain/smcsdk/common/gls"
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/std"
	"blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/utest"
	"common/jsoniter"
	"fmt"
	"gopkg.in/check.v1"
)

//Transfer This is a method of MySuite
func (mysuit *MySuite) TestDemo_Transfer(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)

	gls.Mgr.SetValues(gls.Values{gls.SDKKey: utest.UTP.ISmartContract}, func() {
		test := NewTestObject(contractOwner)
		test.setSender(contractOwner).InitChain()
		mysuit.test_Transfer(contractOwner, test)
	})
}

func (mysuit *MySuite) test_Transfer(owner sdk.IAccount, test *TestObject) {
	fmt.Println("=== Run  UnitTest case: Transfer(donee types.Address, value bn.Number)")

	//prepare
	zero := bn.N(0)
	oneCoin := bn.N(1000000000)
	halfCoin := bn.N(500000000)
	utest.Transfer(nil, owner.Address(), bn.N(2).Mul(oneCoin))
	a1 := utest.NewAccount("TSC", oneCoin)
	a2 := utest.NewAccount("TSC", oneCoin)
	a3 := utest.NewAccount("TSC", oneCoin)
	test.run(types.CodeOK, func(t *TestObject) types.Error {
		t.setSender(owner)
		return t.AddDonee(a1.Address())
	})

	test.run(types.CodeOK, func(t *TestObject) types.Error {
		t.setSender(owner)
		utest.Assert(test.transfer(oneCoin) != nil)
		return t.Donate(a1.Address())
	})

	utest.AssertOK(test.AddDonee(a2.Address()))

	fmt.Println("=== test for authorization")
	test.run(types.ErrNoAuthorization, func(t *TestObject) types.Error {
		t.setSender(a2)
		utest.Assert(test.transfer(oneCoin) != nil)
		return t.Transfer(a1.Address(), halfCoin)
	})

	fmt.Println("=== pass")

	fmt.Println("=== test for parameters")
	var examples = []struct {
		account sdk.IAccount
		addr    types.Address
		amount  bn.Number
		desc    string
		code    uint32
	}{
		{owner, "", halfCoin, "--异常流程--", types.ErrInvalidAddress},
		{owner, "local", halfCoin, "--异常流程--", types.ErrInvalidAddress},
		{owner, "localhshskhjkshfsswtsysyst6t76ddsg7s7w", halfCoin, "--异常流程--", types.ErrInvalidAddress},
		{owner, owner.Address(), halfCoin, "--异常流程--", errDoneeNotExist},
		{owner, utest.GetContract().Address(), halfCoin, "--异常流程---", errDoneeNotExist},
		{owner, utest.GetContract().Account().Address(), halfCoin, "--异常流程--", errDoneeNotExist},
		{owner, a3.Address(), halfCoin, "--异常流程---", errDoneeNotExist},
		{owner, a1.Address(), bn.N(-1), "--异常流程--", types.ErrInvalidParameter},
		{owner, a1.Address(), bn.N(0), "--异常流程---", types.ErrInvalidParameter},
		{owner, a1.Address(), bn.N(1), "--正常流程--", types.CodeOK},
		{owner, a1.Address(), oneCoin, "--异常流程---", errDonationNotEnough},
	}
	type transferDonation struct {
		Donee   types.Address `json:"donee"`
		Value   bn.Number     `json:"value"`
		Balance bn.Number     `json:"balance"`
	}
	balance := bn.N(0)
	var acc sdk.IAccount
	accBanlance := bn.N(0)
	tokenName := test.obj.sdk.Helper().GenesisHelper().Token().Name()
	donationAcc := utest.UTP.Helper().ContractHelper().ContractOfName("mydonation").Account()
	donationBal := bn.N(0)
	for _, example := range examples {

		if example.code == types.CodeOK {
			balance = test.obj._donations(example.addr).Sub(example.amount)
			acc = test.obj.sdk.Helper().AccountHelper().AccountOf(example.addr)
			accBanlance = acc.BalanceOfName(tokenName)
			donationBal = donationAcc.BalanceOfName(tokenName)
		}

		test.run(example.code, func(t *TestObject) types.Error {
			t.setSender(example.account)
			err := t.Transfer(example.addr, example.amount)
			if err.ErrorCode == types.CodeOK {
				//收据检查
				t.assertReceipt(1, transferDonation{example.addr, example.amount, balance})

			}
			return err
		})
		if example.code == types.CodeOK {
			// 用户账户检查
			utest.AssertBalance(acc, tokenName, accBanlance.Add(example.amount))
			// 合约账户检查
			utest.AssertBalance(donationAcc, tokenName, donationBal.Sub(example.amount))
		}

	}

	fmt.Println("=== pass")

	fmt.Println("=== test for business logic")
	x := oneCoin.SubI(1)
	utest.AssertSDB(keyOfDonation(a1.Address()), &x)

	test.run(types.CodeOK, func(t *TestObject) types.Error {
		t.setSender(owner)
		// set block info---start---- //
		stringBlock := "{\"chainID\": \"test\",\"blockHash\": \"8BF02B8B238233453488311BE9B316E58AB7E1356CE948CB90DFEF1AF56992EB\",\"height\": 9,\"time\": 1562928282,\"numTxs\": 1,\"dataHash\": \"E9ED94928FC7C3659DAED85F7C851DDA28E7E8368B594483AED3AEA74507CCBF\",\"proposerAddress\": \"test84RwFFqN8ae9N8U3sE49QLYPS2CRHpSsa\",\"rewardAddress\": \"test3CiR4b8yaHZh91nomkYhQhKr4k3BPf7Bm\",\"randomNumber\": \"E1FCE25ACB03E74C649774EFE08F8929C7C4BC06FEC4B69015B745A2D643DC45\",\"version\": \"\",\"lastBlockHash\": \"04058B18052FD86B2A3032BCC55C823C48BF5810A3726F538A1D01EBB42584C5\",\"lastCommitHash\": \"681965E11A20F569C84449AE4F50FD1AD3DE7AA93C67D25601A7256122C8E8A7\",\"lastAppHash\": \"79D554E3B10760C102E9051E197111ED8D99DC9EDE86E62CD63FEDA422B5F0DC\",\"lastFee\": 1250000}"
		block := &std.Block{}
		jsoniter.Unmarshal([]byte(stringBlock), block)
		t.SetNextBlock(*block)
		// set block info---end---- //
		utest.Assert(test.transfer(bn.N(1)) != nil)
		return t.Donate(a1.Address())
	})

	utest.AssertSDB(keyOfDonation(a1.Address()), &oneCoin)
	test.run(types.CodeOK, func(t *TestObject) types.Error {
		t.setSender(owner)
		return t.Transfer(a1.Address(), halfCoin)
	})
	test.run(types.CodeOK, func(t *TestObject) types.Error {
		t.setSender(owner)
		return t.Transfer(a1.Address(), halfCoin)
	})

	utest.AssertSDB(keyOfDonation(a1.Address()), &zero)
	utest.AssertBalance(a1, "TSC", bn.N(2).Mul(oneCoin).AddI(1))
	fmt.Println("=== pass")
}
