## Gremory
This small library provides common collections like a list, ordered set and ordered dictionary(map). 

## Installation
Gremory does not use any third party libraries. For getting it run on your machine, you just:

```
go get github.com/SealNTibbers/gremory
```
## API and Examples
Every structure holds data wrapped in ValueHolders. This solution is used to operate with untyped data. We assume to use ValueHolder for basic data types like a int, string, bool and float. For user data types one needs to implement CollectionObject interface. For working with all containers you have basic smalltalk-like functions: Do, ReverseDo, Select, Collect. 

### Data container creation
To create a list you can use next code:
```
  valueGen := func(value interface{}) CollectionObject {
		return &ValueHolder{value}
	}
	list := NewSmartList(valueGen)
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
```

To create a dict you can use next code:
```
  keyGen := func(value interface{}) CollectionObject {
		return &ValueHolder{value}
	}
	valueGen := func(value interface{}) CollectionObject {
		return &ValueHolder{value}
	}
	dict := NewSmartODict(keyGen, valueGen)
	dict.AddPair(1, 'a')
	dict.AddPair(2, 'b')
	dict.AddPair(3, 'c')
```

To create a set you can use next code:
```
  valueGen := func(value interface{}) CollectionObject {
		return &ValueHolder{value}
	}
	set := NewSmartOSet(valueGen)
	set.AddValue(1)
	set.AddValue(2)
	set.AddValue(3)
```


### Functions examples
The next examples will be use List for presentation of functions because all containers have familiar way.
#### Do
This function is using for evaluate lambda with each of the elements as the argument:
```
  list.Do(func(each *ListNode) {
		fmt.Println(each.GetValue())
	})
```

#### ReverseDo
This function is used to evaluate lambda with each of the elements as the argument, strating with the last element and taking each in sequence up to the first.
```
  list.ReverseDo(func(each *ListNode) {
		fmt.Println(each.GetValue())
	})
```

#### Select
This function is used to evaluate lambda with each of the elements as the argument. Collect into new collection like a main container, only those element for which lambda evaluates true.
```
selectList := list.Select(func(each *ListNode) bool {
		if each.GetValue().(int) > 1 {
			return true
		}
		return false
	})
// If List is (1,2,3) when result should be (2,3)
```

#### Collect
This function is used to evaluate with each of the elements as the argument. Collect the result values into new collection like a  main container.
```
collectList := list.Collect(func(each *ListNode) CollectionObject {
		return &ValueHolder{each.GetValue().(int) * 10}
	})
// If List is (1,2,3) when result should be (10,20,30)
 
```
