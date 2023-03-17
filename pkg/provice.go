package pkg

// GetProvince 根据识别号获取省份
func GetProvince(code uint) string {
	result, ok := provinceList[code]
	if ok {
		return result
	}
	return "未知省份"
}

var (
	// 省级
	provinceList = make(map[uint]string)

	// 市级
	ctiyList = make(map[uint]map[uint]string)

	// 县、区级
	countyList = make(map[uint]map[uint]string)
)

func init() {
	provinceList[11] = "北京市"
	provinceList[12] = "天津市"
	provinceList[13] = "河北省"
	provinceList[14] = "山西省"
	provinceList[15] = "内蒙古自治区"
	provinceList[21] = "辽宁省"
	provinceList[22] = "吉林省"
	provinceList[23] = "黑龙江省"
	provinceList[31] = "上海市"
	provinceList[32] = "江苏省"
	provinceList[33] = "浙江省"
	provinceList[34] = "安徽省"
	provinceList[35] = "福建省"
	provinceList[36] = "江西省"
	provinceList[37] = "山东省"
	provinceList[41] = "河南省"
	provinceList[42] = "湖北省"
	provinceList[43] = "湖南省"
	provinceList[44] = "广东省"
	provinceList[45] = "广西壮族自治区"
	provinceList[46] = "海南省"
	provinceList[50] = "重庆市"
	provinceList[51] = "四川省"
	provinceList[52] = "贵州省"
	provinceList[53] = "云南省"
	provinceList[54] = "西藏自治区"
	provinceList[61] = "陕西省"
	provinceList[62] = "甘肃省"
	provinceList[63] = "甘肃省"
	provinceList[64] = "宁夏回族自治区"
	provinceList[65] = "新疆维吾尔自治区"
	provinceList[71] = "台湾省"
	provinceList[81] = "香港特别行政区"
	provinceList[82] = "澳门特别行政区"
}
