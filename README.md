# Promissory Note-driven Data Trading Architecture (PNDTA)

**Overview**

The Promissory Note-based Data Trading Architecture (PNDTA) is a pioneering solution designed to address the escalating concerns arising from the personal data collection practices of smart consumer devices. In the last decade, the proliferation of smart devices like internet-connected refrigerators and remote-controlled utilities has become increasingly common, thanks to widespread internet access and advancements in technology. While these innovations have undoubtedly enhanced our quality of life, they have also brought about significant privacy, security, and ethical challenges.

**Key Challenges Addressed:**

 - **Intimate Data Collection**: Smart devices collect highly personal information, such as dietary habits from smart refrigerators, or activity habits from smart watches. This type of data is more intimate than traditional data collection methods.
 - **Security Risks**: The data stored in Original Equipment Manufacturer (OEM) cloud storage is a lucrative target for cyber-attacks, posing a serious security threat.
 - **Economic Disparities**: The economic benefits derived from this data often bypass the individuals who provide it, leading to ethical dilemmas and economic inequities.

**The PNDTA Solution**: 

PNDTA introduces a novel architecture that enhances transparency in data value chains. This architecture is crucial for consumers to understand and control how their data is used and monetized. The key features of PNDTA include:

 - **Consumer Empowerment**: It provides consumers with greater control over their data, allowing them to decide what information is shared and how it is used.
 - **Transparency in Data Usage**: The architecture offers clear insights into the data value chains, revealing how consumer data contributes to various economic activities.
 - **Holistic Approach**: PNDTA is part of a broader initiative that not only focuses on data control but also emphasizes the overall value and transparency of data collected from smart devices.

**Disclaimer**:

This overview of the Promissory Note-driven Data Trading Architecture (PNDTA) is for academic purposes and forms part of a larger research project. It represents a theoretical framework developed within an academic context and is not a commercially available system. The solution discussed are part of this academic exploration and should be considered as such.

# Setup Requisites

This section details all the necessary technologies, tools, and requisites that were used to set up and run the Promissory Note-driven Data Trading Architecture (PNDTA). Please follow these steps carefully to ensure a smooth setup process.

### System Requirements

 - **Operating System**: Ubuntu 22.04.3 LTS

### Technologies and Tools

Ensure you have the following technologies and tools installed. Links to official installation guides are provided for convenience.

**Programming Languages**:

 - Python - 3.8 or higher
 - Typescript
 - Go
 - Javascript

**Databases**:

 - PostgreSQL - 16.1 (when setting up Hyperledger Fabric network)

**Frameworks and Libraries**:

 - Flask - 3.0.0
 - Axios - 1.6.5
 - Web3 - 6.14.0

**Tools and Services**:

 - Docker - Containerization Platform
 - Hardhat - Deploying Smart contracts Tool

**Blockchain Technologies**:

 - [Ocean Protocol V4](https://github.com/oceanprotocol/ocean.py/blob/main/READMEs/install.md)
 - [Hyperledger Fabric 2.5](https://hyperledger-fabric.readthedocs.io/en/release-2.5/)


### How to run the different scripts

Here is how this whole solution should be tested

| Ordem   | Script     | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `1` | `publish_access_remote.py` | Python Script|
| :---------- | :--------- | :---------------------------------- |
| `2` | `network.sh` (from Hyperledger fabric guide) | Fabric network|
| :---------- | :--------- | :---------------------------------- |
| `3` | `command` | Command to deploy Payment smart contract|
| :---------- | :--------- | :---------------------------------- |
| `4` | `app.ts` | Compile and Run Fabric Application|
| :---------- | :--------- | :---------------------------------- |
| `5` | `command` | Commands to interact with Fabric network|
| :---------- | :--------- | :---------------------------------- |

