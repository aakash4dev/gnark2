// Copyright 2020 ConsenSys Software Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by gnark DO NOT EDIT

package mpcsetup

import (
	"testing"

	curve "github.com/consensys/gnark-crypto/ecc/bls12-381"
	cs "github.com/aakash4dev/gnark-fork/constraint/bls12-381"
	"github.com/aakash4dev/gnark-fork/frontend"
	"github.com/aakash4dev/gnark-fork/frontend/cs/r1cs"
	gnarkio "github.com/aakash4dev/gnark-fork/io"
	"github.com/stretchr/testify/require"
)

func TestContributionSerialization(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	assert := require.New(t)

	// Phase 1
	srs1 := InitPhase1(9)
	srs1.Contribute()

	assert.NoError(gnarkio.RoundTripCheck(&srs1, func() interface{} { return new(Phase1) }))

	var myCircuit Circuit
	ccs, err := frontend.Compile(curve.ID.ScalarField(), r1cs.NewBuilder, &myCircuit)
	assert.NoError(err)

	r1cs := ccs.(*cs.R1CS)

	// Phase 2
	srs2, _ := InitPhase2(r1cs, &srs1)
	srs2.Contribute()

	assert.NoError(gnarkio.RoundTripCheck(&srs2, func() interface{} { return new(Phase2) }))
}
