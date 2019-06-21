```
BenchmarkMapPointers_Set                300000000              332 ns/op              32 B/op          1 allocs/op
BenchmarkMapRewrite_Set                 300000000              527 ns/op              32 B/op          1 allocs/op
BenchmarkMapSlice_Set                   300000000              320 ns/op              32 B/op          1 allocs/op
BenchmarkMapSliceSlice_Set              300000000              265 ns/op              32 B/op          1 allocs/op
BenchmarkMapByteBuf_Set                 300000000              356 ns/op              32 B/op          1 allocs/op
BenchmarkMapSliceStruct_Set             300000000              291 ns/op              32 B/op          1 allocs/op
BenchmarkMapPointers_Get                300000000              116 ns/op               0 B/op          0 allocs/op
BenchmarkMapRewrite_Get                 300000000              317 ns/op              32 B/op          1 allocs/op
BenchmarkMapSlice_Get                   300000000              117 ns/op               0 B/op          0 allocs/op
BenchmarkMapSliceSlice_Get              300000000              120 ns/op               0 B/op          0 allocs/op
BenchmarkMapByteBuf_Get                 300000000              301 ns/op              32 B/op          1 allocs/op
BenchmarkMapSliceStruct_Get             300000000              231 ns/op              32 B/op          1 allocs/op
BenchmarkMapPointers_Update             300000000              336 ns/op               0 B/op          0 allocs/op
BenchmarkMapRewrite_Update              300000000              374 ns/op               0 B/op          0 allocs/op
BenchmarkMapSlice_Update                300000000              179 ns/op               0 B/op          0 allocs/op
BenchmarkMapSliceSlice_Update           300000000              178 ns/op               0 B/op          0 allocs/op
BenchmarkMapByteBuf_Update              300000000              244 ns/op               0 B/op          0 allocs/op
BenchmarkMapSliceStruct_Update          300000000              114 ns/op               0 B/op          0 allocs/op
BenchmarkMapPointers_Delete             300000000              168 ns/op               0 B/op          0 allocs/op
BenchmarkMapRewrite_Delete              300000000              211 ns/op               0 B/op          0 allocs/op
BenchmarkMapSlice_Delete                300000000              703 ns/op              37 B/op          0 allocs/op
BenchmarkMapSliceSlice_Delete           300000000              248 ns/op              41 B/op          0 allocs/op
BenchmarkMapByteBuf_Delete              300000000              779 ns/op              43 B/op          0 allocs/op
BenchmarkMapSliceStruct_Delete          300000000              689 ns/op              37 B/op          0 allocs/op
BenchmarkMapPointers_SetGet             300000000              357 ns/op              32 B/op          1 allocs/op
BenchmarkMapRewrite_SetGet              300000000              427 ns/op              64 B/op          2 allocs/op
BenchmarkMapSlice_SetGet                300000000              270 ns/op              32 B/op          1 allocs/op
BenchmarkMapSliceSlice_SetGet           300000000              285 ns/op              32 B/op          1 allocs/op
BenchmarkMapByteBuf_SetGet              300000000              394 ns/op              64 B/op          2 allocs/op                      
BenchmarkNewMapSliceStruct              300000000              320 ns/op              64 B/op          2 allocs/op
BenchmarkMapPointers_SetDelete          300000000              414 ns/op              32 B/op          1 allocs/op
BenchmarkMapRewrite_SetDelete           300000000              312 ns/op              32 B/op          1 allocs/op
BenchmarkMapSlice_SetDelete             300000000              595 ns/op              32 B/op          1 allocs/op
BenchmarkMapSliceSlice_SetDelete        300000000              257 ns/op              32 B/op          1 allocs/op
BenchmarkMapByteBuf_SetDelete           300000000              773 ns/op              37 B/op          1 allocs/op
BenchmarkMapSliceStruct_SetDelete       300000000              590 ns/op              32 B/op          1 allocs/op
BenchmarkMapPointers_GetDelete          300000000              213 ns/op               0 B/op          0 allocs/op
BenchmarkMapRewrite_GetDelete           300000000              352 ns/op              32 B/op          1 allocs/op
BenchmarkMapSlice_GetDelete             300000000              742 ns/op              37 B/op          0 allocs/op
BenchmarkMapSliceSlice_GetDelete        300000000              282 ns/op              41 B/op          0 allocs/op
BenchmarkMapByteBuf_GetDelete           300000000              939 ns/op              75 B/op          1 allocs/op
BenchmarkMapSliceStruct_GetDelete       300000000              801 ns/op              69 B/op          1 allocs/op
```


|Name|Total ms before sweep|Pause ms before sweep|Total ms after sweep|Pause after ms sweep|Mem GB before sweep|Mem GB after sweep|
|----|---------------------|---------------------|--------------------|--------------------|-------------------|------------------|
|A Set MapByteBuf|448.8|0.058|1|0.208|16.4|0.0|
|A Set MapPointers|81738.3|0.058|542|0.006|16.4|0.0|
|A Set MapRewrite|528.9|0.055|1|0.599|20.2|0.0|
|A Set MapSlice|9303.7|0.080|732|0.898|19.3|0.0|
|A Set MapSliceSlice|8990.7|0.050|550|0.074|18.9|0.0|
|A Set MapSliceStruct|456.2|0.061|1|0.256|17.1|0.0|
|B Get MapByteBuf|686.6|0.066|1|0.007|16.4|0.0|
|B Get MapPointers|83541.9|0.044|675|0.011|16.4|0.0|
|B Get MapRewrite|583.4|0.052|2|1.143|20.2|0.0|
|B Get MapSlice|9295.2|0.068|558|0.482|19.3|0.0|
|B Get MapSliceSlice|8954.5|0.082|580|0.549|18.9|0.0|
|B Get MapSliceStruct|764.4|0.071|0|0.007|17.1|0.0|
|C Update MapByteBuf|423.4|0.063|1|0.329|16.4|0.0|
|C Update MapPointers|85956.2|0.058|595|1.487|16.4|0.0|
|C Update MapRewrite|96.5|0.460|0|0.089|20.2|0.0|
|C Update MapSlice|9383.8|0.059|671|0.827|19.3|0.0|
|C Update MapSliceSlice|9549.0|0.060|574|0.427|18.9|0.0|
|C Update MapSliceStruct|362.4|0.076|1|0.559|17.1|0.0|
|D Delete MapByteBuf|118.7|1.739|0|0.082|21.7|0.0|
|D Delete MapPointers|6213.6|0.056|1|0.994|7.4|0.0|
|D Delete MapRewrite|516.9|0.054|1|0.195|20.2|0.0|
|D Delete MapSlice|1904.1|0.560|1|0.483|15.0|0.0|
|D Delete MapSliceSlice|1392.1|0.488|1|0.327|12.3|0.0|
|D Delete MapSliceStruct|123.3|1.033|1|0.432|21.7|0.0|
|E SetGet MapByteBuf|671.8|0.059|1|0.657|16.4|0.0|
|E SetGet MapPointers|84808.7|0.057|540|0.984|16.4|0.0|
|E SetGet MapRewrite|549.9|0.047|2|1.424|20.2|0.0|
|E SetGet MapSlice|9408.2|0.058|576|0.450|19.3|0.0|
|E SetGet MapSliceSlice|9425.4|0.054|581|0.730|18.9|0.0|
|E SetGet MapSliceStruct|714.8|0.063|2|1.365|17.1|0.0|
|F SetDelete MapByteBuf|357.6|0.070|1|0.701|17.0|0.0|
|F SetDelete MapPointers|6930.9|0.048|1|0.381|7.4|0.0|
|F SetDelete MapRewrite|524.5|0.040|0|0.006|20.2|0.0|
|F SetDelete MapSlice|1673.1|0.058|1|0.470|10.3|0.0|
|F SetDelete MapSliceSlice|1745.9|0.055|1|0.432|10.0|0.0|
|F SetDelete MapSliceStruct|249.0|0.070|1|0.594|17.0|0.0|
|G GetDelete MapByteBuf|380.7|1.828|1|0.233|21.7|0.0|
|G GetDelete MapPointers|6621.3|0.092|2|1.441|7.4|0.0|
|G GetDelete MapRewrite|622.4|0.048|2|1.454|20.2|0.0|
|G GetDelete MapSlice|1928.7|0.553|1|0.614|15.0|0.0|
|G GetDelete MapSliceSlice|1862.1|0.598|1|0.388|12.3|0.0|
|G GetDelete MapSliceStruct|511.7|2.017|1|0.284|21.7|0.0|
