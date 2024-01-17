import os
from ocean_lib.example_config import get_config_dict
from ocean_lib.ocean.ocean import Ocean
from eth_account import Account
import json
import time
import requests
from datetime import datetime
from ocean_lib.ocean.ocean import Ocean
from ocean_lib.ocean.util import to_wei
from ocean_lib.ocean.ocean_assets import AssetArguments
from ocean_lib.models.dispenser import DispenserArguments
from ocean_lib.models.datatoken_base import DatatokenArguments, DatatokenBase
from web3 import Web3
from flask import Flask, request
app = Flask(__name__)

w3 = Web3(Web3.HTTPProvider('https://polygon-mumbai-bor.publicnode.com'))

# setup() is necessary to create a data buyer and a data seller, returns an Ocean object, 
# which is necessary to interact with the whole Ocean API
def setup():
    os.environ['REMOTE_TEST_PRIVATE_KEY1'] = '0x54a3bf9b6a3f6206b2191d112400129d6b7e9ee71a4a2f32eafcb0ea34b79e24'
    os.environ['REMOTE_TEST_PRIVATE_KEY2'] = '0x9704de265695c41d0b93ee1c7bf8a6889b8c3babaaa99ed9c9f212e6499d7878'

    #If needed go to https://chainlist.org/chain/80001
    os.environ['MUMBAI_RPC_URL'] = 'https://polygon-mumbai-bor.publicnode.com'

    #If needed go to https://chainlist.org/chain/137
    os.environ['POLYGON_RPC_URL'] = 'https://polygon.llamarpc.com'

    config = get_config_dict("mumbai")
    ocean = Ocean(config)
    
    return ocean

# create_accounts() creates 2 wallets, based on setup() private keys. 
# Returns 2 accounts and their respectively public addresses
def create_accounts(ocean):

    OCEAN = ocean.OCEAN_token

    seller_private_key = os.getenv("REMOTE_TEST_PRIVATE_KEY1")
    seller_account = Account.from_key(private_key=seller_private_key)
    seller_address = seller_account.address
    assert ocean.wallet_balance(seller_account) > 0, "Seller needs MATIC"
    assert OCEAN.balanceOf(seller_account) > 0, "Seller needs OCEAN"

    buyer_private_key = os.getenv("REMOTE_TEST_PRIVATE_KEY2")
    buyer_account= Account.from_key(buyer_private_key)
    buyer_address = buyer_account.address
    assert ocean.wallet_balance(buyer_account) > 0, "Buyer needs MATIC"
    assert OCEAN.balanceOf(buyer_account) > 0, "Buyer needs OCEAN"

    print("Buyer's balance in OCEAN: ", OCEAN.balanceOf(buyer_account), "Seller's balance in OCEAN: ", OCEAN.balanceOf(seller_account))

    return seller_account, seller_address, buyer_account, buyer_address

# getContractValues() is used to interact with the Solidity smart contract (Payment.sol) that was developed in this project.
# It is important to note that the directory can be subject to change. In this project, I created a new folder where this 
# process is done
def getContractValues():
    path_txt_file = os.path.expanduser("~/ocean-custom-contracts/payment_address.txt")
    with open(path_txt_file, 'r') as contract_address:
        payment_address = contract_address.read().strip()

    path_abi_file = os.path.expanduser("~/ocean-custom-contracts/artifacts/contracts/Payment.sol/Payment.json")
    with open(path_abi_file, 'r') as contract_abi:
        payment_abi = json.load(contract_abi)

    abi = payment_abi["abi"]

    payment_contract = w3.eth.contract(address=payment_address, abi=abi)

    return payment_address, payment_contract

# publish_dataset() is used to publish a batch of data from data seller address, making sure only data buyer can have access
# and download the batch
# In this function, it is assumed that the data is stored in IPFS, and the CID might be subject to change, depending on the data batch!
def publishing_dataset(ocean, buyer, seller):
    dataset_name = "Test dataset"
    CID = 'QmT1fzcAEhy8YUJWyohPA3miUo3v4mPjiPqh1q6Ri19rgX'
    dataset_url = f"https://ipfs.io/ipfs/{CID}"

 # This will ensure that only ACME will get access to the data shared by
    credentials = {
        "allow": [{"type":"address", "values": [buyer.address]}],
        "deny": [],
    }

    #Setting the price free
    print("Price is set")
    pricing = DispenserArguments()

    #Metadata of the dataset
    metadata = {
        "description": "A nice data set for testing",
        "copyrightHolder": "",
        "name": "data_gen",
        "type": "dataset",
        "author": "Aliceee",
        "license": "No License Specified",
    }

    #Encapsulate all the arguments of the asset
    print("Arguments are encapsulated")
    asset_args = AssetArguments(
        pricing_schema_args=pricing, 
        metadata=metadata, 
        credentials=credentials
    )

    print("Creates data_nft, datatokens and DDO")
    data_nft, datatoken, ddo = ocean.assets.create_url_asset(
        dataset_name,
        dataset_url,
        {"from": seller}, #Alice
        args = asset_args
    )

    print(f"Here is what I'm looking for: {ddo.datatokens[0]['address']}")

    return data_nft, datatoken, ddo

# accessAndConsume() is used to simulate the query searching for the data batch published, finding it and downloading it. All the services
# used in this function (e.g. Provider) are default.
def accessAndConsume(ocean, ddo, buyer, seller):
    counter = 0
    #Seller grants access for "free"
    config = get_config_dict("mumbai")

    query = {
        "query": {
            "query_string": {
                "query": "Aliceee",
                "fields": ["metadata.author"],
            }
        }
    }

    time.sleep(3)
    #List has the set of DDOs that has Alice as its author
    print("Query for search Alice's dataset")
    list_DDOs = ocean.assets.query(query)

    for ddo in list_DDOs:
        print(ddo.metadata["name"], "-", ddo.metadata["author"], "-", ddo.datatokens[0]['address'])    
        datatoken_address = ddo.datatokens[0]['address']


    datatoken1 = DatatokenBase.get_typed(config, datatoken_address)

    print("Dispenser created")
    datatoken1.create_dispenser({"from": seller})
    
    print("Datatoken dispensed")
    datatoken1.dispense(to_wei(1), {"from": buyer})

    # #Buyer pays for the access
    order_tx = ocean.assets.pay_for_access_service(ddo, {"from": buyer}, consumer_address=buyer.address).hex()

    # print(f"Buyer downloads the dataset with did: {ddo.did}")
    ocean.assets.download_asset(ddo, buyer, './', order_tx)

# getInstanceIERC20Contract() is used to get an instance of a IERC20 contract from Ocean API.
# Note that .ocean directory exists by default after Ocean instalation is done, which means that
# these directories are default.
def getInstanceIERC20Contract():
    path_abi_file = os.path.expanduser("~/.ocean/ocean-contracts/artifacts/contracts/interfaces/IERC20.sol/IERC20.json")
    with open(path_abi_file, 'r') as contract_abi:
        IERC20_abi = json.load(contract_abi)

    abi = IERC20_abi["abi"]

    path_contract_ocean_token = os.path.expanduser("~/.ocean/ocean-contracts/artifacts/address.json")
    with open(path_contract_ocean_token, 'r') as contract_address:
        IERC20_address = json.load(contract_address)
    
    address = IERC20_address["mumbai"]["Ocean"]

    IERC20_contract = w3.eth.contract(address=address, abi=abi)

    return IERC20_contract

# display_tx() is a function that was used for testing purposes, a specific group of fields within a transaction were gathered and printed out
def display_tx(receipt, limit):
    BLUE = "\033[94m"
    RESET = "\033[0m"
    
    tx_hash = receipt['transactionHash'].hex()
    gas_price = receipt.get('effectiveGasPrice', None)
    block_number = receipt['blockNumber']
    gas_used = receipt['gasUsed']

    print("--------------------------------------------------")    
    print(f"Transaction Address: {BLUE}{tx_hash}{RESET}")
    print(f"Gas Price: {BLUE}{gas_price} wei{RESET} Gas Limit: {BLUE}{limit} gwei{RESET}")
    print(f"Transaction Confirmed   Block: {BLUE}{block_number}{RESET} Gas Used: {BLUE}{gas_used}{RESET}")
    print("--------------------------------------------------")

# txProcess() is a function that is developed to process each transaction that happens in payment()
# This function does sign the transaction and sends it. The implementation of this function can be modified
# in order to avoid the explicity of the buyer's private key, which is a bad principle!
def txProcess(_to, _from, data, nonce, chainID, gasPrice):

    buyer_private_key = "0x9704de265695c41d0b93ee1c7bf8a6889b8c3babaaa99ed9c9f212e6499d7878"

    approve_txn = {
        'to': _to,
        'from': _from,
        'data': data,
        'nonce': nonce,
        'chainId': chainID,
        'gasPrice': gasPrice
    }

    estimated_gas_approve = w3.eth.estimate_gas(approve_txn)

    gas_limit_approve = int(estimated_gas_approve * 1.20)

    approve_txn['gas'] = gas_limit_approve

    # Sign the transaction
    signed_approve_txn = w3.eth.account.sign_transaction(approve_txn, buyer_private_key)

    # Send the signed transaction
    tx_hash = w3.eth.send_raw_transaction(signed_approve_txn.rawTransaction)
    receipt = w3.eth.wait_for_transaction_receipt(tx_hash)

    display_tx(receipt, gasPrice)

    return receipt

def isPayed(id, receipt_1, receipt_2):
    if receipt_1['status'] != 1:
        print(f"Approval from promissory note {id} failed")
        return False
    elif receipt_2['status'] != 1:
        print(f"Payment from promissory note {id} failed")
        return False
    else:
        return True

# payment() is used to transfer funds from buyer's wallet to seller's wallet.
def payment(id, buyer, seller, price, payment_address, payment_contract):
    OCEAN_Contract = getInstanceIERC20Contract()

    # Fetch OCEAN/EUR price from CoinGecko API
    response = requests.get('https://api.coingecko.com/api/v3/simple/price?ids=ocean-protocol&vs_currencies=eur')
    data = response.json()
    OCEAN_EUR_price = data['ocean-protocol']['eur']

    # Convert price in OCEAN
    _price_OCEAN = float(price)/float(OCEAN_EUR_price)

    # Conversion of the payment received to wei
    _price = to_wei(float(_price_OCEAN))

    print(f"OCEAN amount in the transfer: {_price_OCEAN}")

    receipt_approval = txProcess(OCEAN_Contract.address,
                                 buyer, OCEAN_Contract.encodeABI(fn_name="approve", args=[payment_address, _price]),
                                 w3.eth.get_transaction_count(buyer),
                                 80001,
                                 w3.to_wei('4', 'gwei'))

    allowance = OCEAN_Contract.functions.allowance(buyer, payment_address).call()
    assert allowance >= _price, "Allowance is less than required"

    buyer_ocean_balance_after = OCEAN_Contract.functions.balanceOf(buyer).call()
    buyer_ocean_balance_after_gwei = buyer_ocean_balance_after / 10 ** 18
    print(f"OCEAN balance of buyer: {buyer_ocean_balance_after_gwei}")

    seller_ocean_balance_after = OCEAN_Contract.functions.balanceOf(seller).call()
    seller_ocean_balance_after_gwei = seller_ocean_balance_after / 10 ** 18
    print(f"OCEAN balance of seller {seller_ocean_balance_after_gwei}")

    receipt_triggerPayment = txProcess( payment_contract.address,
                                        buyer,
                                        payment_contract.encodeABI(fn_name="triggerPayment", args=[id, buyer, seller, _price]),
                                        w3.eth.get_transaction_count(buyer),
                                        80001,
                                        w3.to_wei('4', 'gwei'))

    buyer_ocean_balance_after = OCEAN_Contract.functions.balanceOf(buyer).call()
    buyer_ocean_balance_after_gwei = buyer_ocean_balance_after / 10 ** 18
    print(f"OCEAN balance of buyer: {buyer_ocean_balance_after_gwei}")

    seller_ocean_balance_after = OCEAN_Contract.functions.balanceOf(seller).call()
    seller_ocean_balance_after_gwei = seller_ocean_balance_after / 10 ** 18
    print(f"OCEAN balance of seller {seller_ocean_balance_after_gwei}")

    receipt_allowance_receipt = txProcess(OCEAN_Contract.address,
                                 buyer, OCEAN_Contract.encodeABI(fn_name="approve", args=[payment_address, 0]),
                                 w3.eth.get_transaction_count(buyer),
                                 80001,
                                 w3.to_wei('4', 'gwei'))

    return isPayed(id, receipt_approval, receipt_triggerPayment)

# trigger_payment is the function responsible for calling payment(). This function is triggered
# when a payload, or a set of payloads, are sent to the Flask endpoint, from Fabric's application, using Typescript SDK.
@app.route('/trigger_payment', methods=['POST'])   
def trigger_payment():
    payloads = request.get_json()

    payload_list = []
    response_messages = []

    for payload in payloads:
        ID = payload.get('ID')
        BuyerAddr = payload.get('BuyerAddr')
        SellerAddr = payload.get('SellerAddr')
        Payment = payload.get('Payment')

        payload_list.append({
            'ID':           ID,
            'BuyerAddr':    BuyerAddr,
            'SellerAddr':   SellerAddr,
            'Payment':      Payment,          
        })

    for payload in payload_list:
        payment_address, payment_contract = getContractValues()

        if payment(payload['ID'], payload['BuyerAddr'], payload['SellerAddr'], payload['Payment'], payment_address, payment_contract):
            response_messages.append('Payment Successful for ID: ' + str(payload['ID']))
        else:
            response_messages.append('Payment Unsuccessful for ID: ' + str(payload['ID']))
        time.sleep(5)
        
    if 'Payment Unsuccessful' in response_messages:
        return {"messages": response_messages}, 400
    else:
        return {"messages": response_messages}, 200

if __name__ == "__main__":
    ocean = setup()
    seller_account, seller_address, buyer_account, buyer_address = create_accounts(ocean)

    start_time = time.time()
    data_nft, datatoken, ddo = publishing_dataset(ocean, buyer_account, seller_account)

    accessAndConsume(ocean, ddo, buyer_account, seller_account)
    end_time = time.time()

    ellapsed_time = end_time - start_time

    print(f"Trade Efficiency Time: {ellapsed_time} seconds")

    app.run(port=6000)
