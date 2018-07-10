# strkit
strkit is a Go package that provides some tool kit for making up builtin strings.

## Install
To start using strkit, install Go and run `go get`;

```sh
$ go get -u github.com/jinxjinxagain/strkit
```

## Example
iterates through a slice of string

```go
package main
  
import (
        "github.com/jinxjinxagain/strkit"
        "fmt"
)

var AppendXCharFunc = func(x string) string {
        return x + "x"
}

func main() {
        var st = []string{"1", "2", "3", "4", "5"}
        var tt = strkit.SliceMap(st, AppendXCharFunc)
        for index := range st {
                fmt.Println(tt[index])
        }
}

```

output will be:
```
1x
2x
3x
4x
5x
6x
```

## Contact
Be free to mail [jinxjinxagain](mailto:jinxjinxagain1994@gmail.com)

## License

strkit source code is available under the MIT[License](/LICENSE).
