window.BENCHMARK_DATA = {
  "lastUpdate": 1788771800668,
  "repoUrl": "https://github.com/uponusolutions/go-smtp",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "committer": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "distinct": true,
          "id": "21f6161974c03d3aa8e92cec0d1f4dcfbec83487",
          "message": "Merge branch 'emersion-master'",
          "timestamp": "2026-09-07T00:44:47+02:00",
          "tree_id": "c6d5dd74eb1458edf471a01d02dcfe045ffa5370",
          "url": "https://github.com/uponusolutions/go-smtp/commit/21f6161974c03d3aa8e92cec0d1f4dcfbec83487"
        },
        "date": 1788771800005,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 342083,
            "unit": "ns/op",
            "extra": "35332 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 426871,
            "unit": "ns/op",
            "extra": "28236 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 99083,
            "unit": "ns/op",
            "extra": "120146 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 98425,
            "unit": "ns/op",
            "extra": "120728 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 364826,
            "unit": "ns/op",
            "extra": "32835 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 367681,
            "unit": "ns/op",
            "extra": "32612 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 122807,
            "unit": "ns/op",
            "extra": "97845 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 124640,
            "unit": "ns/op",
            "extra": "95845 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 18408269,
            "unit": "ns/op",
            "extra": "661 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 18060078,
            "unit": "ns/op",
            "extra": "661 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 17042112,
            "unit": "ns/op",
            "extra": "703 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 16189252,
            "unit": "ns/op",
            "extra": "738 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 28085106,
            "unit": "ns/op",
            "extra": "417 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 28270390,
            "unit": "ns/op",
            "extra": "403 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 805120063,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 30188968,
            "unit": "ns/op",
            "extra": "344 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 24300774,
            "unit": "ns/op",
            "extra": "493 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1196696,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 24303661,
            "unit": "ns/op",
            "extra": "493 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1195483,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1158970317,
            "unit": "ns/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 49919993,
            "unit": "ns/op",
            "extra": "240 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1172480372,
            "unit": "ns/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 62017582,
            "unit": "ns/op",
            "extra": "193 times\n4 procs"
          }
        ]
      }
    ]
  }
}