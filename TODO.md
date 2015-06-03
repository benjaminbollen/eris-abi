TODO

Accept JSON --json jsonabi

abi storage/indexing

accept chainid, contract pair (chainid should be able to be an ENV)

Server

Cleanup

use Cobra?

Organization:

/cmd/eris-abi/main.go - app creation
/cmd/eris-abi/cli.go - cli functions (calling on eris-abi)

core.go - Eris-specific ABI management system stuff (package ebi)


Eris ABI Folder structure:

root: .eris/abi/

abi file stored in root named as sha256 hash of abi contents

indexes: .eris/abi/index

Index files are json formatted. the name of the index is the "outer key" in the case of chainid/contractaddr pairs. Index name should be chainid the internal mapping would then map contract addr to abi hash.  

