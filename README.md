## Gremory
This small library provides common collections like a list, ordered set and ordered dictionary(map). 

## Installation
Gremory does not use any third party libraries. For getting it run on your machine, you just:

```
go get github.com/SealNTibbers/gremory
```
## API and Examples
Every structure holds data wrapped in ValueHolders. This solution is used to operate with untyped data. We assume to use ValueHolder for basic data types like a int, string, bool and float. For user data types one needs to implement CollectionObject interface.
