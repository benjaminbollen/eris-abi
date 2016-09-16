package ebi

import (
	"testing"
	"github.com/eris-ltd/common/go/common"
	pm "github.com/eris-ltd/eris-pm/definitions"
)

//To Test:
//Bools, Arrays, Addresses, Hashes
//Test Packing different things
//After that, should be good to go


// quick helper padding
func pad(input []byte, size int, left bool) []byte {
	if left {
		return common.LeftPadBytes(input, size)
	}
	return common.RightPadBytes(input, size)
}

func TestUnpacker(t *testing.T) {
	for _, test := range []struct {
		abi string
		packed []byte
		name string
		expectedOutput []pm.Variable
	}{
		{
			`[{"constant":false,"inputs":[],"name":"get","outputs":[{"name":"retVal","type":"int256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"x","type":"int256"}],"name":"set","outputs":[],"payable":false,"type":"function"}]`,
			[]byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
			"get",
			[]pm.Variable{
				{
					Name: "retVal",
					Value: "-1",
				},
			},
		}, 
		{
			`[{"constant":true,"inputs":[],"name":"x","outputs":[{"name":"","type":"string"}],"payable":false,"type":"function"}]`,
			common.RightPadBytes([]byte("Hello, World"), 32),
			"x",
			[]pm.Variable{
				{
					Name: "1",
					Value: "Hello, World",
				},
			},
		},
	} {
		t.Log(test.name)
		output, err := Unpacker(test.abi, test.name, test.packed)
		if err != nil {
			t.Errorf("Unpacker failed: %v", err)
		}
		for i, expectedOutput := range test.expectedOutput {
			if output[i].Name != expectedOutput.Name {
				t.Errorf("Unpacker failed: Incorrect Name, got %v expected %v", output[i].Name, expectedOutput.Name)
			}
			if output[i].Value != expectedOutput.Value {
				t.Errorf("Unpacker failed: Incorrect value, got %v expected %v", output[i].Value, expectedOutput.Value)
			}
		}
	}
}