./problematic_patterns.go:62:33: inlining call to NewConstantGrow
./problematic_patterns.go:64:20: inlining call to (*ConstantGrow).Grow
./problematic_patterns.go:26:9: &ConstantGrow{...} escapes to heap in NewConstantGrow:
./problematic_patterns.go:26:9:   flow: ~r0 ← &{storage for &ConstantGrow{...}}:
./problematic_patterns.go:26:9:     from &ConstantGrow{...} (spill) at ./problematic_patterns.go:26:9
./problematic_patterns.go:26:9:     from return &ConstantGrow{...} (return) at ./problematic_patterns.go:26:2
./problematic_patterns.go:27:10: make(map[int]*Foo) escapes to heap in NewConstantGrow:
./problematic_patterns.go:27:10:   flow: {storage for &ConstantGrow{...}} ← &{storage for make(map[int]*Foo)}:
./problematic_patterns.go:27:10:     from make(map[int]*Foo) (spill) at ./problematic_patterns.go:27:10
./problematic_patterns.go:27:10:     from ConstantGrow{...} (struct literal element) at ./problematic_patterns.go:26:22
./problematic_patterns.go:28:10: make([]*Bar, 0) escapes to heap in NewConstantGrow:
./problematic_patterns.go:28:10:   flow: {storage for &ConstantGrow{...}} ← &{storage for make([]*Bar, 0)}:
./problematic_patterns.go:28:10:     from make([]*Bar, 0) (spill) at ./problematic_patterns.go:28:10
./problematic_patterns.go:28:10:     from ConstantGrow{...} (struct literal element) at ./problematic_patterns.go:26:22
./problematic_patterns.go:26:9: &ConstantGrow{...} escapes to heap
./problematic_patterns.go:27:10: make(map[int]*Foo) escapes to heap
./problematic_patterns.go:28:10: make([]*Bar, 0) escapes to heap
./problematic_patterns.go:45:16: append(cg.s, f.bar) escapes to heap in (*ConstantGrow).Grow:
./problematic_patterns.go:45:16:   flow: {heap} ← &{storage for append(cg.s, f.bar)}:
./problematic_patterns.go:45:16:     from append(cg.s, f.bar) (spill) at ./problematic_patterns.go:45:16
./problematic_patterns.go:45:16:     from cg.s = append(cg.s, f.bar) (assign) at ./problematic_patterns.go:45:8
./problematic_patterns.go:36:8: &Foo{...} escapes to heap in (*ConstantGrow).Grow:
./problematic_patterns.go:36:8:   flow: f ← &{storage for &Foo{...}}:
./problematic_patterns.go:36:8:     from &Foo{...} (spill) at ./problematic_patterns.go:36:8
--
./problematic_patterns.go:32:7:   flow: {heap} ← {temp}:
./problematic_patterns.go:32:7:     from cg.s = append(cg.s, f.bar) (assign) at ./problematic_patterns.go:45:8
./problematic_patterns.go:39:9: &Bar{...} escapes to heap in (*ConstantGrow).Grow:
./problematic_patterns.go:39:9:   flow: {storage for &Foo{...}} ← &{storage for &Bar{...}}:
./problematic_patterns.go:39:9:     from &Bar{...} (spill) at ./problematic_patterns.go:39:9
./problematic_patterns.go:39:9:     from Foo{...} (struct literal element) at ./problematic_patterns.go:36:12
./problematic_patterns.go:32:7: leaking param content: cg
./problematic_patterns.go:36:8: &Foo{...} escapes to heap
./problematic_patterns.go:39:9: &Bar{...} escapes to heap
./problematic_patterns.go:45:16: append escapes to heap
./problematic_patterns.go:52:5: func literal escapes to heap in rundomNumberGenerator:
./problematic_patterns.go:52:5:   flow: {heap} ← &{storage for func literal}:
./problematic_patterns.go:52:5:     from func literal (spill) at ./problematic_patterns.go:52:5
./problematic_patterns.go:51:2: rundomNumberGenerator capturing by value: ch (addr=false assign=false width=8)
./problematic_patterns.go:53:3: rundomNumberGenerator.func1 capturing by value: .autotmp_0 (addr=false assign=false width=8)
./problematic_patterns.go:52:5: func literal escapes to heap
./problematic_patterns.go:64:20: append(cg.s, f.bar) escapes to heap in main:
./problematic_patterns.go:64:20:   flow: {heap} ← &{storage for append(cg.s, f.bar)}:
./problematic_patterns.go:64:20:     from append(cg.s, f.bar) (spill) at ./problematic_patterns.go:64:20
./problematic_patterns.go:64:20:     from cg.s = append(cg.s, f.bar) (assign) at ./problematic_patterns.go:64:20
./problematic_patterns.go:64:20: &Foo{...} escapes to heap in main:
./problematic_patterns.go:64:20:   flow: f ← &{storage for &Foo{...}}:
./problematic_patterns.go:64:20:     from &Foo{...} (spill) at ./problematic_patterns.go:64:20
--
./problematic_patterns.go:64:20:   flow: {heap} ← f:
./problematic_patterns.go:64:20:     from f.bar.foo = f (assign) at ./problematic_patterns.go:64:20
./problematic_patterns.go:62:33: make(map[int]*Foo) escapes to heap in main:
./problematic_patterns.go:62:33:   flow: {storage for &ConstantGrow{...}} ← &{storage for make(map[int]*Foo)}:
./problematic_patterns.go:62:33:     from make(map[int]*Foo) (spill) at ./problematic_patterns.go:62:33
--
./problematic_patterns.go:62:33:   flow: {heap} ← {temp}:
./problematic_patterns.go:62:33:     from cg.s = append(cg.s, f.bar) (assign) at ./problematic_patterns.go:64:20
./problematic_patterns.go:62:33: make([]*Bar, 0) escapes to heap in main:
./problematic_patterns.go:62:33:   flow: {storage for &ConstantGrow{...}} ← &{storage for make([]*Bar, 0)}:
./problematic_patterns.go:62:33:     from make([]*Bar, 0) (spill) at ./problematic_patterns.go:62:33
--
./problematic_patterns.go:62:33:   flow: {heap} ← {temp}:
./problematic_patterns.go:62:33:     from cg.s = append(cg.s, f.bar) (assign) at ./problematic_patterns.go:64:20
./problematic_patterns.go:64:20: &Bar{...} escapes to heap in main:
./problematic_patterns.go:64:20:   flow: {storage for &Foo{...}} ← &{storage for &Bar{...}}:
./problematic_patterns.go:64:20:     from &Bar{...} (spill) at ./problematic_patterns.go:64:20
./problematic_patterns.go:64:20:     from Foo{...} (struct literal element) at ./problematic_patterns.go:64:20
./problematic_patterns.go:62:33: &ConstantGrow{...} does not escape
./problematic_patterns.go:62:33: make(map[int]*Foo) escapes to heap
./problematic_patterns.go:62:33: make([]*Bar, 0) escapes to heap
./problematic_patterns.go:64:20: &Foo{...} escapes to heap
./problematic_patterns.go:64:20: &Bar{...} escapes to heap
./problematic_patterns.go:64:20: append escapes to heap