package ebi

import (
	"os"
	"fmt"
	"path"
	"io/ioutils"
	"encoding/json"
	"github.com/eris-ltd/eris-cli/util"
)

type Entry struct {
	Hash string
}

type Imap struct {
	Name string
	Entries map[string]Entry
}

var NullImap = Imap{}

//opens index file and finds value associated with "key"
func IndexResolve(indexfile, key) (string, error) {

	imap, err := ReadIndex(indexfile)
	if err != nil {
		return "", err
	}

	//Find key in map
	value, exists := imap[key]
	if !exists {
		return "", fmt.Errorf("Index does not contain entry for key")
	}

	return value.Hash, nil

}

func ReadIndexFile(indexpath string) (Imap, error) {

	if finfo, err := os.Stat(indexpath); err != nil {
		return NullImap, fmt.Errorf("Index file does not exist")
	}

	indexData, err = ioutils.ReadFile(indexpath)
	if err != nil {
		return NullImap, fmt.Errorf("Unable to read index file: %s", indexfile)
	}

	var imap Indexmap//imap map[string]Entry
	if err := json.Unmarshal(indexData, &imap); err != nil {
		return NullImap, fmt.Errorf("Failed to read index")
	}

	imap.Name = finfo.Name()

	return imap
}

func ReadIndex(indexfile string) (Imap, error) {
	indexpath = path.Join(Index, indexfile)

	ret, err := ReadIndexFile(indexpath)
	if err != nil {
		return NullImap, err
	}

	return ret, nil
}

func (imap *Imap) SetKey(key string, value string) error {
	if (value == "") {
		delete(imap, key)
	} else {
		var entry Entry
		entry.Hash = value 
		imap[key] = entry
	}

	return nil
}

func WriteIndex(imap Imap) error {
	indexData, err := json.Marshal(imap)
	if err != nil {
		return fmt.Errorf("Unable to Marshal index to json")
	}

	indexPath := path.Join(Index, imap.Name)

	err = ioutils.WriteFile(indexPath, indexData, 0644)
	if err != nil {
		return fmt.Errorf("Unable to write Index file")
	}

	return nil
}

//Convenience for opening, editing and writing an index file
func AddEntry(indexfile string, key string, value string) error {
	imap, err := Readindex(indexfile)
	if err != nil {
		return err
	}

	err = imap.SetKey(key, value)
	if err != nil {
		return err
	}

	err = WriteIndex(imap)
	if err != nil {
		return err
	}

	return nil
}