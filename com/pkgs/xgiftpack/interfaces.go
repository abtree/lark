package xgiftpack

func GetKey(x int) string {
	if giftPack == nil {
		return ""
	}
	return giftPack.GetKey(x)
}

func ParseKey(key string) int {
	if giftPack == nil {
		return 0
	}
	return giftPack.ParseKey(key)
}

func Size() int {
	if giftPack == nil {
		return 0
	}
	return giftPack.Size()
}

func Vaild(key string) bool {
	if giftPack == nil {
		return false
	}
	return giftPack.Vaild(key)
}
