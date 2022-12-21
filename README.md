## Cosmscan
## What is Cosmscan?
Cosmscan is a indexer engine for cosmos based blockchain.

Builders often want to serve a indexed query such as aggregation, search, and so on.
The native query on Cosmos RPC is not enough for this purpose, basically it stores all the data on the LSM tree [(LevelDB)](https://github.com/google/leveldb) which is efficient to perform high write throughput,
On the other side, it has an inefficiency to answer the following questions.
- How many transactions are there in the last 24 hours?
- What tokens the holder has? and How many tokens the holder has?
- Number of active accounts in the last 24 hours?
- Top 10 holders of the coin?
- and so on

Cosmscan is here to solve this problem. 🚀🚀

## Features
- Store the all data from cosmos based blockchain into PostgreSQL
- Support default useful queries with `gRPC` / `HTTP 2.0`
- Easy installation and configuration

## Contribution
If you are interested in contributing to this project,
Please feel free to open issue or pull request.
