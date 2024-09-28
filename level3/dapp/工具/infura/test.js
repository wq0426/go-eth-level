const ethers = require('ethers');

// Infura 节点的 API 密钥
const INFURA_API_KEY = '';

// Infura 节点的 URL
const INFURA_URL = `https://mainnet.infura.io/v3/${INFURA_API_KEY}`;

// 创建一个新的 ethers.js 提供者
const provider = new ethers.JsonRpcProvider(INFURA_URL);

// 查询最新区块
provider.getBlockNumber().then((blockNumber) => {
  console.log(`最新区块号：${blockNumber}`);
  // 查询最新区块的详细信息
  provider.getBlock(blockNumber).then((block) => {
    console.log(`最新区块的详细信息：`);
    console.log(block);
  });
});