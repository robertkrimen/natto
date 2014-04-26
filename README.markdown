# natto
--
    import "github.com/robertkrimen/natto"

Package natto is an example/offshoot of otto that implements an event loop
(supporting setTimeout/setInterval).

http://godoc.org/github.com/robertkrimen/natto

otto: http://github.com/robertkrimen/otto
(http://godoc.org/github.com/robertkrimen/otto)

## Usage

#### func  Run

```go
func Run(src string) error
```
Run will execute the given JavaScript, continuing to run until all timers have
finished executing (if any). The VM has the following functions available:

    <timer> = setTimeout(<function>, <delay>, [<arguments...>])
    <timer> = setInterval(<function>, <delay>, [<arguments...>])
    clearTimeout(<timer>)
    clearInterval(<timer>)

--
**godocdown** http://github.com/robertkrimen/godocdown
