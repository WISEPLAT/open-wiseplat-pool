package util

import (
	"math/big"
	"regexp"
	"strconv"
	"time"

	"github.com/wiseplat/go-wiseplat/common"
	"github.com/wiseplat/go-wiseplat/common/math"
)

const maxUncleLag = 2

var Wise = math.BigPow(10, 18)
var Shannon = math.BigPow(10, 9)

var pow256 = math.BigPow(2, 256)
var addressPattern = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
var zeroHash = regexp.MustCompile("^0?x?0+$")

func IsValidHexAddress(s string) bool {
	if IsZeroHash(s) || !addressPattern.MatchString(s) {
		return false
	}
	return true
}

func IsZeroHash(s string) bool {
	return zeroHash.MatchString(s)
}

func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func GetTargetHex(diff int64) string {
	difficulty := big.NewInt(diff)
	diff1 := new(big.Int).Div(pow256, difficulty)
	return string(common.ToHex(diff1.Bytes()))
}

func TargetHexToDiff(targetHex string) *big.Int {
	targetBytes := common.FromHex(targetHex)
	return new(big.Int).Div(pow256, new(big.Int).SetBytes(targetBytes))
}

func ToHex(n int64) string {
	return "0x0" + strconv.FormatInt(n, 16)
}

func FormatReward(reward *big.Int) string {
	return reward.String()
}

func FormatRatReward(reward *big.Rat) string {
	wei := new(big.Rat).SetInt(Wise)
	reward = reward.Quo(reward, wei)
	return reward.FloatString(8)
}

// Calculate PPS rate at given block and share height
func GetShareReward(shareDiff, netDiff int64, height, topHeight uint64, fee float64) float64 {
	// Don't reward shares which are too lagging behind the tip
	if topHeight-height > maxUncleLag {
		return 0.0
	}

	// Base reward
	base := new(big.Rat).SetInt(Wise)
	base.Mul(base, new(big.Rat).SetInt64(3))
	feePercent := new(big.Rat).SetFloat64(fee / 100)
	feeValue := new(big.Rat).Mul(base, feePercent)
	base.Sub(base, feeValue)
	
	// Reward with given tip and share height
	R := new(big.Rat).SetInt64(int64(height))
	R.Add(R, new(big.Rat).SetInt64(8))
	R.Sub(R, new(big.Rat).SetInt64(int64(topHeight)))
	R.Mul(R, base)
	R.Quo(R, new(big.Rat).SetInt64(8))
	
	// Actual share reward
	wei := R
	wei.Mul(wei, new(big.Rat).SetInt64(shareDiff))
	wei.Quo(wei, new(big.Rat).SetInt64(netDiff))
	shannon := new(big.Rat).SetInt(Shannon)
	inShannon := new(big.Rat).Quo(wei, shannon)
	ppsRate, _ := inShannon.Float64()
	
	return ppsRate
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func MustParseDuration(s string) time.Duration {
	value, err := time.ParseDuration(s)
	if err != nil {
		panic("util: Can't parse duration `" + s + "`: " + err.Error())
	}
	return value
}

func String2Big(num string) *big.Int {
	n := new(big.Int)
	n.SetString(num, 0)
	return n
}

func Schedule(what func(), delay time.Duration) chan bool {
	stop := make(chan bool)
	go func() {
		for {
			what()
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()
	return stop
}
