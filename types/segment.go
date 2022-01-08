package types

// SegmentComplex64 is a sample segment of type complex64
type SegmentComplex64 struct {
	SampleRate    int
	Frequency     int
	Data          []complex64
	SegmentNumber int
}

// SegmentCS8Raw is a sample segment of type complex64.  Assumes little endian.
type SegmentCS8Raw struct {
	SampleRate    int
	Frequency     int
	Data          []byte
	SegmentNumber int
}

// SegmentFloat32 is a real-valued segment of type float32.
type SegmentFloat32 struct {
	SampleRate    int
	Frequency     int
	Data          []float32
	SegmentNumber int
}

// SegmentBinaryBytes uses one byte per received bit
type SegmentBinaryBytes struct {
	SymbolRate    int
	Frequency     int
	Data          []byte
	SegmentNumber int
}

// SegmentBinaryBits packs bits into bytes.  The number of bits in the segment
// is contained in DataBits
type SegmentBinaryBits struct {
	SymbolRate    int
	Frequency     int
	DataBits      int
	Data          []byte
	SegmentNumber int
}

type SegmentInt16 struct {
	Frequency     int
	SampleRate    int
	Data          []int16
	SegmentNumber int
}

func (s *SegmentBinaryBytes) ToSegmentBinaryBits() *SegmentBinaryBits {
	ret := make([]byte, len(s.Data)/8+1)
	for i := 0; i < len(s.Data); i++ {
		curByteValue := ret[i/8]
		pos := 7 - (i % 8)
		val := s.Data[i]
		ret[i/8] = curByteValue | (val << pos)
	}
	return &SegmentBinaryBits{
		SymbolRate: s.SymbolRate,
		Frequency:  s.Frequency,
		DataBits:   len(s.Data),
		Data:       ret,
	}
}

func (s *SegmentComplex64) ToReal32() *SegmentFloat32 {
	ret := make([]float32, len(s.Data))
	for i := 0; i < len(s.Data); i++ {
		ret[i] = real(s.Data[i])
	}
	return &SegmentFloat32{
		SampleRate: s.SampleRate,
		Frequency:  s.Frequency,
		Data:       ret,
	}
}

func (s *SegmentCS8Raw) ToComplex64() *SegmentComplex64 {
	ret := make([]complex64, len(s.Data)/2)
	for retIndex, byteIndex := 0, 0; byteIndex < len(s.Data); retIndex, byteIndex = retIndex+1, byteIndex+2 {
		// ** Assumes little endian **
		Q := float32(int8(s.Data[byteIndex]))
		I := float32(int8(s.Data[byteIndex+1]))
		ret[retIndex] = complex(I, Q)
	}

	return &SegmentComplex64{
		SampleRate: s.SampleRate,
		Data:       ret,
		Frequency:  s.Frequency,
	}
}
