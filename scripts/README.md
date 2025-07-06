# 🌊 Ocean Protocol – Remote Publishing & Access Script

This Python script automates the publication, access, and payment flow of a dataset on the **Ocean Protocol**, using the **Polygon Mumbai testnet**. It integrates data asset publishing, token-based access control, smart contract payments, and Flask-based API endpoints.

---

## 📄 Script Summary

`publish_access_remote.py` performs the following:

1. **Initializes Ocean Protocol and Web3 providers**.
2. **Creates buyer and seller accounts** with pre-funded test wallets.
3. **Publishes a dataset** to IPFS and registers it as an Ocean asset.
4. **Allows access and consumption** of the dataset by a designated buyer.
5. **Triggers payment** between buyer and seller using a custom smart contract.
6. **Exposes an API endpoint (`/trigger_payment`)** to initiate on-chain payment flows.

---

## 🧩 Key Dependencies

Below are the main Python libraries used, with explanations:

### 🔹 `ocean-lib`
- SDK for interacting with the Ocean Protocol.
- Key classes used:
  - `Ocean` – Core object for asset and token management.
  - `AssetArguments`, `DatatokenArguments`, `DispenserArguments` – For publishing assets.
  - `DatatokenBase` – For interacting with datatokens.
  - `to_wei()` – For unit conversion.

### 🔹 `web3.py`
- Ethereum/Web3 library to send transactions, estimate gas, encode data, and interact with smart contracts.
- Used for:
  - Custom contract interactions (e.g., `approve`, `triggerPayment`)
  - Reading/writing balances and allowances

### 🔹 `eth-account`
- Used to derive account objects from private keys.

### 🔹 `requests`
- For external API calls (e.g., fetching OCEAN/EUR price from CoinGecko).

### 🔹 `flask`
- Creates a lightweight HTTP API endpoint:
  - `/trigger_payment` handles POST requests to process payments between buyers and sellers.

---

## 🔧 Setup Instructions

### 1. **Install Requirements**

```bash
pip install ocean-lib web3 eth-account flask requests


