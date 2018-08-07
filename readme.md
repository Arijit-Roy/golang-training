There are 2 types of packages. 
1. executable: 
  1. must import a package called main. 
  2. must contain a function called main, when compiled creates an executable 
2. reusable
  1. must not mport a package called main. 
  2. when compiled does not creates an executable 

Import Statements :
  Import command is usd to gain access to all functionalities of a reusable package
  https://golang.org/pkg/fmt

main.go organization:
  package declaration 
  import other packages
  declare functions
  := is used for declaration and assignment(both in the same statement)
  := cant be used for assignment

Files in same package can freely call functions defined in other files. 

Array: fixed length of things.
Slice: an array that can grow or shrink
They Must contain data of identical types.
Slice declaration:
  sliceName := []datatype{elem1, elem2, elem3}
append to slice:
  append(sliceName, element)
  (append creates a new slice)
iterating through slices:
  for curr_iteration, element := range slicename{
    fmt.Println(curr_iteration, element)
  }

  slice[startIndexIncluding : uptoNotIncluding]
  startIndexIncluding defaults to the beginning of the slice
  uptoNotIncluding defaults to end of the slice(including)

byte slice is ascci representation of a string. (asciitable.com)

io/ioutil package : for file I/o
os package : platform indepenedent os handling
strings package: for string manipulation 
https://golang.org/pkg/
all package methods are camelcased

// accpting multiple return from function
bs, err := ioutil.ReadFile(fileName)

// typecasting
datatype(variable) -- string(bs) converts a byteslice to string

func (personPointer *person) updateName(newFirstName string) 
  Here *person is just a pointer type. personPointer is the address of the the variable


go routons are created using go function name
channels are way to communicate b/w go routines

sleep pauses current goroutine for at least duration

do not use variable of another subroutine. always use a function literal to pass value to different subroutines