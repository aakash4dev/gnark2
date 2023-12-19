package regressiontests

import (
	"testing"

	"github.com/aakash4dev/gnark-fork/frontend"
	"github.com/aakash4dev/gnark-fork/std/rangecheck"
	"github.com/aakash4dev/gnark-fork/test"
)

type TestRangeCheckCircuit struct {
	I1 frontend.Variable
	N  int
}

func (circuit *TestRangeCheckCircuit) Define(api frontend.API) error {
	rangeChecker := rangecheck.New(api)
	rangeChecker.Check(circuit.I1, circuit.N)
	return nil
}

func TestIssue897(t *testing.T) {
	assert := test.NewAssert(t)
	circuit := TestRangeCheckCircuit{
		N: 7,
	}
	witness := TestRangeCheckCircuit{
		I1: 1 << 7,
		N:  7,
	}
	assert.CheckCircuit(&circuit, test.WithInvalidAssignment(&witness))
}
