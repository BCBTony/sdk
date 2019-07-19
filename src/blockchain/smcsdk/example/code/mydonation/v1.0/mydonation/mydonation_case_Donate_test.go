package mydonation

import (
	"blockchain/smcsdk/common/gls"
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/utest"
	"fmt"
	"gopkg.in/check.v1"
)

//Donate This is a method of MySuite
func (mysuit *MySuite) TestDemo_Donate(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)

	gls.Mgr.SetValues(gls.Values{gls.SDKKey: utest.UTP.ISmartContract}, func() {
		test := NewTestObject(contractOwner)
		test.setSender(contractOwner).InitChain()
		mysuit.test_Donate(contractOwner, test)
	})
}

func (mysuit *MySuite) test_Donate(owner sdk.IAccount, test *TestObject) {
	fmt.Println("=== Run  UnitTest case: Donate(donee types.Address)")

	//prepare
	halfCoin := bn.N(500000000)
	oneCoin := bn.N(1000000000)
	oneHalfCoin := bn.N(1500000000)
	twoCoin := bn.N(2000000000)
	utest.Transfer(nil, owner.Address(), "TSC", twoCoin)
	utest.Transfer(nil, owner.Address(), "BTC", twoCoin)
	a1 := utest.NewAccount("TSC", oneCoin)
	a2 := utest.NewAccount("TSC", oneCoin)
	a3 := utest.NewAccount("TSC", oneCoin)

	test.run(types.CodeOK, func(t *TestObject) types.Error {
		t.setSender(owner)
		return t.AddDonee(a1.Address())
	})

	fmt.Println("=== test for parameters")
	var examples = []struct {
		account sdk.IAccount
		addr    types.Address
		desc    string
		code    uint32
	}{
		{owner, "", "--异常流程--", types.ErrInvalidAddress},
		{owner, "local", "--异常流程--", types.ErrInvalidAddress},
		{owner, "localhshskhjkshfsswtsysyst6t76ddsg7s7w", "--异常流程--", types.ErrInvalidAddress},
		{owner, a2.Address(), "--异常流程--", errDoneeNotExist},
		{owner, a3.Address(), "--异常流程---", errDoneeNotExist},
	}
	for _, example := range examples {
		test.run(example.code, func(t *TestObject) types.Error {
			t.setSender(example.account)
			return t.Donate(example.addr)
		})
	}
	fmt.Println("=== pass")

	fmt.Println("=== test for receipt of transfer")
	test.run(types.ErrInvalidParameter, func(t *TestObject) types.Error {
		t.setSender(owner)
		return t.Donate(a1.Address())
	})

	test.run(types.CodeOK, func(t *TestObject) types.Error {
		t.setSender(owner)
		utest.Assert(t.transfer("TSC", halfCoin) != nil)
		return t.Donate(a1.Address())
	})

	test.run(types.ErrInvalidParameter, func(t *TestObject) types.Error {
		t.setSender(owner)
		utest.Assert(test.transfer("TSC", halfCoin) != nil)
		utest.Assert(test.transfer("TSC", halfCoin) != nil)
		return t.Donate(a1.Address())
	})

	test.run(types.ErrInvalidParameter, func(t *TestObject) types.Error {
		t.setSender(owner)
		utest.Assert(t.transfer("BTC", halfCoin) != nil)
		return t.Donate(a1.Address())
	})

	fmt.Println("=== pass")

	fmt.Println("=== test for business logic")
	utest.AssertSDB(keyOfDonation(a1.Address()), &halfCoin)

	test.run(types.CodeOK, func(t *TestObject) types.Error {
		t.setSender(a2)
		utest.Assert(t.transfer(halfCoin) != nil)
		return t.Donate(a1.Address())
	})

	test.run(types.CodeOK, func(t *TestObject) types.Error {
		t.setSender(a3)
		utest.Assert(t.transfer(halfCoin) != nil)
		return t.Donate(a1.Address())
	})

	utest.AssertSDB(keyOfDonation(a1.Address()), &oneHalfCoin)
	fmt.Println("=== pass")
}
