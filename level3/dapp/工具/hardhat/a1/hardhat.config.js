require("@nomiclabs/hardhat-waffle");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.27",
  networks: {
    localhost: {
      url: 'http://127.0.0.1:7545',
      accounts: ['0x2ABE012E82FB6dF4942330cd709d4E06a734EE84'],
    },
  }
};
