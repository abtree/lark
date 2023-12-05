package xgiftpack

import (
	"bytes"
	"lark/com/xcrypto"
	"lark/pb"
	"math"
)

type GiftPack struct {
	Cfg *pb.Jgiftpack
}

var giftPack *GiftPack

func NewGiftPack(cfg *pb.Jgiftpack) *GiftPack {
	if giftPack != nil {
		giftPack.Cfg = cfg
		return giftPack
	}
	giftPack = &GiftPack{
		Cfg: cfg,
	}
	return giftPack
}

func (t *GiftPack) GetKey(x int) string {
	size := len(t.Cfg.BasePos) + len(t.Cfg.ExtroPos)
	byts := make([]byte, size)
	chars := []byte(t.Cfg.Chars)
	total := len(chars)
	m := 0
	n := x
	dp := []byte{}
	for _, p := range t.Cfg.BasePos {
		m = n % total
		n = n / total
		byts[p] = chars[m]
		dp = append(dp, chars[m])
	}
	str := string(dp) + t.Cfg.Key
	dec := xcrypto.MD5(str)
	for i, p := range t.Cfg.ExtroPos {
		byts[p] = dec[t.Cfg.Mix[i]]
	}
	return string(byts)
}

func (t *GiftPack) ParseKey(key string) int {
	byts := []byte(key)
	chars := []byte(t.Cfg.Chars)
	total := len(chars)
	ret := float64(0)
	for i, pos := range t.Cfg.BasePos {
		c := byts[pos]
		x := bytes.IndexByte(chars, c)
		ret += float64(x) * math.Pow(float64(total), float64(i))
	}
	return int(ret)
}

func (t *GiftPack) Size() int {
	return len(t.Cfg.BasePos) + len(t.Cfg.ExtroPos)
}

func (t *GiftPack) Vaild(key string) bool {
	byts := []byte(key)
	if len(byts) != t.Size() {
		return false //长度不正确
	}
	dp := []byte{}
	for _, p := range t.Cfg.BasePos {
		dp = append(dp, byts[p])
	}
	str := string(dp) + t.Cfg.Key
	dec := xcrypto.MD5(str)
	for i, p := range t.Cfg.ExtroPos {
		if byts[p] != dec[t.Cfg.Mix[i]] {
			return false //验证码位验证失败
		}
	}
	return true
}
