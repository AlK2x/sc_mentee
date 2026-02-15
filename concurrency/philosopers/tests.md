```bash
$ go test -race -v

=== RUN   TestPhilosoperProblem
=== RUN   TestPhilosoperProblem/Resource_Ordering
=== PAUSE TestPhilosoperProblem/Resource_Ordering
=== RUN   TestPhilosoperProblem/Restricted_Parallelism
=== PAUSE TestPhilosoperProblem/Restricted_Parallelism
=== RUN   TestPhilosoperProblem/Central_Coordinator
=== PAUSE TestPhilosoperProblem/Central_Coordinator
=== RUN   TestPhilosoperProblem/Non_Blocking
=== PAUSE TestPhilosoperProblem/Non_Blocking
=== CONT  TestPhilosoperProblem/Resource_Ordering
=== CONT  TestPhilosoperProblem/Central_Coordinator
=== CONT  TestPhilosoperProblem/Restricted_Parallelism
=== CONT  TestPhilosoperProblem/Non_Blocking
--- PASS: TestPhilosoperProblem (0.00s)
    --- PASS: TestPhilosoperProblem/Restricted_Parallelism (5.00s)
    --- PASS: TestPhilosoperProblem/Central_Coordinator (5.00s)
    --- PASS: TestPhilosoperProblem/Non_Blocking (5.00s)
    --- PASS: TestPhilosoperProblem/Resource_Ordering (5.00s)
PASS
ok      philosopers     6.014s
```