package entity

type DeviceInfo struct {
	//获取模块的IMEI码
	IMEI string
	//获取SIM卡的IMSI码
	IMSI string
	//获取信号强度
	//-1 未知或不可测
	//0 小于等于-113 dB
	//1 -111 dBm
	//2...30 -109... -53 dBm
	//31 大于等于-51 dBm
	CSQ string
	//获取运营商信息
	COPS string
	//获取基站的LAC和CI，可通过第三方接口经纬度信息
	CREG string
}
