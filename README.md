An empty select{} statement blocks forever. It is similar to an empty for{} statement.

On most (all?) supported Go architectures, the empty select will yield CPU. An empty for-loop won't, i.e. it will "spin" on 100% CPU.

On Mac OS X, in Go, for { } will cause the CPU% to max, and the process's STATE will be running

select { }, on the other hand, will not cause the CPU% to max, and the process's STATE will be sleeping

*****************************************************************************************

