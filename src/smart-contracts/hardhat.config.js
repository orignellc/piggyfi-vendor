require("@nomicfoundation/hardhat-toolbox");
require("@openzeppelin/hardhat-upgrades");
require("hardhat-deploy");
require("@nomiclabs/hardhat-etherscan");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: {
    compilers: [
      {
        version: "0.8.2",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
        },
      },
      {
        version: "0.8.9",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
        },
      },
    ],
  },
  networks: {
    hardhat: {
      accounts: {
        mnemonic:
          "leg hamster soccer evoke candy garlic penalty novel feed they angle never",
        path: "m/44'/60'/0'/0",
        initialIndex: 0,
        count: 20,
        passphrase: "",
      },
    },
    localhost: {
      url: "http://127.0.0.1:8080",
      accounts: {
        mnemonic:
          "leg hamster soccer evoke candy garlic penalty novel feed they angle never",
        path: "m/44'/60'/0'/0",
        initialIndex: 0,
        count: 20,
        passphrase: "",
      },
    },
    alfajores: {
      url: "https://alfajores-forno.celo-testnet.org",
      accounts: [
        "bef9e118a31a0af39585d7e82cae42a7393476d792e36e183a1227f9e3815c50",
      ],
      chainId: 44787,
    },
    baobab: {
      url: "https://api.baobab.klaytn.net:8651/",
      accounts: [
        "bef9e118a31a0af39585d7e82cae42a7393476d792e36e183a1227f9e3815c50",
      ],
      chainId: 1001,
    },
  },
  etherscan: {
    apiKey: {
      alfajores: "K7KR894CXR3JMRI6H84RIV5WT9X9FKS71K",
    },
    customChains: [
      {
        network: "alfajores",
        chainId: 44787,
        urls: {
          apiURL: "https://api-alfajores.celoscan.io/api",
          browserURL: "https://alfajores.celoscan.io",
        },
      },
      {
        network: "celo",
        chainId: 42220,
        urls: {
          apiURL: "https://api.celoscan.io/api",
          browserURL: "https://celoscan.io/",
        },
      },
    ],
  },
};
