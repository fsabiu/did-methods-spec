# did-methods-spec
Did methods specifications

Functions to be implemented:


- Mapping TxID to index
f([a-z0-9]) -> [0, 35]

- Generate 4-dimensional array of odd integers from seed (seed can be integer or 4-dim array)
f([int]) -> [int, int, int, int] or f([int, int, int, int]) -> [int, int, int, int]

- Function f(string) -> block<sub>1</sub>, block<sub>2</sub>, block<sub>i</sub>, 
...,
block<sub>n</sub> 

- Hash function h(string) -> seed

- Function f(seed) -> [0-31]

- Mapping index to base32
f([0-31]) -> [2-7A-Z]

- Dot product of two equal-length arrays
dp([a, b]) -> [a<sub>1</sub>*b<sub>1</sub>, a<sub>2</sub>*b<sub>2</sub>, ..., a<sub>n</sub>*b<sub>n</sub>]

- Hash function
h(x,a,w,M) -> (dp(x,a) mod 2<sup>2w</sup>)) div 2<sup>(2w-M)</sup>