```
BenchmarkMapPointers_Set-64                     300000000              181 ns/op              32 B/op          1 allocs/op              
BenchmarkMapRewrite_Set-64                      300000000              173 ns/op              32 B/op          1 allocs/op
BenchmarkMapSlice_Set-64                        300000000              155 ns/op              32 B/op          1 allocs/op
BenchmarkMapSliceSlice_Set-64                   300000000              186 ns/op              32 B/op          1 allocs/op
BenchmarkMapByteBuf_Set-64                      300000000              180 ns/op              32 B/op          1 allocs/op
BenchmarkMapSliceStruct_Set-64                  300000000              165 ns/op              32 B/op          1 allocs/op
BenchmarkMapPointers_Get-64                     300000000               76.0 ns/op             0 B/op          0 allocs/op
BenchmarkMapRewrite_Get-64                      300000000              169 ns/op              32 B/op          1 allocs/op
BenchmarkMapSlice_Get-64                        300000000               80.2 ns/op             0 B/op          0 allocs/op
BenchmarkMapSliceSlice_Get-64                   300000000               81.3 ns/op             0 B/op          0 allocs/op
BenchmarkMapByteBuf_Get-64                      300000000              173 ns/op              32 B/op          1 allocs/op
BenchmarkMapSliceStruct_Get-64                  300000000              174 ns/op              32 B/op          1 allocs/op
BenchmarkMapPointers_Update-64                  300000000              109 ns/op               0 B/op          0 allocs/op
BenchmarkMapRewrite_Update-64                   300000000              157 ns/op               0 B/op          0 allocs/op
BenchmarkMapSlice_Update-64                     300000000              116 ns/op               0 B/op          0 allocs/op
BenchmarkMapSliceSlice_Update-64                300000000              115 ns/op               0 B/op          0 allocs/op
BenchmarkMapByteBuf_Update-64                   300000000              122 ns/op               0 B/op          0 allocs/op
BenchmarkMapSliceStruct_Update-64               300000000              122 ns/op               0 B/op          0 allocs/op
BenchmarkMapPointers_Delete-64                  300000000              129 ns/op               0 B/op          0 allocs/op
BenchmarkMapRewrite_Delete-64                   300000000              104 ns/op               0 B/op          0 allocs/op
BenchmarkMapSlice_Delete-64                     300000000              339 ns/op              37 B/op          0 allocs/op
BenchmarkMapSliceSlice_Delete-64                300000000              197 ns/op              41 B/op          0 allocs/op
BenchmarkMapByteBuf_Delete-64                   300000000              409 ns/op              43 B/op          0 allocs/op
BenchmarkMapSliceStruct_Delete-64               300000000              401 ns/op              37 B/op          0 allocs/op
BenchmarkMapPointers_SetGet-64                  300000000              203 ns/op              32 B/op          1 allocs/op
BenchmarkMapRewrite_SetGet-64                   300000000              282 ns/op              64 B/op          2 allocs/op
BenchmarkMapSlice_SetGet-64                     300000000              212 ns/op              32 B/op          1 allocs/op
BenchmarkMapSliceSlice_SetGet-64                300000000              213 ns/op              32 B/op          1 allocs/op
BenchmarkMapByteBuf_SetGet-64                   300000000              252 ns/op              64 B/op          2 allocs/op
BenchmarkNewMapSliceStruct-64                   300000000              252 ns/op              64 B/op          2 allocs/op
BenchmarkMapPointers_SetDelete-64               300000000              181 ns/op              32 B/op          1 allocs/op
BenchmarkMapRewrite_SetDelete-64                300000000              163 ns/op              32 B/op          1 allocs/op
BenchmarkMapSlice_SetDelete-64                  300000000              381 ns/op              32 B/op          1 allocs/op
BenchmarkMapSliceSlice_SetDelete-64             300000000              204 ns/op              32 B/op          1 allocs/op
BenchmarkMapByteBuf_SetDelete-64                300000000              398 ns/op              37 B/op          1 allocs/op
BenchmarkMapSliceStruct_SetDelete-64            300000000              361 ns/op              32 B/op          1 allocs/op
BenchmarkMapPointers_GetDelete-64               300000000              173 ns/op               0 B/op          0 allocs/op
BenchmarkMapRewrite_GetDelete-64                300000000              200 ns/op              32 B/op          1 allocs/op
BenchmarkMapSlice_GetDelete-64                  300000000              429 ns/op              37 B/op          0 allocs/op
BenchmarkMapSliceSlice_GetDelete-64             300000000              227 ns/op              41 B/op          0 allocs/op
BenchmarkMapByteBuf_GetDelete-64                300000000              491 ns/op              75 B/op          1 allocs/op
BenchmarkMapSliceStruct_GetDelete-64            300000000              473 ns/op              69 B/op          1 allocs/op
```


|Name|Total ms before sweep|Pause ms before sweep|Total ms after sweep|Pause after ms sweep|Mem GB before sweep|Mem GB after sweep|
|----|---------------------|---------------------|--------------------|--------------------|-------------------|------------------|
|A Set MapByteBuf|1075.8|0.313|2|0.785|16.4|0.0|
|A Set MapPointers|2771.2|0.354|1071|0.541|16.4|0.0|
|A Set MapRewrite|1965.5|0.242|1|0.068|20.2|0.0|
|A Set MapSlice|721.4|0.261|1609|0.489|19.3|0.0|
|A Set MapSliceSlice|1314.8|0.615|2236|0.512|18.9|0.0|
|A Set MapSliceStruct|202.5|0.375|1|0.305|17.1|0.0|
|B Get MapByteBuf|1374.2|0.834|3|1.681|16.4|0.0|
|B Get MapPointers|3228.8|0.238|2134|0.526|16.4|0.0|
|B Get MapRewrite|726.1|0.372|1|0.056|20.2|0.0|
|B Get MapSlice|819.3|0.289|1522|0.454|19.3|0.0|
|B Get MapSliceSlice|1251.4|0.286|1639|0.469|18.9|0.0|
|B Get MapSliceStruct|2131.0|0.409|1|0.061|17.1|0.0|
|C Update MapByteBuf|1458.0|0.374|1|0.446|16.4|0.0|
|C Update MapPointers|3186.9|0.295|2029|0.608|16.4|0.0|
|C Update MapRewrite|2056.7|0.320|1|0.067|20.2|0.0|
|C Update MapSlice|683.8|0.303|1770|0.501|19.3|0.0|
|C Update MapSliceSlice|1267.6|0.308|1837|0.977|18.9|0.0|
|C Update MapSliceStruct|151.1|0.343|2|1.437|17.1|0.0|
|D Delete MapByteBuf|237.3|0.821|2|1.344|21.7|0.0|
|D Delete MapPointers|1858.0|0.313|2|0.506|7.4|0.0|
|D Delete MapRewrite|1612.2|0.239|1|0.077|20.2|0.0|
|D Delete MapSlice|1784.1|1.123|1|0.490|15.0|0.0|
|D Delete MapSliceSlice|1196.1|1.243|2|0.968|12.3|0.0|
|D Delete MapSliceStruct|74.2|1.871|2|0.951|21.7|0.0|
|E SetGet MapByteBuf|902.8|1.522|2|1.794|16.4|0.0|
|E SetGet MapPointers|3222.3|0.288|2132|0.493|16.4|0.0|
|E SetGet MapRewrite|1702.5|0.310|2|1.672|20.2|0.0|
|E SetGet MapSlice|1258.8|0.262|1852|0.557|19.3|0.0|
|E SetGet MapSliceSlice|1229.4|0.357|1842|0.517|18.9|0.0|
|E SetGet MapSliceStruct|983.8|0.301|2|1.457|17.1|0.0|
|F SetDelete MapByteBuf|1842.6|0.278|2|0.172|17.0|0.0|
|F SetDelete MapPointers|713.5|0.270|2|0.823|7.4|0.0|
|F SetDelete MapRewrite|1597.0|0.241|1|0.069|20.2|0.0|
|F SetDelete MapSlice|1762.2|0.343|2|0.634|10.3|0.0|
|F SetDelete MapSliceSlice|670.5|0.392|2|0.669|10.0|0.0|
|F SetDelete MapSliceStruct|87.5|0.389|1|0.291|17.0|0.0|
|G GetDelete MapByteBuf|456.5|2.392|3|2.163|21.7|0.0|
|G GetDelete MapPointers|1812.5|0.291|2|0.731|7.4|0.0|
|G GetDelete MapRewrite|1720.7|0.289|3|1.647|20.2|0.0|
|G GetDelete MapSlice|1999.7|1.046|1|0.445|15.0|0.0|
|G GetDelete MapSliceSlice|888.7|2.638|1|0.439|12.3|0.0|
|G GetDelete MapSliceStruct|782.8|1.813|3|1.734|21.7|0.0|
