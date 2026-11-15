package visualizer

import (
	"bytes"
	"encoding/binary"
	"math"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestAnalyzerReadsSpectrum(t *testing.T) {
	analyzer := New(8)
	if !analyzer.Available() {
		t.Skip("ffmpeg not available")
	}

	path := filepath.Join(t.TempDir(), "tone.wav")
	writeToneWAV(t, path)

	if err := analyzer.Start(path, 0); err != nil {
		t.Fatalf("Start() error = %v", err)
	}
	defer analyzer.Stop()

	deadline := time.Now().Add(1200 * time.Millisecond)
	for time.Now().Before(deadline) {
		for _, level := range analyzer.Levels() {
			if level > 0.02 {
				return
			}
		}
		time.Sleep(40 * time.Millisecond)
	}
	t.Fatal("analyzer levels never moved")
}

func TestShapeLevelsKeepsBassTransientPresentButControlled(t *testing.T) {
	analyzer := New(12)

	quiet := make([]float64, 12)
	for i := 0; i < 10; i++ {
		_ = analyzer.shapeLevels(quiet)
	}

	thump := make([]float64, 12)
	for i := 0; i < 4; i++ {
		thump[i] = 12000
	}

	levels := analyzer.shapeLevels(thump)
	low := average(levels[:4])
	high := average(levels[7:])

	if low < 0.45 {
		t.Fatalf("low-band thump response = %.3f, want at least 0.45", low)
	}
	if low <= high+0.18 {
		t.Fatalf("low-band response %.3f should dominate high-band response %.3f", low, high)
	}
	if analyzer.beatEnv > 0.95 {
		t.Fatalf("beat envelope = %.3f, want controlled response below saturation", analyzer.beatEnv)
	}
}

func writeToneWAV(t *testing.T, path string) {
	t.Helper()

	const (
		sampleRate    = 11025
		duration      = 600 * time.Millisecond
		frequency     = 440.0
		bitsPerSample = 16
		channels      = 1
	)

	sampleCount := int(duration.Seconds() * sampleRate)
	dataSize := sampleCount * channels * bitsPerSample / 8

	var buf bytes.Buffer
	buf.WriteString("RIFF")
	_ = binary.Write(&buf, binary.LittleEndian, uint32(36+dataSize))
	buf.WriteString("WAVE")
	buf.WriteString("fmt ")
	_ = binary.Write(&buf, binary.LittleEndian, uint32(16))
	_ = binary.Write(&buf, binary.LittleEndian, uint16(1))
	_ = binary.Write(&buf, binary.LittleEndian, uint16(channels))
	_ = binary.Write(&buf, binary.LittleEndian, uint32(sampleRate))
	_ = binary.Write(&buf, binary.LittleEndian, uint32(sampleRate*channels*bitsPerSample/8))
	_ = binary.Write(&buf, binary.LittleEndian, uint16(channels*bitsPerSample/8))
	_ = binary.Write(&buf, binary.LittleEndian, uint16(bitsPerSample))
	buf.WriteString("data")
	_ = binary.Write(&buf, binary.LittleEndian, uint32(dataSize))

	for i := 0; i < sampleCount; i++ {
		phase := 2 * math.Pi * frequency * float64(i) / sampleRate
		sample := int16(math.Sin(phase) * 18000)
		_ = binary.Write(&buf, binary.LittleEndian, sample)
	}

	if err := os.WriteFile(path, buf.Bytes(), 0o644); err != nil {
		t.Fatalf("write wav: %v", err)
	}
}
