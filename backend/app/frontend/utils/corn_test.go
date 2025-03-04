package utils

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

// TestBasicCronJob 测试基本的定时任务功能
func TestBasicCronJob(t *testing.T) {
	// 创建cron管理器
	cm := NewCronManager()
	defer cm.Stop()

	// 用于跟踪任务执行次数
	var counter int32
	// 用于等待任务执行的通道
	done := make(chan bool)

	// 添加每秒执行一次的任务
	jobID := "test_basic_job"
	_, err := cm.AddJob(jobID, "* * * * * *", func() {
		atomic.AddInt32(&counter, 1)
		if atomic.LoadInt32(&counter) >= 3 { // 执行3次后完成
			done <- true
		}
		return
	})

	if err != nil {
		t.Fatalf("添加任务失败: %v", err)
	}

	// 等待任务执行3次或超时
	select {
	case <-done:
		if c := atomic.LoadInt32(&counter); c != 3 {
			t.Errorf("期望任务执行3次，实际执行%d次", c)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("测试超时")
	}
}

// TestConcurrentJobs 测试多个任务并发执行
func TestConcurrentJobs(t *testing.T) {
	cm := NewCronManager()
	defer cm.Stop()

	// 用于跟踪每个任务的执行次数
	counters := make([]int32, 3)
	done := make(chan bool)

	// 添加3个并发任务
	for i := 0; i < 3; i++ {
		jobID := fmt.Sprintf("concurrent_job_%d", i)
		finalI := i

		_, err := cm.AddJob(jobID, "* * * * * *", func() {
			atomic.AddInt32(&counters[finalI], 1)

			// 检查是否所有任务都至少执行了2次
			for _, c := range counters {
				if atomic.LoadInt32(&c) < 2 {
					return
				}
			}
			done <- true
			return
		})

		if err != nil {
			t.Fatalf("添加任务失败: %v", err)
		}
	}

	// 等待所有任务都执行至少2次或超时
	select {
	case <-done:
		for i, c := range counters {
			if atomic.LoadInt32(&c) < 2 {
				t.Errorf("任务%d执行次数不足，期望>=2，实际%d", i, c)
			}
		}
	case <-time.After(5 * time.Second):
		t.Fatal("测试超时")
	}
}

// TestRemoveJob 测试删除任务功能
func TestRemoveJob(t *testing.T) {
	cm := NewCronManager()
	defer cm.Stop()

	var counter int32
	jobID := "test_remove_job"

	// 添加每秒执行一次的任务
	_, err := cm.AddJob(jobID, "* * * * * *", func() {
		atomic.AddInt32(&counter, 1)
	})

	if err != nil {
		t.Fatalf("添加任务失败: %v", err)
	}

	// 等待任务执行几次
	time.Sleep(2 * time.Second)

	// 记录当前计数
	count1 := atomic.LoadInt32(&counter)
	if count1 == 0 {
		t.Fatal("任务未执行")
	}

	// 删除任务
	if !cm.RemoveJob(jobID) {
		t.Fatal("删除任务失败")
	}

	// 等待一段时间，确认任务不再执行
	time.Sleep(2 * time.Second)

	// 检查计数是否保持不变
	count2 := atomic.LoadInt32(&counter)
	if count2 != count1 {
		t.Errorf("任务仍在执行：删除前计数=%d，删除后计数=%d", count1, count2)
	}

	// 验证任务确实被删除了
	if entry := cm.GetEntry(jobID); entry != nil {
		t.Error("任务仍然存在")
	}
}
