# Lab1

## Generate data

```bash
./testdata/gen.sh
```

## Benchmark

```bash
go test -bench . -args $(find testdata -name '*.json') 
```