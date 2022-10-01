// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// You can also run a script with `npx hardhat run <script>`. If you do that, Hardhat
// will compile your contracts, add the Hardhat Runtime Environment's members to the
// global scope, and execute the script.
const { ethers, upgrades } = require("hardhat");

async function main() {
  const WalletLogicV1 = await ethers.getContractFactory("WalletLogicV1");
  const walletLogicV1 = await WalletLogicV1.deploy();

  const logicContract = await walletLogicV1.deployed();

  const MockUSD = await ethers.getContractFactory("MockUSD");
  const mockUSD = await MockUSD.deploy();

  const cUSD = await mockUSD.deployed();

  const GlobalP2P = await ethers.getContractFactory("GlobalP2P");
  const gp2p = await upgrades.deployProxy(GlobalP2P, [
    logicContract.address,
    "0x874069Fa1Eb16D44d622F2e0Ca25eeA172369bC1", //cUSD on Alfajores
  ]);

  const gp2pContract = await gp2p.deployed();

  const admin = await upgrades.erc1967.getAdminAddress(gp2pContract.address);
  const impl = await upgrades.erc1967.getImplementationAddress(
    gp2pContract.address
  );

  console.table([
    { contract: "WalletLogicV1", address: logicContract.address },
    { contract: "GlobalP2P Proxy", address: gp2pContract.address },
    { contract: "Admin", address: admin },
    { contract: "Implementation", address: impl },
    { contract: "MockUSD", address: cUSD.address },
  ]);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});


// ┌─────────┬───────────────────┬──────────────────────────────────────────────┐
// │ (index) │     contract      │                   address                    │
// ├─────────┼───────────────────┼──────────────────────────────────────────────┤
// │    0    │  'WalletLogicV1'  │ '0xadf1446E115Cdf824eBdFD185De50dA84c6878e3' │
// │    1    │ 'GlobalP2P Proxy' │ '0xfD603E169066aB60eeb34343066e0badfe014C64' │
// │    2    │      'Admin'      │ '0x56921Eba54DB2d70A32B5186153680d01bb6f0a2' │
// │    3    │ 'Implementation'  │ '0xF8C90A634a1674CcF5Cc73e44Deb4bab7A741ce3' │
// │    4    │     'MockUSD'     │ '0x625CCf6B18bEF0C62c63BeE34D5862e8fB2a260B' │
// └─────────┴───────────────────┴──────────────────────────────────────────────┘
