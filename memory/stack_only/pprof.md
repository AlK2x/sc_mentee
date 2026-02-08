(pprof) top 5
Showing nodes accounting for 1538kB, 100% of 1538kB total
Showing top 5 nodes out of 23
      flat  flat%   sum%        cum   cum%
     513kB 33.36% 33.36%      513kB 33.36%  runtime.allocm
  512.88kB 33.35% 66.70%   512.88kB 33.35%  syscall.init.func1
  512.12kB 33.30%   100%   512.12kB 33.30%  net/http.ListenAndServe (inline)
         0     0%   100%   512.12kB 33.30%  main.main.func1
         0     0%   100%   512.88kB 33.35%  os.Getenv