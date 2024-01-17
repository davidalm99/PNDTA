pragma solidity 0.8.12;
// SPDX-License-Identifier: (Apache-2.0 AND CC-BY-4.0)
// Code is Apache-2.0 and docs are CC-BY-4.0

import "./interfaces/IERC20.sol";

contract Payment {
    IERC20 public oceanToken;

    constructor(address _oceanTokenAddress) {
        oceanToken = IERC20(_oceanTokenAddress);
    }

    event PaymentSuccess(uint256 agreementId, uint256 paymentAmount);

    function triggerPayment(uint256 agreementId, address buyer_addr, address seller_addr, uint256 payment) public returns (bool) {
        bool paymentDone = false;

        bool transferred = oceanToken.transferFrom(buyer_addr, seller_addr, payment);
        require(transferred, "Payment failed");

        paymentDone = true;
        emit PaymentSuccess(agreementId, payment);

        return paymentDone;
    }

}
