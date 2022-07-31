## Developer Guide to Setup Project

- Author: Oluwafemi Alofe

Note: This setup is platform agnostic. 
My development environment is on Mac OS, so this documentation assumes so for everyone.

Good Luck!

Run
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

Run 
```
//Replace star with version folder as seen on your PC
cd $GOPATH/pkg/mod/github.com/celo-org//celo-blockchain@v*.*.*  
sudo go build -o celo-abigen cmd/abigen/main.go
cp celo-abigen /usr/local/bin 
```

Run
```azure
-- In the root folder
 solc --abi src/smart-contracts/contracts/*.sol -o src/go-vendor/build --bin --include-path node_modules/ --base-path .
```

Now let's generate go binding for our smart contracts

```azure
-- In go-vendor folder
celo-abigen --bin=./build/GlobalP2P.bin --abi=./build/GlobalP2P.abi --pkg=GlobalP2P --out=contracts/celo/GlobalP2P.go
celo-abigen --bin=./build/WalletLogicV1.bin --abi=./build/WalletLogicV1.abi --pkg=WalletLogicV1 --out=contracts/celo/WalletLogicV1.go
celo-abigen --bin=./build/WalletProxy.bin --abi=./build/WalletProxy.abi --pkg=WalletProxy --out=contracts/celo/WalletProxy.go
celo-abigen --bin=./build/CommonWallet.bin --abi=./build/CommonWallet.abi --pkg=CommonWallet --out=contracts/celo/CommonWallet.go
```
