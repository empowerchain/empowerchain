# Concepts

## Proof of Existence

Sometimes also called "constructive proof" or "existence proof" in mathematics and is a method of proof that 
demonstrates the existence of a mathematical object by creating or providing a method for creating the object.
In our case, it is a way for someone to prove that some piece of data (document, text, bytes, etc)
existed at a specific point in time (and has not changed since).

A proof of existence is a combination of a hash and a timestamp where both are immutable (and everyone trusts
that they are in fact immutable). Given a piece of data that was hashed and that hash was stored on an immutable blockchain,
anyone who has the same data can at any time prove to themselves or other that the data existed at the point the proof was added on-chain.

An important property of Proof of Existence is also that you don't need to provide the data itself until 
a verification of proof is needed. In other words, all that is stored on-chain is the proof, which cannot be used 
to re-create the original data.

The typical examples include:
- Demonstrating data/IP ownership
- Proving data has not changed/been corrupted
- Document integrity

## Hashing

A cryptographic hash is the output of a mathematical hashing function that turns (maps) a variably sized piece of data
into a small, practically unique, fixed size unique string ("the hash").
Traditionally these hashes are used for security purposes such as avoiding to store passwords in 
databases for software services (instead they just store a hash of the password that was provided 
when the user registered, and then check the provided password upon login with the stored hash). 

The way hashes are used in Proof of Existence is in combination with the immutability of a Blockchain. 
When someone needs to verify some data they can find the hash on the blockchain and compare it to the data they have themselves. 
They only need to run the data through the same hashing function and compare their output with the hash stored on the blockchain. 
If they are the same, the data has not been altered and existed at the time the transaction was entered into the blockchain.

## Hash collision

Whenever you map something large (or at least larger than the output) into something small there will also be the 
theoretical possibility of collisions, where two different datasets yield the same result. 
In other words it is theoretically possible to have one hash function as the correct proof for more than one set of data. 
To illustrate this we can take as example the widely used SHA-256 hashing algorithm: 
It produces a hash made out of 32 bytes (1 byte is 8 bits, so a total of 256 bits) 
and each byte can each represent 256 different values. The total number of possible 
permutations under SHA-256 then becomes 32^256. That is such an absurdly large number 
(it is a number so large that it has 386 digits and is more than 10^56 (or just about 3 septendecillion) 
times larger than the number of atoms in the universe) that any practical chances of a collision is just about inconceivable.