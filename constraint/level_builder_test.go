package constraint_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/aakash4dev/gnark2/constraint/solver"
	"github.com/aakash4dev/gnark2/frontend"
	"github.com/aakash4dev/gnark2/test"
)

func idHint(_ *big.Int, in []*big.Int, out []*big.Int) error {
	if len(in) != len(out) {
		return fmt.Errorf("in/out length mismatch %d≠%d", len(in), len(out))
	}
	for i := range in {
		out[i].Set(in[i])
	}
	return nil
}

type idHintCircuit struct {
	X frontend.Variable
}

func (c *idHintCircuit) Define(api frontend.API) error {
	x, err := api.Compiler().NewHint(idHint, 1, api.Mul(c.X, c.X))
	if err != nil {
		return err
	}
	api.AssertIsEqual(x[0], api.Mul(c.X, c.X))
	return nil
}

func TestIdHint(t *testing.T) {
	solver.RegisterHint(idHint)
	assignment := idHintCircuit{0}

	test.NewAssert(t).CheckCircuit(&idHintCircuit{}, test.WithValidAssignment(&assignment))
}
