Gopherjs Electron 
-------------------------------

This project implements an API translator and some convenience functions for 
creating Electron APP using Gopherjs.

# Usage

    go get github.com/oskca/gopherjs-electron
    go generate # calling the API translator 

# API Translator

The translator reads electron api json file and output the corresponding 
gopherjs struct and functions into files with `raw_` prefix.

The translator lives in directory: `json2rawApi`.

# Easy binding functions

Hand written binding functions/struct lives in `ex_*.go` files and have
`Ex` postfix.
