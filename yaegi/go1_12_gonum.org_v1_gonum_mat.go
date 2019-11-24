// Code generated by 'goexports gonum.org/v1/gonum/mat'. DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.12,!go1.13

package yaegi

import (
	"reflect"

	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/blas/cblas128"
	"gonum.org/v1/gonum/mat"
)

func init() {
	Symbols["gonum.org/v1/gonum/mat"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"CEqual":                 reflect.ValueOf(mat.CEqual),
		"CEqualApprox":           reflect.ValueOf(mat.CEqualApprox),
		"Col":                    reflect.ValueOf(mat.Col),
		"Cond":                   reflect.ValueOf(mat.Cond),
		"CondNorm":               reflect.ValueOf(mat.CondNorm),
		"CondNormTrans":          reflect.ValueOf(mat.CondNormTrans),
		"ConditionTolerance":     reflect.ValueOf(mat.ConditionTolerance),
		"DenseCopyOf":            reflect.ValueOf(mat.DenseCopyOf),
		"Det":                    reflect.ValueOf(mat.Det),
		"Dot":                    reflect.ValueOf(mat.Dot),
		"DotByte":                reflect.ValueOf(mat.DotByte),
		"EigenBoth":              reflect.ValueOf(mat.EigenBoth),
		"EigenLeft":              reflect.ValueOf(mat.EigenLeft),
		"EigenNone":              reflect.ValueOf(mat.EigenNone),
		"EigenRight":             reflect.ValueOf(mat.EigenRight),
		"Equal":                  reflect.ValueOf(mat.Equal),
		"EqualApprox":            reflect.ValueOf(mat.EqualApprox),
		"ErrBandSet":             reflect.ValueOf(&mat.ErrBandSet).Elem(),
		"ErrColAccess":           reflect.ValueOf(&mat.ErrColAccess).Elem(),
		"ErrColLength":           reflect.ValueOf(&mat.ErrColLength).Elem(),
		"ErrDiagSet":             reflect.ValueOf(&mat.ErrDiagSet).Elem(),
		"ErrFailedEigen":         reflect.ValueOf(&mat.ErrFailedEigen).Elem(),
		"ErrIllegalStride":       reflect.ValueOf(&mat.ErrIllegalStride).Elem(),
		"ErrIndexOutOfRange":     reflect.ValueOf(&mat.ErrIndexOutOfRange).Elem(),
		"ErrNormOrder":           reflect.ValueOf(&mat.ErrNormOrder).Elem(),
		"ErrNotPSD":              reflect.ValueOf(&mat.ErrNotPSD).Elem(),
		"ErrPivot":               reflect.ValueOf(&mat.ErrPivot).Elem(),
		"ErrRowAccess":           reflect.ValueOf(&mat.ErrRowAccess).Elem(),
		"ErrRowLength":           reflect.ValueOf(&mat.ErrRowLength).Elem(),
		"ErrShape":               reflect.ValueOf(&mat.ErrShape).Elem(),
		"ErrSingular":            reflect.ValueOf(&mat.ErrSingular).Elem(),
		"ErrSliceLengthMismatch": reflect.ValueOf(&mat.ErrSliceLengthMismatch).Elem(),
		"ErrSquare":              reflect.ValueOf(&mat.ErrSquare).Elem(),
		"ErrTriangle":            reflect.ValueOf(&mat.ErrTriangle).Elem(),
		"ErrTriangleSet":         reflect.ValueOf(&mat.ErrTriangleSet).Elem(),
		"ErrVectorAccess":        reflect.ValueOf(&mat.ErrVectorAccess).Elem(),
		"ErrZeroLength":          reflect.ValueOf(&mat.ErrZeroLength).Elem(),
		"Excerpt":                reflect.ValueOf(mat.Excerpt),
		"Formatted":              reflect.ValueOf(mat.Formatted),
		"GSVDAll":                reflect.ValueOf(mat.GSVDAll),
		"GSVDNone":               reflect.ValueOf(mat.GSVDNone),
		"GSVDQ":                  reflect.ValueOf(mat.GSVDQ),
		"GSVDU":                  reflect.ValueOf(mat.GSVDU),
		"GSVDV":                  reflect.ValueOf(mat.GSVDV),
		"Inner":                  reflect.ValueOf(mat.Inner),
		"LogDet":                 reflect.ValueOf(mat.LogDet),
		"Lower":                  reflect.ValueOf(mat.Lower),
		"Max":                    reflect.ValueOf(mat.Max),
		"Maybe":                  reflect.ValueOf(mat.Maybe),
		"MaybeComplex":           reflect.ValueOf(mat.MaybeComplex),
		"MaybeFloat":             reflect.ValueOf(mat.MaybeFloat),
		"Min":                    reflect.ValueOf(mat.Min),
		"NewBandDense":           reflect.ValueOf(mat.NewBandDense),
		"NewCDense":              reflect.ValueOf(mat.NewCDense),
		"NewDense":               reflect.ValueOf(mat.NewDense),
		"NewDiagDense":           reflect.ValueOf(mat.NewDiagDense),
		"NewDiagonalRect":        reflect.ValueOf(mat.NewDiagonalRect),
		"NewSymBandDense":        reflect.ValueOf(mat.NewSymBandDense),
		"NewSymDense":            reflect.ValueOf(mat.NewSymDense),
		"NewTriBandDense":        reflect.ValueOf(mat.NewTriBandDense),
		"NewTriDense":            reflect.ValueOf(mat.NewTriDense),
		"NewVecDense":            reflect.ValueOf(mat.NewVecDense),
		"Norm":                   reflect.ValueOf(mat.Norm),
		"Prefix":                 reflect.ValueOf(mat.Prefix),
		"Row":                    reflect.ValueOf(mat.Row),
		"SVDFull":                reflect.ValueOf(mat.SVDFull),
		"SVDFullU":               reflect.ValueOf(mat.SVDFullU),
		"SVDFullV":               reflect.ValueOf(mat.SVDFullV),
		"SVDNone":                reflect.ValueOf(mat.SVDNone),
		"SVDThin":                reflect.ValueOf(mat.SVDThin),
		"SVDThinU":               reflect.ValueOf(mat.SVDThinU),
		"SVDThinV":               reflect.ValueOf(mat.SVDThinV),
		"Squeeze":                reflect.ValueOf(mat.Squeeze),
		"Sum":                    reflect.ValueOf(mat.Sum),
		"Trace":                  reflect.ValueOf(mat.Trace),
		"Upper":                  reflect.ValueOf(mat.Upper),
		"VecDenseCopyOf":         reflect.ValueOf(mat.VecDenseCopyOf),

		// type definitions
		"BandDense":            reflect.ValueOf((*mat.BandDense)(nil)),
		"BandWidther":          reflect.ValueOf((*mat.BandWidther)(nil)),
		"Banded":               reflect.ValueOf((*mat.Banded)(nil)),
		"CDense":               reflect.ValueOf((*mat.CDense)(nil)),
		"CMatrix":              reflect.ValueOf((*mat.CMatrix)(nil)),
		"Cholesky":             reflect.ValueOf((*mat.Cholesky)(nil)),
		"ClonerFrom":           reflect.ValueOf((*mat.ClonerFrom)(nil)),
		"ColNonZeroDoer":       reflect.ValueOf((*mat.ColNonZeroDoer)(nil)),
		"ColViewer":            reflect.ValueOf((*mat.ColViewer)(nil)),
		"Condition":            reflect.ValueOf((*mat.Condition)(nil)),
		"Conjugate":            reflect.ValueOf((*mat.Conjugate)(nil)),
		"Copier":               reflect.ValueOf((*mat.Copier)(nil)),
		"Dense":                reflect.ValueOf((*mat.Dense)(nil)),
		"DiagDense":            reflect.ValueOf((*mat.DiagDense)(nil)),
		"Diagonal":             reflect.ValueOf((*mat.Diagonal)(nil)),
		"Eigen":                reflect.ValueOf((*mat.Eigen)(nil)),
		"EigenKind":            reflect.ValueOf((*mat.EigenKind)(nil)),
		"EigenSym":             reflect.ValueOf((*mat.EigenSym)(nil)),
		"Error":                reflect.ValueOf((*mat.Error)(nil)),
		"ErrorStack":           reflect.ValueOf((*mat.ErrorStack)(nil)),
		"FormatOption":         reflect.ValueOf((*mat.FormatOption)(nil)),
		"GSVD":                 reflect.ValueOf((*mat.GSVD)(nil)),
		"GSVDKind":             reflect.ValueOf((*mat.GSVDKind)(nil)),
		"Grower":               reflect.ValueOf((*mat.Grower)(nil)),
		"HOGSVD":               reflect.ValueOf((*mat.HOGSVD)(nil)),
		"LQ":                   reflect.ValueOf((*mat.LQ)(nil)),
		"LU":                   reflect.ValueOf((*mat.LU)(nil)),
		"Matrix":               reflect.ValueOf((*mat.Matrix)(nil)),
		"Mutable":              reflect.ValueOf((*mat.Mutable)(nil)),
		"MutableBanded":        reflect.ValueOf((*mat.MutableBanded)(nil)),
		"MutableDiagonal":      reflect.ValueOf((*mat.MutableDiagonal)(nil)),
		"MutableSymBanded":     reflect.ValueOf((*mat.MutableSymBanded)(nil)),
		"MutableSymmetric":     reflect.ValueOf((*mat.MutableSymmetric)(nil)),
		"MutableTriBanded":     reflect.ValueOf((*mat.MutableTriBanded)(nil)),
		"MutableTriangular":    reflect.ValueOf((*mat.MutableTriangular)(nil)),
		"NonZeroDoer":          reflect.ValueOf((*mat.NonZeroDoer)(nil)),
		"QR":                   reflect.ValueOf((*mat.QR)(nil)),
		"RawBander":            reflect.ValueOf((*mat.RawBander)(nil)),
		"RawCMatrixer":         reflect.ValueOf((*mat.RawCMatrixer)(nil)),
		"RawColViewer":         reflect.ValueOf((*mat.RawColViewer)(nil)),
		"RawMatrixSetter":      reflect.ValueOf((*mat.RawMatrixSetter)(nil)),
		"RawMatrixer":          reflect.ValueOf((*mat.RawMatrixer)(nil)),
		"RawRowViewer":         reflect.ValueOf((*mat.RawRowViewer)(nil)),
		"RawSymBander":         reflect.ValueOf((*mat.RawSymBander)(nil)),
		"RawSymmetricer":       reflect.ValueOf((*mat.RawSymmetricer)(nil)),
		"RawTriBander":         reflect.ValueOf((*mat.RawTriBander)(nil)),
		"RawTriangular":        reflect.ValueOf((*mat.RawTriangular)(nil)),
		"RawVectorer":          reflect.ValueOf((*mat.RawVectorer)(nil)),
		"Reseter":              reflect.ValueOf((*mat.Reseter)(nil)),
		"RowNonZeroDoer":       reflect.ValueOf((*mat.RowNonZeroDoer)(nil)),
		"RowViewer":            reflect.ValueOf((*mat.RowViewer)(nil)),
		"SVD":                  reflect.ValueOf((*mat.SVD)(nil)),
		"SVDKind":              reflect.ValueOf((*mat.SVDKind)(nil)),
		"SymBandDense":         reflect.ValueOf((*mat.SymBandDense)(nil)),
		"SymBanded":            reflect.ValueOf((*mat.SymBanded)(nil)),
		"SymDense":             reflect.ValueOf((*mat.SymDense)(nil)),
		"Symmetric":            reflect.ValueOf((*mat.Symmetric)(nil)),
		"Tracer":               reflect.ValueOf((*mat.Tracer)(nil)),
		"Transpose":            reflect.ValueOf((*mat.Transpose)(nil)),
		"TransposeBand":        reflect.ValueOf((*mat.TransposeBand)(nil)),
		"TransposeTri":         reflect.ValueOf((*mat.TransposeTri)(nil)),
		"TransposeTriBand":     reflect.ValueOf((*mat.TransposeTriBand)(nil)),
		"TransposeVec":         reflect.ValueOf((*mat.TransposeVec)(nil)),
		"TriBandDense":         reflect.ValueOf((*mat.TriBandDense)(nil)),
		"TriBanded":            reflect.ValueOf((*mat.TriBanded)(nil)),
		"TriDense":             reflect.ValueOf((*mat.TriDense)(nil)),
		"TriKind":              reflect.ValueOf((*mat.TriKind)(nil)),
		"Triangular":           reflect.ValueOf((*mat.Triangular)(nil)),
		"Unconjugator":         reflect.ValueOf((*mat.Unconjugator)(nil)),
		"UntransposeBander":    reflect.ValueOf((*mat.UntransposeBander)(nil)),
		"UntransposeTriBander": reflect.ValueOf((*mat.UntransposeTriBander)(nil)),
		"UntransposeTrier":     reflect.ValueOf((*mat.UntransposeTrier)(nil)),
		"Untransposer":         reflect.ValueOf((*mat.Untransposer)(nil)),
		"VecDense":             reflect.ValueOf((*mat.VecDense)(nil)),
		"Vector":               reflect.ValueOf((*mat.Vector)(nil)),

		// interface wrapper definitions
		"_BandWidther":          reflect.ValueOf((*_gonum_org_v1_gonum_mat_BandWidther)(nil)),
		"_Banded":               reflect.ValueOf((*_gonum_org_v1_gonum_mat_Banded)(nil)),
		"_CMatrix":              reflect.ValueOf((*_gonum_org_v1_gonum_mat_CMatrix)(nil)),
		"_ClonerFrom":           reflect.ValueOf((*_gonum_org_v1_gonum_mat_ClonerFrom)(nil)),
		"_ColNonZeroDoer":       reflect.ValueOf((*_gonum_org_v1_gonum_mat_ColNonZeroDoer)(nil)),
		"_ColViewer":            reflect.ValueOf((*_gonum_org_v1_gonum_mat_ColViewer)(nil)),
		"_Copier":               reflect.ValueOf((*_gonum_org_v1_gonum_mat_Copier)(nil)),
		"_Diagonal":             reflect.ValueOf((*_gonum_org_v1_gonum_mat_Diagonal)(nil)),
		"_Grower":               reflect.ValueOf((*_gonum_org_v1_gonum_mat_Grower)(nil)),
		"_Matrix":               reflect.ValueOf((*_gonum_org_v1_gonum_mat_Matrix)(nil)),
		"_Mutable":              reflect.ValueOf((*_gonum_org_v1_gonum_mat_Mutable)(nil)),
		"_MutableBanded":        reflect.ValueOf((*_gonum_org_v1_gonum_mat_MutableBanded)(nil)),
		"_MutableDiagonal":      reflect.ValueOf((*_gonum_org_v1_gonum_mat_MutableDiagonal)(nil)),
		"_MutableSymBanded":     reflect.ValueOf((*_gonum_org_v1_gonum_mat_MutableSymBanded)(nil)),
		"_MutableSymmetric":     reflect.ValueOf((*_gonum_org_v1_gonum_mat_MutableSymmetric)(nil)),
		"_MutableTriBanded":     reflect.ValueOf((*_gonum_org_v1_gonum_mat_MutableTriBanded)(nil)),
		"_MutableTriangular":    reflect.ValueOf((*_gonum_org_v1_gonum_mat_MutableTriangular)(nil)),
		"_NonZeroDoer":          reflect.ValueOf((*_gonum_org_v1_gonum_mat_NonZeroDoer)(nil)),
		"_RawBander":            reflect.ValueOf((*_gonum_org_v1_gonum_mat_RawBander)(nil)),
		"_RawCMatrixer":         reflect.ValueOf((*_gonum_org_v1_gonum_mat_RawCMatrixer)(nil)),
		"_RawColViewer":         reflect.ValueOf((*_gonum_org_v1_gonum_mat_RawColViewer)(nil)),
		"_RawMatrixSetter":      reflect.ValueOf((*_gonum_org_v1_gonum_mat_RawMatrixSetter)(nil)),
		"_RawMatrixer":          reflect.ValueOf((*_gonum_org_v1_gonum_mat_RawMatrixer)(nil)),
		"_RawRowViewer":         reflect.ValueOf((*_gonum_org_v1_gonum_mat_RawRowViewer)(nil)),
		"_RawSymBander":         reflect.ValueOf((*_gonum_org_v1_gonum_mat_RawSymBander)(nil)),
		"_RawSymmetricer":       reflect.ValueOf((*_gonum_org_v1_gonum_mat_RawSymmetricer)(nil)),
		"_RawTriBander":         reflect.ValueOf((*_gonum_org_v1_gonum_mat_RawTriBander)(nil)),
		"_RawTriangular":        reflect.ValueOf((*_gonum_org_v1_gonum_mat_RawTriangular)(nil)),
		"_RawVectorer":          reflect.ValueOf((*_gonum_org_v1_gonum_mat_RawVectorer)(nil)),
		"_Reseter":              reflect.ValueOf((*_gonum_org_v1_gonum_mat_Reseter)(nil)),
		"_RowNonZeroDoer":       reflect.ValueOf((*_gonum_org_v1_gonum_mat_RowNonZeroDoer)(nil)),
		"_RowViewer":            reflect.ValueOf((*_gonum_org_v1_gonum_mat_RowViewer)(nil)),
		"_SymBanded":            reflect.ValueOf((*_gonum_org_v1_gonum_mat_SymBanded)(nil)),
		"_Symmetric":            reflect.ValueOf((*_gonum_org_v1_gonum_mat_Symmetric)(nil)),
		"_Tracer":               reflect.ValueOf((*_gonum_org_v1_gonum_mat_Tracer)(nil)),
		"_TriBanded":            reflect.ValueOf((*_gonum_org_v1_gonum_mat_TriBanded)(nil)),
		"_Triangular":           reflect.ValueOf((*_gonum_org_v1_gonum_mat_Triangular)(nil)),
		"_Unconjugator":         reflect.ValueOf((*_gonum_org_v1_gonum_mat_Unconjugator)(nil)),
		"_UntransposeBander":    reflect.ValueOf((*_gonum_org_v1_gonum_mat_UntransposeBander)(nil)),
		"_UntransposeTriBander": reflect.ValueOf((*_gonum_org_v1_gonum_mat_UntransposeTriBander)(nil)),
		"_UntransposeTrier":     reflect.ValueOf((*_gonum_org_v1_gonum_mat_UntransposeTrier)(nil)),
		"_Untransposer":         reflect.ValueOf((*_gonum_org_v1_gonum_mat_Untransposer)(nil)),
		"_Vector":               reflect.ValueOf((*_gonum_org_v1_gonum_mat_Vector)(nil)),
	}
}

// _gonum_org_v1_gonum_mat_BandWidther is an interface wrapper for BandWidther type
type _gonum_org_v1_gonum_mat_BandWidther struct {
	WBandWidth func() (k1 int, k2 int)
}

func (W _gonum_org_v1_gonum_mat_BandWidther) BandWidth() (k1 int, k2 int) { return W.WBandWidth() }

// _gonum_org_v1_gonum_mat_Banded is an interface wrapper for Banded type
type _gonum_org_v1_gonum_mat_Banded struct {
	WAt        func(i int, j int) float64
	WBandwidth func() (kl int, ku int)
	WDims      func() (r int, c int)
	WT         func() mat.Matrix
	WTBand     func() mat.Banded
}

func (W _gonum_org_v1_gonum_mat_Banded) At(i int, j int) float64     { return W.WAt(i, j) }
func (W _gonum_org_v1_gonum_mat_Banded) Bandwidth() (kl int, ku int) { return W.WBandwidth() }
func (W _gonum_org_v1_gonum_mat_Banded) Dims() (r int, c int)        { return W.WDims() }
func (W _gonum_org_v1_gonum_mat_Banded) T() mat.Matrix               { return W.WT() }
func (W _gonum_org_v1_gonum_mat_Banded) TBand() mat.Banded           { return W.WTBand() }

// _gonum_org_v1_gonum_mat_CMatrix is an interface wrapper for CMatrix type
type _gonum_org_v1_gonum_mat_CMatrix struct {
	WAt   func(i int, j int) complex128
	WDims func() (r int, c int)
	WH    func() mat.CMatrix
}

func (W _gonum_org_v1_gonum_mat_CMatrix) At(i int, j int) complex128 { return W.WAt(i, j) }
func (W _gonum_org_v1_gonum_mat_CMatrix) Dims() (r int, c int)       { return W.WDims() }
func (W _gonum_org_v1_gonum_mat_CMatrix) H() mat.CMatrix             { return W.WH() }

// _gonum_org_v1_gonum_mat_ClonerFrom is an interface wrapper for ClonerFrom type
type _gonum_org_v1_gonum_mat_ClonerFrom struct {
	WCloneFrom func(a mat.Matrix)
}

func (W _gonum_org_v1_gonum_mat_ClonerFrom) CloneFrom(a mat.Matrix) { W.WCloneFrom(a) }

// _gonum_org_v1_gonum_mat_ColNonZeroDoer is an interface wrapper for ColNonZeroDoer type
type _gonum_org_v1_gonum_mat_ColNonZeroDoer struct {
	WDoColNonZero func(j int, fn func(i int, j int, v float64))
}

func (W _gonum_org_v1_gonum_mat_ColNonZeroDoer) DoColNonZero(j int, fn func(i int, j int, v float64)) {
	W.WDoColNonZero(j, fn)
}

// _gonum_org_v1_gonum_mat_ColViewer is an interface wrapper for ColViewer type
type _gonum_org_v1_gonum_mat_ColViewer struct {
	WColView func(j int) mat.Vector
}

func (W _gonum_org_v1_gonum_mat_ColViewer) ColView(j int) mat.Vector { return W.WColView(j) }

// _gonum_org_v1_gonum_mat_Copier is an interface wrapper for Copier type
type _gonum_org_v1_gonum_mat_Copier struct {
	WCopy func(a mat.Matrix) (r int, c int)
}

func (W _gonum_org_v1_gonum_mat_Copier) Copy(a mat.Matrix) (r int, c int) { return W.WCopy(a) }

// _gonum_org_v1_gonum_mat_Diagonal is an interface wrapper for Diagonal type
type _gonum_org_v1_gonum_mat_Diagonal struct {
	WAt        func(i int, j int) float64
	WBandwidth func() (kl int, ku int)
	WDiag      func() int
	WDims      func() (r int, c int)
	WSymBand   func() (n int, k int)
	WSymmetric func() int
	WT         func() mat.Matrix
	WTBand     func() mat.Banded
	WTTri      func() mat.Triangular
	WTTriBand  func() mat.TriBanded
	WTriBand   func() (n int, k int, kind mat.TriKind)
	WTriangle  func() (int, mat.TriKind)
}

func (W _gonum_org_v1_gonum_mat_Diagonal) At(i int, j int) float64     { return W.WAt(i, j) }
func (W _gonum_org_v1_gonum_mat_Diagonal) Bandwidth() (kl int, ku int) { return W.WBandwidth() }
func (W _gonum_org_v1_gonum_mat_Diagonal) Diag() int                   { return W.WDiag() }
func (W _gonum_org_v1_gonum_mat_Diagonal) Dims() (r int, c int)        { return W.WDims() }
func (W _gonum_org_v1_gonum_mat_Diagonal) SymBand() (n int, k int)     { return W.WSymBand() }
func (W _gonum_org_v1_gonum_mat_Diagonal) Symmetric() int              { return W.WSymmetric() }
func (W _gonum_org_v1_gonum_mat_Diagonal) T() mat.Matrix               { return W.WT() }
func (W _gonum_org_v1_gonum_mat_Diagonal) TBand() mat.Banded           { return W.WTBand() }
func (W _gonum_org_v1_gonum_mat_Diagonal) TTri() mat.Triangular        { return W.WTTri() }
func (W _gonum_org_v1_gonum_mat_Diagonal) TTriBand() mat.TriBanded     { return W.WTTriBand() }
func (W _gonum_org_v1_gonum_mat_Diagonal) TriBand() (n int, k int, kind mat.TriKind) {
	return W.WTriBand()
}
func (W _gonum_org_v1_gonum_mat_Diagonal) Triangle() (int, mat.TriKind) { return W.WTriangle() }

// _gonum_org_v1_gonum_mat_Grower is an interface wrapper for Grower type
type _gonum_org_v1_gonum_mat_Grower struct {
	WCaps func() (r int, c int)
	WGrow func(r int, c int) mat.Matrix
}

func (W _gonum_org_v1_gonum_mat_Grower) Caps() (r int, c int)         { return W.WCaps() }
func (W _gonum_org_v1_gonum_mat_Grower) Grow(r int, c int) mat.Matrix { return W.WGrow(r, c) }

// _gonum_org_v1_gonum_mat_Matrix is an interface wrapper for Matrix type
type _gonum_org_v1_gonum_mat_Matrix struct {
	WAt   func(i int, j int) float64
	WDims func() (r int, c int)
	WT    func() mat.Matrix
}

func (W _gonum_org_v1_gonum_mat_Matrix) At(i int, j int) float64 { return W.WAt(i, j) }
func (W _gonum_org_v1_gonum_mat_Matrix) Dims() (r int, c int)    { return W.WDims() }
func (W _gonum_org_v1_gonum_mat_Matrix) T() mat.Matrix           { return W.WT() }

// _gonum_org_v1_gonum_mat_Mutable is an interface wrapper for Mutable type
type _gonum_org_v1_gonum_mat_Mutable struct {
	WAt   func(i int, j int) float64
	WDims func() (r int, c int)
	WSet  func(i int, j int, v float64)
	WT    func() mat.Matrix
}

func (W _gonum_org_v1_gonum_mat_Mutable) At(i int, j int) float64     { return W.WAt(i, j) }
func (W _gonum_org_v1_gonum_mat_Mutable) Dims() (r int, c int)        { return W.WDims() }
func (W _gonum_org_v1_gonum_mat_Mutable) Set(i int, j int, v float64) { W.WSet(i, j, v) }
func (W _gonum_org_v1_gonum_mat_Mutable) T() mat.Matrix               { return W.WT() }

// _gonum_org_v1_gonum_mat_MutableBanded is an interface wrapper for MutableBanded type
type _gonum_org_v1_gonum_mat_MutableBanded struct {
	WAt        func(i int, j int) float64
	WBandwidth func() (kl int, ku int)
	WDims      func() (r int, c int)
	WSetBand   func(i int, j int, v float64)
	WT         func() mat.Matrix
	WTBand     func() mat.Banded
}

func (W _gonum_org_v1_gonum_mat_MutableBanded) At(i int, j int) float64         { return W.WAt(i, j) }
func (W _gonum_org_v1_gonum_mat_MutableBanded) Bandwidth() (kl int, ku int)     { return W.WBandwidth() }
func (W _gonum_org_v1_gonum_mat_MutableBanded) Dims() (r int, c int)            { return W.WDims() }
func (W _gonum_org_v1_gonum_mat_MutableBanded) SetBand(i int, j int, v float64) { W.WSetBand(i, j, v) }
func (W _gonum_org_v1_gonum_mat_MutableBanded) T() mat.Matrix                   { return W.WT() }
func (W _gonum_org_v1_gonum_mat_MutableBanded) TBand() mat.Banded               { return W.WTBand() }

// _gonum_org_v1_gonum_mat_MutableDiagonal is an interface wrapper for MutableDiagonal type
type _gonum_org_v1_gonum_mat_MutableDiagonal struct {
	WAt        func(i int, j int) float64
	WBandwidth func() (kl int, ku int)
	WDiag      func() int
	WDims      func() (r int, c int)
	WSetDiag   func(i int, v float64)
	WSymBand   func() (n int, k int)
	WSymmetric func() int
	WT         func() mat.Matrix
	WTBand     func() mat.Banded
	WTTri      func() mat.Triangular
	WTTriBand  func() mat.TriBanded
	WTriBand   func() (n int, k int, kind mat.TriKind)
	WTriangle  func() (int, mat.TriKind)
}

func (W _gonum_org_v1_gonum_mat_MutableDiagonal) At(i int, j int) float64     { return W.WAt(i, j) }
func (W _gonum_org_v1_gonum_mat_MutableDiagonal) Bandwidth() (kl int, ku int) { return W.WBandwidth() }
func (W _gonum_org_v1_gonum_mat_MutableDiagonal) Diag() int                   { return W.WDiag() }
func (W _gonum_org_v1_gonum_mat_MutableDiagonal) Dims() (r int, c int)        { return W.WDims() }
func (W _gonum_org_v1_gonum_mat_MutableDiagonal) SetDiag(i int, v float64)    { W.WSetDiag(i, v) }
func (W _gonum_org_v1_gonum_mat_MutableDiagonal) SymBand() (n int, k int)     { return W.WSymBand() }
func (W _gonum_org_v1_gonum_mat_MutableDiagonal) Symmetric() int              { return W.WSymmetric() }
func (W _gonum_org_v1_gonum_mat_MutableDiagonal) T() mat.Matrix               { return W.WT() }
func (W _gonum_org_v1_gonum_mat_MutableDiagonal) TBand() mat.Banded           { return W.WTBand() }
func (W _gonum_org_v1_gonum_mat_MutableDiagonal) TTri() mat.Triangular        { return W.WTTri() }
func (W _gonum_org_v1_gonum_mat_MutableDiagonal) TTriBand() mat.TriBanded     { return W.WTTriBand() }
func (W _gonum_org_v1_gonum_mat_MutableDiagonal) TriBand() (n int, k int, kind mat.TriKind) {
	return W.WTriBand()
}
func (W _gonum_org_v1_gonum_mat_MutableDiagonal) Triangle() (int, mat.TriKind) { return W.WTriangle() }

// _gonum_org_v1_gonum_mat_MutableSymBanded is an interface wrapper for MutableSymBanded type
type _gonum_org_v1_gonum_mat_MutableSymBanded struct {
	WAt         func(i int, j int) float64
	WBandwidth  func() (kl int, ku int)
	WDims       func() (r int, c int)
	WSetSymBand func(i int, j int, v float64)
	WSymBand    func() (n int, k int)
	WSymmetric  func() int
	WT          func() mat.Matrix
	WTBand      func() mat.Banded
}

func (W _gonum_org_v1_gonum_mat_MutableSymBanded) At(i int, j int) float64     { return W.WAt(i, j) }
func (W _gonum_org_v1_gonum_mat_MutableSymBanded) Bandwidth() (kl int, ku int) { return W.WBandwidth() }
func (W _gonum_org_v1_gonum_mat_MutableSymBanded) Dims() (r int, c int)        { return W.WDims() }
func (W _gonum_org_v1_gonum_mat_MutableSymBanded) SetSymBand(i int, j int, v float64) {
	W.WSetSymBand(i, j, v)
}
func (W _gonum_org_v1_gonum_mat_MutableSymBanded) SymBand() (n int, k int) { return W.WSymBand() }
func (W _gonum_org_v1_gonum_mat_MutableSymBanded) Symmetric() int          { return W.WSymmetric() }
func (W _gonum_org_v1_gonum_mat_MutableSymBanded) T() mat.Matrix           { return W.WT() }
func (W _gonum_org_v1_gonum_mat_MutableSymBanded) TBand() mat.Banded       { return W.WTBand() }

// _gonum_org_v1_gonum_mat_MutableSymmetric is an interface wrapper for MutableSymmetric type
type _gonum_org_v1_gonum_mat_MutableSymmetric struct {
	WAt        func(i int, j int) float64
	WDims      func() (r int, c int)
	WSetSym    func(i int, j int, v float64)
	WSymmetric func() int
	WT         func() mat.Matrix
}

func (W _gonum_org_v1_gonum_mat_MutableSymmetric) At(i int, j int) float64        { return W.WAt(i, j) }
func (W _gonum_org_v1_gonum_mat_MutableSymmetric) Dims() (r int, c int)           { return W.WDims() }
func (W _gonum_org_v1_gonum_mat_MutableSymmetric) SetSym(i int, j int, v float64) { W.WSetSym(i, j, v) }
func (W _gonum_org_v1_gonum_mat_MutableSymmetric) Symmetric() int                 { return W.WSymmetric() }
func (W _gonum_org_v1_gonum_mat_MutableSymmetric) T() mat.Matrix                  { return W.WT() }

// _gonum_org_v1_gonum_mat_MutableTriBanded is an interface wrapper for MutableTriBanded type
type _gonum_org_v1_gonum_mat_MutableTriBanded struct {
	WAt         func(i int, j int) float64
	WBandwidth  func() (kl int, ku int)
	WDims       func() (r int, c int)
	WSetTriBand func(i int, j int, v float64)
	WT          func() mat.Matrix
	WTBand      func() mat.Banded
	WTTri       func() mat.Triangular
	WTTriBand   func() mat.TriBanded
	WTriBand    func() (n int, k int, kind mat.TriKind)
	WTriangle   func() (n int, kind mat.TriKind)
}

func (W _gonum_org_v1_gonum_mat_MutableTriBanded) At(i int, j int) float64     { return W.WAt(i, j) }
func (W _gonum_org_v1_gonum_mat_MutableTriBanded) Bandwidth() (kl int, ku int) { return W.WBandwidth() }
func (W _gonum_org_v1_gonum_mat_MutableTriBanded) Dims() (r int, c int)        { return W.WDims() }
func (W _gonum_org_v1_gonum_mat_MutableTriBanded) SetTriBand(i int, j int, v float64) {
	W.WSetTriBand(i, j, v)
}
func (W _gonum_org_v1_gonum_mat_MutableTriBanded) T() mat.Matrix           { return W.WT() }
func (W _gonum_org_v1_gonum_mat_MutableTriBanded) TBand() mat.Banded       { return W.WTBand() }
func (W _gonum_org_v1_gonum_mat_MutableTriBanded) TTri() mat.Triangular    { return W.WTTri() }
func (W _gonum_org_v1_gonum_mat_MutableTriBanded) TTriBand() mat.TriBanded { return W.WTTriBand() }
func (W _gonum_org_v1_gonum_mat_MutableTriBanded) TriBand() (n int, k int, kind mat.TriKind) {
	return W.WTriBand()
}
func (W _gonum_org_v1_gonum_mat_MutableTriBanded) Triangle() (n int, kind mat.TriKind) {
	return W.WTriangle()
}

// _gonum_org_v1_gonum_mat_MutableTriangular is an interface wrapper for MutableTriangular type
type _gonum_org_v1_gonum_mat_MutableTriangular struct {
	WAt       func(i int, j int) float64
	WDims     func() (r int, c int)
	WSetTri   func(i int, j int, v float64)
	WT        func() mat.Matrix
	WTTri     func() mat.Triangular
	WTriangle func() (n int, kind mat.TriKind)
}

func (W _gonum_org_v1_gonum_mat_MutableTriangular) At(i int, j int) float64        { return W.WAt(i, j) }
func (W _gonum_org_v1_gonum_mat_MutableTriangular) Dims() (r int, c int)           { return W.WDims() }
func (W _gonum_org_v1_gonum_mat_MutableTriangular) SetTri(i int, j int, v float64) { W.WSetTri(i, j, v) }
func (W _gonum_org_v1_gonum_mat_MutableTriangular) T() mat.Matrix                  { return W.WT() }
func (W _gonum_org_v1_gonum_mat_MutableTriangular) TTri() mat.Triangular           { return W.WTTri() }
func (W _gonum_org_v1_gonum_mat_MutableTriangular) Triangle() (n int, kind mat.TriKind) {
	return W.WTriangle()
}

// _gonum_org_v1_gonum_mat_NonZeroDoer is an interface wrapper for NonZeroDoer type
type _gonum_org_v1_gonum_mat_NonZeroDoer struct {
	WDoNonZero func(a0 func(i int, j int, v float64))
}

func (W _gonum_org_v1_gonum_mat_NonZeroDoer) DoNonZero(a0 func(i int, j int, v float64)) {
	W.WDoNonZero(a0)
}

// _gonum_org_v1_gonum_mat_RawBander is an interface wrapper for RawBander type
type _gonum_org_v1_gonum_mat_RawBander struct {
	WRawBand func() blas64.Band
}

func (W _gonum_org_v1_gonum_mat_RawBander) RawBand() blas64.Band { return W.WRawBand() }

// _gonum_org_v1_gonum_mat_RawCMatrixer is an interface wrapper for RawCMatrixer type
type _gonum_org_v1_gonum_mat_RawCMatrixer struct {
	WRawCMatrix func() cblas128.General
}

func (W _gonum_org_v1_gonum_mat_RawCMatrixer) RawCMatrix() cblas128.General { return W.WRawCMatrix() }

// _gonum_org_v1_gonum_mat_RawColViewer is an interface wrapper for RawColViewer type
type _gonum_org_v1_gonum_mat_RawColViewer struct {
	WRawColView func(j int) []float64
}

func (W _gonum_org_v1_gonum_mat_RawColViewer) RawColView(j int) []float64 { return W.WRawColView(j) }

// _gonum_org_v1_gonum_mat_RawMatrixSetter is an interface wrapper for RawMatrixSetter type
type _gonum_org_v1_gonum_mat_RawMatrixSetter struct {
	WSetRawMatrix func(a blas64.General)
}

func (W _gonum_org_v1_gonum_mat_RawMatrixSetter) SetRawMatrix(a blas64.General) { W.WSetRawMatrix(a) }

// _gonum_org_v1_gonum_mat_RawMatrixer is an interface wrapper for RawMatrixer type
type _gonum_org_v1_gonum_mat_RawMatrixer struct {
	WRawMatrix func() blas64.General
}

func (W _gonum_org_v1_gonum_mat_RawMatrixer) RawMatrix() blas64.General { return W.WRawMatrix() }

// _gonum_org_v1_gonum_mat_RawRowViewer is an interface wrapper for RawRowViewer type
type _gonum_org_v1_gonum_mat_RawRowViewer struct {
	WRawRowView func(i int) []float64
}

func (W _gonum_org_v1_gonum_mat_RawRowViewer) RawRowView(i int) []float64 { return W.WRawRowView(i) }

// _gonum_org_v1_gonum_mat_RawSymBander is an interface wrapper for RawSymBander type
type _gonum_org_v1_gonum_mat_RawSymBander struct {
	WRawSymBand func() blas64.SymmetricBand
}

func (W _gonum_org_v1_gonum_mat_RawSymBander) RawSymBand() blas64.SymmetricBand {
	return W.WRawSymBand()
}

// _gonum_org_v1_gonum_mat_RawSymmetricer is an interface wrapper for RawSymmetricer type
type _gonum_org_v1_gonum_mat_RawSymmetricer struct {
	WRawSymmetric func() blas64.Symmetric
}

func (W _gonum_org_v1_gonum_mat_RawSymmetricer) RawSymmetric() blas64.Symmetric {
	return W.WRawSymmetric()
}

// _gonum_org_v1_gonum_mat_RawTriBander is an interface wrapper for RawTriBander type
type _gonum_org_v1_gonum_mat_RawTriBander struct {
	WRawTriBand func() blas64.TriangularBand
}

func (W _gonum_org_v1_gonum_mat_RawTriBander) RawTriBand() blas64.TriangularBand {
	return W.WRawTriBand()
}

// _gonum_org_v1_gonum_mat_RawTriangular is an interface wrapper for RawTriangular type
type _gonum_org_v1_gonum_mat_RawTriangular struct {
	WRawTriangular func() blas64.Triangular
}

func (W _gonum_org_v1_gonum_mat_RawTriangular) RawTriangular() blas64.Triangular {
	return W.WRawTriangular()
}

// _gonum_org_v1_gonum_mat_RawVectorer is an interface wrapper for RawVectorer type
type _gonum_org_v1_gonum_mat_RawVectorer struct {
	WRawVector func() blas64.Vector
}

func (W _gonum_org_v1_gonum_mat_RawVectorer) RawVector() blas64.Vector { return W.WRawVector() }

// _gonum_org_v1_gonum_mat_Reseter is an interface wrapper for Reseter type
type _gonum_org_v1_gonum_mat_Reseter struct {
	WReset func()
}

func (W _gonum_org_v1_gonum_mat_Reseter) Reset() { W.WReset() }

// _gonum_org_v1_gonum_mat_RowNonZeroDoer is an interface wrapper for RowNonZeroDoer type
type _gonum_org_v1_gonum_mat_RowNonZeroDoer struct {
	WDoRowNonZero func(i int, fn func(i int, j int, v float64))
}

func (W _gonum_org_v1_gonum_mat_RowNonZeroDoer) DoRowNonZero(i int, fn func(i int, j int, v float64)) {
	W.WDoRowNonZero(i, fn)
}

// _gonum_org_v1_gonum_mat_RowViewer is an interface wrapper for RowViewer type
type _gonum_org_v1_gonum_mat_RowViewer struct {
	WRowView func(i int) mat.Vector
}

func (W _gonum_org_v1_gonum_mat_RowViewer) RowView(i int) mat.Vector { return W.WRowView(i) }

// _gonum_org_v1_gonum_mat_SymBanded is an interface wrapper for SymBanded type
type _gonum_org_v1_gonum_mat_SymBanded struct {
	WAt        func(i int, j int) float64
	WBandwidth func() (kl int, ku int)
	WDims      func() (r int, c int)
	WSymBand   func() (n int, k int)
	WSymmetric func() int
	WT         func() mat.Matrix
	WTBand     func() mat.Banded
}

func (W _gonum_org_v1_gonum_mat_SymBanded) At(i int, j int) float64     { return W.WAt(i, j) }
func (W _gonum_org_v1_gonum_mat_SymBanded) Bandwidth() (kl int, ku int) { return W.WBandwidth() }
func (W _gonum_org_v1_gonum_mat_SymBanded) Dims() (r int, c int)        { return W.WDims() }
func (W _gonum_org_v1_gonum_mat_SymBanded) SymBand() (n int, k int)     { return W.WSymBand() }
func (W _gonum_org_v1_gonum_mat_SymBanded) Symmetric() int              { return W.WSymmetric() }
func (W _gonum_org_v1_gonum_mat_SymBanded) T() mat.Matrix               { return W.WT() }
func (W _gonum_org_v1_gonum_mat_SymBanded) TBand() mat.Banded           { return W.WTBand() }

// _gonum_org_v1_gonum_mat_Symmetric is an interface wrapper for Symmetric type
type _gonum_org_v1_gonum_mat_Symmetric struct {
	WAt        func(i int, j int) float64
	WDims      func() (r int, c int)
	WSymmetric func() int
	WT         func() mat.Matrix
}

func (W _gonum_org_v1_gonum_mat_Symmetric) At(i int, j int) float64 { return W.WAt(i, j) }
func (W _gonum_org_v1_gonum_mat_Symmetric) Dims() (r int, c int)    { return W.WDims() }
func (W _gonum_org_v1_gonum_mat_Symmetric) Symmetric() int          { return W.WSymmetric() }
func (W _gonum_org_v1_gonum_mat_Symmetric) T() mat.Matrix           { return W.WT() }

// _gonum_org_v1_gonum_mat_Tracer is an interface wrapper for Tracer type
type _gonum_org_v1_gonum_mat_Tracer struct {
	WTrace func() float64
}

func (W _gonum_org_v1_gonum_mat_Tracer) Trace() float64 { return W.WTrace() }

// _gonum_org_v1_gonum_mat_TriBanded is an interface wrapper for TriBanded type
type _gonum_org_v1_gonum_mat_TriBanded struct {
	WAt        func(i int, j int) float64
	WBandwidth func() (kl int, ku int)
	WDims      func() (r int, c int)
	WT         func() mat.Matrix
	WTBand     func() mat.Banded
	WTTri      func() mat.Triangular
	WTTriBand  func() mat.TriBanded
	WTriBand   func() (n int, k int, kind mat.TriKind)
	WTriangle  func() (n int, kind mat.TriKind)
}

func (W _gonum_org_v1_gonum_mat_TriBanded) At(i int, j int) float64     { return W.WAt(i, j) }
func (W _gonum_org_v1_gonum_mat_TriBanded) Bandwidth() (kl int, ku int) { return W.WBandwidth() }
func (W _gonum_org_v1_gonum_mat_TriBanded) Dims() (r int, c int)        { return W.WDims() }
func (W _gonum_org_v1_gonum_mat_TriBanded) T() mat.Matrix               { return W.WT() }
func (W _gonum_org_v1_gonum_mat_TriBanded) TBand() mat.Banded           { return W.WTBand() }
func (W _gonum_org_v1_gonum_mat_TriBanded) TTri() mat.Triangular        { return W.WTTri() }
func (W _gonum_org_v1_gonum_mat_TriBanded) TTriBand() mat.TriBanded     { return W.WTTriBand() }
func (W _gonum_org_v1_gonum_mat_TriBanded) TriBand() (n int, k int, kind mat.TriKind) {
	return W.WTriBand()
}
func (W _gonum_org_v1_gonum_mat_TriBanded) Triangle() (n int, kind mat.TriKind) { return W.WTriangle() }

// _gonum_org_v1_gonum_mat_Triangular is an interface wrapper for Triangular type
type _gonum_org_v1_gonum_mat_Triangular struct {
	WAt       func(i int, j int) float64
	WDims     func() (r int, c int)
	WT        func() mat.Matrix
	WTTri     func() mat.Triangular
	WTriangle func() (n int, kind mat.TriKind)
}

func (W _gonum_org_v1_gonum_mat_Triangular) At(i int, j int) float64             { return W.WAt(i, j) }
func (W _gonum_org_v1_gonum_mat_Triangular) Dims() (r int, c int)                { return W.WDims() }
func (W _gonum_org_v1_gonum_mat_Triangular) T() mat.Matrix                       { return W.WT() }
func (W _gonum_org_v1_gonum_mat_Triangular) TTri() mat.Triangular                { return W.WTTri() }
func (W _gonum_org_v1_gonum_mat_Triangular) Triangle() (n int, kind mat.TriKind) { return W.WTriangle() }

// _gonum_org_v1_gonum_mat_Unconjugator is an interface wrapper for Unconjugator type
type _gonum_org_v1_gonum_mat_Unconjugator struct {
	WUnconjugate func() mat.CMatrix
}

func (W _gonum_org_v1_gonum_mat_Unconjugator) Unconjugate() mat.CMatrix { return W.WUnconjugate() }

// _gonum_org_v1_gonum_mat_UntransposeBander is an interface wrapper for UntransposeBander type
type _gonum_org_v1_gonum_mat_UntransposeBander struct {
	WUntransposeBand func() mat.Banded
}

func (W _gonum_org_v1_gonum_mat_UntransposeBander) UntransposeBand() mat.Banded {
	return W.WUntransposeBand()
}

// _gonum_org_v1_gonum_mat_UntransposeTriBander is an interface wrapper for UntransposeTriBander type
type _gonum_org_v1_gonum_mat_UntransposeTriBander struct {
	WUntransposeTriBand func() mat.TriBanded
}

func (W _gonum_org_v1_gonum_mat_UntransposeTriBander) UntransposeTriBand() mat.TriBanded {
	return W.WUntransposeTriBand()
}

// _gonum_org_v1_gonum_mat_UntransposeTrier is an interface wrapper for UntransposeTrier type
type _gonum_org_v1_gonum_mat_UntransposeTrier struct {
	WUntransposeTri func() mat.Triangular
}

func (W _gonum_org_v1_gonum_mat_UntransposeTrier) UntransposeTri() mat.Triangular {
	return W.WUntransposeTri()
}

// _gonum_org_v1_gonum_mat_Untransposer is an interface wrapper for Untransposer type
type _gonum_org_v1_gonum_mat_Untransposer struct {
	WUntranspose func() mat.Matrix
}

func (W _gonum_org_v1_gonum_mat_Untransposer) Untranspose() mat.Matrix { return W.WUntranspose() }

// _gonum_org_v1_gonum_mat_Vector is an interface wrapper for Vector type
type _gonum_org_v1_gonum_mat_Vector struct {
	WAt    func(i int, j int) float64
	WAtVec func(a0 int) float64
	WDims  func() (r int, c int)
	WLen   func() int
	WT     func() mat.Matrix
}

func (W _gonum_org_v1_gonum_mat_Vector) At(i int, j int) float64 { return W.WAt(i, j) }
func (W _gonum_org_v1_gonum_mat_Vector) AtVec(a0 int) float64    { return W.WAtVec(a0) }
func (W _gonum_org_v1_gonum_mat_Vector) Dims() (r int, c int)    { return W.WDims() }
func (W _gonum_org_v1_gonum_mat_Vector) Len() int                { return W.WLen() }
func (W _gonum_org_v1_gonum_mat_Vector) T() mat.Matrix           { return W.WT() }