package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/eris-ltd/eris-abi"
)

//------------------------------------------------------------------------
// http server exports same commands as the cli
// all request arguments are keyed and passed through header
// body is ignored

func ListenAndServe(host, port string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/pack",packHandler)
	return http.ListenAndServe(host+":"+port, mux)
}

// dead simple response struct
type HTTPResponse struct {
	Response string
	Error    string
}

func WriteResult(w http.ResponseWriter, result string) {
	resp := HTTPResponse{result, ""}
	b, _ := json.Marshal(resp)
	w.Write(b)
}

func WriteError(w http.ResponseWriter, err error) {
	resp := HTTPResponse{"", err.Error()}
	b, _ := json.Marshal(resp)
	w.Write(b)
}

//------------------------------------------------------------------------
// handlers

func packHandler(w http.ResponseWriter, r *http.Request) {
	input := r.Header.Get("input") //json, hash, or index. "FILE" NOT VALID

	//tx argument parsing
	argStr := r.Header.Get("args")
	if argStr == "" {
		argStr = `[]`
	}

	var args []string
	err := json.Unmarshal([]byte(argStr), &args)
	if err != nil {
		WriteError(w, err)
		return
	}

	//input method switch
	if (input == "json") {
		jsonabi := []byte(r.Header.Get("json"))

		tx, err := ebi.Packer(jsonabi, args...)
		if err != nil {
			WriteError(w, err)
			return
		}

		WriteResult(w, fmt.Sprintf("%s",tx))

	} else if (input == "hash") {
		hash := r.Header.Get("hash")

		tx, err := ebi.HashPack(hash, args...)
		if err != nil {
			WriteError(w, err)
			return
		}

		WriteResult(w, fmt.Sprintf("%s", tx))
		return		

	} else if (input == "index") {
		index := r.Header.Get("index")
		if index == "" {
			index = DefaultIndex
		}

		key := r.Header.Get("key")
		if key == "" {
			WriteError(w, fmt.Errorf("A key for the index MUST be specified"))
		}

		tx, err := ebi.IndexPack(index, key, args...)
		if err != nil {
			WriteError(w, err)
			return
		}

		WriteResult(w, fmt.Sprintf("%s", tx))
		return

	} else {
		WriteError(w, fmt.Errorf("Unrecoginized abi specification method"))
	} 
}

/*
func genHandler(w http.ResponseWriter, r *http.Request) {
	typ, dir, auth := typeDirAuth(r)
	addr, err := coreKeygen(dir, auth, typ)
	if err != nil {
		WriteError(w, err)
		return
	}
	WriteResult(w, fmt.Sprintf("%X", addr))
}

func pubHandler(w http.ResponseWriter, r *http.Request) {
	_, dir, auth := typeDirAuth(r)
	addr := r.Header.Get("addr")
	if addr == "" {
		WriteError(w, fmt.Errorf("must provide an address with the `addr` key"))
		return
	}
	pub, err := corePub(dir, auth, addr)
	if err != nil {
		WriteError(w, err)
		return
	}
	WriteResult(w, fmt.Sprintf("%X", pub))
}

func signHandler(w http.ResponseWriter, r *http.Request) {
	_, dir, auth := typeDirAuth(r)
	addr := r.Header.Get("addr")
	if addr == "" {
		WriteError(w, fmt.Errorf("must provide an address with the `addr` key"))
		return
	}
	hash := r.Header.Get("hash")
	if hash == "" {
		WriteError(w, fmt.Errorf("must provide a message hash with the `hash` key"))
		return
	}
	sig, err := coreSign(dir, auth, hash, addr)
	if err != nil {
		WriteError(w, err)
		return
	}
	WriteResult(w, fmt.Sprintf("%X", sig))
}

func verifyHandler(w http.ResponseWriter, r *http.Request) {
	_, dir, auth := typeDirAuth(r)
	addr := r.Header.Get("addr")
	if addr == "" {
		WriteError(w, fmt.Errorf("must provide an address with the `addr` key"))
		return
	}
	hash := r.Header.Get("hash")
	if hash == "" {
		WriteError(w, fmt.Errorf("must provide a message hash with the `hash` key"))
		return
	}
	sig := r.Header.Get("sig")
	if sig == "" {
		WriteError(w, fmt.Errorf("must provide a signature with the `sig` key"))
		return
	}

	res, err := coreVerify(dir, auth, addr, hash, sig)
	if err != nil {
		WriteError(w, err)
		return
	}
	WriteResult(w, fmt.Sprintf("%v", res))
}

func hashHandler(w http.ResponseWriter, r *http.Request) {
	typ, _, _ := typeDirAuth(r)
	data := r.Header.Get("data")

	hash, err := coreHash(typ, data)
	if err != nil {
		WriteError(w, err)
		return
	}
	WriteResult(w, fmt.Sprintf("%X", hash))
}

// convenience function
func typeDirAuth(r *http.Request) (string, string, string) {
	typ := r.Header.Get("type")
	if typ == "" {
		typ = DefaultKeyType
	}
	dir := r.Header.Get("dir")
	if dir == "" {
		dir = DefaultDir
	}
	auth := r.Header.Get("dir")
	if auth == "" {
		auth = DefaultAuth
	}
	return typ, dir, auth
}
*/