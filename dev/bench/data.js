window.BENCHMARK_DATA = {
  "lastUpdate": 1794847356545,
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
          "id": "23c7d6f3053f7ad060b54cff0860dd2828a4f6d6",
          "message": "rework: remove integrated way to wrap conn",
          "timestamp": "2026-09-16T03:59:47+02:00",
          "tree_id": "d594afcc6d77ad7bdc3c7d313a31e8507294d2ec",
          "url": "https://github.com/uponusolutions/go-smtp/commit/23c7d6f3053f7ad060b54cff0860dd2828a4f6d6"
        },
        "date": 1789524573118,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 387031,
            "unit": "ns/op",
            "extra": "30604 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 378958,
            "unit": "ns/op",
            "extra": "31435 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 117531,
            "unit": "ns/op",
            "extra": "101024 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 114773,
            "unit": "ns/op",
            "extra": "105166 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 322633,
            "unit": "ns/op",
            "extra": "37160 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 320579,
            "unit": "ns/op",
            "extra": "37422 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 60980,
            "unit": "ns/op",
            "extra": "195717 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 57765,
            "unit": "ns/op",
            "extra": "208474 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 462217,
            "unit": "ns/op",
            "extra": "25945 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 362030,
            "unit": "ns/op",
            "extra": "33008 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 93494,
            "unit": "ns/op",
            "extra": "128004 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 93578,
            "unit": "ns/op",
            "extra": "127424 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 483035,
            "unit": "ns/op",
            "extra": "24655 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 329981,
            "unit": "ns/op",
            "extra": "36079 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 55637,
            "unit": "ns/op",
            "extra": "214082 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 56841,
            "unit": "ns/op",
            "extra": "208771 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 26794241,
            "unit": "ns/op",
            "extra": "448 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 27316697,
            "unit": "ns/op",
            "extra": "414 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 29510687,
            "unit": "ns/op",
            "extra": "386 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 29782442,
            "unit": "ns/op",
            "extra": "364 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 27771563,
            "unit": "ns/op",
            "extra": "448 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 26997128,
            "unit": "ns/op",
            "extra": "452 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 28113757,
            "unit": "ns/op",
            "extra": "445 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 29745077,
            "unit": "ns/op",
            "extra": "380 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 18735910,
            "unit": "ns/op",
            "extra": "627 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 17516598,
            "unit": "ns/op",
            "extra": "675 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 15894475,
            "unit": "ns/op",
            "extra": "720 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 16883557,
            "unit": "ns/op",
            "extra": "696 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 17249984,
            "unit": "ns/op",
            "extra": "688 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 17427836,
            "unit": "ns/op",
            "extra": "700 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 16637267,
            "unit": "ns/op",
            "extra": "702 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 16766569,
            "unit": "ns/op",
            "extra": "703 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21526162,
            "unit": "ns/op",
            "extra": "558 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1129598,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21572591,
            "unit": "ns/op",
            "extra": "558 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1132712,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21547119,
            "unit": "ns/op",
            "extra": "558 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 927944,
            "unit": "ns/op",
            "extra": "12886 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21786018,
            "unit": "ns/op",
            "extra": "552 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 911053,
            "unit": "ns/op",
            "extra": "13404 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16115662,
            "unit": "ns/op",
            "extra": "746 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 748750,
            "unit": "ns/op",
            "extra": "16003 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16339953,
            "unit": "ns/op",
            "extra": "734 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 841066,
            "unit": "ns/op",
            "extra": "14264 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16343200,
            "unit": "ns/op",
            "extra": "735 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 946202,
            "unit": "ns/op",
            "extra": "12679 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16579389,
            "unit": "ns/op",
            "extra": "721 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1046027,
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
          "id": "3543cc33c8b51429fb7337826c497a2d0e101183",
          "message": "feat: split textproto",
          "timestamp": "2026-09-16T23:09:05+02:00",
          "tree_id": "b28a38d3f5459591420d48ffd9a299160e49c2c8",
          "url": "https://github.com/uponusolutions/go-smtp/commit/3543cc33c8b51429fb7337826c497a2d0e101183"
        },
        "date": 1789593545418,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 387849,
            "unit": "ns/op",
            "extra": "31406 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 373639,
            "unit": "ns/op",
            "extra": "32019 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 117156,
            "unit": "ns/op",
            "extra": "102318 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 112866,
            "unit": "ns/op",
            "extra": "106286 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 322492,
            "unit": "ns/op",
            "extra": "37216 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 318177,
            "unit": "ns/op",
            "extra": "37374 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 59528,
            "unit": "ns/op",
            "extra": "199214 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 56533,
            "unit": "ns/op",
            "extra": "212175 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 449723,
            "unit": "ns/op",
            "extra": "26739 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 353114,
            "unit": "ns/op",
            "extra": "33982 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 91422,
            "unit": "ns/op",
            "extra": "129753 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 91422,
            "unit": "ns/op",
            "extra": "130924 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 411608,
            "unit": "ns/op",
            "extra": "28862 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 313913,
            "unit": "ns/op",
            "extra": "38236 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 53844,
            "unit": "ns/op",
            "extra": "223988 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 54425,
            "unit": "ns/op",
            "extra": "218540 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 27550234,
            "unit": "ns/op",
            "extra": "436 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 25556504,
            "unit": "ns/op",
            "extra": "476 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 32763418,
            "unit": "ns/op",
            "extra": "364 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 26847291,
            "unit": "ns/op",
            "extra": "404 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 26367079,
            "unit": "ns/op",
            "extra": "463 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 25723794,
            "unit": "ns/op",
            "extra": "463 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 30495310,
            "unit": "ns/op",
            "extra": "394 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 27475765,
            "unit": "ns/op",
            "extra": "398 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 19482685,
            "unit": "ns/op",
            "extra": "613 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 17434044,
            "unit": "ns/op",
            "extra": "690 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 15941256,
            "unit": "ns/op",
            "extra": "757 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 16079317,
            "unit": "ns/op",
            "extra": "727 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 19439566,
            "unit": "ns/op",
            "extra": "625 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 18221398,
            "unit": "ns/op",
            "extra": "648 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 16476913,
            "unit": "ns/op",
            "extra": "736 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 17595154,
            "unit": "ns/op",
            "extra": "676 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21495611,
            "unit": "ns/op",
            "extra": "558 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1126723,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21608174,
            "unit": "ns/op",
            "extra": "558 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1127338,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21516240,
            "unit": "ns/op",
            "extra": "558 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 875373,
            "unit": "ns/op",
            "extra": "13671 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 21458256,
            "unit": "ns/op",
            "extra": "560 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 874036,
            "unit": "ns/op",
            "extra": "13753 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16153944,
            "unit": "ns/op",
            "extra": "741 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 746746,
            "unit": "ns/op",
            "extra": "16058 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16268595,
            "unit": "ns/op",
            "extra": "736 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 837861,
            "unit": "ns/op",
            "extra": "14323 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16334674,
            "unit": "ns/op",
            "extra": "733 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 944152,
            "unit": "ns/op",
            "extra": "12691 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16518578,
            "unit": "ns/op",
            "extra": "726 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1041632,
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
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "6488dbfff461f6c66a2b356271d3000b3a0a1615",
          "message": "Merge pull request #12 from uponusolutions/feature/ms-1784-dial-control\n\nfeat: dial control for client connections",
          "timestamp": "2026-11-05T11:24:52+01:00",
          "tree_id": "9f9f86a8b32aa9ab6b96077fbe0d86c460e2d6c8",
          "url": "https://github.com/uponusolutions/go-smtp/commit/6488dbfff461f6c66a2b356271d3000b3a0a1615"
        },
        "date": 1793874896976,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 290050,
            "unit": "ns/op",
            "extra": "41840 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 286499,
            "unit": "ns/op",
            "extra": "41992 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 96489,
            "unit": "ns/op",
            "extra": "124700 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 93921,
            "unit": "ns/op",
            "extra": "128118 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 240507,
            "unit": "ns/op",
            "extra": "49347 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 237363,
            "unit": "ns/op",
            "extra": "49760 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 48488,
            "unit": "ns/op",
            "extra": "249390 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 46208,
            "unit": "ns/op",
            "extra": "257718 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 345212,
            "unit": "ns/op",
            "extra": "34540 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 268304,
            "unit": "ns/op",
            "extra": "44590 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 75171,
            "unit": "ns/op",
            "extra": "159274 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 75778,
            "unit": "ns/op",
            "extra": "158660 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 312869,
            "unit": "ns/op",
            "extra": "38210 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 234431,
            "unit": "ns/op",
            "extra": "51084 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 43503,
            "unit": "ns/op",
            "extra": "256699 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 43514,
            "unit": "ns/op",
            "extra": "273261 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 29039602,
            "unit": "ns/op",
            "extra": "428 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 29635123,
            "unit": "ns/op",
            "extra": "404 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 33829602,
            "unit": "ns/op",
            "extra": "372 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 33310748,
            "unit": "ns/op",
            "extra": "374 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 28705654,
            "unit": "ns/op",
            "extra": "424 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 28742546,
            "unit": "ns/op",
            "extra": "426 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 38380950,
            "unit": "ns/op",
            "extra": "333 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 37730339,
            "unit": "ns/op",
            "extra": "332 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 21947332,
            "unit": "ns/op",
            "extra": "548 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 19009251,
            "unit": "ns/op",
            "extra": "621 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 17497772,
            "unit": "ns/op",
            "extra": "662 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 17670025,
            "unit": "ns/op",
            "extra": "674 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 20453476,
            "unit": "ns/op",
            "extra": "620 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 18741408,
            "unit": "ns/op",
            "extra": "577 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 16059739,
            "unit": "ns/op",
            "extra": "727 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 18983922,
            "unit": "ns/op",
            "extra": "654 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17745965,
            "unit": "ns/op",
            "extra": "675 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 943782,
            "unit": "ns/op",
            "extra": "12709 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17699453,
            "unit": "ns/op",
            "extra": "678 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 930671,
            "unit": "ns/op",
            "extra": "12909 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17681590,
            "unit": "ns/op",
            "extra": "679 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 708747,
            "unit": "ns/op",
            "extra": "16958 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17680710,
            "unit": "ns/op",
            "extra": "676 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 687038,
            "unit": "ns/op",
            "extra": "17478 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16922423,
            "unit": "ns/op",
            "extra": "686 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 588728,
            "unit": "ns/op",
            "extra": "20348 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17505176,
            "unit": "ns/op",
            "extra": "675 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 679396,
            "unit": "ns/op",
            "extra": "17637 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17633324,
            "unit": "ns/op",
            "extra": "673 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 660779,
            "unit": "ns/op",
            "extra": "18133 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17497979,
            "unit": "ns/op",
            "extra": "717 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 765823,
            "unit": "ns/op",
            "extra": "15615 times\n4 procs"
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
          "id": "132f2ecc857f433645e80c225573b4b84d3fd0f2",
          "message": "docs: security note on prefer starttls",
          "timestamp": "2026-11-12T15:57:00+01:00",
          "tree_id": "1f0efec3dafc57ecfbb5f93482ac2bfd9bb2710f",
          "url": "https://github.com/uponusolutions/go-smtp/commit/132f2ecc857f433645e80c225573b4b84d3fd0f2"
        },
        "date": 1794496449053,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 387240,
            "unit": "ns/op",
            "extra": "31153 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 381220,
            "unit": "ns/op",
            "extra": "31470 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 120311,
            "unit": "ns/op",
            "extra": "98980 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 115307,
            "unit": "ns/op",
            "extra": "103462 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 332484,
            "unit": "ns/op",
            "extra": "34929 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 323995,
            "unit": "ns/op",
            "extra": "36877 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 62658,
            "unit": "ns/op",
            "extra": "189832 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 58640,
            "unit": "ns/op",
            "extra": "203966 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 503358,
            "unit": "ns/op",
            "extra": "23974 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 365733,
            "unit": "ns/op",
            "extra": "31761 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 93674,
            "unit": "ns/op",
            "extra": "126826 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 94192,
            "unit": "ns/op",
            "extra": "125800 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 483546,
            "unit": "ns/op",
            "extra": "24598 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 323078,
            "unit": "ns/op",
            "extra": "36912 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 56456,
            "unit": "ns/op",
            "extra": "210726 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 56464,
            "unit": "ns/op",
            "extra": "213105 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 37835639,
            "unit": "ns/op",
            "extra": "322 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 36850682,
            "unit": "ns/op",
            "extra": "320 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 49845891,
            "unit": "ns/op",
            "extra": "255 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 45704409,
            "unit": "ns/op",
            "extra": "254 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 36719020,
            "unit": "ns/op",
            "extra": "316 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 36329299,
            "unit": "ns/op",
            "extra": "319 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 47729108,
            "unit": "ns/op",
            "extra": "256 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 43073060,
            "unit": "ns/op",
            "extra": "274 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 25155091,
            "unit": "ns/op",
            "extra": "469 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 25931811,
            "unit": "ns/op",
            "extra": "475 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 22914834,
            "unit": "ns/op",
            "extra": "540 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 25746861,
            "unit": "ns/op",
            "extra": "463 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 23627978,
            "unit": "ns/op",
            "extra": "513 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 26296017,
            "unit": "ns/op",
            "extra": "439 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 22790311,
            "unit": "ns/op",
            "extra": "530 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 23546822,
            "unit": "ns/op",
            "extra": "489 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 22811032,
            "unit": "ns/op",
            "extra": "525 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1154879,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 22880164,
            "unit": "ns/op",
            "extra": "524 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1154275,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 22824443,
            "unit": "ns/op",
            "extra": "525 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 926220,
            "unit": "ns/op",
            "extra": "13000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 22865468,
            "unit": "ns/op",
            "extra": "524 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 968894,
            "unit": "ns/op",
            "extra": "12415 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 20004257,
            "unit": "ns/op",
            "extra": "597 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 745865,
            "unit": "ns/op",
            "extra": "16083 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 20326644,
            "unit": "ns/op",
            "extra": "590 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 845520,
            "unit": "ns/op",
            "extra": "14191 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 20159074,
            "unit": "ns/op",
            "extra": "595 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 953701,
            "unit": "ns/op",
            "extra": "12574 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 20465318,
            "unit": "ns/op",
            "extra": "585 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1063544,
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
          "id": "fc9d9d58c8a0dc15bcc26a7172a3ff7bbd0280ce",
          "message": "feat: add meaningful defaults, improve deadline handling",
          "timestamp": "2026-11-13T21:40:36+01:00",
          "tree_id": "fe8dac76b267483128360a083a34a1c2237fb44b",
          "url": "https://github.com/uponusolutions/go-smtp/commit/fc9d9d58c8a0dc15bcc26a7172a3ff7bbd0280ce"
        },
        "date": 1794603051671,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 293666,
            "unit": "ns/op",
            "extra": "40782 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 289497,
            "unit": "ns/op",
            "extra": "41252 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 96426,
            "unit": "ns/op",
            "extra": "124404 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 94360,
            "unit": "ns/op",
            "extra": "127173 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 243808,
            "unit": "ns/op",
            "extra": "49096 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 242795,
            "unit": "ns/op",
            "extra": "49245 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 47363,
            "unit": "ns/op",
            "extra": "249874 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 46194,
            "unit": "ns/op",
            "extra": "260827 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 338154,
            "unit": "ns/op",
            "extra": "35443 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 270842,
            "unit": "ns/op",
            "extra": "44220 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 75198,
            "unit": "ns/op",
            "extra": "157848 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 75179,
            "unit": "ns/op",
            "extra": "159346 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 311186,
            "unit": "ns/op",
            "extra": "39114 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 241921,
            "unit": "ns/op",
            "extra": "49278 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 43574,
            "unit": "ns/op",
            "extra": "273708 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 44326,
            "unit": "ns/op",
            "extra": "269450 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 28949249,
            "unit": "ns/op",
            "extra": "415 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 30375141,
            "unit": "ns/op",
            "extra": "403 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 39025439,
            "unit": "ns/op",
            "extra": "326 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 38163055,
            "unit": "ns/op",
            "extra": "328 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 30249954,
            "unit": "ns/op",
            "extra": "378 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 30358110,
            "unit": "ns/op",
            "extra": "399 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 39573586,
            "unit": "ns/op",
            "extra": "319 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 39080005,
            "unit": "ns/op",
            "extra": "320 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 20951844,
            "unit": "ns/op",
            "extra": "558 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 17786442,
            "unit": "ns/op",
            "extra": "673 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 17687156,
            "unit": "ns/op",
            "extra": "715 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 17446263,
            "unit": "ns/op",
            "extra": "708 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 20917493,
            "unit": "ns/op",
            "extra": "571 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 17857177,
            "unit": "ns/op",
            "extra": "642 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 18025809,
            "unit": "ns/op",
            "extra": "664 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 16881495,
            "unit": "ns/op",
            "extra": "728 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17736919,
            "unit": "ns/op",
            "extra": "674 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 935005,
            "unit": "ns/op",
            "extra": "12848 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17690680,
            "unit": "ns/op",
            "extra": "676 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 930348,
            "unit": "ns/op",
            "extra": "12859 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17633794,
            "unit": "ns/op",
            "extra": "680 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 688956,
            "unit": "ns/op",
            "extra": "17420 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17691819,
            "unit": "ns/op",
            "extra": "681 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 689127,
            "unit": "ns/op",
            "extra": "17419 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16382462,
            "unit": "ns/op",
            "extra": "732 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 587453,
            "unit": "ns/op",
            "extra": "20401 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17500424,
            "unit": "ns/op",
            "extra": "679 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 671767,
            "unit": "ns/op",
            "extra": "17829 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17189722,
            "unit": "ns/op",
            "extra": "680 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 660831,
            "unit": "ns/op",
            "extra": "18146 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17654256,
            "unit": "ns/op",
            "extra": "687 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 758615,
            "unit": "ns/op",
            "extra": "15818 times\n4 procs"
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
          "id": "5df8198261be5cb896158cbac7204d709bb93373",
          "message": "fix: missing connection close in some cases of the client, pipeline\nimprovement",
          "timestamp": "2026-11-15T14:19:59+01:00",
          "tree_id": "728958240b9c7b8565f6e44ddbcb0a779c9a01df",
          "url": "https://github.com/uponusolutions/go-smtp/commit/5df8198261be5cb896158cbac7204d709bb93373"
        },
        "date": 1794749398471,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 290178,
            "unit": "ns/op",
            "extra": "41527 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 289369,
            "unit": "ns/op",
            "extra": "41848 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 96758,
            "unit": "ns/op",
            "extra": "123249 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 95076,
            "unit": "ns/op",
            "extra": "123628 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 242197,
            "unit": "ns/op",
            "extra": "48507 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 238424,
            "unit": "ns/op",
            "extra": "50223 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 48889,
            "unit": "ns/op",
            "extra": "248275 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 46295,
            "unit": "ns/op",
            "extra": "253838 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 351250,
            "unit": "ns/op",
            "extra": "36030 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 272662,
            "unit": "ns/op",
            "extra": "44042 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 76148,
            "unit": "ns/op",
            "extra": "157526 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 77192,
            "unit": "ns/op",
            "extra": "155151 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 307863,
            "unit": "ns/op",
            "extra": "39252 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 236733,
            "unit": "ns/op",
            "extra": "49926 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 43083,
            "unit": "ns/op",
            "extra": "273658 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 43252,
            "unit": "ns/op",
            "extra": "275790 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 27399193,
            "unit": "ns/op",
            "extra": "432 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 26170042,
            "unit": "ns/op",
            "extra": "469 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 36689273,
            "unit": "ns/op",
            "extra": "358 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 30807693,
            "unit": "ns/op",
            "extra": "416 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 26356901,
            "unit": "ns/op",
            "extra": "450 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 26412879,
            "unit": "ns/op",
            "extra": "447 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 37038282,
            "unit": "ns/op",
            "extra": "351 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 34862791,
            "unit": "ns/op",
            "extra": "394 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 19907122,
            "unit": "ns/op",
            "extra": "590 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 16272520,
            "unit": "ns/op",
            "extra": "770 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 15106425,
            "unit": "ns/op",
            "extra": "744 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 15076464,
            "unit": "ns/op",
            "extra": "723 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 20144817,
            "unit": "ns/op",
            "extra": "561 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 16321620,
            "unit": "ns/op",
            "extra": "724 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 15659759,
            "unit": "ns/op",
            "extra": "768 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 15152453,
            "unit": "ns/op",
            "extra": "794 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17676208,
            "unit": "ns/op",
            "extra": "679 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 940245,
            "unit": "ns/op",
            "extra": "12746 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17724851,
            "unit": "ns/op",
            "extra": "679 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 930281,
            "unit": "ns/op",
            "extra": "12856 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17792054,
            "unit": "ns/op",
            "extra": "679 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 693063,
            "unit": "ns/op",
            "extra": "17352 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17614298,
            "unit": "ns/op",
            "extra": "686 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 688739,
            "unit": "ns/op",
            "extra": "17408 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16432067,
            "unit": "ns/op",
            "extra": "733 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 590746,
            "unit": "ns/op",
            "extra": "20242 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17410189,
            "unit": "ns/op",
            "extra": "680 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 673741,
            "unit": "ns/op",
            "extra": "17790 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 16507232,
            "unit": "ns/op",
            "extra": "726 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 660660,
            "unit": "ns/op",
            "extra": "18138 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 17715526,
            "unit": "ns/op",
            "extra": "675 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 758173,
            "unit": "ns/op",
            "extra": "15825 times\n4 procs"
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
          "id": "d9bfc6d06d634898dee1ab1285eaf166e708e2ea",
          "message": "fix: transaction reset after failed data transfer",
          "timestamp": "2026-11-15T14:27:38+01:00",
          "tree_id": "935b7dfd9111bccde188e6553a36516f419d50d7",
          "url": "https://github.com/uponusolutions/go-smtp/commit/d9bfc6d06d634898dee1ab1285eaf166e708e2ea"
        },
        "date": 1794749860150,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 168826,
            "unit": "ns/op",
            "extra": "69918 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 171332,
            "unit": "ns/op",
            "extra": "72385 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 43988,
            "unit": "ns/op",
            "extra": "275734 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 43447,
            "unit": "ns/op",
            "extra": "282021 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 152385,
            "unit": "ns/op",
            "extra": "78205 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 148193,
            "unit": "ns/op",
            "extra": "79738 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 26959,
            "unit": "ns/op",
            "extra": "452218 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 24623,
            "unit": "ns/op",
            "extra": "487135 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 318723,
            "unit": "ns/op",
            "extra": "36530 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 162105,
            "unit": "ns/op",
            "extra": "67388 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 33473,
            "unit": "ns/op",
            "extra": "353469 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 33836,
            "unit": "ns/op",
            "extra": "359967 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 314650,
            "unit": "ns/op",
            "extra": "38426 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 145689,
            "unit": "ns/op",
            "extra": "81573 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 23758,
            "unit": "ns/op",
            "extra": "527347 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 23248,
            "unit": "ns/op",
            "extra": "538987 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 22415075,
            "unit": "ns/op",
            "extra": "534 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 22590102,
            "unit": "ns/op",
            "extra": "542 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 27616420,
            "unit": "ns/op",
            "extra": "464 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 24126430,
            "unit": "ns/op",
            "extra": "494 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 21882327,
            "unit": "ns/op",
            "extra": "546 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 22489023,
            "unit": "ns/op",
            "extra": "552 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 24663882,
            "unit": "ns/op",
            "extra": "496 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 24686313,
            "unit": "ns/op",
            "extra": "469 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 16810743,
            "unit": "ns/op",
            "extra": "709 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 15928038,
            "unit": "ns/op",
            "extra": "784 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 16901138,
            "unit": "ns/op",
            "extra": "681 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 15211586,
            "unit": "ns/op",
            "extra": "775 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 16512763,
            "unit": "ns/op",
            "extra": "730 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 16015904,
            "unit": "ns/op",
            "extra": "746 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 16369647,
            "unit": "ns/op",
            "extra": "715 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 16189766,
            "unit": "ns/op",
            "extra": "718 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 10948474,
            "unit": "ns/op",
            "extra": "1104 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 961502,
            "unit": "ns/op",
            "extra": "12424 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 11101118,
            "unit": "ns/op",
            "extra": "1101 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 925678,
            "unit": "ns/op",
            "extra": "12883 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 11185596,
            "unit": "ns/op",
            "extra": "1083 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 707656,
            "unit": "ns/op",
            "extra": "16935 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 11159144,
            "unit": "ns/op",
            "extra": "1087 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 717037,
            "unit": "ns/op",
            "extra": "16623 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 8809731,
            "unit": "ns/op",
            "extra": "1384 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 515797,
            "unit": "ns/op",
            "extra": "23446 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 9042226,
            "unit": "ns/op",
            "extra": "1356 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 620854,
            "unit": "ns/op",
            "extra": "19431 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 9839934,
            "unit": "ns/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 546063,
            "unit": "ns/op",
            "extra": "21940 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 10146761,
            "unit": "ns/op",
            "extra": "1174 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 661713,
            "unit": "ns/op",
            "extra": "18372 times\n4 procs"
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
          "id": "aa8bff12edd6c8e75ab5e030c3d87fa74d688eea",
          "message": "fix: Mailer leaves the transaction open after a failed DATA/BDAT",
          "timestamp": "2026-11-15T14:52:19+01:00",
          "tree_id": "2a901b7e1f85ba996f1255c981d67bdf904a16f6",
          "url": "https://github.com/uponusolutions/go-smtp/commit/aa8bff12edd6c8e75ab5e030c3d87fa74d688eea"
        },
        "date": 1794751339964,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 248099,
            "unit": "ns/op",
            "extra": "46908 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 242573,
            "unit": "ns/op",
            "extra": "49413 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 90506,
            "unit": "ns/op",
            "extra": "132723 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 87611,
            "unit": "ns/op",
            "extra": "137382 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 200721,
            "unit": "ns/op",
            "extra": "59704 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 198636,
            "unit": "ns/op",
            "extra": "61058 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 45475,
            "unit": "ns/op",
            "extra": "268698 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 42856,
            "unit": "ns/op",
            "extra": "276639 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 290413,
            "unit": "ns/op",
            "extra": "41534 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 227221,
            "unit": "ns/op",
            "extra": "52292 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 70432,
            "unit": "ns/op",
            "extra": "168969 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 71026,
            "unit": "ns/op",
            "extra": "169582 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 259382,
            "unit": "ns/op",
            "extra": "46266 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 199947,
            "unit": "ns/op",
            "extra": "59995 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 41367,
            "unit": "ns/op",
            "extra": "289879 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 41052,
            "unit": "ns/op",
            "extra": "275408 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 20685747,
            "unit": "ns/op",
            "extra": "590 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 20894445,
            "unit": "ns/op",
            "extra": "553 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 30528756,
            "unit": "ns/op",
            "extra": "444 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 30859197,
            "unit": "ns/op",
            "extra": "421 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 20622297,
            "unit": "ns/op",
            "extra": "582 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 20356116,
            "unit": "ns/op",
            "extra": "586 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 30572480,
            "unit": "ns/op",
            "extra": "427 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 30501940,
            "unit": "ns/op",
            "extra": "429 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 14974651,
            "unit": "ns/op",
            "extra": "787 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 12446564,
            "unit": "ns/op",
            "extra": "955 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 13237382,
            "unit": "ns/op",
            "extra": "946 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 12590632,
            "unit": "ns/op",
            "extra": "938 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 14631298,
            "unit": "ns/op",
            "extra": "819 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 13018984,
            "unit": "ns/op",
            "extra": "901 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 12921434,
            "unit": "ns/op",
            "extra": "943 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 12446266,
            "unit": "ns/op",
            "extra": "967 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 11972962,
            "unit": "ns/op",
            "extra": "1017 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 814677,
            "unit": "ns/op",
            "extra": "14704 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 11843955,
            "unit": "ns/op",
            "extra": "1026 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 820560,
            "unit": "ns/op",
            "extra": "14674 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 12089390,
            "unit": "ns/op",
            "extra": "994 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 550800,
            "unit": "ns/op",
            "extra": "21949 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 12235385,
            "unit": "ns/op",
            "extra": "994 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 555337,
            "unit": "ns/op",
            "extra": "21630 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 11266345,
            "unit": "ns/op",
            "extra": "1039 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 453614,
            "unit": "ns/op",
            "extra": "26612 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 11190595,
            "unit": "ns/op",
            "extra": "1118 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 521045,
            "unit": "ns/op",
            "extra": "22917 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 11606670,
            "unit": "ns/op",
            "extra": "1081 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 525633,
            "unit": "ns/op",
            "extra": "22594 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 11476628,
            "unit": "ns/op",
            "extra": "1041 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 596925,
            "unit": "ns/op",
            "extra": "20078 times\n4 procs"
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
          "id": "93de3f54e27220436547a82fddb9917abdcc9e0a",
          "message": "fix: support multi-digit enhanced codes",
          "timestamp": "2026-11-15T15:00:05+01:00",
          "tree_id": "9c540e81a6aba5b01c169b15b63f302247e4a4c1",
          "url": "https://github.com/uponusolutions/go-smtp/commit/93de3f54e27220436547a82fddb9917abdcc9e0a"
        },
        "date": 1794751808589,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 246523,
            "unit": "ns/op",
            "extra": "48481 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 242573,
            "unit": "ns/op",
            "extra": "49262 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 90789,
            "unit": "ns/op",
            "extra": "133083 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 89023,
            "unit": "ns/op",
            "extra": "135950 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 207828,
            "unit": "ns/op",
            "extra": "59178 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 204396,
            "unit": "ns/op",
            "extra": "57889 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 45946,
            "unit": "ns/op",
            "extra": "262537 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 44194,
            "unit": "ns/op",
            "extra": "277162 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 306028,
            "unit": "ns/op",
            "extra": "38192 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 232099,
            "unit": "ns/op",
            "extra": "51106 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 71321,
            "unit": "ns/op",
            "extra": "168350 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 71587,
            "unit": "ns/op",
            "extra": "170396 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 261695,
            "unit": "ns/op",
            "extra": "47151 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 198469,
            "unit": "ns/op",
            "extra": "60788 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 41729,
            "unit": "ns/op",
            "extra": "280268 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 42200,
            "unit": "ns/op",
            "extra": "293815 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 20685927,
            "unit": "ns/op",
            "extra": "549 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 20425082,
            "unit": "ns/op",
            "extra": "607 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 30744224,
            "unit": "ns/op",
            "extra": "428 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 30929602,
            "unit": "ns/op",
            "extra": "409 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 20069627,
            "unit": "ns/op",
            "extra": "588 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 20010313,
            "unit": "ns/op",
            "extra": "585 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 31001972,
            "unit": "ns/op",
            "extra": "414 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 31983334,
            "unit": "ns/op",
            "extra": "409 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 16174686,
            "unit": "ns/op",
            "extra": "700 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 12287211,
            "unit": "ns/op",
            "extra": "991 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 12959680,
            "unit": "ns/op",
            "extra": "996 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 11803645,
            "unit": "ns/op",
            "extra": "948 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 15762187,
            "unit": "ns/op",
            "extra": "777 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 12231260,
            "unit": "ns/op",
            "extra": "942 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 12959797,
            "unit": "ns/op",
            "extra": "866 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 12014325,
            "unit": "ns/op",
            "extra": "936 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 12375176,
            "unit": "ns/op",
            "extra": "1004 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 829868,
            "unit": "ns/op",
            "extra": "14432 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 12451449,
            "unit": "ns/op",
            "extra": "957 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 824486,
            "unit": "ns/op",
            "extra": "14557 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 12551302,
            "unit": "ns/op",
            "extra": "976 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 550575,
            "unit": "ns/op",
            "extra": "21808 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 12087583,
            "unit": "ns/op",
            "extra": "997 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 554936,
            "unit": "ns/op",
            "extra": "21344 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 12012849,
            "unit": "ns/op",
            "extra": "1071 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 460488,
            "unit": "ns/op",
            "extra": "26319 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 11127920,
            "unit": "ns/op",
            "extra": "1098 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 516113,
            "unit": "ns/op",
            "extra": "23630 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 11476429,
            "unit": "ns/op",
            "extra": "1015 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 526519,
            "unit": "ns/op",
            "extra": "22609 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 11454046,
            "unit": "ns/op",
            "extra": "1015 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 590813,
            "unit": "ns/op",
            "extra": "20338 times\n4 procs"
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
          "id": "59f508f2306dfe797661692f4099f6b9ff2900da",
          "message": "fix: improve handling of invalid response",
          "timestamp": "2026-11-15T15:07:55+01:00",
          "tree_id": "269c908548aead79d8e1151c42b52fcaf3fd5a93",
          "url": "https://github.com/uponusolutions/go-smtp/commit/59f508f2306dfe797661692f4099f6b9ff2900da"
        },
        "date": 1794752270278,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 170132,
            "unit": "ns/op",
            "extra": "71722 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 163354,
            "unit": "ns/op",
            "extra": "72418 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 43142,
            "unit": "ns/op",
            "extra": "273000 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 40071,
            "unit": "ns/op",
            "extra": "300615 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 150215,
            "unit": "ns/op",
            "extra": "80529 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 150536,
            "unit": "ns/op",
            "extra": "78631 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 27795,
            "unit": "ns/op",
            "extra": "414766 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 24705,
            "unit": "ns/op",
            "extra": "488971 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 286569,
            "unit": "ns/op",
            "extra": "41672 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 161452,
            "unit": "ns/op",
            "extra": "75945 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 32445,
            "unit": "ns/op",
            "extra": "363279 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 32621,
            "unit": "ns/op",
            "extra": "363774 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 276205,
            "unit": "ns/op",
            "extra": "43341 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 140685,
            "unit": "ns/op",
            "extra": "85194 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 21837,
            "unit": "ns/op",
            "extra": "536840 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 22124,
            "unit": "ns/op",
            "extra": "545607 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 27003868,
            "unit": "ns/op",
            "extra": "448 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 26716803,
            "unit": "ns/op",
            "extra": "436 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 27425179,
            "unit": "ns/op",
            "extra": "429 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 27305972,
            "unit": "ns/op",
            "extra": "442 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 28127662,
            "unit": "ns/op",
            "extra": "411 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 27432804,
            "unit": "ns/op",
            "extra": "429 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 27688432,
            "unit": "ns/op",
            "extra": "441 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 29077708,
            "unit": "ns/op",
            "extra": "402 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 26940297,
            "unit": "ns/op",
            "extra": "442 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 21795137,
            "unit": "ns/op",
            "extra": "548 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 20995822,
            "unit": "ns/op",
            "extra": "568 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 21050833,
            "unit": "ns/op",
            "extra": "565 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 27058721,
            "unit": "ns/op",
            "extra": "439 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 21974532,
            "unit": "ns/op",
            "extra": "544 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 20678991,
            "unit": "ns/op",
            "extra": "578 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 21189513,
            "unit": "ns/op",
            "extra": "555 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 11842738,
            "unit": "ns/op",
            "extra": "1009 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 911710,
            "unit": "ns/op",
            "extra": "13152 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 11832540,
            "unit": "ns/op",
            "extra": "1015 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 910696,
            "unit": "ns/op",
            "extra": "13153 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 11969320,
            "unit": "ns/op",
            "extra": "1000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 679853,
            "unit": "ns/op",
            "extra": "17616 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 11999635,
            "unit": "ns/op",
            "extra": "999 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 679781,
            "unit": "ns/op",
            "extra": "17622 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 9369359,
            "unit": "ns/op",
            "extra": "1282 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 534328,
            "unit": "ns/op",
            "extra": "22460 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 9605429,
            "unit": "ns/op",
            "extra": "1244 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 649421,
            "unit": "ns/op",
            "extra": "18418 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 10923978,
            "unit": "ns/op",
            "extra": "1105 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 583666,
            "unit": "ns/op",
            "extra": "20605 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 11053306,
            "unit": "ns/op",
            "extra": "1081 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 687453,
            "unit": "ns/op",
            "extra": "17439 times\n4 procs"
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
          "id": "35b311f8d0ce47ee44ee62c286aa262d1f3abfcc",
          "message": "fix: deadline handling",
          "timestamp": "2026-11-15T16:00:51+01:00",
          "tree_id": "b9dd34556843322a28660e61b0e755b3d186c5b4",
          "url": "https://github.com/uponusolutions/go-smtp/commit/35b311f8d0ce47ee44ee62c286aa262d1f3abfcc"
        },
        "date": 1794755452393,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 399908,
            "unit": "ns/op",
            "extra": "30450 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 388464,
            "unit": "ns/op",
            "extra": "30674 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 122905,
            "unit": "ns/op",
            "extra": "97891 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 118850,
            "unit": "ns/op",
            "extra": "100842 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 333835,
            "unit": "ns/op",
            "extra": "35907 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 335291,
            "unit": "ns/op",
            "extra": "36313 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 64367,
            "unit": "ns/op",
            "extra": "183944 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 61359,
            "unit": "ns/op",
            "extra": "195484 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 476140,
            "unit": "ns/op",
            "extra": "25291 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 364595,
            "unit": "ns/op",
            "extra": "32833 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 96854,
            "unit": "ns/op",
            "extra": "123364 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 96232,
            "unit": "ns/op",
            "extra": "125074 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 436569,
            "unit": "ns/op",
            "extra": "27234 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 326011,
            "unit": "ns/op",
            "extra": "36632 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 57196,
            "unit": "ns/op",
            "extra": "206632 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 57999,
            "unit": "ns/op",
            "extra": "207272 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 35441151,
            "unit": "ns/op",
            "extra": "331 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 35992972,
            "unit": "ns/op",
            "extra": "327 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 46670693,
            "unit": "ns/op",
            "extra": "283 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 41827220,
            "unit": "ns/op",
            "extra": "310 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 36331935,
            "unit": "ns/op",
            "extra": "306 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 34645829,
            "unit": "ns/op",
            "extra": "356 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 46656480,
            "unit": "ns/op",
            "extra": "282 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 47189032,
            "unit": "ns/op",
            "extra": "273 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 26687207,
            "unit": "ns/op",
            "extra": "468 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 23934217,
            "unit": "ns/op",
            "extra": "444 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 21305092,
            "unit": "ns/op",
            "extra": "556 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 21025298,
            "unit": "ns/op",
            "extra": "571 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 26728910,
            "unit": "ns/op",
            "extra": "446 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 23542006,
            "unit": "ns/op",
            "extra": "486 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 21522199,
            "unit": "ns/op",
            "extra": "584 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 20790837,
            "unit": "ns/op",
            "extra": "585 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 22906207,
            "unit": "ns/op",
            "extra": "520 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1151529,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 22817124,
            "unit": "ns/op",
            "extra": "524 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1155294,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 22829093,
            "unit": "ns/op",
            "extra": "524 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 907697,
            "unit": "ns/op",
            "extra": "13291 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 22739716,
            "unit": "ns/op",
            "extra": "526 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 921796,
            "unit": "ns/op",
            "extra": "13083 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 19991696,
            "unit": "ns/op",
            "extra": "600 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 749616,
            "unit": "ns/op",
            "extra": "15972 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 20562710,
            "unit": "ns/op",
            "extra": "585 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 846795,
            "unit": "ns/op",
            "extra": "14158 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 20178908,
            "unit": "ns/op",
            "extra": "595 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 958717,
            "unit": "ns/op",
            "extra": "12510 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 20396289,
            "unit": "ns/op",
            "extra": "578 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1055192,
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
          "id": "1dfa66cbe93dfca50b58a9506ea50f1d6b136d10",
          "message": "chore: cleanup deadline handling",
          "timestamp": "2026-11-15T16:09:25+01:00",
          "tree_id": "1fa77723532fee28963e0eb1a39672a2099abeb2",
          "url": "https://github.com/uponusolutions/go-smtp/commit/1dfa66cbe93dfca50b58a9506ea50f1d6b136d10"
        },
        "date": 1794755965844,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 403904,
            "unit": "ns/op",
            "extra": "29722 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 389671,
            "unit": "ns/op",
            "extra": "30643 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 120769,
            "unit": "ns/op",
            "extra": "98128 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 117109,
            "unit": "ns/op",
            "extra": "102801 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 331568,
            "unit": "ns/op",
            "extra": "36086 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 334832,
            "unit": "ns/op",
            "extra": "35918 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 62866,
            "unit": "ns/op",
            "extra": "191614 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 59937,
            "unit": "ns/op",
            "extra": "198884 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 480171,
            "unit": "ns/op",
            "extra": "24740 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 367302,
            "unit": "ns/op",
            "extra": "32823 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 95500,
            "unit": "ns/op",
            "extra": "123888 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 95726,
            "unit": "ns/op",
            "extra": "125607 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 438884,
            "unit": "ns/op",
            "extra": "27502 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 327320,
            "unit": "ns/op",
            "extra": "36639 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 55983,
            "unit": "ns/op",
            "extra": "209596 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 57750,
            "unit": "ns/op",
            "extra": "210751 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 36104617,
            "unit": "ns/op",
            "extra": "324 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 35451953,
            "unit": "ns/op",
            "extra": "339 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 48124843,
            "unit": "ns/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 46896196,
            "unit": "ns/op",
            "extra": "265 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 36189820,
            "unit": "ns/op",
            "extra": "350 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 35101511,
            "unit": "ns/op",
            "extra": "337 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 48151840,
            "unit": "ns/op",
            "extra": "266 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 47227687,
            "unit": "ns/op",
            "extra": "271 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 24338103,
            "unit": "ns/op",
            "extra": "500 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 23381748,
            "unit": "ns/op",
            "extra": "512 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 22473033,
            "unit": "ns/op",
            "extra": "469 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 21940572,
            "unit": "ns/op",
            "extra": "566 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 27135511,
            "unit": "ns/op",
            "extra": "445 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 23542340,
            "unit": "ns/op",
            "extra": "516 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 23373277,
            "unit": "ns/op",
            "extra": "532 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 20398000,
            "unit": "ns/op",
            "extra": "597 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 22921203,
            "unit": "ns/op",
            "extra": "523 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1162108,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 22935841,
            "unit": "ns/op",
            "extra": "524 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1158701,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 22776714,
            "unit": "ns/op",
            "extra": "522 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 921777,
            "unit": "ns/op",
            "extra": "13029 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 22856109,
            "unit": "ns/op",
            "extra": "526 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 910224,
            "unit": "ns/op",
            "extra": "13140 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 19998264,
            "unit": "ns/op",
            "extra": "598 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 747433,
            "unit": "ns/op",
            "extra": "16032 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 20186820,
            "unit": "ns/op",
            "extra": "595 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 838827,
            "unit": "ns/op",
            "extra": "14304 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 20207022,
            "unit": "ns/op",
            "extra": "594 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 954693,
            "unit": "ns/op",
            "extra": "12560 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 20339170,
            "unit": "ns/op",
            "extra": "590 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1055462,
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
          "id": "bc78d737ccf08cf250d90e58903800e983f23357",
          "message": "chore: remove unused code",
          "timestamp": "2026-11-15T16:12:07+01:00",
          "tree_id": "36ce6f7f5b1c51aa8489bd7e696a559e05738521",
          "url": "https://github.com/uponusolutions/go-smtp/commit/bc78d737ccf08cf250d90e58903800e983f23357"
        },
        "date": 1794756129102,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 242228,
            "unit": "ns/op",
            "extra": "50041 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 238843,
            "unit": "ns/op",
            "extra": "49405 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 86990,
            "unit": "ns/op",
            "extra": "136417 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 86034,
            "unit": "ns/op",
            "extra": "139466 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 197475,
            "unit": "ns/op",
            "extra": "62086 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 199550,
            "unit": "ns/op",
            "extra": "60033 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 42998,
            "unit": "ns/op",
            "extra": "268627 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 41866,
            "unit": "ns/op",
            "extra": "289706 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 280812,
            "unit": "ns/op",
            "extra": "43041 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 228037,
            "unit": "ns/op",
            "extra": "53666 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 70703,
            "unit": "ns/op",
            "extra": "170869 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 70270,
            "unit": "ns/op",
            "extra": "166876 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 261994,
            "unit": "ns/op",
            "extra": "47334 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 197446,
            "unit": "ns/op",
            "extra": "60229 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 39853,
            "unit": "ns/op",
            "extra": "301982 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 39651,
            "unit": "ns/op",
            "extra": "299538 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 20339115,
            "unit": "ns/op",
            "extra": "588 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 20051041,
            "unit": "ns/op",
            "extra": "590 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 30531444,
            "unit": "ns/op",
            "extra": "418 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 30688143,
            "unit": "ns/op",
            "extra": "409 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 20280265,
            "unit": "ns/op",
            "extra": "588 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 20163479,
            "unit": "ns/op",
            "extra": "591 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 26543421,
            "unit": "ns/op",
            "extra": "470 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 30322579,
            "unit": "ns/op",
            "extra": "421 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 15041955,
            "unit": "ns/op",
            "extra": "781 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 12868654,
            "unit": "ns/op",
            "extra": "994 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 13244002,
            "unit": "ns/op",
            "extra": "906 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 12465009,
            "unit": "ns/op",
            "extra": "967 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 14841536,
            "unit": "ns/op",
            "extra": "796 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 13462262,
            "unit": "ns/op",
            "extra": "920 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 12977957,
            "unit": "ns/op",
            "extra": "921 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 12958833,
            "unit": "ns/op",
            "extra": "927 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 12381272,
            "unit": "ns/op",
            "extra": "997 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 826591,
            "unit": "ns/op",
            "extra": "14367 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 12259403,
            "unit": "ns/op",
            "extra": "968 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 827294,
            "unit": "ns/op",
            "extra": "14551 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 12119880,
            "unit": "ns/op",
            "extra": "991 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 555372,
            "unit": "ns/op",
            "extra": "21482 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 12422044,
            "unit": "ns/op",
            "extra": "896 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 551863,
            "unit": "ns/op",
            "extra": "21608 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 10915130,
            "unit": "ns/op",
            "extra": "1136 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 459019,
            "unit": "ns/op",
            "extra": "26407 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 11470772,
            "unit": "ns/op",
            "extra": "976 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 518993,
            "unit": "ns/op",
            "extra": "22963 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 11429942,
            "unit": "ns/op",
            "extra": "1048 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 531856,
            "unit": "ns/op",
            "extra": "22575 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 11390208,
            "unit": "ns/op",
            "extra": "1075 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 586767,
            "unit": "ns/op",
            "extra": "20226 times\n4 procs"
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
          "id": "2fe39fdd437ee8a1ea1c71b8a67b3773ad00641e",
          "message": "feat: add connection limit per ip",
          "timestamp": "2026-11-16T17:32:38+01:00",
          "tree_id": "b5af59b1e6b8ac9dbc15c07ab5edda58e4461000",
          "url": "https://github.com/uponusolutions/go-smtp/commit/2fe39fdd437ee8a1ea1c71b8a67b3773ad00641e"
        },
        "date": 1794847355736,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 395264,
            "unit": "ns/op",
            "extra": "29852 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 386035,
            "unit": "ns/op",
            "extra": "30999 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 120396,
            "unit": "ns/op",
            "extra": "99958 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 115567,
            "unit": "ns/op",
            "extra": "104025 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 335330,
            "unit": "ns/op",
            "extra": "35700 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 330247,
            "unit": "ns/op",
            "extra": "35985 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 62531,
            "unit": "ns/op",
            "extra": "188628 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 58784,
            "unit": "ns/op",
            "extra": "202681 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 458178,
            "unit": "ns/op",
            "extra": "26516 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 362826,
            "unit": "ns/op",
            "extra": "33045 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 92408,
            "unit": "ns/op",
            "extra": "128818 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 92666,
            "unit": "ns/op",
            "extra": "129525 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 440471,
            "unit": "ns/op",
            "extra": "27166 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 329589,
            "unit": "ns/op",
            "extra": "36100 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 54918,
            "unit": "ns/op",
            "extra": "215761 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Small/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 56555,
            "unit": "ns/op",
            "extra": "212845 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 32858813,
            "unit": "ns/op",
            "extra": "366 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 33774873,
            "unit": "ns/op",
            "extra": "380 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 45787387,
            "unit": "ns/op",
            "extra": "273 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 45639814,
            "unit": "ns/op",
            "extra": "282 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 31947794,
            "unit": "ns/op",
            "extra": "381 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 31957565,
            "unit": "ns/op",
            "extra": "374 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 47414834,
            "unit": "ns/op",
            "extra": "252 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/NoChunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 45263607,
            "unit": "ns/op",
            "extra": "280 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 26045002,
            "unit": "ns/op",
            "extra": "456 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 20467425,
            "unit": "ns/op",
            "extra": "570 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 20832103,
            "unit": "ns/op",
            "extra": "583 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/NoPipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 19320752,
            "unit": "ns/op",
            "extra": "574 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 26601602,
            "unit": "ns/op",
            "extra": "446 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/CloseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 20973259,
            "unit": "ns/op",
            "extra": "536 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/MinimalReader (github.com/uponusolutions/go-smtp)",
            "value": 20534601,
            "unit": "ns/op",
            "extra": "562 times\n4 procs"
          },
          {
            "name": "BenchmarkMailer/Large/Chunking/Pipelining/ReuseConn/BufferReader (github.com/uponusolutions/go-smtp)",
            "value": 18175503,
            "unit": "ns/op",
            "extra": "648 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 22775585,
            "unit": "ns/op",
            "extra": "526 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1152582,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 22990366,
            "unit": "ns/op",
            "extra": "498 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1150105,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 22723286,
            "unit": "ns/op",
            "extra": "528 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 922149,
            "unit": "ns/op",
            "extra": "12920 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 23164889,
            "unit": "ns/op",
            "extra": "528 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Read/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 950257,
            "unit": "ns/op",
            "extra": "12576 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 19972440,
            "unit": "ns/op",
            "extra": "598 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 743108,
            "unit": "ns/op",
            "extra": "16125 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 20225180,
            "unit": "ns/op",
            "extra": "594 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Binary/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 839354,
            "unit": "ns/op",
            "extra": "14305 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 20163751,
            "unit": "ns/op",
            "extra": "589 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/MinimalReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 952709,
            "unit": "ns/op",
            "extra": "12609 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Upstream (github.com/uponusolutions/go-smtp/internal)",
            "value": 20463008,
            "unit": "ns/op",
            "extra": "574 times\n4 procs"
          },
          {
            "name": "BenchmarkDot/Write/Text/BufferReader/Fork (github.com/uponusolutions/go-smtp/internal)",
            "value": 1063179,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          }
        ]
      }
    ]
  }
}