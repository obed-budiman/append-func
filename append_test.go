package main

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestNewAppend(t *testing.T) {
	g := NewGomegaWithT(t)

	module := NewAppend()

	g.Expect(module.NextIndex).Should(Equal(0))
	g.Expect(module.Data).ShouldNot(BeEmpty())
}

func TestAppend_Insert(t *testing.T) {
	g := NewGomegaWithT(t)

	module := NewAppend()
	module.Insert(1)

	g.Expect(module.NextIndex).Should(Equal(1))
	g.Expect(module.Data[module.NextIndex-1]).Should(Equal(1))
}

func TestAppend_RemoveLast(t *testing.T) {
	g := NewGomegaWithT(t)

	module := NewAppend()
	module.Insert(1)
	module.Insert(2)
	module.Insert(3)
	module.RemoveLast()

	g.Expect(module.NextIndex).Should(Equal(2))
	g.Expect(module.Data[module.NextIndex]).Should(BeZero())
}

func TestAppend_CheckArray(t *testing.T) {
	g := NewGomegaWithT(t)

	module := NewAppend()
	module.CheckArray()

	g.Expect(module.NextIndex).Should(BeZero())
}
