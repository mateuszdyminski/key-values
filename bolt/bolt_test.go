package bolt

import "testing"

func Benchmark1000000(b *testing.B)  { Insert(1000000) }