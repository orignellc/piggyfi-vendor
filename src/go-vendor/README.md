## Developer Guide to Setup Project

- Author: Oluwafemi Alofe

Note: This setup is platform agnostic.
My development environment is on Mac OS, so this documentation assumes so for everyone.

Good Luck!

Run:

```azure
go get github.com/celo-org/celo-blockchain@v1.5.8
```

Running without tagging version give issues.

Installed the latest version of solidity compiler `solc`

```azure
brew update
brew tap ethereum/ethereum
brew install solidity
```

If you have an existing version less than `0.8.9`. Run `brew install solidity@9`.

In my own case when that proved to be an issue,
I ran `brew install solidity` which showed I have `0.8.15` installed but not linked, running `brew link --overwrite solidity` solved the issue.

More details here https://docs.soliditylang.org/en/latest/installing-solidity.html.

Run:

```
//Replace star with version folder as seen on your PC
cd $GOPATH/pkg/mod/github.com/celo-org//celo-blockchain@v*.*.*
sudo go build -o celo-abigen cmd/abigen/main.go
cp celo-abigen /usr/local/bin
```

Run:

```azure
-- In the root folder
 solc --abi src/smart-contracts/contracts/*.sol -o src/go-vendor/build --bin --include-path node_modules/ --base-path .
```

Now let's generate go binding for our smart contracts, I have taken time to create a convenience script that is re-usable for every project

Run:

```azure
-- In the go-vendor/utils folder
go run compile_contracts.go
```

Confirm inside `go-vendor/contracts/celo` for the output bindings.

**NB:** By default only two binding are generated, you can add more contracts on `line 17` of the script to generate its binding, e.g. `WalletProxy`

## Start Using

Open a first terminal and start ganache cli.

Run: 

```azure
ganache-cli
```

Open the second terminal in the project root dir (i.e. `piggyfi-vendor`) and deploy the contracts locally.

Run :

```
yarn delpoy:locally
```