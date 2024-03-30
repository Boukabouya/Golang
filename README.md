# Golang

golang tutorial
1 - in golang everythig is modular.
gonalg basics :

- complied  language and go tool run file directly, without VM
- Executable are different for OS

## what and where i can use it

system apps to web apps -Cloud
Already in production

## Don't bring baggage

- The mindset of this programming language is different  from other languages.

- but it similare to c and java and pascal

## Object Oriented

- Yes and No
what you see in the screen is the code

## missing

- it really missing but not lagging
- No try catch
- lexer does a lot of work

by installing go intellisense we don't need to write import we just write function directly

## Lexer

- In Go, semicolons are automatically removed in most cases but it's not the valide syntax , but there are some exceptions where we put it like loop.
- After installing the go intellisense The lexer in programming languages that ensures correct syntax and variable definitions and automatically inserting semicolons in code.

## Types

- case insensitive : capital for the first letter -> exported and public
small for the first letter -> private

- Variables in golang should be known in advance and every thing is type in glonag .

### simple types

string integers float bool complex

### Advenced types

Arrat Slices Maps Structs Pointers functions channels ... almost everything

## Build for windows , linux and mac

- we just type

```bash
go env
```

go env is a tool to view and temporarily modify Go environment variables that configure aspects of the Go build process and from it we could see GOOS that mean the os that we go build and we could change it by add GOOS before the go build cli

```bash
# GOOS="windows" ND GOOS="linux" 
GOOS="darwin" go build # for macos
```
