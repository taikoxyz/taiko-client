// SPDX-License-Identifier: GPL-3.0

pragma solidity 0.8.24;

contract BallotTest {
    error L1_BLOB_NOT_FOUND();
    bytes32 blobHash;
    function storeBlobHash () external {
        blobHash = blobhash(0);
        if (blobHash == 0) revert L1_BLOB_NOT_FOUND();
    }
}