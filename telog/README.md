# Tamper Evident Log

## Design
A tamper evident log consists of a head of type hash pointer.

A block consists of a hash pointer field and a data field.

A hash pointer consists of a pointer to the previous block and a hash digest of the previous block.

When an adversary attempts to change the data in a block, the hash pointer pointing to the changed block will have a different hash than the new hash of the changed block. Editing the hash of the hash pointer will itself change the hash of the next block. The only way to tamper with the log undetected is to change every hash pointer, all the way to the head.

<img src="/images/tamper-evident-log.png" alt="Tamper-evident Log Design">

## Side Note
We added an Attack function that modifies data at an arbitrary block. This helps us to make sure we have implemented Check correctly.

## References

1. Hash: [sha256](https://golang.org/pkg/crypto/sha256/)

2. [How to hash a struct?](https://blog.8bitzen.com/posts/22-08-2019-how-to-hash-a-struct-in-go)
