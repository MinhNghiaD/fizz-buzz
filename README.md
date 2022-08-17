# Fizz buzz Server

## About 
Fizz buzz service provides users the REST APIs to run a generic fizz buzz program. The program exposes the following endpoints:

* POST /fizzbuzz endpoint accepts five parameters: three integers int1, int2, limit, and two strings str1 and str2. It returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, and all multiples of int1 and int2 are replaced by str1str2.
* GET /most-frequent endpoint accepts no parameter and returns the parameters corresponding to the most used request, as well as the number of hits for this request.

## Table of contents
> * [Fizz buzz](#fizz-buzz-server)
>   * [About](#about)
>   * [Table of contents](#table-of-contents)
>   * [Building from source](#building-from-source)

## Building from source

The Fizz buzz server source code was written in Golang. In order to build the application locally, make sure you have Golang v1.18 or newer, then run:

```bash
# get the source & build:
$ git clone https://github.com/MinhNghiaD/fizz-buzz.git
$ cd fizzbuzz
$ make build
```

If the build succeeds, the binaries can be found in the following directory: `./bin`.

For containerized usage, we can generate the docker image of the fizz buzz server with command `make docker`.
