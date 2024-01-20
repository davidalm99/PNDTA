# Payment Mechanism
This application uses a straightforward and custom-made payment system to manage micro-transactions. At the heart of this system is a smart contract named "Payment.sol." Written in Solidity, a programming language for blockchain applications, this contract embodies the simple yet effective logic governing our payment process.

The smart contract "Payment.sol" is deployed on the Mumbai testnet, a network used for testing blockchain technologies. Deployment is conducted through a script named "deploy_Payment.js." This script ensures that our smart contract is correctly integrated into the blockchain. Once deployed, the Payment.sol contract automates the transfer of funds between data buyers and sellers. This system is designed to be efficient and reliable, prioritizing ease of use while ensuring secure and transparent transactions. It's a straightforward solution tailored to our specific transactional needs, offering a hassle-free experience for users.

## Setup

- Before running deploy_payment.js, clone a directory called "contracts" (custom-contracts e.g.) that will be installed after following the Ocean Protocol framework instalation guide. This file should be copied to "~/custom-contracts/scripts" directory
- Copy Payment.sol to "~/custom-contracts/contracts" directory
- Run the following command

```bash
cd ~/custom-contracts

npx hardhat run --network mumbai scripts/deploy_payment.js
```
