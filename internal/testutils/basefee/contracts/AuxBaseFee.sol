// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "./Lib1559Math.sol";

contract AuxBaseFee {
    function baseFee(uint256 gasExcess_, uint8 basefeeAdjustmentQuotient, uint32 gasTargetPerL1Block) public pure returns (uint256 basefee_) {
        basefee_ = Lib1559Math.basefee(
            gasExcess_, uint256(basefeeAdjustmentQuotient) * gasTargetPerL1Block
        );
    }
}