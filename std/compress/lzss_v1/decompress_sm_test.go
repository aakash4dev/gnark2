package lzss_v1

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test1ZeroSm(t *testing.T) {
	testCompressionRoundTripSm(t, 1, []byte{0})
	testCompressionRoundTripSm(t, 2, []byte{0, 0, 0, 0, 0, 0, 0, 0})
}

func Test2ZeroSm(t *testing.T) {
	testCompressionRoundTripSm(t, 1, []byte{0, 0, 0})
	testCompressionRoundTripSm(t, 2, []byte{0, 0, 0, 0, 0, 0, 0, 0})
}

func Test8ZerosSm(t *testing.T) {
	testCompressionRoundTripSm(t, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0})
	testCompressionRoundTripSm(t, 2, []byte{0, 0, 0, 0, 0, 0, 0, 0})
}

func TestTwoConsecutiveBackrefs(t *testing.T) {
	c := make([]byte, 9)
	d := c[:2]
	dBack := make([]byte, 4)
	settings := Settings{
		BackRefSettings: BackRefSettings{
			NbBytesAddress: 1,
			NbBytesLength:  1,
			Symbol:         0,
		},
		Logger:   nil,
		LogHeads: new([]LogHeads),
	}

	dLen, err := decompressStateMachine(c, 6, dBack, settings)
	require.NoError(t, err)
	require.Equal(t, d, dBack[:dLen])
}
func Test300ZerosSm(t *testing.T) { // probably won't happen in our calldata
	testCompressionRoundTripSm(t, 1, make([]byte, 300))
	testCompressionRoundTripSm(t, 2, make([]byte, 300))
}

func TestNoCompressionSm(t *testing.T) {
	testCompressionRoundTripSm(t, 1, []byte{'h', 'i'})
	testCompressionRoundTripSm(t, 2, []byte{'h', 'i'})
}

func TestZeroAfterNonzeroSm(t *testing.T) { // probably won't happen in our calldata
	testCompressionRoundTripSm(t, 1, []byte{1, 0})
	testCompressionRoundTripSm(t, 2, []byte{1, 0})
}

func TestTwoZerosAfterNonzeroSm(t *testing.T) { // probably won't happen in our calldata
	testCompressionRoundTripSm(t, 1, append([]byte{1}, make([]byte, 8)...))
	//testCompressionRoundTripSm(t, 2, append([]byte{1}, make([]byte, 2)...))
}

func Test8ZerosAfterNonzeroSm(t *testing.T) { // probably won't happen in our calldata
	//testCompressionRoundTripSm(t, 1, append([]byte{1}, make([]byte, 8)...))
	testCompressionRoundTripSm(t, 2, append([]byte{1}, make([]byte, 8)...))
}

func Test300ZerosAfterNonzeroSm(t *testing.T) { // probably won't happen in our calldata
	testCompressionRoundTripSm(t, 1, append([]byte{'h', 'i'}, make([]byte, 300)...))
	testCompressionRoundTripSm(t, 2, append([]byte{'h', 'i'}, make([]byte, 300)...))
}

func TestRepeatedNonzeroSm(t *testing.T) {
	testCompressionRoundTripSm(t, 1, []byte{'h', 'i', 'h', 'i', 'h', 'i'})
	testCompressionRoundTripSm(t, 2, []byte{'h', 'i', 'h', 'i', 'h', 'i'})
}

func TestCalldataSm(t *testing.T) {
	t.Parallel()
	folders := []string{
		"3c2943",
	}
	for _, folder := range folders {
		d, err := os.ReadFile("../" + folder + "/data.bin")
		require.NoError(t, err)
		t.Run(folder, func(t *testing.T) {
			testCompressionRoundTripSm(t, 2, d)
		})
	}
}

func testCompressionRoundTripSm(t *testing.T, nbBytesOffset uint, d []byte) {
	settings := Settings{
		BackRefSettings: BackRefSettings{
			NbBytesAddress: nbBytesOffset,
			NbBytesLength:  1,
			Symbol:         0,
			ReferenceTo:    false,
			AddressingMode: false,
		},
		Logger:   nil,
		LogHeads: new([]LogHeads),
	}

	c, err := Compress(d, settings)
	c = append(c, make([]byte, 2*len(c))...)
	require.NoError(t, err)
	dBack := make([]byte, len(d)*2)

	T := func(padCoeff int) {
		dLength, err := decompressStateMachine(c, len(c)/3, dBack, settings)
		require.NoError(t, err)
		require.Equal(t, d, dBack[:dLength])
	}

	T(2)
	T(1)
}
