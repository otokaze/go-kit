package progressbar

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Bar struct {
	refresh time.Duration
	size    uint8
	total   int64
	current int64
	prefix  string
	suffix  string
	stop    bool
	sync.RWMutex
	wg sync.WaitGroup
}

type Config struct {
	Size    uint8
	Total   int64
	Current int64
	Refresh time.Duration
}

func New(c *Config) *Bar {
	if c == nil {
		c = &Config{}
	}
	if c.Size == 0 {
		c.Size = 50
	}
	if c.Total == 0 {
		c.Total = 100
	}
	if c.Refresh == 0 {
		c.Refresh = time.Millisecond * 500
	}
	return &Bar{
		size:    c.Size,
		total:   c.Total,
		refresh: c.Refresh,
		current: c.Current,
	}
}

func (b *Bar) SetPrefix(prefix string) {
	b.Lock()
	b.prefix = prefix
	b.Unlock()
}

func (b *Bar) SetSuffix(suffix string) {
	b.Lock()
	b.suffix = suffix
	b.Unlock()
}

func (b *Bar) Add(i int64) {
	atomic.AddInt64(&b.current, i)
}

func (b *Bar) Set(i int64) {
	atomic.StoreInt64(&b.current, i)
}

func (b *Bar) ListenDir(dir string) {
	go func() {
		for {
			b.RLock()
			if b.stop {
				b.RUnlock()
				println()
				return
			}
			b.RUnlock()
			var dirSzie int64
			filepath.Walk(dir, func(_ string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() {
					dirSzie += info.Size()
				}
				return nil
			})
			if dirSzie > 0 {
				b.Lock()
				b.current = dirSzie
				b.Unlock()
			}
			time.Sleep(b.refresh)
		}
	}()
}

func (b *Bar) Run() {
	b.wg.Add(1)
	go func() {
		defer b.wg.Done()
		for {
			b.RLock()
			if b.stop {
				println()
				b.RUnlock()
				return
			}
			percent := float64(b.current) / float64(b.total)
			if percent > 1 {
				percent = 1
			}
			step := uint8(float64(b.size) * percent)
			str := strings.Repeat("=", int(step))
			if step < b.size {
				str += ">"
				str += strings.Repeat("-", int(b.size-step)-1)
			}
			fmt.Printf("%s [%s]%2.f%% %s\r", b.prefix, str, percent*100, b.suffix)
			b.RUnlock()
			time.Sleep(b.refresh)
		}
	}()
}

func (b *Bar) Stop() {
	b.Lock()
	b.stop = true
	b.Unlock()
	b.wg.Wait()
}
