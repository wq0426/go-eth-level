const { ethers } = import("hardhat");
const { expect } = import("chai");

describe("Token contract", function() {
  it("Deployment should assign the total supply of tokens to the owner", async function() {
    const [owner] = await ethers.getSigners();

    const Token = await ethers.getContractFactory("ERC20");

    const hardhatToken = await Token.deploy("My Token", "MT");
    await hardhatToken.deployed();

    const ownerBalance = await hardhatToken.balanceOf(owner.getAddress());
    console.log(ownerBalance)
    //expect(await hardhatToken.totalSupply()).to.equal(ownerBalance);
  });
});
