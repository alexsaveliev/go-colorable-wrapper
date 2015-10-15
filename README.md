# go-colorable-wrapper

Wrapper over github.com/mattn/go-colorable that provides helper functions and variables

## fmtc
Mimics functions from standard `fmt` package but makes them colorable-aware. Includes
* `fmtc.Print`
* `fmtc.Printf`
* `fmtc.Println`

## streamc
Provides `streamc.Stderr` and `streamc.Stdout` variables. They have the same capabilities as `os.Stderr` and `os.Stdout`, but colorable-aware

## ansicodes
Provides functions to output colored text, change background color, or add some text effect such as bold or underline
