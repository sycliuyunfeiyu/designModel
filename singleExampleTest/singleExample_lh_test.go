package singleExample

import (
	singleExa "design/singleExample"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLazyInstance(t *testing.T) {
	assert.Equal(t, singleExa.GetInstance_lh(), singleExa.GetInstance_lh())
}

func BenchmarkGetLazyInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleExa.GetInstance_lh() != singleExa.GetInstance_lh() {
				b.Errorf("test fail")
			}
		}
	})
}
