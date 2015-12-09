package ebi

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

var (
	EVcopy  = os.Getenv("ERIS_ABI_ROOT")
	DIR     = path.Join(os.Getenv("HOME"), ".testing_eris-abi")
	testabi = "[{\"constant\":false,\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"input\",\"type\":\"bytes32\"}],\"name\":\"set\",\"outputs\":[],\"type\":\"function\"},{\"inputs\":[],\"type\":\"constructor\"}]"
	testtx1 = []string{
		"set",
		"0xDEADBEEF",
	}
	tx1     = "db80813f00000000000000000000000000000000000000000000000000000000deadbeef"
	tx2     = "db80813f68656c6c6f000000000000000000000000000000000000000000000000000000"
	testtx2 = []string{
		"set",
		"hello",
	}
	file  = path.Join(DIR, "example.abi")
	hash  = "3c49d66fe61179b31260ff006e401d489b0a73db175980ee269147b12a0041f8"
	index = "test"
	key   = "bob"
)

func setup() error {
	err := os.Setenv("ERIS_ABI_ROOT", DIR)
	if err != nil {
		return err
	}

	InitPaths()

	err = BuildDirTree()
	if err != nil {
		return err
	}

	//make test files
	err = ioutil.WriteFile(file, []byte(testabi), 0644)
	if err != nil {
		return err
	}
	hash, err = ImportAbi(file)
	if err != nil {
		return err
	}

	err = NewIndex(index)
	if err != nil {
		return err
	}

	err = AddEntry(index, key, hash)
	if err != nil {
		return err
	}

	return nil
}

func cleanup() {
	if EVcopy == "" {
		//Wasn't set in the first place, unset
		err := os.Unsetenv("ERIS_ABI_ROOT")
		if err != nil {
			fmt.Println("There was an error during cleanup")
		}
	} else {
		err := os.Setenv("ERIS_ABI_ROOT", EVcopy)
		if err != nil {
			fmt.Println("There was an error during cleanup")
		}
	}

	//Delete testing directory
	err := os.RemoveAll(DIR)
	if err != nil {
		fmt.Println("There was an error during cleanup")
	}

	return
}

func TestMain(m *testing.M) {
	err := setup()
	result := 0
	if err == nil {
		fmt.Println("Running Tests")
		result = m.Run()
	}
	cleanup()
	os.Exit(result)
}

func TestPacker(t *testing.T) {
	abidata := []byte(testabi)
	txout, err := Packer(abidata, testtx1...)
	if err != nil {
		t.Fatal(err)
	}

	if txout != tx1 {
		t.Fatalf("Output transaction did not match expected")
	}
}

func TestFilePack(t *testing.T) {
	txout, err := FilePack(file, testtx1...)
	if err != nil {
		t.Fatal(err)
	}

	if txout != tx1 {
		t.Fatalf("Output transaction did not match expected")
	}
}

func TestHashPack(t *testing.T) {
	txout, err := HashPack(hash, testtx1...)
	if err != nil {
		t.Fatal(err)
	}

	if txout != tx1 {
		t.Fatalf("Output transaction did not match expected")
	}
}

func TestIndexPack(t *testing.T) {
	txout, err := IndexPack(index, key, testtx1...)
	if err != nil {
		t.Fatal(err)
	}

	if txout != tx1 {
		t.Fatalf("Output transaction did not match expected")
	}
}

func TestNewIndex(t *testing.T) {
	err := NewIndex("asdf")
	if err != nil {
		t.Fatal(err)
	}
}
