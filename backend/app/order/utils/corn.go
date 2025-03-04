package utils

import (
	"fmt"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

// CronManager cron管理器
type CronManager struct {
	cron     *cron.Cron
	jobMap   sync.Map // jobID -> EntryID 映射
	stopChan chan struct{}
}

// getDefaultOptions 获取默认的cron选项
func (c *CronManager) getDefaultOptions() []cron.Option {
	return []cron.Option{
		cron.WithLocation(time.Local),
		cron.WithSeconds(),
		cron.WithChain(
			cron.Recover(cron.DefaultLogger),
		),
	}
}

// NewCronManager 创建新的cron管理器
func NewCronManager() *CronManager {
	c := &CronManager{
		stopChan: make(chan struct{}),
	}

	c.cron = cron.New(c.getDefaultOptions()...)
	c.cron.Start()
	return c
}

// AddJob 添加定时任务
func (c *CronManager) AddJob(jobID string, spec string, cmd func()) (cron.EntryID, error) {

	if _, ok := c.jobMap.Load(jobID); ok {
		return 0, fmt.Errorf("jobID %s already exists", jobID)
	}

	// 添加任务到cron
	entryID, err := c.cron.AddFunc(spec, cmd)

	if err != nil {
		return 0, err
	}

	c.jobMap.Store(jobID, entryID)
	return entryID, nil
}

// RemoveJob 移除定时任务
func (c *CronManager) RemoveJob(jobID string) bool {
	if entryIDVal, exists := c.jobMap.Load(jobID); exists {
		entryID := entryIDVal.(cron.EntryID)
		c.cron.Remove(entryID)
		c.jobMap.Delete(jobID)
		return true
	}
	return false
}

// Stop 停止所有定时任务
func (c *CronManager) Stop() {
	c.cron.Stop()
	close(c.stopChan)
}

// GetEntries 获取所有定时任务
func (c *CronManager) GetEntries() []cron.Entry {
	return c.cron.Entries()
}

// GetEntry 获取指定任务
func (c *CronManager) GetEntry(jobID string) *cron.Entry {
	if entryIDVal, exists := c.jobMap.Load(jobID); exists {
		entryID := entryIDVal.(cron.EntryID)
		entry := c.cron.Entry(entryID)
		return &entry
	}
	return nil
}

// Clear 清空所有定时任务
func (c *CronManager) Clear() {
	c.jobMap.Range(func(key, value interface{}) bool {
		c.jobMap.Delete(key)
		return true
	})
	c.cron.Stop()
	c.cron = cron.New(c.getDefaultOptions()...)
	c.cron.Start()
}
