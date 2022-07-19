# State

The `proofofexistence` module keeps track of the proofs of existence uploaded to the chain.

## Proof

A Proof consists of a hash, a timestamp and a reporter address. This gives you enough information to prove that
a hash existed at a specific point in time. You can prove you knew the hash at said time (if you can also prove 
that you are the owner of the reporter address).

- Proof `0x01 | hash | -> ProtocolBuffer(Proof)`

The hash is not made on-chain to avoid the privacy issues that would entail. Instead, the hash is just checked to
be a valid SHA-256 hash.

The timestamp is simply taken from current block's header (i.e. the time of the block).

The reporter address is the address of the account that is creating the transaction.