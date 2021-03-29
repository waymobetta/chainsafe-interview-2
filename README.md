chainsafe interview task #2

---
### Instructions

> Pure backend Task
> 
> Implement a Go command line tool that takes a local file as a CLI argument, uploads it into IPFS, and puts the resulting cid into a smart contract that will store it inside an array.
> 
> This smart contract needs to be created and deployed to the Ethereum network (whichever you choose based on option in overview). The set of smart contract methods, details of implementation and way of organizing cid array inside it completely arbitrary and designed by candidate


### Steps to get going

1. Run IPFS daemon
2. Run Ganache
3. Open Remix (desktop/web); set environment to web3 provider for local Ganache (likely localhost:7545)
4. Deploy contract using Remix (Desktop/Web)
5. Update .env with the following: (storing in git is obviously a bad practice and is ordinarily added to .gitignore but is dismissable for the demo)
    - IPFS API Server URL (likely localhost:5001)
    - Ganache URL to sim ethereum client (likely localhost:7545)
    - Private key from the same address used to deploy the contract (as the modifier in the contract will prevent non-owners from storing CIDs)
    - Deployed contract address
6. ```make build```
7. run CLI


### Examples

Store file in contract
```zsh
./chainsafe-interview-2 store ilth.txt
```

Get all stored hashes
```zsh
./chainsafe-interview-2 list
```

### Screenshots

Contract Deployed
![DeployContract](https://user-images.githubusercontent.com/17755587/112772751-6d5d9900-8fe7-11eb-80fc-38e6233ddcfe.png)

Store CID Contract Call
![StoreHash](https://user-images.githubusercontent.com/17755587/112772800-ac8bea00-8fe7-11eb-972a-ea3bf9607e5a.png)

Example CLI Usage
![CLI Usage](https://user-images.githubusercontent.com/17755587/112772791-9e3dce00-8fe7-11eb-8959-98930d68b16e.png)

Tests
![Tests](https://user-images.githubusercontent.com/17755587/112772539-3fc42000-8fe6-11eb-9148-8462f8ca9cdd.png)


### Tests

* Test CLI
```zsh
make test
```

* Test Store.sol steps:
1. Run Ganache
2. Open Remix (desktop/web); set environment to web3 provider for local Ganache (likely localhost:7545)
3. Deploy TestStore.sol
4. Run `beforeEach` method to instantiate a new instance of Store.sol
5. Run `testAddCID` method to test adding a sample IPFS hash
    - this will add a sample hash, retrieve the hash, then compare the added/retrieved hashes and assert they are equal


### Notes

*Thought Process*
- pro/cons of ipfs cli v. go-ipfs-api
- pro/cons of urfave/cli v. spf13/cobra
- pro/cons of storing strings in contract v. bytes; strings for speed of development
- pro/cons of actually storing a private key in a file; dismissable for speed of development/test network
- write smart contract (Remix), likely use Ganache for speed of development (no rinkeby faucet)

*Known Issues (TODO)*
- viper should be used for env support over godotenv (reduce redundancy)
- strings in solidity should be avoided; utilize storing bytes instead
- splitting the IPFS multihash into two, stripping and utilizing a base58 encoding/decoding package seems like the correct path (more research needed)
- likely better/more tests can be written
- add/emit events to contract

_Useful Commands:_
- init ipfs daemon
```zsh
    ipfs daemon
```

- generate contract bindings
```zsh
    abigen --sol contracts/Store.sol --pkg eth --out eth/Store.go
```

### Post-mortem

__Summary__: Project took longer than anticipated due to the environment setup
1. Remix not loading in Brave/chrome; forced to download desktop version
2. Xcode not updated; comp ran out of disk space necessary; switched to new laptop
3. Researching best practices (Solidity, IPFS, CLI libraries) took more time; even still there is room for improvement


__allocated time to complete:__ ~4hrs

__actual time to complete:__ ~10hrs

__biggest like:__ real-world exercise (I really appreciated the practicality of this)

__biggest miss:__ not being able to run my design thoughts by a colleague before turning it in

---

Thank you for the opportunity!


[LICENSE](https://github.com/waymobetta/chainsafe-interview-2/blob/master/LICENSE)
