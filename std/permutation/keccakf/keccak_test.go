package keccakf_test

import (
	"testing"

	"github.com/aakash4dev/gnark2/std/math/uints"

	"github.com/aakash4dev/gnark2/backend"
	"github.com/aakash4dev/gnark2/frontend"
	"github.com/aakash4dev/gnark2/std/permutation/keccakf"
	"github.com/aakash4dev/gnark2/test"
	"github.com/consensys/gnark-crypto/ecc"
)

type keccakfCircuit struct {
	In       [25]uints.U64
	Expected [25]uints.U64 `gnark:",public"`
}

func (c *keccakfCircuit) Define(api frontend.API) error {
	var res [25]uints.U64
	for i := range res {
		res[i] = c.In[i]
	}
	uapi, err := uints.New[uints.U64](api)
	if err != nil {
		return err
	}
	for i := 0; i < 2; i++ {
		res = keccakf.Permute(uapi, res)
	}
	for i := range res {
		uapi.AssertEq(res[i], c.Expected[i])
	}
	return nil
}

func TestKeccakf(t *testing.T) {
	var nativeIn [25]uint64
	var res [25]uint64
	for i := range nativeIn {
		nativeIn[i] = 2
		res[i] = 2
	}
	for i := 0; i < 2; i++ {
		res = keccakF1600(res)
	}
	witness := keccakfCircuit{}
	for i := range nativeIn {
		witness.In[i] = uints.NewU64(nativeIn[i])
		witness.Expected[i] = uints.NewU64(res[i])
	}
	assert := test.NewAssert(t)
	assert.ProverSucceeded(&keccakfCircuit{}, &witness,
		test.WithCurves(ecc.BN254),
		test.WithBackends(backend.GROTH16, backend.PLONK),
		test.NoProverChecks(),
		test.NoFuzzing())
}
