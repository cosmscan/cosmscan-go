## Project Structure
This page describes how project layout is organized.

```shell
.
├── cmd
│    └── cosmscan : (main package)
├── internel : (internal packages)
│    ├── db : defines postgres database object
│    ├── client : defines cosmos & tendermint client
│    └── config : defines global configuration object
├── proto : protobuf files are defined in here
├── modules 
│    ├── server : represents API server based on gRPC & HTTP protocol
│    └── indexer : represents indexer engine 
├── pkg : useful util functions are defined in here
└── example : anyone can launch cosmscan with this example in local machine.
```