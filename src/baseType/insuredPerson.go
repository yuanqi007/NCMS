/*
	参保人员信息结构 2015.11.28 created by yq
*/
package baseType

type InsuredPerson struct {
	NcmsId                             string  //新农合编号
	Name                               string  //姓名
	Id                                 string  //身份证号
	IdentityType                       int8    //人员身份
	Account                            float32 //账户余额
	CumulativeTotalPayment             float32 //基本统筹支付累计
	CumulativeTotalOfOutpatientService float32 //门诊统筹支付累计
	ReliefPayment                      float32 //救助支付累计
	SelfPayWithinTheScopeOfInsurance   float32 //保险范围内自付累计
}
