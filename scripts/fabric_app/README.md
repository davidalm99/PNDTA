# Fabric Application
This application serves as a specialized observer within the Fabric blockchain network. Its main role is to keenly monitor a specific type of event occurring on the network. Imagine it as a diligent sentinel, always on the lookout for a particular signal. Once this event is detected, the application springs into action, gathering important details related to the event.

One of the key functionalities of this application is its ability to gather relevant information from Promissory Notes. These notes act like digital contracts, which are used as a promise between a data buyer and a data seller and a guarantee that data seller will be rewarded. Upon the occurrence of the monitored event, our application automatically compiles the necessary information from these Promissory Notes and dispatches it via an HTTP payload to the Python Flask endpoint. This process facilitates a seamless and secure exchange of value between those who seek data and those who provide it.

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
