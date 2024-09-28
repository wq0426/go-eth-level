// This is a script for deploying your contracts. You can adapt it to deploy
// yours, or create new ones.

const path = require("path");

async function main() {
  // This is just a convenience check
  if (network.name === "hardhat") {
    console.warn(
      "You are trying to deploy a contract to the Hardhat Network, which" +
        "gets automatically created and destroyed every time. Use the Hardhat" +
        " option '--network localhost'"
    );
  }

  // ethers is available in the global scope
  const [deployer1, deployer2] = await ethers.getSigners();
  // const privateKey = "YOUR_PRIVATE_KEY";
  // const provider = new ethers.providers.JsonRpcProvider("http://127.0.0.1:8545");  // 使用本地网络或其他网络
  // const wallet = new ethers.Wallet(privateKey, provider);
  console.log(
    "Deploying the contracts with the account:",
    await deployer1.getAddress(),
    ", deployer2:",
    await deployer2.getAddress(),
  );

  console.log("Account balance:", (await deployer1.getBalance()).toString());

  const Token = await ethers.getContractFactory("ERC20");
  const token = await Token.deploy("My Token", "MT");
  await token.deployed();

  console.log("Token address:", token.address);

  // We also save the contract's artifacts and address in the frontend directory
  saveFrontendFiles(token);
}

function saveFrontendFiles(token) {
  const fs = require("fs");
  const contractsDir = path.join(__dirname, "..", "frontend", "src", "contracts");

  if (!fs.existsSync(contractsDir)) {
    fs.mkdirSync(contractsDir);
  }

  fs.writeFileSync(
    path.join(contractsDir, "contract-address.json"),
    JSON.stringify({ Token: token.address }, undefined, 2)
  );

  const TokenArtifact = artifacts.readArtifactSync("ERC20");

  fs.writeFileSync(
    path.join(contractsDir, "Token.json"),
    JSON.stringify(TokenArtifact, null, 2)
  );
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
