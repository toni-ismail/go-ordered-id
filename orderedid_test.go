package go_ordered_id

import (
	"testing"
)

// go test -bench=. -benchmem
func BenchmarkGenerate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Generate()
	}
}

func BenchmarkGenerateBase36(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GenerateBase36()
	}
}

// go test -v
func TestGenerate(t *testing.T) {
	t.Run("Test Base 62", func(t *testing.T) {
		base62Ids := map[string]bool{}
		for range 1000 {
			gen := GenerateBase62()
			t.Log(gen)

			exist := base62Ids[gen]
			if exist {
				t.Errorf("Base62 is duplicated id %s\n", gen)
			} else {
				base62Ids[gen] = true
			}
		}
	})
	t.Run("Test Base 36", func(t *testing.T) {
		base36Ids := map[string]bool{}
		for range 1000 {
			gen := GenerateBase36()
			t.Log(gen)

			exist := base36Ids[gen]
			if exist {
				t.Errorf("Base36 is duplicated id %s\n", gen)
			} else {
				base36Ids[gen] = true
			}
		}
	})
}
