window.BENCHMARK_DATA = {
  "lastUpdate": 1788778048660,
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
      },
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
          "id": "f02688e05e0e7f067059c1372c15446d9163f620",
          "message": "feat: remove redundant upper case and improve starttls handling",
          "timestamp": "2026-09-07T11:59:15+02:00",
          "tree_id": "41ca33b3e65f218d7c64bfda9ca7dec9b28a4136",
          "url": "https://github.com/uponusolutions/go-smtp/commit/f02688e05e0e7f067059c1372c15446d9163f620"
        },
        "date": 1788775459660,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 344465,
            "unit": "ns/op",
            "extra": "34398 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 418903,
            "unit": "ns/op",
            "extra": "28329 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 87058,
            "unit": "ns/op",
            "extra": "137946 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 85712,
            "unit": "ns/op",
            "extra": "140560 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 356625,
            "unit": "ns/op",
            "extra": "33630 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 357402,
            "unit": "ns/op",
            "extra": "33338 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 108245,
            "unit": "ns/op",
            "extra": "114649 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 113846,
            "unit": "ns/op",
            "extra": "106302 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 16382346,
            "unit": "ns/op",
            "extra": "738 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 16917868,
            "unit": "ns/op",
            "extra": "688 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 15098475,
            "unit": "ns/op",
            "extra": "781 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 15324049,
            "unit": "ns/op",
            "extra": "794 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 24128177,
            "unit": "ns/op",
            "extra": "499 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 24262691,
            "unit": "ns/op",
            "extra": "502 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 25941927,
            "unit": "ns/op",
            "extra": "410 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 27091355,
            "unit": "ns/op",
            "extra": "411 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 22978291,
            "unit": "ns/op",
            "extra": "528 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1083917,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 23489616,
            "unit": "ns/op",
            "extra": "522 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1076306,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1049173889,
            "unit": "ns/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 45264369,
            "unit": "ns/op",
            "extra": "266 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1054297169,
            "unit": "ns/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 54800909,
            "unit": "ns/op",
            "extra": "217 times\n4 procs"
          }
        ]
      },
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
          "id": "e7a5913b16d0a860827e07e1fbf096b6a345ae44",
          "message": "fix: Args parsing now splits only by whitespace and returns an array\ninstead of a map, reverse path removes leading whitespaces from S,\nimproved error handling of Args calls",
          "timestamp": "2026-09-07T12:42:22+02:00",
          "tree_id": "c44283a9e7065ec6661723f62acaa5793b1b455f",
          "url": "https://github.com/uponusolutions/go-smtp/commit/e7a5913b16d0a860827e07e1fbf096b6a345ae44"
        },
        "date": 1788778048086,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 264228,
            "unit": "ns/op",
            "extra": "45412 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 325674,
            "unit": "ns/op",
            "extra": "36915 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 75667,
            "unit": "ns/op",
            "extra": "155265 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 75350,
            "unit": "ns/op",
            "extra": "158689 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 281649,
            "unit": "ns/op",
            "extra": "42723 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 283576,
            "unit": "ns/op",
            "extra": "42296 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 94089,
            "unit": "ns/op",
            "extra": "127352 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 95455,
            "unit": "ns/op",
            "extra": "125744 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 14197231,
            "unit": "ns/op",
            "extra": "817 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 15391279,
            "unit": "ns/op",
            "extra": "772 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 13671575,
            "unit": "ns/op",
            "extra": "870 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 13165309,
            "unit": "ns/op",
            "extra": "892 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 21407760,
            "unit": "ns/op",
            "extra": "544 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 21562695,
            "unit": "ns/op",
            "extra": "553 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 23144853,
            "unit": "ns/op",
            "extra": "454 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 26087052,
            "unit": "ns/op",
            "extra": "493 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 18811650,
            "unit": "ns/op",
            "extra": "636 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 926396,
            "unit": "ns/op",
            "extra": "12946 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 19469944,
            "unit": "ns/op",
            "extra": "626 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 925553,
            "unit": "ns/op",
            "extra": "12963 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 897794906,
            "unit": "ns/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 38811415,
            "unit": "ns/op",
            "extra": "308 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 860109365,
            "unit": "ns/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 50988882,
            "unit": "ns/op",
            "extra": "235 times\n4 procs"
          }
        ]
      }
    ]
  }
}