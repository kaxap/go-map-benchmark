```
BenchmarkMapPointers_Set-32                     300000000              222 ns/op              32 B/op          1 allocs/op
BenchmarkMapRewrite_Set-32                      300000000              247 ns/op              32 B/op          1 allocs/op
BenchmarkMapSlice_Set-32                        300000000              276 ns/op              32 B/op          1 allocs/op
BenchmarkMapSliceSlice_Set-32                   300000000              253 ns/op              32 B/op          1 allocs/op
BenchmarkMapByteBuf_Set-32                      300000000              303 ns/op              32 B/op          1 allocs/op
BenchmarkMapSliceStruct_Set-32                  300000000              242 ns/op              32 B/op          1 allocs/op
BenchmarkMapPointers_Get-32                     300000000               91.2 ns/op             0 B/op          0 allocs/op
BenchmarkMapRewrite_Get-32                      300000000              225 ns/op              32 B/op          1 allocs/op
BenchmarkMapSlice_Get-32                        300000000               89.0 ns/op             0 B/op          0 allocs/op
BenchmarkMapSliceSlice_Get-32                   300000000               97.5 ns/op             0 B/op          0 allocs/op
BenchmarkMapByteBuf_Get-32                      300000000              174 ns/op              32 B/op          1 allocs/op
BenchmarkMapSliceStruct_Get-32                  300000000              224 ns/op              32 B/op          1 allocs/op
BenchmarkMapPointers_Update-32                  300000000              118 ns/op               0 B/op          0 allocs/op
BenchmarkMapRewrite_Update-32                   300000000              192 ns/op               0 B/op          0 allocs/op
BenchmarkMapSlice_Update-32                     300000000              138 ns/op               0 B/op          0 allocs/op
BenchmarkMapSliceSlice_Update-32                300000000              123 ns/op               0 B/op          0 allocs/op
BenchmarkMapByteBuf_Update-32                   300000000              165 ns/op               0 B/op          0 allocs/op
BenchmarkMapSliceStruct_Update-32               300000000               87.0 ns/op             0 B/op          0 allocs/op
BenchmarkMapPointers_Delete-32                  300000000              127 ns/op               0 B/op          0 allocs/op
BenchmarkMapRewrite_Delete-32                   300000000              196 ns/op               0 B/op          0 allocs/op
BenchmarkMapSlice_Delete-32                     300000000              466 ns/op              37 B/op          0 allocs/op
BenchmarkMapSliceSlice_Delete-32                300000000              227 ns/op              41 B/op          0 allocs/op
BenchmarkMapByteBuf_Delete-32                   300000000              449 ns/op              43 B/op          0 allocs/op
BenchmarkMapSliceStruct_Delete-32               300000000              397 ns/op              37 B/op          0 allocs/op
BenchmarkMapPointers_SetGet-32                  300000000              199 ns/op              32 B/op          1 allocs/op
BenchmarkMapRewrite_SetGet-32                   300000000              379 ns/op              64 B/op          2 allocs/op
BenchmarkMapSlice_SetGet-32                     300000000              252 ns/op              32 B/op          1 allocs/op
BenchmarkMapSliceSlice_SetGet-32                300000000              237 ns/op              32 B/op          1 allocs/op
BenchmarkMapByteBuf_SetGet-32                   300000000              294 ns/op              64 B/op          2 allocs/op
BenchmarkNewMapSliceStruct-32                   300000000              278 ns/op              64 B/op          2 allocs/op
BenchmarkMapPointers_SetDelete-32               300000000              233 ns/op              32 B/op          1 allocs/op
BenchmarkMapRewrite_SetDelete-32                300000000              209 ns/op              32 B/op          1 allocs/op
BenchmarkMapSlice_SetDelete-32                  300000000              537 ns/op              32 B/op          1 allocs/op
BenchmarkMapSliceSlice_SetDelete-32             300000000              193 ns/op              32 B/op          1 allocs/op
BenchmarkMapByteBuf_SetDelete-32                300000000              523 ns/op              37 B/op          1 allocs/op
BenchmarkMapSliceStruct_SetDelete-32            300000000              533 ns/op              32 B/op          1 allocs/op
BenchmarkMapPointers_GetDelete-32               300000000              204 ns/op               0 B/op          0 allocs/op
BenchmarkMapRewrite_GetDelete-32                300000000              218 ns/op              32 B/op          1 allocs/op
BenchmarkMapSlice_GetDelete-32                  300000000              465 ns/op              37 B/op          0 allocs/op
BenchmarkMapSliceSlice_GetDelete-32             300000000              249 ns/op              41 B/op          0 allocs/op
BenchmarkMapByteBuf_GetDelete-32                300000000              542 ns/op              75 B/op          1 allocs/op
BenchmarkMapSliceStruct_GetDelete-32            300000000              542 ns/op              69 B/op          1 allocs/op
```

|Name|Total ms before sweep|Pause ms before sweep|Total ms after sweep|Pause after ms sweep|Mem GB before sweep|Mem GB after sweep|
|----|---------------------|---------------------|--------------------|--------------------|-------------------|------------------|
|A Set MapByteBuf|1380.9|0.206|2|0.311|16.4|0.0|
|A Set MapPointers|78612.5|0.267|1626|0.028|16.4|0.0|
|A Set MapRewrite|1028.1|0.693|1|0.032|20.2|0.0|
|A Set MapSlice|4867.7|0.259|1780|0.367|19.3|0.0|
|A Set MapSliceSlice|1562.1|0.188|939|0.353|18.9|0.0|
|A Set MapSliceStruct|190.6|0.252|3|1.643|17.1|0.0|
|B Get MapByteBuf|1426.2|0.589|2|1.704|16.4|0.0|
|B Get MapPointers|5529.8|0.180|1519|0.455|16.4|0.0|
|B Get MapRewrite|1522.2|0.214|1|0.025|20.2|0.0|
|B Get MapSlice|1569.0|0.242|1802|0.413|19.3|0.0|
|B Get MapSliceSlice|1426.4|0.233|1808|0.471|18.9|0.0|
|B Get MapSliceStruct|1813.3|0.265|2|1.137|17.1|0.0|
|C Update MapByteBuf|1341.4|0.213|1|0.334|16.4|0.0|
|C Update MapPointers|5583.5|0.205|1715|0.709|16.4|0.0|
|C Update MapRewrite|87.6|0.300|3|2.486|20.2|0.0|
|C Update MapSlice|1560.8|0.202|1525|0.379|19.3|0.0|
|C Update MapSliceSlice|1558.9|0.203|1863|0.803|18.9|0.0|
|C Update MapSliceStruct|246.6|0.253|4|3.031|17.1|0.0|
|D Delete MapByteBuf|274.7|0.530|2|1.407|21.7|0.0|
|D Delete MapPointers|1854.7|0.246|1|0.355|7.4|0.0|
|D Delete MapRewrite|73.6|0.203|3|2.076|20.2|0.0|
|D Delete MapSlice|2446.9|0.862|1|0.486|15.0|0.0|
|D Delete MapSliceSlice|1074.8|0.815|1|0.379|12.3|0.0|
|D Delete MapSliceStruct|156.4|1.532|1|0.504|21.7|0.0|
|E SetGet MapByteBuf|1380.2|0.561|2|1.384|16.4|0.0|
|E SetGet MapPointers|5533.9|0.182|1515|0.317|16.4|0.0|
|E SetGet MapRewrite|1632.1|0.210|2|1.389|20.2|0.0|
|E SetGet MapSlice|1505.6|0.187|1659|0.355|19.3|0.0|
|E SetGet MapSliceSlice|1464.8|0.465|1943|0.387|18.9|0.0|
|E SetGet MapSliceStruct|1616.2|0.226|2|1.149|17.1|0.0|
|F SetDelete MapByteBuf|215.7|0.425|2|1.347|17.0|0.0|
|F SetDelete MapPointers|1729.4|0.220|1|0.567|7.4|0.0|
|F SetDelete MapRewrite|1799.5|0.149|1|0.026|20.2|0.0|
|F SetDelete MapSlice|2274.7|0.235|2|0.695|10.3|0.0|
|F SetDelete MapSliceSlice|1190.7|0.179|2|1.098|10.0|0.0|
|F SetDelete MapSliceStruct|464.1|0.298|2|1.216|17.0|0.0|
|G GetDelete MapByteBuf|802.1|1.841|2|1.427|21.7|0.0|
|G GetDelete MapPointers|2180.7|0.239|1|0.895|7.4|0.0|
|G GetDelete MapRewrite|1362.0|0.188|3|2.218|20.2|0.0|
|G GetDelete MapSlice|2117.2|0.741|9|0.788|15.0|0.0|
|G GetDelete MapSliceSlice|1499.4|1.362|1|0.421|12.3|0.0|
|G GetDelete MapSliceStruct|626.8|1.455|5|3.464|21.7|0.0|
