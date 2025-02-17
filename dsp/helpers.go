package dsp

import (
	"math"

	"github.com/racerxdl/segdsp/dsp/native"
	"github.com/racerxdl/segdsp/tools"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ComplexDotProductResult performs the Dot Product between two complex vectors and returns result
var ComplexDotProductResult func(input []complex64, taps []complex64) complex64

// DotProductResult performs the Dot Product between a complex vector and a float32 vector and returns result
var DotProductResult func(input []complex64, taps []float32) complex64

// DotProductFloatResult performs the Dot Product between two float32 vectors and returns result
var DotProductFloatResult func(input []float32, taps []float32) float32

// MultiplyConjugate performs the Multply by conjugate for each item from vecA and vecB
// output[i] = vecA[i] * Conj(vecB[i])
var MultiplyConjugate func(vecA, vecB []complex64, length int) []complex64

// MultiplyConjugateInline performs the Multply by conjugate for each item from vecA and vecB with the result in vecA
// vecA[i] = vecA[i] * Conj(vecB[i])
var MultiplyConjugateInline func(vecA, vecB []complex64, length int)

// RotateComplex performs the phase rotation of each input items by phase and then increments phase
// out[i] = input[i] * phase
// phase = phase * phaseIncrement
var RotateComplex func(input []complex64, phase *complex64, phaseIncrement complex64, length int) []complex64

// RotateComplex performs the phase rotation of each input items by phase and then increments phase
// out[i] = input[i] * phase
// phase = phase * phaseIncrement
var RotateComplexBuffer func(input, output []complex64, phase *complex64, phaseIncrement complex64, length int) int

// MultiplyFloatFloatVectors performs multiplication of each element from A by the same element in B
// A[i] = A[i] * B[i]
var MultiplyFloatFloatVectors func(A, B []float32)

// DivideFloatFloatVectors performs division of each element from A by the same element in B
// A[i] = A[i] / B[i]
var DivideFloatFloatVectors func(A, B []float32)

// AddFloatFloatVectors performs addition of each element from A by the same element in B
// A[i] = A[i] + B[i]
var AddFloatFloatVectors func(A, B []float32)

// SubtractFloatFloatVectors performs addition of each element from A by the same element in B
// A[i] = A[i] - B[i]
var SubtractFloatFloatVectors func(A, B []float32)

// MultiplyComplexComplexVectors performs multiplication of each element from A by the same element in B
// A[i] = A[i] * B[i]
var MultiplyComplexComplexVectors func(A, B []complex64)

// DivideComplexComplexVectors performs division of each element from A by the same element in B
// A[i] = A[i] / B[i]
var DivideComplexComplexVectors func(A, B []complex64)

// AddComplexComplexVectors performs addition of each element from A by the same element in B
// A[i] = A[i] + B[i]
var AddComplexComplexVectors func(A, B []complex64)

// SubtractComplexComplexVectors performs addition of each element from A by the same element in B
// A[i] = A[i] - B[i]
var SubtractComplexComplexVectors func(A, B []complex64)

// ComplexDotProduct performs the Dot Product between two complex vectors and store the result at *result
func ComplexDotProduct(result *complex64, input []complex64, taps []complex64) {
	*result = ComplexDotProductResult(input, taps)
}

// DotProduct performs the Dot Product between a complex vector and a float vector and store the result at *result
func DotProduct(result *complex64, input []complex64, taps []float32) {
	*result = DotProductResult(input, taps)
}

// DotProductFloat performs the Dot Product between two float vectors and store the result at *result
func DotProductFloat(result *float32, input []float32, taps []float32) {
	*result = DotProductFloatResult(input, taps)
}

func Modulus(c complex64) float32 {
	return float32(math.Sqrt(float64(real(c)*real(c) + imag(c)*imag(c))))
}

func Divide(c complex64, f float32) complex64 {
	var b = 1 / f
	return complex(real(c)*b, imag(c)*b)
}

func Argument(c complex64) float32 {
	return float32(math.Atan2(float64(imag(c)), float64(real(c))))
}

// GetSIMDMode returns a string containg the current SIMD mode used.
func GetSIMDMode() string {
	return native.GetSIMDMode()
}

// region Private Functions

// genericMultiplyFloatFloatVectors performs multiplication of each element from A by the same element in B
// This is the Generic Function in case no SIMD alternative is available
func genericMultiplyFloatFloatVectors(A, B []float32) {
	for i, v := range B {
		A[i] = A[i] * v
	}
}

// genericDivideFloatFloatVectors performs division of each element from A by the same element in B
// This is the Generic Function in case no SIMD alternative is available
func genericDivideFloatFloatVectors(A, B []float32) {
	for i, v := range B {
		A[i] = A[i] / v
	}
}

// genericAddFloatFloatVectors performs addition of each element from A by the same element in B
// This is the Generic Function in case no SIMD alternative is available
func genericAddFloatFloatVectors(A, B []float32) {
	for i, v := range B {
		A[i] = A[i] + v
	}
}

// genericSubtractFloatFloatVectors performs subtraction of each element from A by the same element in B
// This is the Generic Function in case no SIMD alternative is available
func genericSubtractFloatFloatVectors(A, B []float32) {
	for i, v := range B {
		A[i] = A[i] - v
	}
}

// genericMultiplyComplexComplexVectors performs multiplication of each element from A by the same element in B
// This is the Generic Function in case no SIMD alternative is available
func genericMultiplyComplexComplexVectors(A, B []complex64) {
	for i, v := range B {
		A[i] = A[i] * v
	}
}

// genericDivideComplexComplexVectors performs division of each element from A by the same element in B
// This is the Generic Function in case no SIMD alternative is available
func genericDivideComplexComplexVectors(A, B []complex64) {
	for i, v := range B {
		A[i] = A[i] / v
	}
}

// genericAddComplexComplexVectors performs addition of each element from A by the same element in B
// This is the Generic Function in case no SIMD alternative is available
func genericAddComplexComplexVectors(A, B []complex64) {
	for i, v := range B {
		A[i] = A[i] + v
	}
}

// genericSubtractComplexComplexVectors performs subtraction of each element from A by the same element in B
// This is the Generic Function in case no SIMD alternative is available
func genericSubtractComplexComplexVectors(A, B []complex64) {
	for i, v := range B {
		A[i] = A[i] - v
	}
}

// genericComplexDotProductResult performs the Dot Product between two complex vectors and returns the result
// This is the Generic Function in case no SIMD alternative is available
func genericComplexDotProductResult(input []complex64, taps []complex64) complex64 {
	var length = Min(len(taps), len(input))

	var res = complex64(complex(0, 0))

	for i := 0; i < length; i++ {
		var r = real(input[i])*real(taps[i]) - imag(input[i])*imag(taps[i])
		var i = real(input[i])*imag(taps[i]) + imag(input[i])*real(taps[i])

		res += complex(r, i)
	}

	return res
}

// genericDotProductResult performs the Dot Product between a complex vector and a float vector and returns the result
// This is the Generic Function in case no SIMD alternative is available
func genericDotProductResult(input []complex64, taps []float32) complex64 {
	var length = Min(len(taps), len(input))
	var res [2]float32

	for i := 0; i < length; i++ {
		res[0] += real(input[i]) * taps[i]
		res[1] += imag(input[i]) * taps[i]
	}

	return complex(res[0], res[1])
}

// genericDotProductFloatResult performs the Dot Product between two float vectors and returns the result
// This is the Generic Function in case no SIMD alternative is available
func genericDotProductFloatResult(input []float32, taps []float32) float32 {
	var res = float32(0.0)
	var length = Min(len(taps), len(input))

	for i := 0; i < length; i++ {
		res += input[i] * taps[i]
	}

	return res
}

// genericMultiplyConjugate performs the Multply by conjugate for each item from vecA and vecB
// output[i] = vecA[i] * Conj(vecB[i])
// This is the Generic Function in case no SIMD alternative is available
func genericMultiplyConjugate(vecA, vecB []complex64, length int) []complex64 {
	var output = make([]complex64, length)
	for i := 0; i < length; i++ {
		output[i] = vecA[i] * tools.Conj(vecB[i])
	}

	return output
}

// genericMultiplyConjugateInline performs the Multply by conjugate for each item from vecA and vecB with the result in vecA
// vecA[i] = vecA[i] * Conj(vecB[i])
// This is the Generic Function in case no SIMD alternative is available
func genericMultiplyConjugateInline(vecA, vecB []complex64, length int) {
	for i := 0; i < length; i++ {
		vecA[i] = vecA[i] * tools.Conj(vecB[i])
	}
}

// genericRotateComplex performs the phase rotation of each input items by phase and then increments phase
// out[i] = input[i] * phase
// phase = phase * phaseIncrement
// This is the Generic Function in case no SIMD alternative is available
func genericRotateComplex(input []complex64, phase *complex64, phaseIncrement complex64, length int) []complex64 {
	var out = make([]complex64, length)
	var counter = 0

	for i := 0; i < length; i++ {
		counter++
		out[i] = input[i] * (*phase)
		*phase = *phase * phaseIncrement
		if counter%512 == 0 {
			*phase = tools.ComplexNormalize(*phase)
		}
	}

	return out
}

// genericRotateComplex performs the phase rotation of each input items by phase and then increments phase
// out[i] = input[i] * phase
// phase = phase * phaseIncrement
// This is the Generic Function in case no SIMD alternative is available
func genericRotateComplexBuffer(input, output []complex64, phase *complex64, phaseIncrement complex64, length int) int {
	var counter = 0

	if len(input) > len(output) {
		panic("There is not enough space in output buffer")
	}

	for i := 0; i < length; i++ {
		counter++
		output[i] = input[i] * (*phase)
		*phase = *phase * phaseIncrement
		if counter%512 == 0 {
			*phase = tools.ComplexNormalize(*phase)
		}
	}

	return len(input)
}

// init initializes the Helper function placeholders with SIMD Alternatives when available
func init() {
	if native.GetNativeDotProductComplex() != nil {
		DotProductResult = native.GetNativeDotProductComplex()
	} else {
		DotProductResult = genericDotProductResult
	}

	if native.GetNativeDotProductFloat() != nil {
		DotProductFloatResult = native.GetNativeDotProductFloat()
	} else {
		DotProductFloatResult = genericDotProductFloatResult
	}

	// SIMD Multiply Conjugate is actually slower

	//if native.GetMultiplyConjugate() != nil {
	//	MultiplyConjugate = native.GetMultiplyConjugate()
	//} else {
	MultiplyConjugate = genericMultiplyConjugate
	//}

	//if native.GetMultiplyConjugateInline() != nil {
	//	MultiplyConjugateInline = native.GetMultiplyConjugateInline()
	//} else {
	MultiplyConjugateInline = genericMultiplyConjugateInline
	//}

	if native.GetNativeDotProductComplexComplex() != nil {
		ComplexDotProductResult = native.GetNativeDotProductComplexComplex()
	} else {
		ComplexDotProductResult = genericComplexDotProductResult
	}

	if native.GetNativeRotateComplex() != nil {
		RotateComplex = native.GetNativeRotateComplex()
	} else {
		RotateComplex = genericRotateComplex
	}

	if native.GetNativeRotateComplexBuffer() != nil {
		RotateComplexBuffer = native.GetNativeRotateComplexBuffer()
	} else {
		RotateComplexBuffer = genericRotateComplexBuffer
	}

	if native.GetNativeAddFloatFloatVectors() != nil {
		AddFloatFloatVectors = native.GetNativeAddFloatFloatVectors()
	} else {
		AddFloatFloatVectors = genericAddFloatFloatVectors
	}

	// region Float-Float Vector Operations
	if native.GetNativeSubtractFloatFloatVectors() != nil {
		SubtractFloatFloatVectors = native.GetNativeSubtractFloatFloatVectors()
	} else {
		SubtractFloatFloatVectors = genericSubtractFloatFloatVectors
	}

	// The difference between native and golang is not that huge, but native is a bit faster
	if native.GetNativeMultiplyFloatFloatVectors() != nil {
		MultiplyFloatFloatVectors = native.GetNativeMultiplyFloatFloatVectors()
	} else {
		MultiplyFloatFloatVectors = genericMultiplyFloatFloatVectors
	}

	if native.GetNativeDivideFloatFloatVectors() != nil {
		DivideFloatFloatVectors = native.GetNativeDivideFloatFloatVectors()
	} else {
		DivideFloatFloatVectors = genericDivideFloatFloatVectors
	}
	// endregion
	// region Complex-Complex Vector Operations
	if native.GetNativeAddComplexComplexVectors() != nil {
		AddComplexComplexVectors = native.GetNativeAddComplexComplexVectors()
	} else {
		AddComplexComplexVectors = genericAddComplexComplexVectors
	}

	if native.GetNativeSubtractComplexComplexVectors() != nil {
		SubtractComplexComplexVectors = native.GetNativeSubtractComplexComplexVectors()
	} else {
		SubtractComplexComplexVectors = genericSubtractComplexComplexVectors
	}

	// Multiply ComplexComplex in golang is actually faster than native
	//if native.GetNativeMultiplyComplexComplexVectors() != nil {
	//	MultiplyComplexComplexVectors = native.GetNativeMultiplyComplexComplexVectors()
	//} else {
	MultiplyComplexComplexVectors = genericMultiplyComplexComplexVectors
	//}

	if native.GetNativeDivideComplexComplexVectors() != nil {
		DivideComplexComplexVectors = native.GetNativeDivideComplexComplexVectors()
	} else {
		DivideComplexComplexVectors = genericDivideComplexComplexVectors
	}
	// endregion
}

// endregion
