# gosmartmap

**gosmartmap** is a Golang library that provides a smart implementation of a map. 
It utilizes size hints to dynamically estimate the optimal size for map creation, reducing memory allocation overhead and potentially improving runtime performance for map-heavy applications.


## Benchmarks against builtin map[TKey]TValue

|                        |simple-map.txt       |      |smart-map.txt        |      |           |            |
|--------------------    |---------------------|------|---------------------|------|-------    |------------|
|                        |sec/op               |CI    |sec/op               |CI    |vs base    |P           |
|Map/smartmap-0-8        |552.6n               |3%    |1033.0n              |3%    |+86.92%    |p=0.000 n=10|
|Map/smartmap-128-8      |5.749µ               |2%    |5.844µ               |2%    |~          |p=0.066 n=10|
|**Map/smartmap-1024-8** |41.50µ               |3%    |35.70µ               |2%    |**-13.98%**|p=0.000 n=10|
|**Map/smartmap-65536-8**|3.147m               |8%    |2.801m               |9%    |**-11.00%**|p=0.001 n=10|


|                        |simple-map.txt       |      |smart-map.txt        |      |           |            |
|--------------------    |---------------------|------|---------------------|------|-------    |------------|
|                        |B/op                 |CI    |B/op                 |CI    |vs base    |P           |
|Map/smartmap-0-8        |48.00                |0%    |112.00               |0%    |+133.33%   |p=0.000 n=10|
|**Map/smartmap-128-8**  |8.875Ki              |0%    |8.191Ki              |0%    |**-7.71%** |p=0.000 n=10|
|**Map/smartmap-1024-8** |74.55Ki              |1%    |65.41Ki              |1%    |**-12.27%**|p=0.000 n=10|
|**Map/smartmap-65536-8**|4.482Mi              |5%    |3.970Mi              |5%    |**-11.43%**|p=0.000 n=10|

|                        |simple-map.txt       |      |smart-map.txt        |      |            |            |
|--------------------    |---------------------|------|---------------------|------|-------     |------------|
|                        |allocs/op            |CI    |allocs/op            |CI    |vs base     |P           |
|Map/smartmap-0-8        |1                    |0%    |3                    |0%    |+200.00%    |p=0.000 n=10|
|Map/smartmap-128-8      |8                    |0%    |8                    |0%    |~           |p=1.000 n=10|
|**Map/smartmap-1024-8** |23                   |0%    |15                   |0%    |**-34.78%** |p=0.000 n=10|
|**Map/smartmap-65536-8**|1130                 |6%    |954                  |5%    |**-15.58%** |p=0.000 n=10|
