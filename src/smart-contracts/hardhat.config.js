require("@nomicfoundation/hardhat-toolbox");
require('@openzeppelin/hardhat-upgrades');

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.9",
  networks: {
    "alfajores": {
      url: "https://alfajores-forno.celo-testnet.org",
      accounts: ["bef9e118a31a0af39585d7e82cae42a7393476d792e36e183a1227f9e3815c50"],
      chainId: 44787,
    },
  },
};
