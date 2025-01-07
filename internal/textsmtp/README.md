# Golang's net/textproto but 𝙁/𝘼/𝙎/𝙏/𝙀/𝙍

## `import "gopkg.in/textproto.v0"`

![Go Import](https://img.shields.io/badge/import-gopkg.in/textproto.v0-9cf?logo=go&style=for-the-badge)
[![Go Reference](https://img.shields.io/badge/reference-go.dev-007d9c?logo=go&style=for-the-badge)](https://pkg.go.dev/gopkg.in/textproto.v0)

The Golang standard library `net/textproto` suffers in data throughput performance when using its `DotReader` or `DotWriter` implementations. It's caused by iterating each input output byte using the `ReadByte` or `WriteByte` functions. This increases a lot of CPU cycles for accessing memory for each byte. This module aim to provide an optimized implementation by leveraging the SIMD optimized `ReadSlice` and `IndexByte` functions. Here are the resulting numbers:

```
BenchmarkDotReader/Legacy-10                  72          15078444 ns/op         278.17 MB/s
BenchmarkDotReader/Optimized-10             1752            653339 ns/op        6419.79 MB/s
BenchmarkDotWriter/Legacy-10                 100          10760824 ns/op         389.78 MB/s
BenchmarkDotWriter/Optimized-10             2116            553493 ns/op        7577.88 MB/s
```
