# hashwalker

## testing

go test ./...

## build

go build -o hashwalker

## usage

```
  -dir string
    	walks the file tree rooted at dir
  -ignore string
    	ignore sub folder or file
  -out string
    	save result to out file
```

## example

``` shell
./hashwalker -dir=$PWD -ignore="file.go,*_test.go" -out=out.txt
```

`cat out.txt`

    ~/hashwalker/LICENSE,92170cdc034b2ff819323ff670d3b7266c8bffcd,11357
    ~/hashwalker/README.md,fb784341f675ed0ee1f364c5e7c5365b7984539e,12
    ~/hashwalker/hash.go,641aeb884057a59d241716c127fe0c31e0f0c708,498
    ~/hashwalker/main.go,6423fb66c1b21226f6f2f4522d3e25ab1e43c969,1228
    ~/hashwalker/out.txt,286365e81cd4e3f16c83bbf3d8cac6b0a355cc36,0
    ~/hashwalker/walker.go,87c580002126d9df6fa1e64e667ded2c3a3170b8,775
    ~/hashwalker/writer.go,295d7b8d7522e14ebc7b48bd0cc75ae61aeaf86e,368
    ~/hashwalker/hashwalker,76b0b0595d789bd3694134bf605ff783ac9d3d63,2745616
