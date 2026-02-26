package health

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

func Run() {
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Println("Error fetching CPU:", err)
		return
	}
	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("/")

	fmt.Printf("CPU: %.2f%%\n", cpuPercent[0])
	fmt.Printf("Memory: %.2f%%\n", vmStat.UsedPercent)
	fmt.Printf("Disk: %.2f%%\n", diskStat.UsedPercent)
}
