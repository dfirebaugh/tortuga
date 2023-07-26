package stream

import (
	"math"
	"math/rand"
)

type stream struct {
	position       int64
	remaining      []byte
	frequency      float32
	sampleRate     int
	fadeInSamples  int
	fadeOutSamples int
}

func New(sampleRate int, frequency float32, fadeInMs int, fadeOutMs int) *stream {
	return &stream{
		sampleRate:     sampleRate,
		frequency:      frequency,
		fadeInSamples:  fadeInMs * sampleRate / 1000,
		fadeOutSamples: fadeOutMs * sampleRate / 1000,
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

	var length = int64(s.sampleRate / int(s.frequency))
	p := s.position / 4
	for i := 0; i < len(buf)/4; i++ {
		const max = 32767
		// b := int16(math.Sin(2*math.Pi*float64(p)/float64(length)) * max)
		b := s.applyEnvelope(p, length, max)
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

func (s *stream) applyEnvelope(p int64, length int64, max int) int16 {
	envelope := 1.0
	if p < int64(s.fadeInSamples) {
		envelope = float64(p) / float64(s.fadeInSamples)
	} else if p > length-int64(s.fadeOutSamples) {
		envelope = float64(length-p) / float64(s.fadeOutSamples)
	}
	return s.applyDither(int16(0.5 * envelope * math.Sin(2*math.Pi*float64(p)/float64(length)) * float64(max)))
}

func (s *stream) applyDither(sample int16) int16 {
	dither := (rand.Float64() * 1) - .01
	sampleWithDither := float64(sample) + dither

	// Check for clipping
	if sampleWithDither > 32767 {
		sampleWithDither = 32767
	} else if sampleWithDither < -32768 {
		sampleWithDither = -32768
	}

	return int16(sampleWithDither)
}

// Close is io.Closer's Close.
func (s *stream) Close() error {
	return nil
}
