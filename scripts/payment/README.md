# 💸 Ocean Protocol – Payment Smart Contract & Deployment

This module includes a Solidity smart contract and a deployment script to handle payments between buyers and sellers using the **Ocean Protocol ERC20 token** on the **Polygon Mumbai testnet**.

It is designed to work seamlessly with the Ocean Protocol stack and is typically used to facilitate payments for data access in a decentralized data marketplace setup.

---

## 📁 File Descriptions

### `contracts/Payment.sol`
- Solidity smart contract that manages payment execution logic.
- It likely exposes a `triggerPayment(...)` function to transfer tokens from a buyer to a seller.
- Integrates with Ocean Protocol's ERC20 token and supports modular use with existing datasets and asset agreements.
- ✅ **Recommended Location**: Place this file under the `contracts/` folder of your **Ocean Protocol Barge stack** (e.g., `~/ocean-protocol/barge/contracts/Payment.sol`).

### `scripts/deploy_payment.js`
- JavaScript deployment script using **Hardhat** and **ethers.js**.
- Reads the deployed Ocean token address from the Mumbai testnet config file (`address.json`).
- Deploys the `Payment` contract and writes its address to `payment_address.txt` for use by backend services or smart contract consumers.

---

## 🌐 Polygon Mumbai Testnet

The deployment targets the **Mumbai testnet**, and it uses the Ocean token contract deployed by the Ocean team.

Ensure you have the Mumbai RPC endpoint set:

```bash
https://polygon-mumbai-bor.publicnode.com
```

## Setup

- Before running deploy_payment.js, clone a directory called "contracts" (custom-contracts e.g.) that will be installed after following the Ocean Protocol framework instalation guide. This file should be copied to "~/custom-contracts/scripts" directory
- Copy Payment.sol to "~/custom-contracts/contracts" directory
- Run the following command

```bash
cd ~/custom-contracts

npx hardhat run --network mumbai scripts/deploy_payment.js
```
