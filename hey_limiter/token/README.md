$  hey -c 10 -n 50 http://localhost:9090
//hey命令-c表示并发数，我设为10，-n表示总共发送多少条，我发50条。
Summary:
  Total:        5.0309 secs
  Slowest:      1.0080 secs
  Fastest:      0.0016 secs
  Average:      0.1034 secs
  Requests/sec: 9.9385


Response time histogram:
  0.002 [1]     |■
  0.102 [44]    |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.203 [0]     |
  0.304 [0]     |
  0.404 [0]     |
  0.505 [0]     |
  0.605 [0]     |
  0.706 [0]     |
  0.807 [0]     |
  0.907 [0]     |
  1.008 [5]     |■■■■■


Latency distribution:
  10% in 0.0020 secs
  25% in 0.0023 secs
  50% in 0.0027 secs
  75% in 0.0054 secs
  90% in 1.0051 secs
  95% in 1.0066 secs
  0% in 0.0000 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0024 secs, 0.0016 secs, 1.0080 secs
  DNS-lookup:   0.0015 secs, 0.0007 secs, 0.0031 secs
  req write:    0.0001 secs, 0.0000 secs, 0.0014 secs
  resp wait:    0.1042 secs, 0.0000 secs, 1.0043 secs
  resp read:    0.0001 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [200] 5 responses
  [404] 45 responses



