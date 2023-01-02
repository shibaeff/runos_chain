# runoschain
**runoschain** is a blockchain built using Cosmos SDK and Tendermint. Blockchain stores configurations of the 
networking devices backed by the RUNOS SDN controller. The interaction between blockchain and RUNOS happens through the 
special adapter.

## Get started - docker

```bash
#build tmkms binary
cargo install tmkms --features=yubihsm
cp /home/golanger/.cargo/bin/tmkms tmkms

#build tmkms 
docker build -f Dockerfile-ubuntu-tmkms . -t tmkms_i:v0.12.2

# create volumes
mkdir -p docker/kms-alice
mkdir -p docker/node-carol
mkdir -p docker/sentry-alice
mkdir -p docker/sentry-bob
mkdir -p docker/val-alice
mkdir -p docker/val-bob

# init nodes
echo -e node-carol'\n'sentry-alice'\n'sentry-bob'\n'val-alice'\n'val-bob \
    | xargs -I {} \
    docker run --rm -i \
    -v $(pwd)/docker/{}:/root/.runos_chain \
    runos_chaind_i \
    init runos_chain
    
# making sure that chain_id is right
echo -e node-carol'\n'sentry-alice'\n'sentry-bob'\n'val-alice'\n'val-bob \
    | xargs -I {} \
    docker run --rm -i \
    -v $(pwd)/docker/{}:/root/.runos_chain \
    --entrypoint sed \
    runos_chaind_i \
    -Ei 's/^chain-id = .*$/chain-id = "runos_chain-1"/g' \
    /root/.runos_chain/config/client.toml
# create key for val-alice
docker run --rm -it \
    -v $(pwd)/docker/val-alice:/root/.runos_chain \
    runos_chaind_i \
    keys \
    --keyring-backend file --keyring-dir /root/.runos_chain/keys \
    add alice
# imit tmkms
docker run --rm -it \
    -v $(pwd)/docker/kms-alice:/root/tmkms \
    tmkms_i:v0.12.2 \
    init /root/tmkms
# Make sure that you use the right protocol version. In your case:
docker run --rm -i \
  -v $(pwd)/docker/kms-alice:/root/tmkms \
  --entrypoint sed \
  tmkms_i:v0.12.2 \
  -Ei 's/^protocol_version = .*$/protocol_version = "v0.34"/g' \
  /root/tmkms/tmkms.toml

# Pick an expressive name for the file that will contain the softsign key for val-alice
docker run --rm -i \
  -v $(pwd)/docker/kms-alice:/root/tmkms \
  --entrypoint sed \
  tmkms_i:v0.12.2 \
  -Ei 's/path = "\/root\/tmkms\/secrets\/cosmoshub-3-consensus.key"/path = "\/root\/tmkms\/secrets\/val-alice-consensus.key"/g' \
  /root/tmkms/tmkms.toml

# Replace cosmoshub-3 with runos_chain-1, the name of your blockchain, wherever the former appears
docker run --rm -i \
    -v $(pwd)/docker/kms-alice:/root/tmkms \
    --entrypoint sed \
    tmkms_i:v0.12.2 \
    -Ei 's/cosmoshub-3/runos_chain-1/g' /root/tmkms/tmkms.toml
# import alice consensus key
docker run --rm -t \
    -v $(pwd)/docker/val-alice:/root/.runos_chain \
    runos_chaind_i \
    tendermint show-validator \
    | tr -d '\n' | tr -d '\r' \
    > docker/val-alice/config/pub_validator_key.json

# move key away from validator
mv docker/val-alice/config/priv_validator_key.json   docker/kms-alice/secrets/priv_validator_key-val-alice.json

# sample password for val-alice
echo -n password > docker/val-alice/keys/passphrase.txt
# sample password for val-bob
mkdir -p docker/val-bob/keys
echo -n password > docker/val-bob/keys/passphrase.txt


docker run --rm -it \
    -v $(pwd)/docker/val-bob:/root/.runos_chain \
    runos_chaind_i \
    keys \
    --keyring-backend file --keyring-dir /root/.runos_chain/keys \
    add bob

# import softsign device  
docker run --rm -i \
    -v $(pwd)/docker/kms-alice:/root/tmkms \
    -w /root/tmkms \
    tmkms_i:v0.12.2 \
    softsign import secrets/priv_validator_key-val-alice.json \
    secrets/val-alice-consensus.key
# On start, val-alice may still recreate a missing private key file due to how defaults are handled in the code. 
# To prevent that, you can instead copy it from sentry-alice where it has no value.
cp docker/sentry-alice/config/priv_validator_key.json \
    docker/val-alice/config/

# Choose a port unused on val-alice, for instance 26659, and inform kms-alice
docker run --rm -i \
    -v $(pwd)/docker/kms-alice:/root/tmkms \
    --entrypoint sed \
    tmkms_i:v0.12.2 \
    -Ei 's/^addr = "tcp:.*$/addr = "tcp:\/\/val-alice:26659"/g' /root/tmkms/tmkms.toml

# Do not forget, you must inform Alice's validator that it should indeed listen on port 26659. 
# In val-alice/config/config.toml
docker run --rm -i \
  -v $(pwd)/docker/val-alice:/root/.runos_chain \
  --entrypoint sed \
  runos_chaind_i \
  -Ei 's/priv_validator_laddr = ""/priv_validator_laddr = "tcp:\/\/0.0.0.0:26659"/g' \
  /root/.runos_chain/config/config.toml

# Make sure it will not look for the consensus key on file:
docker run --rm -i \
  -v $(pwd)/docker/val-alice:/root/.runos_chain \
  --entrypoint sed \
  runos_chaind_i \
  -Ei 's/^priv_validator_key_file/# priv_validator_key_file/g' \
  /root/.runos_chain/config/config.toml

# Make sure it will not look for the consensus state file either, 
# as this is taken care of by the KMS
docker run --rm -i \
  -v $(pwd)/docker/val-alice:/root/.runos_chain \
  --entrypoint sed \
  runos_chaind_i \
  -Ei 's/^priv_validator_state_file/# priv_validator_state_file/g' \
  /root/.runos_chain/config/config.toml

# copy empty file from sentry for the code to work
cp docker/sentry-alice/config/priv_validator_key.json \
    docker/val-alice/config
```
### Genesis
```bash
# set up chain id
docker run --rm -i \
    -v $(pwd)/docker/val-alice:/root/.runos_chain \
    --entrypoint sed \
    runos_chaind_i \
    -Ei 's/"chain_id": "runos_chain"/"chain_id": "runos_chain-1"/g' \
    /root/.runos_chain/config/genesis.json

# initial balances
# initial balance for Alice
 ALICE=$(echo password | docker run --rm -i \
    -v $(pwd)/docker/val-alice:/root/.runos_chain \
    runos_chaind_i \
    keys \
    --keyring-backend file --keyring-dir /root/.runos_chain/keys \
    show alice --address)
# add alice balance to genesis
docker run --rm -it \
    -v $(pwd)/docker/val-alice:/root/.runos_chain \
    runos_chaind_i \
    add-genesis-account $ALICE 1000000000stake
# move genesis to bob
mv docker/val-alice/config/genesis.json \
    docker/val-bob/config/
# do the same for Bob
BOB=$(echo password | docker run --rm -i \
    -v $(pwd)/docker/val-bob:/root/.runos_chain \
    runos_chaind_i \
    keys \
    --keyring-backend file --keyring-dir /root/.runos_chain/keys \
    show bob --address)
docker run --rm -it \
    -v $(pwd)/docker/val-bob:/root/.runos_chain \
    runos_chaind_i \
    add-genesis-account $BOB 500000000stake
```

### Initial stakes
```bash
# bob
echo password | docker run --rm -i \
    -v $(pwd)/docker/val-bob:/root/.runos_chain \
    runos_chaind_i \
    gentx bob 40000000stake \
    --keyring-backend file --keyring-dir /root/.runos_chain/keys \
    --account-number 0 --sequence 0 \
    --chain-id runos_chain-1 \
    --gas 1000000 \
    --gas-prices 0.1stake
# alice
# return genesis to alice
mv docker/val-bob/config/genesis.json \
    docker/val-alice/config/
# alice stake 
echo password | docker run --rm -i \
    -v $(pwd)/docker/val-alice:/root/.runos_chain \
    runos_chaind_i \
    gentx alice 60000000stake\
    --keyring-backend file --keyring-dir /root/.runos_chain/keys \
    --account-number 0 --sequence 0 \
    --pubkey $(cat docker/val-alice/config/pub_validator_key.json) \
    --chain-id runos_chain-1 \
    --gas 1000000 \
    --gas-prices 0.1stake
```

### Genesis assembly and distribution
```bash
cp docker/val-bob/config/gentx/gentx-* \
docker/val-alice/config/gentx

docker run --rm -it \
-v $(pwd)/docker/val-alice:/root/.runos_chain \
runos_chaind_i collect-gentxs

docker run --rm -it \
-v $(pwd)/docker/val-alice:/root/.runos_chain \
runos_chaind_i \
validate-genesis

echo -e node-carol'\n'sentry-alice'\n'sentry-bob'\n'val-bob \
    | xargs -I {} \
    cp docker/val-alice/config/genesis.json docker/{}/config
```

### Network preparation
```bash
# alice node id
docker run --rm -i \
    -v $(pwd)/docker/val-alice:/root/.runos_chain \
    runos_chaind_i \
    tendermint show-node-id
# alice-val node id 805d682cc7317d25505c03140a45518a42392380

docker run --rm -i -v $(pwd)/docker/sentry-bob:/root/.runos_chain \
    runos_chaind_i \
    tendermint show-node-id
# sentry-bob is 72bd0e024ba97259ea73d79ba47bd03d0c7a5da7
docker run --rm -i -v $(pwd)/docker/node-carol:/root/.runos_chain \
    runos_chaind_i \
    tendermint show-node-id
# carol-node is d19c0e5a16a3f18b0f3c4df62c390fc2428d8919
docker run --rm -i \
    -v $(pwd)/docker/sentry-alice:/root/.runos_chain \
    runos_chaind_i \
    tendermint show-node-id
72bd0e024ba97259ea73d79ba47bd03d0c7a5da7@sentry-bob:26656
d19c0e5a16a3f18b0f3c4df62c390fc2428d8919@node-carol:26656
805d682cc7317d25505c03140a45518a42392380@val-alice:26656
79306d5be3a9aeb8db8238178b99c2c7d2731484@sentry-alice:26656
72bd0e024ba97259ea73d79ba47bd03d0c7a5da7@sentry-bob:26656
8a58cfa0daf04797d8012d0a76fbc1bc19e9a099@val-bob:26656

```


## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts the blockchain in development.

### Configure

The blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release policy
To release a new version of the blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make the final changes from the release page and publish it.

### Install
To install the latest version of the blockchain node's binary, execute the following command on the machine:

```
curl https://get.ignite.com/username/runos_chain@latest! | sudo bash
```
`username/runos_chain` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

