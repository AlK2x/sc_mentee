./stack_only.go:40:15: inlining call to cleanNewBoard
./stack_only.go:85:17: inlining call to NewGame
./stack_only.go:13:14: startPoints does not escape
./stack_only.go:39:10: NextGeneration ignoring self-assignment in g.board = g.newBoard
./stack_only.go:26:21: g does not escape
./stack_only.go:43:20: g does not escape
./stack_only.go:51:23: g does not escape
./stack_only.go:85:17: ... argument does not escape

Ничего не эскейпит в Heap