package utils

import "github.com/shopspring/decimal"

var Zero = decimal.NewFromFloat(0)

func CastDecimalFromInt(val int) decimal.Decimal {
	return decimal.New(int64(val), 0)
}

func CastYuanToCentInt64(val float64) int64 {
	return decimal.NewFromFloat(val).Mul(CastDecimalFromInt(100)).IntPart()
}

func CastCentToYuanFloat64(val float64) float64 {
	res, _ := decimal.NewFromFloat(val).Div(CastDecimalFromInt(100)).Float64()
	return res
}

func Float64Add(f1, f2 float64) float64 {
	res, _ := decimal.NewFromFloat(f1).Add(decimal.NewFromFloat(f2)).Float64()
	return res
}

func Float64Sub(f1, f2 float64) float64 {
	res, _ := decimal.NewFromFloat(f1).Sub(decimal.NewFromFloat(f2)).Float64()
	return res
}

func Float64Mul(f1, f2 float64) float64 {
	res, _ := decimal.NewFromFloat(f1).Mul(decimal.NewFromFloat(f2)).Float64()
	return res
}

func Float64Div(f1, f2 float64) float64 {
	res, _ := decimal.NewFromFloat(f1).DivRound(decimal.NewFromFloat(f2), 2).Float64()
	return res
}

func Float64MulInt(f1 float64, f2 int) float64 {
	res, _ := decimal.NewFromFloat(f1).Mul(CastDecimalFromInt(f2)).Float64()
	return res
}

func Float64DivInt(f1 float64, f2 int) float64 {
	res, _ := decimal.NewFromFloat(f1).DivRound(CastDecimalFromInt(f2), 2).Float64()
	return res
}

func Float64MulIntToInt64(f1 float64, f2 int) int64 {
	return decimal.NewFromFloat(f1).Mul(CastDecimalFromInt(f2)).IntPart()
}
