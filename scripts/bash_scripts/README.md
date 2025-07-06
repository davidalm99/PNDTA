# Automation Shell Scripts for DT Network Operations

This folder contains automation scripts designed to streamline and standardize interactions with the Digital Twin (DT) network and the underlying Hyperledger Fabric infrastructure. These scripts help reduce command-line errors and save time during repetitive operations.

## 📄 Script Overview

### **`dtnetwork.sh`** 
- **Purpose**: Automates the setup and initialization of the Digital Twin network.
- **Functionality**: 
  - Creates and registers devices on the network.
  - These devices represent data-producing entities whose output can later be monetized.

### **`claimDevice.sh`**
- **Purpose**: Facilitates interaction with the Hyperledger Fabric network to claim a device.
- **Functionality**:
  - Executes a command that asserts ownership or access rights over a registered device.

### **`create_agree`**
- **Purpose**: Automates the creation and agreement of a promissory note between two entities.
- **Functionality**:
  - Uses the 'promiseAndClaim.go' smart contract.
  - Establishes mutual agreement on defined terms between two parties.

### **`registerAndBuy`**
- **Purpose**: Automates device registration and subsequent purchase actions.
- **Functionality**:
  - Registers a device according to the defined taxonomy and metadata.
  - Initiates a purchase transaction for the registered device.

## ⚙️ Usage

Ensure each script has execution permissions before running:

```bash
chmod +x <script_name>.sh
```
## 🔁 Recommended Workflow

To ensure proper interaction with the Digital Twin network and avoid execution errors, we recommend the following script execution order:

1. **'dtnetwork.sh'**  
   _→ Initializes the network and registers the devices._  
   Start by creating the network environment and registering devices that will be part of the data economy.

2. **'create_agree.sh'**  
   _→ Establishes a promissory note (agreement) between the involved entities._  
   This sets the legal/contractual context for future transactions.

3. **'registerAndBuy.sh'**  
   _→ Registers the device according to the pre-defined taxonomy and executes a purchase._  
   This step finalizes the registration and simulates an economic transaction over the registered Digital Twin.

4. **`'claimDevice.sh'**  
   _→ Claims ownership or access rights over the purchased device._  
   This step ensures that the entity now officially holds usage rights over the Digital Twin.

### Example Execution

```bash
./dtnetwork.sh
./create_agree.sh
./registerAndBuy.sh
./claimDevice.sh

