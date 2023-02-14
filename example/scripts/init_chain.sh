#!/bin/sh

MONIKER=testchainer
CHAIN_ID=testchain
KEYRING=test
HOME_PATH=/data/chain/.gaiad

rm -rf $HOME_PATH
mkdir -p $HOME_PATH

# init all three validators
gaiad init $MONIKER --chain-id=$CHAIN_ID --home=$HOME_PATH

# create keys for all three validators
gaiad keys add validator --keyring-backend=$KEYRING --home=$HOME_PATH

# create validator node with tokens to transfer to the three other nodes
gaiad add-genesis-account $(gaiad keys show validator -a --keyring-backend=$KEYRING --home=$HOME_PATH) 100000000000uatom,100000000000stake --home=$HOME_PATH
gaiad gentx validator 500000000stake --keyring-backend=$KEYRING --home=$HOME_PATH --chain-id=$CHAIN_ID
gaiad collect-gentxs --home=$HOME_PATH

# validator
# enable rest api server & unsafe cors
sed -i -E 's|enable = false|enable = true|g' $HOME_PATH/config/app.toml
sed -i -E 's|enabled-unsafe-cors = false|enabled-unsafe-cors = true|g' $HOME_PATH/config/app.toml

# allow duplicate ip
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME_PATH/config/config.toml
sed -i -E 's|tcp://127.0.0.1:26657|tcp://0.0.0.0:26657|g' $HOME_PATH/config/config.toml

gaiad start --home=$HOME_PATH
