# Eris-ABI
Eris ABI tool

A simple tool for constructing transaction call data for ABI-enabled contracts.

Features:
- Command-line interface
- Http interface
- Automated abi access in eris directory
- Import file
- General index system built in

Planned:
See TODO.md

#Directory Structure:
```
//Directory Structure
//.eris (or ERIS_ROOT)
//   +-abi (or ERIS_ABI_ROOT)
//      +-index
//      |    + abi indexing jsons
//      |
//      +-raw
//           + abi files (hash named)
```

#ENV Vars
There are two environment variables which can be used for customization:
`ERIS_ABI_ROOT`: Specifies a directory for the "abi" root directory. If not set will default to `ERIS_ROOT/abi`
`ERIS_HEAD`: Used as default name for index (think checked out chain id)

##WARNING: This is completely untested software. It should not yet be presumed operational.

#CLI

```
>> eris-abi pack --input file example1.abi set hello

>> eris-abi import example1.abi

>> eris-abi new indexname

>> eris-abi add indexname key value

>> eris-abi server
```

##Commands    
```
NAME:
   ebi - Tool for using ABI's to construct transaction data

USAGE:
   ebi [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR(S): 
   Dennis Mckinnon <contact@erisindustries.com> 

COMMANDS:
   pack		generate a transaction
   import	Import an existing ABI file into abi directory
   add		Add an entry to index
   new		Create new index
   server	Starts a packing server
   help, h	Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```


#Pack:
```
NAME:
   pack - generate a transaction

USAGE:
   eris-abi pack [command options] [arguments...]

OPTIONS:
   --input "index"	Specify input method of ABI data.
   --index, -i 		Specify Chainid to use as look-up
 ```
Note: Options to --input flag are ["file","json","hash","index"]. Meaning of supplied arguments differ depending on which you choose.

file:
```
eris-abi pack --input file <Filename> <TxArgs...>
```

json:
```
eris-abi pack --input json <JSON-STRING> <TxArgs...>
```

hash:
```
eris-abi pack --input hash <ABI-Hash-STRING (no 0x)> <TxArgs...>
```

index:
```
eris-abi pack --input index [--index IndexName (Default: EnvVar - "ERIS_HEAD")] <key> <TxArgs...>
```
----------------------------------------------------------------     
#import:

```
NAME:
   import - Import an existing ABI file into abi directory

USAGE:
   eris-abi import [--input <inputtype, file/json>]  <abi filepath/abi-json string>
```
TODO: Allow import to accept --index and --key flags to supply index and key names for automatic entry into to an index (instead of having to use `eris-abi add` afterwards.

----------------------------------------------------------------    
#new:

```
NAME:
   new - Create new index

USAGE:
   eris-abi new <IndexName>
```

No Tricks here. Creates a blank index which you can add key-value pairs too. When using the index input method the abi hash is fetched from the index's value for supplied key. New entries are added to index using the "add" command

----------------------------------------------------------------     
#add:
```
NAME:
   add - Add an entry to index

USAGE:
   eris-abi add [--index IndexName (Default: EnvVar - "ERIS_HEAD")] <key> <abi-Hash>
```

----------------------------------------------------------------
#server:
```
NAME:
   server - Starts a packing server

USAGE:
   eris-abi server [command options]

OPTIONS:
   --host "localhost"	set the host for key daemon to listen on
   --port "4592"	set the port for key daemon to listen on

```

The server has not yet been tested. Gets arguments similar to the cli pack command through the header. Use of "file" input method is not permitted 

# Contributions

Are Welcome! Before submitting a pull request please:

* read up on [How The Marmots Git](https://github.com/eris-ltd/coding/wiki/How-The-Marmots-Git)
* fork from `develop`
* go fmt your changes
* have tests
* pull request
* be awesome

That's pretty much it. 

See our [CONTRIBUTING.md](.github/CONTRIBUTING.md) and [PULL_REQUEST_TEMPLATE.md](.github/PULL_REQUEST_TEMPLATE.md) for more details.

Please note that this repository is GPLv3.0 per the LICENSE file. Any code which is contributed via pull request shall be deemed to have consented to GPLv3.0 via submission of the code (were such code accepted into the repository).

# Bug Reporting

Found a bug in our stack? Make an issue!

The [issue template](.github/ISSUE_TEMPLATE.md] specifies what needs to be included in your issue and will autopopulate the issue.

# License

[Proudly GPL-3](http://www.gnu.org/philosophy/enforcing-gpl.en.html). See [license file](https://github.com/eris-ltd/eris-cli/blob/master/LICENSE.md).
