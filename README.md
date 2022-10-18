# DID Methods specification

**About**
<br/>
The Oracle DID method specification conforms to the requirements specified in the DID Specification v1.0 published by the W3C Credentials Community Group.

**Abstract**
<br/>
TODO

**Status of This Document**
<br/>
This the first draft.

**DID Method Syntax**
<br/>
The Oracle DID scheme is defined by the following ABNF:

```
orcl-did = "did:orcl:" idstring
idstring = 16*(base32char)
base32char = "2" / "3" / "4" / "5" / "6" / "7" / "A" / "B" / "C"
        / "D" / "E" / "F" / "G" / "H" / "I" / "J" / "K" / "L" / "M" / "N" / "O" / "P" / "Q"
        / "R" / "S" / "T" / "U" / "V" / "W" / "X" / "Y" / "Z"
```

The Oracle DID Method name is `orcl`. A DID that uses this method must begin with the following prefix: `did:orcl`.

The `idstring` is base32 encoded with length of 16 characters, and is computed through these steps:

1. The HLF `TxID` of 64 characters is split into 16 blocks of 4 characters each;
2. For each block i, we apply a function `f(block_i)->[2-7A-Z]`;
3. The `idstring` is composed of the concatenation of the results of the  function f applied to the 16 blocks.

An example of the resulting DID is: `did:orcl:QC5S3KGCFN37Z5VP`.

The Regex matching the `idstring` is the following:

`^did:orcl:[2-7A-Z]{16}$`.



**DID Method Operations**
<br/>
DIDs are managed through chaincode in HLF, implementing the typical CRUD methods. In what follows, we assume that a DID controller has access to the network, has obtained a HLF client identity (X.509 certificate), and is running a Fabric Client.

***Create***
<br/>
1. A subject/controller calls the method `did.create([public_key])` of the Fabric Chaincode.

2. The Fabric Chaincode generates a DID document identified by a DID generated according to the Syntax defined in "Did Method Syntax".
If the parameter `public_key` is not null, it is stored in the `public_key` field of the generated DID document.

3. The Fabric Chaincode returns the pair `<DID, DID document>` to the client.

***Read***
1. Anyone with access to the blockchain calls the method `did.read(DID)` of the Fabric Chaincode.

2. The Fabric Chaincode returns to the client the DID document associated to the given DID.

***Update***
1. The controller of the document calls the method `did.update(DID, [*args])` of the Fabric Chaincode.

2. The Fabric Chaincode returns to the client the DID document updated with the new parameters as well as `versionId` and `lastModified` fields.

***Delete***
1. The controller of the document calls the method `did.delete(DID)` of the Fabric Chaincode.

2. The Fabric Chaincode returns to the client the DID document with all the verification methods deleted and a flag in the registry set to `deactivated`. According to this, no authentication method can be used to verify the holder's identity and the did is deprecated. This implies that the DID cannot be registered or reactivated.
<br />
<br />

**Implementation**
<br />
The above mentioned operations are implemented in Golang.
In the following, their key aspects are detailed.

<br />

***idstring generation***
<br />
As per <i>DID Method Syntax</i>, dids are made up of the concatenation of the string `orcl:did:` and a unique `idstring`. Such idstrings need to be generated in a deterministic manner, and their generation must guarantee the lowest number of duplicates as possible. <br />
In this section, two methods for the idstring generations are proposed.
To achieve determinism, both of them make use of the Transaction ID (TxID), while uniqueness is guaranteed by exploiting the key concepts of the hash functions.
<br /><br />
Both methods assume that:
- Transaction IDs are unique strings of length 64 with characters in [a-z0-9];
- ...

****Method 1**** 
<br />
This methods considers the TxID as a sequence of 16 blocks of 4 characters each. <br />
For each block j, the algorithm computes its md5 checksum and uses it as a seed for a randomic function returning a value `i` in the range [0, 32). Then, the j<sup>th</sup> character of the idstring is obtained by fetching the i<sup>th</sup> character of the base32char alphabet. <br />
The resulting idstring is made up of the concatenation of the 16 characters thus obtained. <br />

Pseudo-code:
```
Input: TxID // len(TxID)= 64

blocks:= split(TxID, 4) // Split TxID in 16 blocks of size 4
idstring:= ""
for block in blocks:
  checksum:= md5(block)
  index:= RandInt(seed=checksum) // RandInt returns a rand value in [0, 32)
  idstring:= idstring + base32char[index]

return idstring
```

****Method 2**** <br />
Similarly, Method 2 also exploits the TxIDs to generate the idstrings.
The main difference with respect to the previous one resides in the conversion of the TxID blocks into base32 characters. In facts, after splitting the TxID into 16 blocks of 4 characters each, the algorithm transforms each block into a vector `v` of length 4, whose elements are the ASCII representation of its characters.
Then, the checksum of the block's characters is used as a seed for the pseudo-random generation of an odd-integer array `a`. Finally, according to the Hashing Vector technique, the dot product between the vectors `v` and `a` is computed, and its congruent modulo 32 represents the base32 index of the character obteined from the block. <br />
These steps are iterated over the 16 blocks to obtain the 16 characters of the idstring.

Pseudo-code:
```
Input: TxID // len(TxID)= 64
blocks:= split(TxID, 4) // Split TxID in 16 blocks of size 4
idstring:= ""
for block in blocks:
  checksum:=md5(block)
  a:= RandOddVect(seed=checksum)  // Returns 4-dim odd integer vector
  v:= ASCII(block)                // Returns 4-dim vect of ASCII integers
  index:= DotProduct(a,v) mod 32  
  idstring:= idstring + base32char[index]

return idstring
```
It is worth noting that, given its polinomial pattern, the latter method is proven to behave like a uniform hash function, meaning that every output value (i.e. base32 character) is generated with roughly the same probability.

# TODO Section

Things to consider in a future:
1. Add substring definining the environment
2. Implement methods to compute the `idstring`. Two proposals:

    1.  
    ```
    step = ceil(len(txid)/len(idstring))
    string idstring = ""
    for (j=0, j<len(txid), j+=step):
        seed_j = txid[j:j+step]
        idstring_j = Rand(seed_j) mod 32
        idstring+=idstring_j
    return idstring
    ```
    2.  
    ```
    p = rand(seed)
    step = ceil(len(txid)/len(idstring))
    string idstring = ""
    for (j=0, j<len(txid), j+=step):
      d = 1
      for char in txid[j:j+step]:
        int i = getASCII(char) //a-zA-Z0-9
        int d = d*i
      b_i = f(p,d) mod 32
      idstring+=base32char(b_i)
    return idstring
   ```


# did-methods-spec
Did methods specifications

Functions to be implemented:

Function 1

- Function f(string) -> block<sub>1</sub>, block<sub>2</sub>, block<sub>i</sub>, 
...,
block<sub>n</sub> 

- Hash function h(block<sub>i</sub>) -> seed

- Function f(seed) -> [0-31]

- Mapping index to base32
f([0-31]) -> [2-7A-Z]

-----
Function 2

- Function f(string) -> block<sub>1</sub>, block<sub>2</sub>, block<sub>i</sub>, 
...,
block<sub>n</sub> 

- Hash function h(block<sub>i</sub>) -> seed

- Function f(seed) -> [0-31]

- Generate 4-dimensional array of odd integers from seed (seed can be integer or 4-dim array)
f([int]) -> [int, int, int, int] or f([int, int, int, int]) -> [int, int, int, int]


- Dot product of two equal-length arrays
dp([a, b]) -> [a<sub>1</sub>*b<sub>1</sub>, a<sub>2</sub>*b<sub>2</sub>, ..., a<sub>n</sub>*b<sub>n</sub>]

- Hash function
h(x,a,w,M) -> (dp(x,a) mod 2<sup>2w</sup>)) div 2<sup>(2w-M)</sup>
