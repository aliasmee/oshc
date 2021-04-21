package main

import (
	"fmt"
	"syscall"

)

type DiskStatus struct {
	All uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
	UsedPercent uint64 `json:"usedpct"`
}

const (
	B = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
)

func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	disk.UsedPercent = disk.Used / disk.All
	return
}

func main() {
	disk := DiskUsage("/")
	fmt.Printf("All: %0.2f GB\n", float64(disk.All)/float64(GB))
	fmt.Printf("Used: %.0f GB\n", float64(disk.Used)/float64(GB))
	fmt.Println(disk.Used)
	fmt.Println(disk.All)
	fmt.Printf("Percent %.0f%\n", float64(disk.Used)/float64(disk.All)*100)
}