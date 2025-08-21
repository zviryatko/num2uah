# num2uah

A Go library for converting numbers to Ukrainian currency representation (hryvnias and kopiyok) with proper grammatical forms.

## Installation

```bash
go get github.com/zviryatko/num2uah
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/zviryatko/num2uah"
)

func main() {
    // Convert float64 to Ukrainian currency text
    result := num2uah.Convert(1234.56)
    fmt.Println(result)
    // Output: одна тисяча двісті тридцять чотири гривні та п'ятдесят шість копійок
    
    // More examples
    fmt.Println(num2uah.Convert(1.01))
    // Output: одна гривня та одна копійка
    
    fmt.Println(num2uah.Convert(2222.22))
    // Output: дві тисячі двісті двадцять дві гривні та двадцять дві копійки
}
```

## License

This project is available under the MIT License.
