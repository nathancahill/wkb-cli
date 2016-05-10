
package main

import (
    "fmt"
    "os"

    "github.com/davecgh/go-spew/spew"
    "github.com/nathancahill/wkb"
)

func main() {
    path := os.Args[1]
    file, err := os.Open(path)

    scs := spew.ConfigState{MaxDepth: 3, Indent: "\t"}

    if err != nil {
        panic(err)
    }

    out, err := wkb.ReadWKBRaster(file)

    if err != nil {
        panic(err)
    }

    fmt.Print("Headers:\n")

    scs.Dump(out)

    for i, band := range out.Bands {
        r := len(band.Data[0]) - 1
        c := len(band.Data[0]) - 1

        fmt.Printf("\n")
        fmt.Printf("Band %d:\n", i + 1)
        fmt.Printf("%v.....%v\n", band.Data[0][0], band.Data[0][c])
        fmt.Print(".           .\n")
        fmt.Print(".           .\n")
        fmt.Print(".           .\n")
        fmt.Print(".           .\n")
        fmt.Printf("%v.....%v\n", band.Data[r][0], band.Data[r][c])
    }
}
