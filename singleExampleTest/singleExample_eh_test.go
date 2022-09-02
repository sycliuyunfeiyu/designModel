package singleExample

import (
	singleExa "design/singleExample"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInstance_eh(t *testing.T) {
	assert.Equal(t, singleExa.GetInstance_eh(), singleExa.GetInstance_eh())
}

func BenchmarkGetInstanceParallel_eh(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {

		for pb.Next() {
			if singleExa.GetInstance_eh() != singleExa.GetInstance_eh() {
				b.Errorf("单例模式实例化出现不相等。。。")

			}
		}
	})
}
