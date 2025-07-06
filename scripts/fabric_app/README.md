# 🔐 Identity Management & Chaincode Interaction for Hyperledger Fabric Application

This module contains scripts and an application entry point for identity enrollment, user registration, and chaincode interaction within a **Hyperledger Fabric** network using **TypeScript** and **Node.js**.

---

## 📄 File Overview

### `enrollAdmin.js`
- Enrolls the default **admin identity** with the Fabric Certificate Authority (CA).
- Stores the resulting X.509 credentials in a local file-based wallet.
- Must be run before registering other users.

### `registerUser.js`
- Registers and enrolls a new user with the Fabric CA using the enrolled admin.
- Also stores the resulting identity in the wallet.
- Requires a username as a command-line argument.

### `app.ts`
- TypeScript-based application that likely performs transactions or queries on the Fabric network.
- Uses the registered user identity from the wallet to connect to the network and interact with chaincode.

---

## 📦 Prerequisites

Make sure your environment has the following installed:

### 🔧 Core Tools
- Node.js (v14 or later recommended)
- TypeScript (`npm install -g typescript`)
- Fabric CA Client SDKs:
  ```bash
  npm install fabric-network fabric-ca-client

## Setup

Before compiling the app.ts, we first need to ensure that the Admin (of the Fabric network) and Users (of the Fabric network) are registred. All these files should be in the same directory

```bash
node enrollAdmin.js

node registerUser.js appUser

node registerUser.js appUser2 (optional, if we want to test 2 different applications instances)
```

To compile the app.tsc

```bash
tsc app.ts

node app.js
```
