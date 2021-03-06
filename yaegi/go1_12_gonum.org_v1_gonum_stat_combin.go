// Code generated by 'goexports gonum.org/v1/gonum/stat/combin'. DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.12,!go1.13

package yaegi

import (
	"reflect"

	"gonum.org/v1/gonum/stat/combin"
)

func init() {
	Symbols["gonum.org/v1/gonum/stat/combin"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Binomial":                reflect.ValueOf(combin.Binomial),
		"Cartesian":               reflect.ValueOf(combin.Cartesian),
		"Combinations":            reflect.ValueOf(combin.Combinations),
		"GeneralizedBinomial":     reflect.ValueOf(combin.GeneralizedBinomial),
		"IdxFor":                  reflect.ValueOf(combin.IdxFor),
		"LogGeneralizedBinomial":  reflect.ValueOf(combin.LogGeneralizedBinomial),
		"NewCombinationGenerator": reflect.ValueOf(combin.NewCombinationGenerator),
		"SubFor":                  reflect.ValueOf(combin.SubFor),

		// type definitions
		"CombinationGenerator": reflect.ValueOf((*combin.CombinationGenerator)(nil)),
	}
}
