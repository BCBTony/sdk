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

//AddDonee This is a method of MySuite
func (mysuit *MySuite) TestDemo_AddDonee(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)

	gls.Mgr.SetValues(gls.Values{gls.SDKKey: utest.UTP.ISmartContract}, func() {
		test := NewTestObject(contractOwner)
		test.setSender(contractOwner).InitChain()
		mysuit.test_AddDonee(contractOwner, test)
	})
}

func (mysuit *MySuite) test_AddDonee(owner sdk.IAccount, test *TestObject) {
	fmt.Println("=== Run  UnitTest case: AddDonee(donee types.Address)")

	//prepare
	zero := bn.N(0)
	oneCoin := bn.N(1000000000)
	a1 := utest.NewAccount("TSC", oneCoin)
	a2 := utest.NewAccount("TSC", oneCoin)

	fmt.Println("=== test for authorization")
	test.run(types.ErrNoAuthorization, func(t *TestObject) types.Error {
		t.setSender(a1)
		return t.AddDonee(a2.Address())
	})

	fmt.Println("=== pass")

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
		{owner, owner.Address(), "--异常流程--", errDoneeCannotBeOwner},
		{owner, utest.GetContract().Address(), "--异常流程---", errDoneeCannotBeSmc},
		{owner, utest.GetContract().Account().Address(), "--异常流程---", errDoneeCannotBeSmc},
	}
	for _, example := range examples {
		test.run(example.code, func(t *TestObject) types.Error {
			t.setSender(example.account)
			return t.AddDonee(example.addr)
		})
	}
	fmt.Println("=== pass")

	fmt.Println("=== test for business logic")
	utest.AssertSDB(keyOfDonation(a1.Address()), nil)
	var examples2 = []struct {
		account sdk.IAccount
		addr    string
		desc    string
		code    uint32
	}{
		{owner, a1.Address(), "--正常流程--", types.CodeOK},
		{owner, a1.Address(), "--异常流程--", errDoneeAlreadyExist},
	}
	for _, example := range examples2 {
		test.run(example.code, func(t *TestObject) types.Error {
			t.setSender(example.account)
			return t.AddDonee(example.addr)
		})
	}

	utest.AssertSDB(keyOfDonation(a1.Address()), &zero)
	utest.AssertSDB(keyOfDonation(a2.Address()), nil)
	fmt.Println("=== pass")
}

func keyOfDonation(addr types.Address) string {
	return fmt.Sprintf("/donations/%v", addr)
}
