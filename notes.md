# Getting Started with Golang
# By UCI in Coursera

Start: 09/06/2024
End: 

Hopefully I will finish this one in less than a month.

Sessions:
- 09/06/2024

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

## Variables

Golang has 3 ways to define a variable:

Simple declaration, supports multiple asignment:
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

## Types


## Type declaration

With the basic or madeup types we can create alises for them, for example:

type Celsius float64
type IDnumber uint32

Then, we can use them as a normal type in our code.
