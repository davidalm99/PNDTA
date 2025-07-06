# 📘 Chaincode Files Overview

This directory includes three Go-based chaincode files developed for a Hyperledger Fabric network, facilitating interactions within a Digital Twin and asset trading ecosystem.

## 🔹 Files Description

### `createPromise.go`
- Defines the creation logic for promissory notes between a buyer and a seller.
- Used to generate and store agreement terms regarding asset usage and monetization.

### `promiseAndClaim.go`
- A merged smart contract handling full lifecycle operations:
  - Creating and signing promissory notes,
  - Registering and claiming devices,
  - Managing payment distribution based on agreement terms.

### `registerAsset.go`
- Provides foundational asset and device registration capabilities.
- Includes support for registering agents, devices, device models, and data batches on the ledger.

---

Each file implements `contractapi.Contract` and operates on Fabric’s world state. These contracts are key building blocks of the Digital Twin monetization network.
