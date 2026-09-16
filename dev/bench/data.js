window.BENCHMARK_DATA = {
  "lastUpdate": 1789522612123,
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
          "id": "4fdbefaecd774f033b23d571106f300cfd55b336",
          "message": "chore: makefile",
          "timestamp": "2026-09-14T02:04:18+02:00",
          "tree_id": "07cf44b8ce474df1332e2d39b15050c04b4342c6",
          "url": "https://github.com/uponusolutions/go-smtp/commit/4fdbefaecd774f033b23d571106f300cfd55b336"
        },
        "date": 1789345004200,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 383320,
            "unit": "ns/op",
            "extra": "31824 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 380523,
            "unit": "ns/op",
            "extra": "31238 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 116690,
            "unit": "ns/op",
            "extra": "101851 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 114332,
            "unit": "ns/op",
            "extra": "102392 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 324698,
            "unit": "ns/op",
            "extra": "36855 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 319829,
            "unit": "ns/op",
            "extra": "37396 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 60049,
            "unit": "ns/op",
            "extra": "199702 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 57466,
            "unit": "ns/op",
            "extra": "207488 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 456972,
            "unit": "ns/op",
            "extra": "26059 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 355485,
            "unit": "ns/op",
            "extra": "33754 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 91298,
            "unit": "ns/op",
            "extra": "131799 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 93498,
            "unit": "ns/op",
            "extra": "129015 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 445768,
            "unit": "ns/op",
            "extra": "26097 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 320650,
            "unit": "ns/op",
            "extra": "37389 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 54281,
            "unit": "ns/op",
            "extra": "221016 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 53624,
            "unit": "ns/op",
            "extra": "222066 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 26408832,
            "unit": "ns/op",
            "extra": "468 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 25592055,
            "unit": "ns/op",
            "extra": "476 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 27200765,
            "unit": "ns/op",
            "extra": "388 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 206531742,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 26458258,
            "unit": "ns/op",
            "extra": "476 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 27764429,
            "unit": "ns/op",
            "extra": "436 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 26969124,
            "unit": "ns/op",
            "extra": "458 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 1494943011,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 17589716,
            "unit": "ns/op",
            "extra": "687 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 16639216,
            "unit": "ns/op",
            "extra": "711 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 15680504,
            "unit": "ns/op",
            "extra": "778 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 16417840,
            "unit": "ns/op",
            "extra": "727 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 17714111,
            "unit": "ns/op",
            "extra": "678 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 16612237,
            "unit": "ns/op",
            "extra": "741 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 15134811,
            "unit": "ns/op",
            "extra": "796 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 15695657,
            "unit": "ns/op",
            "extra": "766 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 21546438,
            "unit": "ns/op",
            "extra": "550 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1146299,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 21553090,
            "unit": "ns/op",
            "extra": "555 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1143929,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 21435785,
            "unit": "ns/op",
            "extra": "558 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 957130,
            "unit": "ns/op",
            "extra": "12552 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 21487931,
            "unit": "ns/op",
            "extra": "554 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 972484,
            "unit": "ns/op",
            "extra": "12337 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 16073801,
            "unit": "ns/op",
            "extra": "745 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 741721,
            "unit": "ns/op",
            "extra": "16159 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 16248707,
            "unit": "ns/op",
            "extra": "735 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 833724,
            "unit": "ns/op",
            "extra": "14398 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 16272838,
            "unit": "ns/op",
            "extra": "736 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 942332,
            "unit": "ns/op",
            "extra": "12733 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 16663419,
            "unit": "ns/op",
            "extra": "728 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1039613,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
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
          "id": "438dbc3c87efc2421f9f658de32427c49dfd7e80",
          "message": "chore: improve readme",
          "timestamp": "2026-09-14T02:10:47+02:00",
          "tree_id": "1ee4fca12677effb3e4ffc30a50b45f0e8f1bbdd",
          "url": "https://github.com/uponusolutions/go-smtp/commit/438dbc3c87efc2421f9f658de32427c49dfd7e80"
        },
        "date": 1789345257353,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 159542,
            "unit": "ns/op",
            "extra": "74038 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 159867,
            "unit": "ns/op",
            "extra": "75373 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 42271,
            "unit": "ns/op",
            "extra": "275859 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 39562,
            "unit": "ns/op",
            "extra": "307785 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 146229,
            "unit": "ns/op",
            "extra": "71840 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 138387,
            "unit": "ns/op",
            "extra": "86449 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 25416,
            "unit": "ns/op",
            "extra": "429020 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 22190,
            "unit": "ns/op",
            "extra": "543429 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 295086,
            "unit": "ns/op",
            "extra": "40166 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 155144,
            "unit": "ns/op",
            "extra": "75374 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 31179,
            "unit": "ns/op",
            "extra": "380026 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 31821,
            "unit": "ns/op",
            "extra": "383053 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 286863,
            "unit": "ns/op",
            "extra": "41364 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 135884,
            "unit": "ns/op",
            "extra": "87974 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 20147,
            "unit": "ns/op",
            "extra": "578724 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 20398,
            "unit": "ns/op",
            "extra": "591494 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 18646482,
            "unit": "ns/op",
            "extra": "646 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 18332080,
            "unit": "ns/op",
            "extra": "656 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 144042432,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 19571958,
            "unit": "ns/op",
            "extra": "571 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 18043529,
            "unit": "ns/op",
            "extra": "660 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 17742040,
            "unit": "ns/op",
            "extra": "678 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 20637971,
            "unit": "ns/op",
            "extra": "504 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 147844786,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 15887057,
            "unit": "ns/op",
            "extra": "770 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 14404053,
            "unit": "ns/op",
            "extra": "837 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 14108902,
            "unit": "ns/op",
            "extra": "868 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 14135821,
            "unit": "ns/op",
            "extra": "838 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 16127653,
            "unit": "ns/op",
            "extra": "740 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 14802281,
            "unit": "ns/op",
            "extra": "829 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 14667593,
            "unit": "ns/op",
            "extra": "829 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 14860601,
            "unit": "ns/op",
            "extra": "765 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 10930934,
            "unit": "ns/op",
            "extra": "1124 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 937828,
            "unit": "ns/op",
            "extra": "12738 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 10909445,
            "unit": "ns/op",
            "extra": "1058 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 929927,
            "unit": "ns/op",
            "extra": "12970 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 11167950,
            "unit": "ns/op",
            "extra": "1105 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 726424,
            "unit": "ns/op",
            "extra": "16292 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 11033638,
            "unit": "ns/op",
            "extra": "1104 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 719961,
            "unit": "ns/op",
            "extra": "16447 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 8921569,
            "unit": "ns/op",
            "extra": "1354 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 501524,
            "unit": "ns/op",
            "extra": "24496 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 9115261,
            "unit": "ns/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 619062,
            "unit": "ns/op",
            "extra": "19621 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 9049053,
            "unit": "ns/op",
            "extra": "1321 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 539806,
            "unit": "ns/op",
            "extra": "22396 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 9418427,
            "unit": "ns/op",
            "extra": "1274 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 639104,
            "unit": "ns/op",
            "extra": "18794 times\n4 procs"
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
          "id": "d35e9096d9469cefcc535e49f05a16d4a1b6496a",
          "message": "fix: dotreader fragmented",
          "timestamp": "2026-09-14T11:33:34+02:00",
          "tree_id": "048edcf14b37f4cbd4f07c0624f0325459ab5497",
          "url": "https://github.com/uponusolutions/go-smtp/commit/d35e9096d9469cefcc535e49f05a16d4a1b6496a"
        },
        "date": 1789379010571,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 390708,
            "unit": "ns/op",
            "extra": "29798 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 381887,
            "unit": "ns/op",
            "extra": "31374 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 119082,
            "unit": "ns/op",
            "extra": "100741 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 115379,
            "unit": "ns/op",
            "extra": "103681 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 328622,
            "unit": "ns/op",
            "extra": "36415 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 324774,
            "unit": "ns/op",
            "extra": "36889 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 63485,
            "unit": "ns/op",
            "extra": "189381 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 59224,
            "unit": "ns/op",
            "extra": "201436 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 503799,
            "unit": "ns/op",
            "extra": "23991 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 362151,
            "unit": "ns/op",
            "extra": "33062 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 93955,
            "unit": "ns/op",
            "extra": "127568 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 94359,
            "unit": "ns/op",
            "extra": "127221 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 473095,
            "unit": "ns/op",
            "extra": "25178 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 323757,
            "unit": "ns/op",
            "extra": "37059 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 55391,
            "unit": "ns/op",
            "extra": "213926 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 55474,
            "unit": "ns/op",
            "extra": "216681 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 26820859,
            "unit": "ns/op",
            "extra": "463 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 26673588,
            "unit": "ns/op",
            "extra": "454 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 30759857,
            "unit": "ns/op",
            "extra": "379 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 28990396,
            "unit": "ns/op",
            "extra": "378 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 26606918,
            "unit": "ns/op",
            "extra": "459 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 25707158,
            "unit": "ns/op",
            "extra": "436 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 26860446,
            "unit": "ns/op",
            "extra": "447 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 28716754,
            "unit": "ns/op",
            "extra": "364 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 17949915,
            "unit": "ns/op",
            "extra": "674 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 17127362,
            "unit": "ns/op",
            "extra": "696 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 15469495,
            "unit": "ns/op",
            "extra": "782 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 15962504,
            "unit": "ns/op",
            "extra": "759 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 18013776,
            "unit": "ns/op",
            "extra": "663 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 17038408,
            "unit": "ns/op",
            "extra": "708 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 14901137,
            "unit": "ns/op",
            "extra": "776 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 15696786,
            "unit": "ns/op",
            "extra": "765 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 22753026,
            "unit": "ns/op",
            "extra": "526 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1138145,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 22690966,
            "unit": "ns/op",
            "extra": "526 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1138801,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 23004936,
            "unit": "ns/op",
            "extra": "529 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 936424,
            "unit": "ns/op",
            "extra": "12854 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 22631497,
            "unit": "ns/op",
            "extra": "528 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 905924,
            "unit": "ns/op",
            "extra": "13063 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 16115964,
            "unit": "ns/op",
            "extra": "745 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 741711,
            "unit": "ns/op",
            "extra": "16168 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 16309593,
            "unit": "ns/op",
            "extra": "735 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 831794,
            "unit": "ns/op",
            "extra": "14442 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 16302815,
            "unit": "ns/op",
            "extra": "734 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 939125,
            "unit": "ns/op",
            "extra": "12774 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 16616480,
            "unit": "ns/op",
            "extra": "710 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1038220,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
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
          "id": "4fcc5e78d8d9a2c9184e21bf78a2a3a5761b2b5d",
          "message": "fix: improve dotreader, revert some changes",
          "timestamp": "2026-09-14T12:22:06+02:00",
          "tree_id": "23e93e02a0aebb9cf803201174d9015c55be7ec9",
          "url": "https://github.com/uponusolutions/go-smtp/commit/4fcc5e78d8d9a2c9184e21bf78a2a3a5761b2b5d"
        },
        "date": 1789381918695,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 386501,
            "unit": "ns/op",
            "extra": "31081 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 385028,
            "unit": "ns/op",
            "extra": "31210 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 118812,
            "unit": "ns/op",
            "extra": "100112 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 114203,
            "unit": "ns/op",
            "extra": "103677 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 325700,
            "unit": "ns/op",
            "extra": "37008 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 322841,
            "unit": "ns/op",
            "extra": "37176 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 61787,
            "unit": "ns/op",
            "extra": "189613 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 57889,
            "unit": "ns/op",
            "extra": "205231 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 474086,
            "unit": "ns/op",
            "extra": "25376 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 359547,
            "unit": "ns/op",
            "extra": "33420 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 92523,
            "unit": "ns/op",
            "extra": "128544 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 92607,
            "unit": "ns/op",
            "extra": "127584 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 439319,
            "unit": "ns/op",
            "extra": "27636 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 325608,
            "unit": "ns/op",
            "extra": "36836 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 54986,
            "unit": "ns/op",
            "extra": "216813 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 55581,
            "unit": "ns/op",
            "extra": "215348 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 27227845,
            "unit": "ns/op",
            "extra": "450 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 27658968,
            "unit": "ns/op",
            "extra": "416 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 30301741,
            "unit": "ns/op",
            "extra": "331 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 27526837,
            "unit": "ns/op",
            "extra": "410 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 27589934,
            "unit": "ns/op",
            "extra": "454 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 27522769,
            "unit": "ns/op",
            "extra": "451 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 28096484,
            "unit": "ns/op",
            "extra": "408 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 29837354,
            "unit": "ns/op",
            "extra": "380 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 19970943,
            "unit": "ns/op",
            "extra": "618 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 18365064,
            "unit": "ns/op",
            "extra": "663 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 17040646,
            "unit": "ns/op",
            "extra": "699 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 17033230,
            "unit": "ns/op",
            "extra": "702 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 20394830,
            "unit": "ns/op",
            "extra": "576 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 17880092,
            "unit": "ns/op",
            "extra": "675 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 15884378,
            "unit": "ns/op",
            "extra": "756 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 17290718,
            "unit": "ns/op",
            "extra": "678 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 21528335,
            "unit": "ns/op",
            "extra": "555 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1140031,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 21572990,
            "unit": "ns/op",
            "extra": "559 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1140294,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 21650320,
            "unit": "ns/op",
            "extra": "559 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 901620,
            "unit": "ns/op",
            "extra": "13419 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 21467290,
            "unit": "ns/op",
            "extra": "558 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 890793,
            "unit": "ns/op",
            "extra": "13468 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 16123252,
            "unit": "ns/op",
            "extra": "744 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 741521,
            "unit": "ns/op",
            "extra": "16191 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 16390127,
            "unit": "ns/op",
            "extra": "736 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 834545,
            "unit": "ns/op",
            "extra": "14362 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 16418685,
            "unit": "ns/op",
            "extra": "736 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 932219,
            "unit": "ns/op",
            "extra": "12874 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 16496222,
            "unit": "ns/op",
            "extra": "728 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1033968,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
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
          "id": "34e0b305b00adb381482d1bc95e5d7f0dfb50ad5",
          "message": "test: add loose cr lf test",
          "timestamp": "2026-09-14T12:32:07+02:00",
          "tree_id": "6e2a96329757ec63a4742c51f3a8a9c6198fcd89",
          "url": "https://github.com/uponusolutions/go-smtp/commit/34e0b305b00adb381482d1bc95e5d7f0dfb50ad5"
        },
        "date": 1789382701637,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 377235,
            "unit": "ns/op",
            "extra": "32130 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 368709,
            "unit": "ns/op",
            "extra": "32407 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 113853,
            "unit": "ns/op",
            "extra": "104484 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 111173,
            "unit": "ns/op",
            "extra": "108186 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 317100,
            "unit": "ns/op",
            "extra": "37818 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 315023,
            "unit": "ns/op",
            "extra": "37988 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 58494,
            "unit": "ns/op",
            "extra": "202959 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 55628,
            "unit": "ns/op",
            "extra": "214086 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 437916,
            "unit": "ns/op",
            "extra": "27537 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 347716,
            "unit": "ns/op",
            "extra": "34396 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 89791,
            "unit": "ns/op",
            "extra": "132798 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 90598,
            "unit": "ns/op",
            "extra": "132675 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 404704,
            "unit": "ns/op",
            "extra": "29709 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 311333,
            "unit": "ns/op",
            "extra": "38348 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 52088,
            "unit": "ns/op",
            "extra": "223588 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 52897,
            "unit": "ns/op",
            "extra": "224922 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 25858667,
            "unit": "ns/op",
            "extra": "429 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 26773115,
            "unit": "ns/op",
            "extra": "480 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 1172092643,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 872789291,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 25619860,
            "unit": "ns/op",
            "extra": "475 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 25340069,
            "unit": "ns/op",
            "extra": "474 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 26640207,
            "unit": "ns/op",
            "extra": "420 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 27514056,
            "unit": "ns/op",
            "extra": "399 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 18175058,
            "unit": "ns/op",
            "extra": "663 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 17973220,
            "unit": "ns/op",
            "extra": "664 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 15693423,
            "unit": "ns/op",
            "extra": "753 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 16862375,
            "unit": "ns/op",
            "extra": "711 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 17912231,
            "unit": "ns/op",
            "extra": "654 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 17888304,
            "unit": "ns/op",
            "extra": "684 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 15393435,
            "unit": "ns/op",
            "extra": "774 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 16479228,
            "unit": "ns/op",
            "extra": "724 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 22675445,
            "unit": "ns/op",
            "extra": "529 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1139525,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 22682183,
            "unit": "ns/op",
            "extra": "528 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1135993,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 22738384,
            "unit": "ns/op",
            "extra": "528 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 936284,
            "unit": "ns/op",
            "extra": "12890 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 22658304,
            "unit": "ns/op",
            "extra": "529 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 885255,
            "unit": "ns/op",
            "extra": "13658 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 16078230,
            "unit": "ns/op",
            "extra": "745 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 739346,
            "unit": "ns/op",
            "extra": "16225 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 16253754,
            "unit": "ns/op",
            "extra": "734 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 830918,
            "unit": "ns/op",
            "extra": "14433 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 16289264,
            "unit": "ns/op",
            "extra": "734 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 932294,
            "unit": "ns/op",
            "extra": "12868 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 16437675,
            "unit": "ns/op",
            "extra": "727 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1028709,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
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
          "id": "eee87b0d36517748d6a533599196fc638d7ee23a",
          "message": "test: imrove dot-cr case",
          "timestamp": "2026-09-14T12:42:27+02:00",
          "tree_id": "f2e4c7249601f7390efe1cc253bf4a6420c1718f",
          "url": "https://github.com/uponusolutions/go-smtp/commit/eee87b0d36517748d6a533599196fc638d7ee23a"
        },
        "date": 1789383136978,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 385644,
            "unit": "ns/op",
            "extra": "30579 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 380867,
            "unit": "ns/op",
            "extra": "31500 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 117241,
            "unit": "ns/op",
            "extra": "102516 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 114118,
            "unit": "ns/op",
            "extra": "104924 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 327425,
            "unit": "ns/op",
            "extra": "36540 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 324046,
            "unit": "ns/op",
            "extra": "36919 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 60219,
            "unit": "ns/op",
            "extra": "197965 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 57214,
            "unit": "ns/op",
            "extra": "208800 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 452530,
            "unit": "ns/op",
            "extra": "26522 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 358909,
            "unit": "ns/op",
            "extra": "33321 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 91976,
            "unit": "ns/op",
            "extra": "129316 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 91771,
            "unit": "ns/op",
            "extra": "129536 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 417332,
            "unit": "ns/op",
            "extra": "28748 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 320036,
            "unit": "ns/op",
            "extra": "37501 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 53483,
            "unit": "ns/op",
            "extra": "215935 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 53853,
            "unit": "ns/op",
            "extra": "223436 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 28274287,
            "unit": "ns/op",
            "extra": "439 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 27496609,
            "unit": "ns/op",
            "extra": "434 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 29035137,
            "unit": "ns/op",
            "extra": "355 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 28829980,
            "unit": "ns/op",
            "extra": "400 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 26899338,
            "unit": "ns/op",
            "extra": "442 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 29049063,
            "unit": "ns/op",
            "extra": "409 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 30194464,
            "unit": "ns/op",
            "extra": "391 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 28974888,
            "unit": "ns/op",
            "extra": "380 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 20345189,
            "unit": "ns/op",
            "extra": "600 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 18235925,
            "unit": "ns/op",
            "extra": "680 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 16285989,
            "unit": "ns/op",
            "extra": "732 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 16656354,
            "unit": "ns/op",
            "extra": "705 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 20153785,
            "unit": "ns/op",
            "extra": "606 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 18018159,
            "unit": "ns/op",
            "extra": "651 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 16608730,
            "unit": "ns/op",
            "extra": "734 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 17155968,
            "unit": "ns/op",
            "extra": "675 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 22743939,
            "unit": "ns/op",
            "extra": "524 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1142739,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 22770623,
            "unit": "ns/op",
            "extra": "526 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1143880,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 22718019,
            "unit": "ns/op",
            "extra": "528 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 902953,
            "unit": "ns/op",
            "extra": "13267 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 22714013,
            "unit": "ns/op",
            "extra": "526 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 931645,
            "unit": "ns/op",
            "extra": "12984 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 16085460,
            "unit": "ns/op",
            "extra": "746 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 744879,
            "unit": "ns/op",
            "extra": "16101 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 16249255,
            "unit": "ns/op",
            "extra": "736 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 834210,
            "unit": "ns/op",
            "extra": "14374 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 16276498,
            "unit": "ns/op",
            "extra": "738 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 938748,
            "unit": "ns/op",
            "extra": "12778 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 16491298,
            "unit": "ns/op",
            "extra": "728 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1032480,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
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
          "id": "af69d2d458cec8f98b2920bd68597368d1a1132e",
          "message": "rework: internal folder structure",
          "timestamp": "2026-09-14T15:53:33+02:00",
          "tree_id": "8b9c503e78d712818cc2ca4827d2ef4740c764c3",
          "url": "https://github.com/uponusolutions/go-smtp/commit/af69d2d458cec8f98b2920bd68597368d1a1132e"
        },
        "date": 1789394650594,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 362290,
            "unit": "ns/op",
            "extra": "33079 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 355968,
            "unit": "ns/op",
            "extra": "33936 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 124144,
            "unit": "ns/op",
            "extra": "95865 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 121710,
            "unit": "ns/op",
            "extra": "98637 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 295151,
            "unit": "ns/op",
            "extra": "40506 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 291091,
            "unit": "ns/op",
            "extra": "41307 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 61628,
            "unit": "ns/op",
            "extra": "199832 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 59112,
            "unit": "ns/op",
            "extra": "198733 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 427523,
            "unit": "ns/op",
            "extra": "27943 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 327967,
            "unit": "ns/op",
            "extra": "36804 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 97740,
            "unit": "ns/op",
            "extra": "122872 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 97992,
            "unit": "ns/op",
            "extra": "121540 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 399405,
            "unit": "ns/op",
            "extra": "29716 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 287724,
            "unit": "ns/op",
            "extra": "41588 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 56091,
            "unit": "ns/op",
            "extra": "214155 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 56417,
            "unit": "ns/op",
            "extra": "211471 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 27904494,
            "unit": "ns/op",
            "extra": "417 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 28313449,
            "unit": "ns/op",
            "extra": "432 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 34390808,
            "unit": "ns/op",
            "extra": "309 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 502901178,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 28261382,
            "unit": "ns/op",
            "extra": "430 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 27529769,
            "unit": "ns/op",
            "extra": "441 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 178340610,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 36809039,
            "unit": "ns/op",
            "extra": "368 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 18760230,
            "unit": "ns/op",
            "extra": "632 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 17992341,
            "unit": "ns/op",
            "extra": "664 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 16458796,
            "unit": "ns/op",
            "extra": "726 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 17392710,
            "unit": "ns/op",
            "extra": "693 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 18586610,
            "unit": "ns/op",
            "extra": "657 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 18224537,
            "unit": "ns/op",
            "extra": "652 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 16499878,
            "unit": "ns/op",
            "extra": "722 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 17265905,
            "unit": "ns/op",
            "extra": "681 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21592498,
            "unit": "ns/op",
            "extra": "564 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1189926,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21912471,
            "unit": "ns/op",
            "extra": "555 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1188948,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21361967,
            "unit": "ns/op",
            "extra": "564 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 891678,
            "unit": "ns/op",
            "extra": "13477 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21358594,
            "unit": "ns/op",
            "extra": "562 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 892347,
            "unit": "ns/op",
            "extra": "13432 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16643118,
            "unit": "ns/op",
            "extra": "721 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 761652,
            "unit": "ns/op",
            "extra": "15734 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16848322,
            "unit": "ns/op",
            "extra": "711 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 869390,
            "unit": "ns/op",
            "extra": "13788 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16841481,
            "unit": "ns/op",
            "extra": "711 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 851741,
            "unit": "ns/op",
            "extra": "14092 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17375663,
            "unit": "ns/op",
            "extra": "698 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 980796,
            "unit": "ns/op",
            "extra": "12240 times\n4 procs"
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
          "id": "58e5658c8d14109ec008020ac40a815c21c9a6b6",
          "message": "feat: debug client/server",
          "timestamp": "2026-09-14T16:31:16+02:00",
          "tree_id": "80cd9fd5d96e845148b125fd4179ad2acc81f118",
          "url": "https://github.com/uponusolutions/go-smtp/commit/58e5658c8d14109ec008020ac40a815c21c9a6b6"
        },
        "date": 1789396867071,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 369169,
            "unit": "ns/op",
            "extra": "32457 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 362609,
            "unit": "ns/op",
            "extra": "33078 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 124030,
            "unit": "ns/op",
            "extra": "96446 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 121217,
            "unit": "ns/op",
            "extra": "98952 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 304486,
            "unit": "ns/op",
            "extra": "39436 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 300688,
            "unit": "ns/op",
            "extra": "39801 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 61527,
            "unit": "ns/op",
            "extra": "193614 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 58997,
            "unit": "ns/op",
            "extra": "204050 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 440288,
            "unit": "ns/op",
            "extra": "27276 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 339307,
            "unit": "ns/op",
            "extra": "35366 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 97794,
            "unit": "ns/op",
            "extra": "120393 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 98144,
            "unit": "ns/op",
            "extra": "122793 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 397807,
            "unit": "ns/op",
            "extra": "30306 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 298400,
            "unit": "ns/op",
            "extra": "40252 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 56311,
            "unit": "ns/op",
            "extra": "211490 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 56436,
            "unit": "ns/op",
            "extra": "210553 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 29644029,
            "unit": "ns/op",
            "extra": "388 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 28632691,
            "unit": "ns/op",
            "extra": "414 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 30563647,
            "unit": "ns/op",
            "extra": "354 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 29713726,
            "unit": "ns/op",
            "extra": "361 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 29821296,
            "unit": "ns/op",
            "extra": "410 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 28583036,
            "unit": "ns/op",
            "extra": "420 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 31651562,
            "unit": "ns/op",
            "extra": "375 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 31620859,
            "unit": "ns/op",
            "extra": "328 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 20902883,
            "unit": "ns/op",
            "extra": "562 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 19286802,
            "unit": "ns/op",
            "extra": "610 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 17806362,
            "unit": "ns/op",
            "extra": "652 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 18869745,
            "unit": "ns/op",
            "extra": "601 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 20010071,
            "unit": "ns/op",
            "extra": "585 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 19668871,
            "unit": "ns/op",
            "extra": "594 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 17701602,
            "unit": "ns/op",
            "extra": "663 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 18132263,
            "unit": "ns/op",
            "extra": "660 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21691180,
            "unit": "ns/op",
            "extra": "560 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1188015,
            "unit": "ns/op",
            "extra": "9998 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21640562,
            "unit": "ns/op",
            "extra": "547 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1185215,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21331867,
            "unit": "ns/op",
            "extra": "561 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 890795,
            "unit": "ns/op",
            "extra": "13483 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21329750,
            "unit": "ns/op",
            "extra": "562 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 890942,
            "unit": "ns/op",
            "extra": "13461 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16642563,
            "unit": "ns/op",
            "extra": "718 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 760481,
            "unit": "ns/op",
            "extra": "15747 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16889436,
            "unit": "ns/op",
            "extra": "710 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 872158,
            "unit": "ns/op",
            "extra": "13782 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16830531,
            "unit": "ns/op",
            "extra": "711 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 853845,
            "unit": "ns/op",
            "extra": "14071 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17135006,
            "unit": "ns/op",
            "extra": "699 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 980404,
            "unit": "ns/op",
            "extra": "12236 times\n4 procs"
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
          "id": "ed9a003d14b244b5c8ab0c64e3351cd1161cf679",
          "message": "feat: String method for smtp.Status",
          "timestamp": "2026-09-15T15:46:39+02:00",
          "tree_id": "4390f21c1abbca468bc7bd47ec4ceb0885f9caac",
          "url": "https://github.com/uponusolutions/go-smtp/commit/ed9a003d14b244b5c8ab0c64e3351cd1161cf679"
        },
        "date": 1789480715980,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 392575,
            "unit": "ns/op",
            "extra": "30184 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 381944,
            "unit": "ns/op",
            "extra": "31360 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 120025,
            "unit": "ns/op",
            "extra": "100166 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 115654,
            "unit": "ns/op",
            "extra": "103864 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 328416,
            "unit": "ns/op",
            "extra": "36561 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 326951,
            "unit": "ns/op",
            "extra": "36843 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 62272,
            "unit": "ns/op",
            "extra": "189558 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 58123,
            "unit": "ns/op",
            "extra": "203384 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 489100,
            "unit": "ns/op",
            "extra": "24424 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 361440,
            "unit": "ns/op",
            "extra": "32826 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 93140,
            "unit": "ns/op",
            "extra": "125491 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 93084,
            "unit": "ns/op",
            "extra": "128298 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 451088,
            "unit": "ns/op",
            "extra": "27208 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 326525,
            "unit": "ns/op",
            "extra": "36698 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 54720,
            "unit": "ns/op",
            "extra": "217324 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 54885,
            "unit": "ns/op",
            "extra": "218626 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 26509970,
            "unit": "ns/op",
            "extra": "454 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 26990707,
            "unit": "ns/op",
            "extra": "442 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 645341964,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 30087616,
            "unit": "ns/op",
            "extra": "417 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 27396656,
            "unit": "ns/op",
            "extra": "451 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 27650443,
            "unit": "ns/op",
            "extra": "446 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 763787999,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 30676399,
            "unit": "ns/op",
            "extra": "374 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 18773670,
            "unit": "ns/op",
            "extra": "661 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 17902599,
            "unit": "ns/op",
            "extra": "657 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 16591858,
            "unit": "ns/op",
            "extra": "741 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 16970369,
            "unit": "ns/op",
            "extra": "700 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 17638595,
            "unit": "ns/op",
            "extra": "676 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 18106808,
            "unit": "ns/op",
            "extra": "649 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 16708242,
            "unit": "ns/op",
            "extra": "717 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 16187112,
            "unit": "ns/op",
            "extra": "705 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21618998,
            "unit": "ns/op",
            "extra": "553 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1130397,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21554340,
            "unit": "ns/op",
            "extra": "556 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1131147,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21586553,
            "unit": "ns/op",
            "extra": "556 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 876411,
            "unit": "ns/op",
            "extra": "13653 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21975673,
            "unit": "ns/op",
            "extra": "553 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 877758,
            "unit": "ns/op",
            "extra": "13779 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16118833,
            "unit": "ns/op",
            "extra": "744 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 750010,
            "unit": "ns/op",
            "extra": "15978 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16745304,
            "unit": "ns/op",
            "extra": "727 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 845745,
            "unit": "ns/op",
            "extra": "14174 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16304792,
            "unit": "ns/op",
            "extra": "735 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 947265,
            "unit": "ns/op",
            "extra": "12675 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16673606,
            "unit": "ns/op",
            "extra": "723 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1055085,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
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
          "id": "60916c97bdb2a353dd82280fec5b8c3c65b6224d",
          "message": "docs: comment String function",
          "timestamp": "2026-09-15T15:49:33+02:00",
          "tree_id": "b7afc052ef484d823b4cc4d7d6c4b11d0ea4bb3f",
          "url": "https://github.com/uponusolutions/go-smtp/commit/60916c97bdb2a353dd82280fec5b8c3c65b6224d"
        },
        "date": 1789480778882,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 365979,
            "unit": "ns/op",
            "extra": "32311 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 362991,
            "unit": "ns/op",
            "extra": "32991 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 127368,
            "unit": "ns/op",
            "extra": "93838 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 125126,
            "unit": "ns/op",
            "extra": "95662 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 303113,
            "unit": "ns/op",
            "extra": "39634 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 299108,
            "unit": "ns/op",
            "extra": "39873 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 63539,
            "unit": "ns/op",
            "extra": "187832 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 61113,
            "unit": "ns/op",
            "extra": "197080 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 447301,
            "unit": "ns/op",
            "extra": "26713 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 339758,
            "unit": "ns/op",
            "extra": "35265 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 100233,
            "unit": "ns/op",
            "extra": "118864 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 100919,
            "unit": "ns/op",
            "extra": "119068 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 400302,
            "unit": "ns/op",
            "extra": "29936 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 297335,
            "unit": "ns/op",
            "extra": "40483 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 57321,
            "unit": "ns/op",
            "extra": "206288 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 57516,
            "unit": "ns/op",
            "extra": "207900 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 28612009,
            "unit": "ns/op",
            "extra": "428 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 29235526,
            "unit": "ns/op",
            "extra": "388 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 232340016,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 29473121,
            "unit": "ns/op",
            "extra": "403 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 30672060,
            "unit": "ns/op",
            "extra": "339 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 28405244,
            "unit": "ns/op",
            "extra": "422 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 33283482,
            "unit": "ns/op",
            "extra": "338 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 28572680,
            "unit": "ns/op",
            "extra": "405 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 18953168,
            "unit": "ns/op",
            "extra": "638 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 18415172,
            "unit": "ns/op",
            "extra": "637 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 15994255,
            "unit": "ns/op",
            "extra": "715 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 16879535,
            "unit": "ns/op",
            "extra": "722 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 18694008,
            "unit": "ns/op",
            "extra": "640 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 18080693,
            "unit": "ns/op",
            "extra": "662 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 15787072,
            "unit": "ns/op",
            "extra": "781 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 16996097,
            "unit": "ns/op",
            "extra": "684 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21757624,
            "unit": "ns/op",
            "extra": "529 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1196161,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21944915,
            "unit": "ns/op",
            "extra": "546 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1190062,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 22483713,
            "unit": "ns/op",
            "extra": "561 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 896165,
            "unit": "ns/op",
            "extra": "13401 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21751144,
            "unit": "ns/op",
            "extra": "562 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 890176,
            "unit": "ns/op",
            "extra": "13460 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16642445,
            "unit": "ns/op",
            "extra": "721 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 760605,
            "unit": "ns/op",
            "extra": "15778 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16908550,
            "unit": "ns/op",
            "extra": "708 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 872044,
            "unit": "ns/op",
            "extra": "13756 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16865603,
            "unit": "ns/op",
            "extra": "711 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 856866,
            "unit": "ns/op",
            "extra": "13971 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17129247,
            "unit": "ns/op",
            "extra": "702 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 977278,
            "unit": "ns/op",
            "extra": "12280 times\n4 procs"
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
          "id": "42129096792893e9273d762afe6a68a6c7e68fe9",
          "message": "fix: only send * to abort AUTH if negotiation is still running",
          "timestamp": "2026-09-16T03:24:58+02:00",
          "tree_id": "092334ab5437a1a88a7da8ea945e6c06eec57236",
          "url": "https://github.com/uponusolutions/go-smtp/commit/42129096792893e9273d762afe6a68a6c7e68fe9"
        },
        "date": 1789522611088,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 362951,
            "unit": "ns/op",
            "extra": "32638 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 363167,
            "unit": "ns/op",
            "extra": "33210 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 125127,
            "unit": "ns/op",
            "extra": "95972 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 121970,
            "unit": "ns/op",
            "extra": "98584 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 301892,
            "unit": "ns/op",
            "extra": "39746 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 300031,
            "unit": "ns/op",
            "extra": "39944 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 60846,
            "unit": "ns/op",
            "extra": "196286 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 58897,
            "unit": "ns/op",
            "extra": "203052 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 422772,
            "unit": "ns/op",
            "extra": "28430 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 336193,
            "unit": "ns/op",
            "extra": "35772 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 97611,
            "unit": "ns/op",
            "extra": "122635 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 98411,
            "unit": "ns/op",
            "extra": "122188 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 394618,
            "unit": "ns/op",
            "extra": "30460 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 297145,
            "unit": "ns/op",
            "extra": "40460 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 56105,
            "unit": "ns/op",
            "extra": "212672 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 56771,
            "unit": "ns/op",
            "extra": "211216 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 29150111,
            "unit": "ns/op",
            "extra": "426 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 29462715,
            "unit": "ns/op",
            "extra": "418 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 29613863,
            "unit": "ns/op",
            "extra": "373 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 633431969,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 28353007,
            "unit": "ns/op",
            "extra": "420 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 27808929,
            "unit": "ns/op",
            "extra": "435 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 231717873,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 565550651,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 18325564,
            "unit": "ns/op",
            "extra": "670 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 19036347,
            "unit": "ns/op",
            "extra": "630 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 17361075,
            "unit": "ns/op",
            "extra": "709 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 18799919,
            "unit": "ns/op",
            "extra": "606 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 19776547,
            "unit": "ns/op",
            "extra": "630 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 20405431,
            "unit": "ns/op",
            "extra": "597 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 18237174,
            "unit": "ns/op",
            "extra": "668 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 18536695,
            "unit": "ns/op",
            "extra": "679 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21550039,
            "unit": "ns/op",
            "extra": "558 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1188752,
            "unit": "ns/op",
            "extra": "9993 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21523783,
            "unit": "ns/op",
            "extra": "562 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1198645,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21464366,
            "unit": "ns/op",
            "extra": "560 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 899523,
            "unit": "ns/op",
            "extra": "13344 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21707532,
            "unit": "ns/op",
            "extra": "560 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 901876,
            "unit": "ns/op",
            "extra": "13320 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16728218,
            "unit": "ns/op",
            "extra": "720 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 760335,
            "unit": "ns/op",
            "extra": "15759 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16890340,
            "unit": "ns/op",
            "extra": "712 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 872127,
            "unit": "ns/op",
            "extra": "13734 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16845464,
            "unit": "ns/op",
            "extra": "711 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 854406,
            "unit": "ns/op",
            "extra": "14044 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17087970,
            "unit": "ns/op",
            "extra": "705 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 974396,
            "unit": "ns/op",
            "extra": "12307 times\n4 procs"
          }
        ]
      }
    ]
  }
}