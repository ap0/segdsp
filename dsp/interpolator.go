package dsp

type Interpolator struct {
	sampleHistory      []complex64
	fir                *FirFilter
	interpolationRatio int
}

func MakeInterpolator(interpolationRatio int) *Interpolator {
	return &Interpolator{
		fir:                MakeFirFilter(MakeLowPassFixed(1, 1, 1/float64(interpolationRatio*2), 63)),
		interpolationRatio: interpolationRatio,
		sampleHistory:      make([]complex64, 1),
	}
}

func (f *Interpolator) Work(data []complex64) []complex64 {
	samples := append(f.sampleHistory, data...)

	var output = make([]complex64, len(data)*f.interpolationRatio)

	for i := 0; i < len(data); i++ {
		var idx = i * f.interpolationRatio
		output[idx] = samples[i]
		next := samples[i+1]

		dr := float64(real(next) - real(samples[i]))
		di := float64(imag(next) - imag(samples[i]))

		for j := 1; j < f.interpolationRatio; j++ {
			mult := float64(j) / float64(f.interpolationRatio)
			output[idx+j] = complex(
				real(samples[i])+float32(dr*mult),
				imag(samples[i])+float32(di*mult))
		}
	}

	f.fir.Filter(output, len(output))
	f.sampleHistory = samples[len(data):]
	return output
}

func (f *Interpolator) WorkBuffer(input, output []complex64) int {
	var oLen = len(input) * f.interpolationRatio
	if len(output) < oLen {
		panic("Output buffer does not have enough length")
	}
	for i := 0; i < len(input); i++ {
		var idx = i * f.interpolationRatio
		output[idx] = input[i]
		for j := 1; j < f.interpolationRatio; j++ {
			output[idx+j] = complex(0, 0)
		}
	}

	f.fir.Filter(output, oLen)

	return oLen
}

func (f *Interpolator) PredictOutputSize(inputLength int) int {
	return inputLength * f.interpolationRatio
}

type FloatInterpolator struct {
	fir                *FloatFirFilter
	interpolationRatio int
}

func MakeFloatInterpolator(interpolationRatio int) *FloatInterpolator {
	return &FloatInterpolator{
		fir:                MakeFloatFirFilter(MakeLowPassFixed(1, 1, 1/float64(interpolationRatio*2), 63)),
		interpolationRatio: interpolationRatio,
	}
}

func (f *FloatInterpolator) Work(data []float32) []float32 {
	var output = make([]float32, len(data)*f.interpolationRatio)

	for i := 0; i < len(data); i++ {
		var idx = i * f.interpolationRatio
		output[idx] = data[i]
		for j := 1; j < f.interpolationRatio; j++ {
			output[idx+j] = 0
		}
	}

	f.fir.Filter(output, len(output))
	return output
}

func (f *FloatInterpolator) WorkBuffer(input, output []float32) int {
	var oLen = len(input) * f.interpolationRatio
	if len(output) < oLen {
		panic("There is not enough space in output buffer")
	}
	for i := 0; i < len(input); i++ {
		var idx = i * f.interpolationRatio
		output[idx] = input[i]
		for j := 1; j < f.interpolationRatio; j++ {
			output[idx+j] = complex(0, 0)
		}
	}

	f.fir.Filter(output, oLen)
	return oLen
}

func (f *FloatInterpolator) PredictOutputSize(inputLength int) int {
	return inputLength * f.interpolationRatio
}
