# gosmartmap

**gosmartmap** is a Golang library that provides a smart implementation of a map. 
It utilizes size hints to dynamically estimate the optimal size for map creation, reducing memory allocation overhead and potentially improving runtime performance for map-heavy applications.


## Some benchmarks (that you should not trust)

|     | simple-map (sec / op) | smart-map (sec / op) | delta |
| --- | --- | ----------- | --- |
| Map/map-128-8   | 7.817µ ± 11% | 4.520µ ± 1% | -42.18% (p=0.000 n=10) |
| Map/map-1024-8  | 57.49µ ± 9% | 40.88µ ± 6% | -28.89% (p=0.000 n=10) |
| Map/map-65536-8 | 3.686m ± 6% | 3.570m ± 5% | ~ (p=0.165 n=10) |
| geomean         | 118.3µ | 87.05µ | -26.43% |

|     | simple-map (B / op) | smart-map (B / op) | delta |
| --- | --- | ----------- | --- |
| Map/map-128-8   | 8.784Ki ± 0% | 9.578Ki ± 0% |  +9.04% (p=0.000 n=10) |
| Map/map-1024-8  | 74.46Ki ± 1% | 65.38Ki ± 5% | -12.20% (p=0.000 n=10) |
| Map/map-65536-8 | 4.496Mi ± 3% | 4.416Mi ± 3% | ~ (p=0.165 n=10) |
| geomean         | 144.4Ki | 141.5Ki | -2.03% |

|     | simple-map (allocs / op) | smart-map (allocs / op) | delta |
| --- | --- | ----------- | --- |
| Map/map-128-8   | 8.000 ± 0% | 7.000 ±  0% |  -12.50% (p=0.000 n=10) |
| Map/map-1024-8  | 23.00 ± 0% | 17.00 ± 12% | -26.09% (p=0.000 n=10) |
| Map/map-65536-8 | 1.125k ± 5% | 1.117k ±  3% | ~ (p=0.591 n=10) |
| geomean         | 59.16 | 51.03 | -2.03% |
