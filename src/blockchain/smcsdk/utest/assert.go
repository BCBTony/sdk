/*
 * assert.go 实现各种断言方法，判断执行结果是否符合预期
 */

package utest

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/jsoniter"
	"blockchain/smcsdk/sdk/std"
	"blockchain/smcsdk/sdk/types"
	"gopkg.in/check.v1"
)

// Assert assert true
func Assert(b bool) {
	UTP.c.Assert(b, check.Equals, true)
}

//AssertEquals assert a equals b
func AssertEquals(a, b interface{}) {
	UTP.c.Assert(a, check.Equals, b)
}

//AssertError assert error code
func AssertError(err types.Error, expected uint32) {
	UTP.c.Assert(err.ErrorCode, check.Equals, expected)
}

//AssertErrors assert errors
func AssertErrors(err types.Error, arg ...uint32) {
	for _, item := range arg {
		AssertError(err, item)
	}
}

//AssertOK assert errcode is CodeOK
func AssertOK(err types.Error) {
	UTP.c.Assert(err.ErrorCode, check.Equals, uint32(types.CodeOK))
}

//AssertErrorMsg assert error message
func AssertErrorMsg(err types.Error, msg string) {
	UTP.c.Assert(err.Error(), check.Matches, "*"+msg+"*")
}

//AssertBalance assert balance
func AssertBalance(account sdk.IAccount, tokenName string, value bn.Number) {

	_token := UTP.Helper().TokenHelper().TokenOfName(tokenName)
	key := std.KeyOfAccountToken(account.Address(), _token.Address())
	b := sdbGet(0, 0, key)

	v := std.AccountInfo{}
	err := jsoniter.Unmarshal(b, &v)
	if err != nil {
		panic(err.Error())
	}
	UTP.c.Assert(_token.Address(), check.Equals, v.Address)
	UTP.c.Assert(value.V.String(), check.Equals, v.Balance.V.String())
}

//AssertSDB assert key's value in SDB
//判断状态数据库中某一Key的值，匹配完整格式，可以为结构体
func AssertSDB(key string, interf interface{}) {

	_v, err := jsoniter.Marshal(interf)
	if err != nil {
		panic(err.Error())
	}

	b := sdbGet(0, 0, key)

	UTP.c.Assert(b, check.DeepEquals, _v)
}

//AssertReceipt assert a receipt is existing
//判断测试结果包含某一特定收据，匹配完整收据格式
func AssertReceipt(interf interface{}) {
	_r := std.Receipt{}
	_s, err := jsoniter.Marshal(interf)
	if err != nil {
		panic(err.Error())
	}
	bMatch := false
Loop:
	for _, v := range UTP.Message().InputReceipts() {
		err := jsoniter.Unmarshal(v.Value, &_r)
		if err != nil {
			panic(err.Error())
		}
		if len(_r.Bytes) != len(_s) {
			continue Loop
		}

		for i := 0; i < len(_s); i++ {
			if _r.Bytes[i] != _s[i] {
				continue Loop
			}
		}
		//Find receipt
		bMatch = true
		break
	}

	UTP.c.Assert(true, check.Equals, bMatch)
}

//CheckError check error code is expected or not
func CheckError(err types.Error, expected int) {
	UTP.c.Check(err.ErrorCode, check.Equals, uint32(expected))
}

//CheckOK check error code is CodeOK or not
func CheckOK(err types.Error) {
	UTP.c.Check(err.ErrorCode, check.Equals, uint32(types.CodeOK))
}

//CheckErrorMsg check error message is expected or not
func CheckErrorMsg(err types.Error, msg string) {
	UTP.c.Check(err.Error(), check.Matches, "*"+msg+"*")
}
