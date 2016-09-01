# hashwalker

## testing

```
go test ./...
```

## build

```
go build -o hashwalker
```

## usage

```
  -dir string
    	walks the file tree rooted at dir
  -ignore string
    	ignore sub directory or file
  -out string
    	save result to out file
```

## example

``` shell
./hashwalker -dir=$PWD -ignore="out.txt,*_test.go,.git*" -out=out.txt
```

`cat out.txt`

    LICENSE,92170cdc034b2ff819323ff670d3b7266c8bffcd,11357
    README.md,0b6a0cd170aaf4e014b1ab9aa752cfc26706e190,881
    file.go,202469a510ceead49cb36afc81931a7cbfcce055,184
    hash.go,641aeb884057a59d241716c127fe0c31e0f0c708,498
    main.go,957645e724b495cbc7538e8a723bf24a15be28d3,1231
    walker.go,b764ad909b580a1a44fc260c38c3d55e08949f10,982
    write.go,c78429393f4c38f5013ec124dd7eebc6b501f44e,246
    hashwalker,0a69188cfe54441785ffb6dada1dc45efc4f71f8,2751032
