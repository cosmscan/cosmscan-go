#!/bin/sh
MONIKER=novatest
CHAIN_ID=testchain
KEYRING=test
HOME_PATH=/data/chain/.gaiad

INDEX=1
SEQUENCE=1

while :;
do
        echo "generating keys"
        gaiad keys add validator$INDEX --keyring-backend test --home $HOME_PATH > /dev/null
        ADDR=$(gaiad keys show validator$INDEX -a --keyring-backend test --home $HOME_PATH)

        echo $ADDR

        echo ""
        echo "sent 100uatom from validator to the $ADDR"
        gaiad tx bank send validator $ADDR 100uatom \
                --chain-id $CHAIN_ID \
                --keyring-backend test \
                --home $HOME_PATH \
                --sequence $SEQUENCE \
                --node http://testchain:26657 \
                -y

        INDEX=$(( INDEX + 1))
        SEQUENCE=$(( SEQUENCE + 1))
        sleep 1
done