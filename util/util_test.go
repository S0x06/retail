package util

import (
	"testing"
)

func TestAdd(t *testing.T) {
	_, err := Add()
	if err != nil {
		t.Error("GenShortId failed!")
	}

	t.Log("GenShortId test pass")
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add()
	}
}

func BenchmarkAddTimeConsuming(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	_, err := Add()
	if err != nil {
		b.Error(err)
	}

	b.StartTimer() //重新开始时间

	for i := 0; i < b.N; i++ {
		Add()
	}
}
