window.BENCHMARK_DATA = {
  "lastUpdate": 1760988129573,
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
          "id": "4df5c6bb1e42d477d37ab0ab7755ae7f59e8e0a3",
          "message": "fix: benchmark",
          "timestamp": "2025-10-20T21:21:36+02:00",
          "tree_id": "612a5ed262da13b1e66f9a6813b599cb6d2f1e6f",
          "url": "https://github.com/uponusolutions/go-smtp/commit/4df5c6bb1e42d477d37ab0ab7755ae7f59e8e0a3"
        },
        "date": 1760988129067,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking",
            "value": 429066,
            "unit": "ns/op\t   8.24 MB/s",
            "extra": "2779 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunking - ns/op",
            "value": 429066,
            "unit": "ns/op",
            "extra": "2779 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunking - MB/s",
            "value": 8.24,
            "unit": "MB/s",
            "extra": "2779 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection",
            "value": 141489,
            "unit": "ns/op\t  24.98 MB/s",
            "extra": "8851 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection - ns/op",
            "value": 141489,
            "unit": "ns/op",
            "extra": "8851 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection - MB/s",
            "value": 24.98,
            "unit": "MB/s",
            "extra": "8851 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking",
            "value": 390043,
            "unit": "ns/op\t   9.06 MB/s",
            "extra": "2764 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking - ns/op",
            "value": 390043,
            "unit": "ns/op",
            "extra": "2764 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking - MB/s",
            "value": 9.06,
            "unit": "MB/s",
            "extra": "2764 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection",
            "value": 132334,
            "unit": "ns/op\t  26.71 MB/s",
            "extra": "8404 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection - ns/op",
            "value": 132334,
            "unit": "ns/op",
            "extra": "8404 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection - MB/s",
            "value": 26.71,
            "unit": "MB/s",
            "extra": "8404 times\n4 procs"
          }
        ]
      }
    ]
  }
}