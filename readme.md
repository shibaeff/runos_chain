# runoschain
**runoschain** is a blockchain built using Cosmos SDK and Tendermint. Blockchain stores configurations of the 
networking devices backed by the RUNOS SDN controller. The interaction between blockchain and RUNOS happens through the 
special adapter.

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

