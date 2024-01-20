# Scripts Description

### promiseAndClaim.go
  This Golang script acts as a custom chaincode deployed within the Hyperledger Fabric network. Its primary function is to introduce additional logic to the blockchain operations. The script enhances the network's capabilities by enabling specific business logic or rules to be executed as part of the blockchain transactions. This includes validating, modifying, or appending data on the ledger. It's designed to ensure that all operations adhere to the defined protocols and maintain the integrity and security of the blockchain network.

## Setup

- After launching the Fabric network in test-network directory with the following command:

```bash
./network.sh up createChannel -d dtnetwork -ca -s couchdb
```

- Here is the command to deploy the chaincode promiseAndClaim.go

```bash
./network.sh deployCC -c dtnetwork -ccn promiseAndClaim -ccp [directory of the chaincode] -ccl go
```

- Define the following env variables:

`PATH` = ${PWD}/../bin:$PATH

`FABRIC_CFG_PATH` = $PWD/../config/

`CORE_PEER_TLS_ENABLED` = true

`CORE_PEER_LOCALMSPID` = "Org1MSP"

`CORE_PEER_TLS_ROOTCERT_FILE` = ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt

`CORE_PEER_MSPCONFIGPATH` = ${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp

`CORE_PEER_ADDRESS` = localhost:7051

### publish_access_remote.py
  This Python script is tailored for managing interactions with the Ocean Protocol framework. It facilitates the integration of blockchain technology with data sharing and monetization services provided by Ocean Protocol. The script sets up a robust endpoint for HTTP requests, enabling seamless communication between the Hyperledger Fabric application and Ocean Protocol's ecosystem. This allows users to securely publish, share, and monetize their data while ensuring compliance with the protocol's standards for data privacy and security. The endpoint is essential for bridging the Fabric application with Ocean Protocol's decentralized data exchange functionalities.

 - After following the Ocean Protocol instalation guide, create a python virtual environment and run the script

```bash
python publish_access_remote.py
```
