# Lab1

## Generate data

```bash
./testdata/gen.sh
```

## Test

```bash
go test -v ./...
```

## Benchmark

```bash
go test -bench . -args $(find testdata -name '*.json') 
```
