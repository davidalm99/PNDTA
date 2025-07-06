# Define the directory where your wallets are stored
WALLET_DIR="/home/david/fabric/fabric-samples/test-network/client-app/wallet"

# Delete the existing wallets
rm -rf $WALLET_DIR/*

# Start the network
./network.sh up createChannel -c dtnetwork -ca -s couchdb

./network.sh deployCC -c dtnetwork -ccn promiseAndClaim -ccp ~/fabric/fabric-samples/promiseAndClaim -ccl go

# Define environment variables
export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

# Enroll the admin and register the users
cd $PWD/client-app
node enrollAdmin.js
node registerUser.js appUser2
node registerUser.js appUser
