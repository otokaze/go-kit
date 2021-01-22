package progressbar

import (
	"fmt"
	"testing"
	"time"
)

func TestMonitor(t *testing.T) {
	p := New(&Config{Size: 50, Total: 100000})
	p.SetPrefix("下载文件1")
	p.SetSuffix("正在下载...")
	p.Run()
	go func() {
		for {
			p.Add(1000)
			time.Sleep(time.Millisecond * 200)
		}
	}()
	time.Sleep(3 * time.Second)
	p.Stop()
	fmt.Println("=======已经暂停")
}
