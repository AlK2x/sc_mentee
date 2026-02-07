./heap_only.go:61:10: inlining call to main.(*LargeObject).DataUpdaterFunc.func1
./heap_only.go:59:12: inlining call to (*LargeObject).generateRundomData
./heap_only.go:12:33: make([]byte, lo.Size) escapes to heap in NewLargeObject:
./heap_only.go:12:33:   flow: data ← &{storage for make([]byte, lo.Size)}:
./heap_only.go:12:33:     from make([]byte, lo.Size) (spill) at ./heap_only.go:12:33
--
./heap_only.go:12:33:   flow: {heap} ← ~r0:
./heap_only.go:12:33:     from lo.Data = ~r0 (assign) at ./heap_only.go:12:10
./heap_only.go:9:8: &LargeObject{...} escapes to heap in NewLargeObject:
./heap_only.go:9:8:   flow: lo ← &{storage for &LargeObject{...}}:
./heap_only.go:9:8:     from &LargeObject{...} (spill) at ./heap_only.go:9:8
--
./heap_only.go:9:8:   flow: ~r0 ← lo:
./heap_only.go:9:8:     from return lo (return) at ./heap_only.go:14:2
./heap_only.go:9:8: &LargeObject{...} escapes to heap
./heap_only.go:12:33: make([]byte, lo.Size) escapes to heap
./heap_only.go:22:7: lo does not escape
./heap_only.go:23:34: make([]byte, lo.Size) does not escape
./heap_only.go:27:7: (*LargeObject).DataUpdaterFunc capturing by value: lo (addr=false assign=false width=8)
./heap_only.go:27:40: (*LargeObject).DataUpdaterFunc capturing by value: val (addr=false assign=false width=1)
./heap_only.go:28:13: func literal escapes to heap in (*LargeObject).DataUpdaterFunc:
./heap_only.go:28:13:   flow: resetFn ← &{storage for func literal}:
./heap_only.go:28:13:     from func literal (spill) at ./heap_only.go:28:13
--
./heap_only.go:27:40:     from val (captured by a closure) at ./heap_only.go:30:17
./heap_only.go:27:7: leaking param: lo
./heap_only.go:28:13: func literal escapes to heap
./heap_only.go:37:14: make([]byte, lo.Size) escapes to heap in (*LargeObject).generateRundomData:
./heap_only.go:37:14:   flow: data ← &{storage for make([]byte, lo.Size)}:
./heap_only.go:37:14:     from make([]byte, lo.Size) (spill) at ./heap_only.go:37:14
--
./heap_only.go:37:14:   flow: ~r0 ← data:
./heap_only.go:37:14:     from return data (return) at ./heap_only.go:39:2
./heap_only.go:36:7: lo does not escape
./heap_only.go:37:14: make([]byte, lo.Size) escapes to heap
./heap_only.go:48:24: append(largeObjects, obj) escapes to heap in fillGlobalObjects:
./heap_only.go:48:24:   flow: {heap} ← &{storage for append(largeObjects, obj)}:
./heap_only.go:48:24:     from append(largeObjects, obj) (spill) at ./heap_only.go:48:24
./heap_only.go:48:24:     from largeObjects = append(largeObjects, obj) (assign) at ./heap_only.go:48:16
./heap_only.go:48:24: append escapes to heap
./heap_only.go:54:15: make(map[int]*LargeObject) escapes to heap in main:
./heap_only.go:54:15:   flow: {heap} ← &{storage for make(map[int]*LargeObject)}:
./heap_only.go:54:15:     from make(map[int]*LargeObject) (spill) at ./heap_only.go:54:15
./heap_only.go:54:15:     from bigMap = make(map[int]*LargeObject) (assign) at ./heap_only.go:54:9
./heap_only.go:55:21: make([]*LargeObject, 0) escapes to heap in main:
./heap_only.go:55:21:   flow: {heap} ← &{storage for make([]*LargeObject, 0)}:
./heap_only.go:55:21:     from make([]*LargeObject, 0) (spill) at ./heap_only.go:55:21
./heap_only.go:55:21:     from largeObjects = make([]*LargeObject, 0) (assign) at ./heap_only.go:55:15
./heap_only.go:56:19: append(largeObjects, obj) escapes to heap in main:
./heap_only.go:56:19:   flow: {heap} ← &{storage for append(largeObjects, obj)}:
./heap_only.go:56:19:     from append(largeObjects, obj) (spill) at ./heap_only.go:56:19
--
./heap_only.go:60:33: main capturing by value: lo (addr=false assign=false width=8)
./heap_only.go:60:33: main capturing by value: val (addr=false assign=false width=1)
./heap_only.go:54:15: make(map[int]*LargeObject) escapes to heap
./heap_only.go:55:21: make([]*LargeObject, 0) escapes to heap
./heap_only.go:56:19: append escapes to heap
./heap_only.go:59:12: make([]byte, lo.Size) does not escape
./heap_only.go:60:33: func literal does not escape


make([]byte, lo.Size) escapes to heap in NewLargeObject: - размер может меняться
&LargeObject{...} escapes to heap - время жизни больше времени жизни функции
28:13: func literal escapes to heap - замыкание
