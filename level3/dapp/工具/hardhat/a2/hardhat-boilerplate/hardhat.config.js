require("@nomicfoundation/hardhat-toolbox");

// The next line is part of the sample project, you don't need it in your
// project. It imports a Hardhat task definition, that can be used for
// testing the frontend.
require("./tasks/faucet");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.17",
  networks: {
    localhost: {
      url: "http://127.0.0.1:7545",
      //accounts: [`0x204b54b23422b9cfe598f23f2d49334bce06fd6e38cc4e3e3c0b95bcdb1a0e3b`, `0xc540b747e2d28606d16b3af50d5f9cd088fa9cfc6bda4ebeba5a4a5b9f752ab6`],
    } 
  }
};
