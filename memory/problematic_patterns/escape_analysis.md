./problematic_patterns.go:77:20: inlining call to (*ConstantGrow).Grow
./problematic_patterns.go:78:34: inlining call to (*ConstantGrow).HeapAllocation
./problematic_patterns.go:31:9: &ConstantGrow{...} escapes to heap in NewConstantGrow:
./problematic_patterns.go:31:9:   flow: ~r0 ← &{storage for &ConstantGrow{...}}:
./problematic_patterns.go:31:9:     from &ConstantGrow{...} (spill) at ./problematic_patterns.go:31:9
./problematic_patterns.go:31:9:     from return &ConstantGrow{...} (return) at ./problematic_patterns.go:31:2
./problematic_patterns.go:32:10: make(map[int]*Foo) escapes to heap in NewConstantGrow:
./problematic_patterns.go:32:10:   flow: {storage for &ConstantGrow{...}} ← &{storage for make(map[int]*Foo)}:
./problematic_patterns.go:32:10:     from make(map[int]*Foo) (spill) at ./problematic_patterns.go:32:10
./problematic_patterns.go:32:10:     from ConstantGrow{...} (struct literal element) at ./problematic_patterns.go:31:22
./problematic_patterns.go:33:10: make([]*Bar, 0) escapes to heap in NewConstantGrow:
./problematic_patterns.go:33:10:   flow: {storage for &ConstantGrow{...}} ← &{storage for make([]*Bar, 0)}:
./problematic_patterns.go:33:10:     from make([]*Bar, 0) (spill) at ./problematic_patterns.go:33:10
./problematic_patterns.go:33:10:     from ConstantGrow{...} (struct literal element) at ./problematic_patterns.go:31:22
./problematic_patterns.go:31:9: &ConstantGrow{...} escapes to heap
./problematic_patterns.go:32:10: make(map[int]*Foo) escapes to heap
./problematic_patterns.go:33:10: make([]*Bar, 0) escapes to heap
./problematic_patterns.go:51:16: append(cg.s, f.bar) escapes to heap in (*ConstantGrow).Grow:
./problematic_patterns.go:51:16:   flow: {heap} ← &{storage for append(cg.s, f.bar)}:
./problematic_patterns.go:51:16:     from append(cg.s, f.bar) (spill) at ./problematic_patterns.go:51:16
./problematic_patterns.go:51:16:     from cg.s = append(cg.s, f.bar) (assign) at ./problematic_patterns.go:51:8
./problematic_patterns.go:42:8: &Foo{...} escapes to heap in (*ConstantGrow).Grow:
./problematic_patterns.go:42:8:   flow: f ← &{storage for &Foo{...}}:
./problematic_patterns.go:42:8:     from &Foo{...} (spill) at ./problematic_patterns.go:42:8
--
./problematic_patterns.go:38:7:   flow: {heap} ← {temp}:
./problematic_patterns.go:38:7:     from cg.s = append(cg.s, f.bar) (assign) at ./problematic_patterns.go:51:8
./problematic_patterns.go:45:9: &Bar{...} escapes to heap in (*ConstantGrow).Grow:
./problematic_patterns.go:45:9:   flow: {storage for &Foo{...}} ← &{storage for &Bar{...}}:
./problematic_patterns.go:45:9:     from &Bar{...} (spill) at ./problematic_patterns.go:45:9
./problematic_patterns.go:45:9:     from Foo{...} (struct literal element) at ./problematic_patterns.go:42:12
./problematic_patterns.go:38:7: leaking param content: cg
./problematic_patterns.go:42:8: &Foo{...} escapes to heap
./problematic_patterns.go:45:9: &Bar{...} escapes to heap
./problematic_patterns.go:51:16: append escapes to heap
./problematic_patterns.go:57:10: &Stat{} escapes to heap in (*ConstantGrow).HeapAllocation:
./problematic_patterns.go:57:10:   flow: stat ← &{storage for &Stat{}}:
./problematic_patterns.go:57:10:     from &Stat{} (spill) at ./problematic_patterns.go:57:10
--
./problematic_patterns.go:57:10:   flow: ~r0 ← stat:
./problematic_patterns.go:57:10:     from return stat (return) at ./problematic_patterns.go:60:2
./problematic_patterns.go:56:7: cg does not escape
./problematic_patterns.go:57:10: &Stat{} escapes to heap
./problematic_patterns.go:65:5: func literal escapes to heap in rundomNumberGenerator:
./problematic_patterns.go:65:5:   flow: {heap} ← &{storage for func literal}:
./problematic_patterns.go:65:5:     from func literal (spill) at ./problematic_patterns.go:65:5
./problematic_patterns.go:64:2: rundomNumberGenerator capturing by value: ch (addr=false assign=false width=8)
./problematic_patterns.go:66:3: rundomNumberGenerator.func1 capturing by value: .autotmp_0 (addr=false assign=false width=8)
./problematic_patterns.go:65:5: func literal escapes to heap
./problematic_patterns.go:77:20: append(cg.s, f.bar) escapes to heap in main:
./problematic_patterns.go:77:20:   flow: {heap} ← &{storage for append(cg.s, f.bar)}:
./problematic_patterns.go:77:20:     from append(cg.s, f.bar) (spill) at ./problematic_patterns.go:77:20
./problematic_patterns.go:77:20:     from cg.s = append(cg.s, f.bar) (assign) at ./problematic_patterns.go:77:20
./problematic_patterns.go:77:20: &Foo{...} escapes to heap in main:
./problematic_patterns.go:77:20:   flow: f ← &{storage for &Foo{...}}:
./problematic_patterns.go:77:20:     from &Foo{...} (spill) at ./problematic_patterns.go:77:20
--
./problematic_patterns.go:77:20:   flow: {heap} ← f:
./problematic_patterns.go:77:20:     from f.bar.foo = f (assign) at ./problematic_patterns.go:77:20
./problematic_patterns.go:75:33: make(map[int]*Foo) escapes to heap in main:
./problematic_patterns.go:75:33:   flow: {storage for &ConstantGrow{...}} ← &{storage for make(map[int]*Foo)}:
./problematic_patterns.go:75:33:     from make(map[int]*Foo) (spill) at ./problematic_patterns.go:75:33
--
./problematic_patterns.go:75:33:   flow: {heap} ← {temp}:
./problematic_patterns.go:75:33:     from cg.s = append(cg.s, f.bar) (assign) at ./problematic_patterns.go:77:20
./problematic_patterns.go:75:33: make([]*Bar, 0) escapes to heap in main:
./problematic_patterns.go:75:33:   flow: {storage for &ConstantGrow{...}} ← &{storage for make([]*Bar, 0)}:
./problematic_patterns.go:75:33:     from make([]*Bar, 0) (spill) at ./problematic_patterns.go:75:33
--
./problematic_patterns.go:75:33:   flow: {heap} ← {temp}:
./problematic_patterns.go:75:33:     from cg.s = append(cg.s, f.bar) (assign) at ./problematic_patterns.go:77:20
./problematic_patterns.go:77:20: &Bar{...} escapes to heap in main:
./problematic_patterns.go:77:20:   flow: {storage for &Foo{...}} ← &{storage for &Bar{...}}:
./problematic_patterns.go:77:20:     from &Bar{...} (spill) at ./problematic_patterns.go:77:20
./problematic_patterns.go:77:20:     from Foo{...} (struct literal element) at ./problematic_patterns.go:77:20
./problematic_patterns.go:75:33: &ConstantGrow{...} does not escape
./problematic_patterns.go:75:33: make(map[int]*Foo) escapes to heap
./problematic_patterns.go:75:33: make([]*Bar, 0) escapes to heap
./problematic_patterns.go:77:20: &Foo{...} escapes to heap
./problematic_patterns.go:77:20: &Bar{...} escapes to heap
./problematic_patterns.go:77:20: append escapes to heap
./problematic_patterns.go:78:34: &Stat{} does not escape