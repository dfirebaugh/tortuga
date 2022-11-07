package stream

import (
	"math"
)

type stream struct {
	position   int64
	remaining  []byte
	frequency  int
	sampleRate int
}

func New(sampleRate int, frequency int) *stream {
	return &stream{
		sampleRate: sampleRate,
		frequency:  frequency,
	}
}

// Read is io.Reader's Read.
//
// Read fills the data with sine wave samples.
func (s *stream) Read(buf []byte) (int, error) {
	if len(s.remaining) > 0 {
		n := copy(buf, s.remaining)
		s.remaining = s.remaining[n:]
		return n, nil
	}

	var origBuf []byte
	if len(buf)%4 > 0 {
		origBuf = buf
		buf = make([]byte, len(origBuf)+4-len(origBuf)%4)
	}

	var length = int64(s.sampleRate / s.frequency)
	p := s.position / 4
	for i := 0; i < len(buf)/4; i++ {
		const max = 32767
		b := int16(math.Sin(2*math.Pi*float64(p)/float64(length)) * max)
		buf[4*i] = byte(b)
		buf[4*i+1] = byte(b >> 8)
		buf[4*i+2] = byte(b)
		buf[4*i+3] = byte(b >> 8)
		p++
	}

	s.position += int64(len(buf))
	s.position %= length * 4

	if origBuf != nil {
		n := copy(origBuf, buf)
		s.remaining = buf[n:]
		return n, nil
	}
	return len(buf), nil
}

// Close is io.Closer's Close.
func (s *stream) Close() error {
	return nil
}
