window.BENCHMARK_DATA = {
  "lastUpdate": 1760989907258,
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
          "id": "7be00cda88053c1b127966f3556a1a124f3ceb76",
          "message": "fix: output",
          "timestamp": "2025-10-20T21:46:36+02:00",
          "tree_id": "b59c9ce01d7faafb666c4fc95e46513063c80d16",
          "url": "https://github.com/uponusolutions/go-smtp/commit/7be00cda88053c1b127966f3556a1a124f3ceb76"
        },
        "date": 1760989619273,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking",
            "value": 393111,
            "unit": "ns/op",
            "extra": "2970 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection",
            "value": 132776,
            "unit": "ns/op",
            "extra": "9219 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking",
            "value": 396636,
            "unit": "ns/op",
            "extra": "3091 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection",
            "value": 132871,
            "unit": "ns/op",
            "extra": "8907 times\n4 procs"
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
          "id": "74311b207da15ad3d9334c274abb9f9231d86759",
          "message": "feat: set minimum coverage",
          "timestamp": "2025-10-20T21:51:18+02:00",
          "tree_id": "512e41066913b4b6bc41fb1fd9f87404acc92aa6",
          "url": "https://github.com/uponusolutions/go-smtp/commit/74311b207da15ad3d9334c274abb9f9231d86759"
        },
        "date": 1760989906185,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking",
            "value": 258104,
            "unit": "ns/op",
            "extra": "4446 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection",
            "value": 83088,
            "unit": "ns/op",
            "extra": "14460 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking",
            "value": 260168,
            "unit": "ns/op",
            "extra": "4568 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection",
            "value": 83591,
            "unit": "ns/op",
            "extra": "14332 times\n4 procs"
          }
        ]
      }
    ]
  }
}