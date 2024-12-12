## SinGeth: A blockchain to Social and Envinronmental Impact Applications

Official golang implementation of the Sintrop protocol.

We are building a decentralized p2p blockchain infrastructure to power projects with social and environmental impact purpose.

## CoreGeth: An Ethereum Protocol Provider

Forked from CoreGeth protocol.

## Run with docker

### Build sin_geth

```
docker build -t sin_geth .
```

### Run sin_geth

```
docker run -p=30303:30303 -p=8545:8545 -it -v /home/user/sequoia_volume:/go-sintrop/sequoia_node  sin_geth

## Change /home/user/sequoia_volume to your dir
```

### Start a sequoia node
```
geth --identity Sequoia --datadir ./sequoia_node \
  --sequoia \
  --syncmode "full" \
  --networkid 1500 \
  --cache=1024 \
  --port 30303 \
  -authrpc.addr localhost --authrpc.port 8551 \
  --http.vhosts=* --http.addr "0.0.0.0" --http.port 8545 --http=true \
  --miner.threads=1 \
  --miner.etherbase=0x0000000000000000000000000000000000000000 \
  console
```

### Operate network

```
balance = web3.fromWei(eth.getBalance("0x0000000000000000000000000000000000000000), "ether");
eth.blockNumber
web3.eth.getBlock(eth.blockNumber)
```

### Start and stop mining
```
  miner.start()
  miner.stop()
```

## Documentation

- CoreGeth documentation is available [here](https://etclabscore.github.io/core-geth).
  + Getting Started [Installation](https://etclabscore.github.io/core-geth/getting-started/installation) and [CLI](https://etclabscore.github.io/core-geth/getting-started/run-cli)
  + [JSONRPC API](https://etclabscore.github.io/core-geth/JSON-RPC-API)
  + [Developers](https://etclabscore.github.io/core-geth/developers/build-from-source)
  + [Tutorials](https://etclabscore.github.io/core-geth/tutorials/private-network)
- Further [ethereum/go-ethereum](https://github.com/ethereum/go-ethereum) documentation about can be found [here](https://geth.ethereum.org/docs/).
- Documentation about documentation lives [here](./docs/developers/documentation.md).

## Contribution

Thank you for considering to help out with the source code! We welcome contributions
from anyone on the internet, and are grateful for even the smallest of fixes!

If you'd like to contribute to sin-geth, please fork, fix, commit and send a pull request
for the maintainers to review and merge into the main code base. 

Please make sure your contributions adhere to our coding guidelines:

 * Code must adhere to the official Go [formatting](https://golang.org/doc/effective_go.html#formatting)
   guidelines (i.e. uses [gofmt](https://golang.org/cmd/gofmt/)).
 * Code must be documented adhering to the official Go [commentary](https://golang.org/doc/effective_go.html#commentary)
   guidelines.
 * Pull requests need to be based on and opened against the `master` branch.
 * Commit messages should be prefixed with the package(s) they modify.
   * E.g. "eth, rpc: make trace configs optional"

Please see the [Developers' Guide](https://github.com/ethereum/go-ethereum/wiki/Developers'-Guide)
for more details on configuring your environment, managing project dependencies, and
testing procedures.

## License

The sin-geth library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html),
also included in our repository in the `COPYING.LESSER` file.

The sin-geth binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also
included in our repository in the `COPYING` file.
