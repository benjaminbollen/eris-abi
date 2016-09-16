package ebi

import (
	"testing"
	//"github.com/eris-ltd/eris-abi/abi"
	"github.com/eris-ltd/common/go/common"
)

var negABIStr string = `[{"constant":false,"inputs":[],"name":"get","outputs":[{"name":"retVal","type":"int256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"x","type":"int256"}],"name":"set","outputs":[],"payable":false,"type":"function"}]`
var name string = "get"


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
	output := []byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}
	out, err := Unpacker(negABIStr, name, output)
	if err != nil {
		t.Errorf("Unpacker failed: %v", err)
	}
	if out[0].Name != "retVal" {
		t.Errorf("Unpacker failed: Incorrect Name, got %v expected %v", out[0].Name, "retVal")
	}
	if out[0].Value != "-1" {
		t.Errorf("Unpacker failed: Incorrect value, got %v expected %v", out[0].Value, "-2")
	} 
}