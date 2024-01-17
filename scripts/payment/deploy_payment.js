const { ethers } = require("hardhat");
const fs = require("fs");

async function main() {
  const [deployer] = await ethers.getSigners();

  console.log("Deploying the Payment contract with the account:", deployer.address);

  // Read Ocean Token address from address.json
  const fromAddressJson = JSON.parse(fs.readFileSync("/home/david/.ocean/ocean-contracts/artifacts/address.json", "utf-8"));
  const oceanTokenAddress = fromAddressJson['mumbai']['Ocean'];
  
  const Payment = await ethers.getContractFactory("Payment");
  //const promissoryNote = await PromissoryNote.deploy();
  const payment = await Payment.deploy(oceanTokenAddress);

  console.log("Payment contract address:", payment.address);

  // Write the contract address to a .txt file
  fs.writeFileSync("payment_address.txt", payment.address);

}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
