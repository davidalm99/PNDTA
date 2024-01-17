import { Gateway, Wallets, ContractEvent } from 'fabric-network';
import * as path from 'path';
import * as fs from 'fs';
import axios from 'axios';

interface Agreement {
    ID:         string;
    BuyerAddr:  string;
    SellerAddr: string;
    AssetType:  string;
    Payment:    string;
}

async function main() {
    try {
        // Load the network configuration
        const ccpPath = path.resolve(__dirname, '..', 'organizations', 'peerOrganizations', 'org1.example.com', 'connection-org1.json');
        const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = await Wallets.newFileSystemWallet(walletPath);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get('appUser');
        if (!identity) {
            console.log('Identity for the user does not exist in the wallet');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'appUser', discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('dtnetwork');

        // Get the contract from the network.
        const contract = network.getContract('promiseAndClaim');

        // How to interact with contract functions
        await contract.submitTransaction('InitLedger');
        console.log('Transaction for promiseAndClaim has been submitted');

        // Define your listener function
        const listener = async (event: ContractEvent) => {
            // Only handle 'ClaimedDevice' events
            if (event.eventName === 'ClaimedDevice') {
                console.log(`Received event: ${event.eventName}`);

                //console.log('Here is the payload: ' + event.payload.toString())
                // Parse the event payload as JSON
                const payloadString = event.payload.toString();
                const eventPayload = JSON.parse(payloadString);

                const deviceModelID = eventPayload[' DeviceModelID '];

                    // Call GetAllAgreements
                const resultBuffer = await contract.evaluateTransaction('GetAllAgreements');
                const rawAgreements = JSON.parse(resultBuffer.toString());

                const agreements: Agreement[] = rawAgreements.map((rawAgreement: any) => {
                console.log(rawAgreement.payment.toString())
                    return {
                        ID:         rawAgreement.id,
                        BuyerAddr:  rawAgreement.eth_buyer_addr,
                        SellerAddr: rawAgreement.eth_seller_addr,
                        AssetType:  rawAgreement.assetType,
                        Payment:    rawAgreement.payment.toString(),
                    };
                });

                    const filteredAgreements = agreements.filter((agreement: Agreement) => {
                        return agreement.AssetType === deviceModelID;
                    });

                    console.log(filteredAgreements);

                    // Send a POST request to the Flask server
                    const response = await axios.post('http://localhost:6000/trigger_payment', filteredAgreements, {
                        headers: { 
                            'Content-Type': 'application/json'
                        }
                    });
            }
        };

        // Add a listener for the event
        await contract.addContractListener(listener);

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        process.exit(1);
    }
}

main();
