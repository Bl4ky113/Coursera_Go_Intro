# Getting Started with Golang
# By UCI in Coursera

Start: 09/06/2024
End: 

Hopefully I will finish this one in less than a month.

Sessions:
- 09/06/2024
- 09/07/2024 Two in a row!!!

## Why Golang?

- Fast
- Garbage Collection
- Simple OOP
- Efficient Concurrency
- Compiled

## Golang and "Simple OOP"

We call golang implementation of oop a leaner way since we don't have a lot of stuff and barely rely on structs 
and we associate methods to thoose structs. Meaning that we don't need to always have a system for constructors, 
generics or inheritance.

## Go and it's Efficient Concurrency

First, we can remeber Moore's 'Law': "The num of transistors in processors doubles every 18 months".
But there's a phisical limit to this law, since this big number of transistors require more power and then
by thermodinamics generate way more heat. And since we can just keep the processors away from melting themselfs 
at a certian way. We started to make more 'cores' or mini processors so we can execute multiple things at the same 
time in a single joined processor. Now days there's processors with over 32 cores and double the threads.
But there's a catch, we need to make our programs use theese new cores, threads and more. 

Golang and it's aproach to concurrency and parallelism make our development with theese features more easy.
We will talk in more detail about them in a next course.

## Download Golang

Just install it from pacman. Duh. Also set some ENV\_VAR for some interpreters in NVIM.

## Golang workspaces and packages

It's the way that Go organizes it's projects in a standarized way.
The workspaces are mainly made by:
/src/ -> source files
/pkg/ -> used packages
/bin/ -> program binaries

Altough this is not enforced, just recomended.
Also, a workspace might be used for multiple projects but if you need, declare one just for you project by using a GOPATH ENV\_VAR.

Packages can be downloaded or std libraries. Internal packages from source or anything in general. 
But it should always be defined in the begining of the file. 

Also for any project, there must be a main package with a main function where it start it execution.

## Go tool

Is the cli for go. That we can use to compile, make docs for the program, format the source files, install new pacakges from online and list them,
and run or test the program 'without' compilation.

## Comments in Golang

The simple C and C++ comments
// foo
/\* bar \*/

## Variables

Golang has 3 ways to define a variable:

Simple declaration, supports multiple assignment:
var foo type = val
var foo, foo2, foo3, ..., fooN type

Type on declaration:
var foo = 100
This might have some problemns with some types that look like they are the same but might become a problemn if they are mismanaged.

Asignment and declaration:
foo := 100
This is a simple way to make variables. But it's only available inside functions and their scope, so be careful.

It doesn't matter what type any variable has, it should always have a default value, 
for example numbers, ints or floats is 0, strings an empty string ("") and such.

If you don't use a declared variable in your code, Go and it's compiler will take it as an error. 
So watch out for some not-taken trash in the code.

### Var Scope 

The scope is defined by blocks, generally the {} in the code. 
There's some more types of implicit blocks like the global or universe block, which is just like a normal global context in the source.
There's a implicit block for the scope for a package and file.
Also, conditional logic blocks can also have their own scope, so not only functions and others structures have them.

The scope works like it generally does, so no big trouble.

## Types

Some basic types are:

### Pointers:
Address to something in memory. 
Can use the reference and dereference (&, \*).
We can use new() so we can create a pointer to a value type:

foo := new(int) // foo is a pointer to a int
\*foo = 3

### Integers:
There's a lot of them but mainly varies on the size and type (signed or unsigned)
- int
- int8
- int16
- int32
- int64
- uint8
- uint16
- uint32
- uint64

### Floats / Doubles
Same stuff as ints
- float32
- float64

You can use scientific notation:
var x float32 = 123.123
var y float32 = 1.234e2

### Complex numbers
Made up from a two floats, one is the real part and the other the imaginary:
var cplx complex128 = complex(123.1, 456.123)

### Strings 

Golang uses UTF-8 for it's strings. Each char in utf-8 has a respective Rune 
that is just the number in binary or hex of that char in utf-8.

Generally declared only by simple assignment using double quotes:
foo = "hello"

They are READ only, and mainly ment to show and print to the user.
There's some packages in the stdlib that will help us have more usage of strings.

## Type declaration

With the basic or madeup types we can create alises for them, for example:

type Celsius float64
type IDnumber uint32

Then, we can use them as a normal type in our code.

## Type casting

Geerally with just the type and using it like a function, passing the variable to cast into is more than enough to 
cast a variable into a equivalent type, for example the ints and their variations:

var foo int32 = 10
var bar int64 = 64;

bar = int64(foo)

## Contants

Are just normal constants in general, there's a big change when defining them.
We can doit for just one or multiple and we can use a constant auto increment.
In a single or multiple assignment we don't have to declare the type, just like type on declaration.

const foo = 123
const (
    foo = 123
    foo2 = 321
    foo3 = 456
    foo4 = 789
    ...
    fooN = 999
)

But with iota we can assign mutliple constants with a constant increment, for example the days of the week.
It will just increment in 1 by each constant defined, which we can use to our advantage since we can use it in operations.
But if we want to skip one constant using iota we can define a constant with \_, to 'skip' and forget about it
there's a simple increment example, and a more advance usage taken from the effective go guide:

const (
    A = iota
    B
    C
    D
)

const (
    _           = iota // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)

## Operators

The general ones:
- Arithmetic: +, -, \*, /, %, \<\<, \>\>
- Comparison: ==, !=, \< (\<=), \> (\>=)
- Boolean: !, &&, ||

## Garbage Collector in Golang

Generally a GC works in a stack for small var scopes like functions which is just deallocated after.
But for a scope like a global, package or file they are stored in a heap. Which we generally have to manualy deallocate in most programming languages.
This can be a headache since if we allocate some memory and then keep using it in other scopes we might have some trouble remebering 
when to deallocate that memory. So there comes the GC, so it detects when a reference isn't beeing used, it dereferences them.
The thing is, GC thends to be a very slow thing and mainly available on interpreted languages like Java with JVM or Python.

But in Golang, the compiler has a buildin GC, not as fast as the languages without one but way way faster as the general GC.

## Control flow

The general ones, but we dont use () in the contiditonals, we could use them to group them but it's not necesary

If flow is just normal, but we can assign and create variables before the condition, which later on will be useful.
if \<assignment\>; \<condition\> {
    ...
}

For loops can be used in multiple ways.
We can use the break and continue keywords just like any other language.
- standard one:
for \<init\>; \<cond\>; \<update\> {
    ...
}

- while:
for \<cond\> {
    ...
}

- while True:
for {
    ...
}

Switch and case look kinda like a normal one, but it doesn't need break points necesaraly:
swtich foo {
case 1:
    ...
case N:
    ...
default:
    ...
}

There's also a special Switch which doesn't need a initial value to compare to other cases, but kinda works 
as checking which case condition is true:
switch {
case x == 1:
    ...
case x < 1:
    ...
case x > 1:
    ...
default:
    ...
}

## stdlib: ftm

The format package has a good variarity of formating varaibles and mainly printing them to stdout. 
The format is just like the stdio from C. But there's some fancy stuff that we can do since Strings in Go are different than the strings in C.

Some output and format functions are:
- fmt.Printf(formated\_str, args...): Equivalent of printf()
- fmt.Println
- fmt.Format

Scan is a function that reads user input in the stdin
It takes a pointer where the value recieved will be stored,
after geting the data, the pointer will be rewritten and will return the number of 
values recieved, which could be more than 1 btw.

## stdlib: unicode

Is a package for strings, mainly working for the runes. So we can handle, see and use the categories of them, mainly linguistical ones like:
- letter
- number
- symbol
- punctiation
- space
and such

## stdlib: string

Is a package for stings that allow us to use two strings to see different stuff like:
- compare
- contains
- see prefix
- index of a substring in a string

This one also let us 'modify' out strings,
which works by taking that string and then returns the modified copy of the string.
And the functions are just like the orthers, but there's some single string functions
- replace
- to upper
- to lower
- trim spaces

## stdlib: strconv

Is a package so we can convert some values from strings to other kinds of types.
Some of it's functions names are weird, for example:
- Atoi: string to int
- Itoa: int to string

And there's 
- ParseFloat: string to float
- FormatFloat: float to string

