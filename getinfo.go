package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
        "net"
)

func GetServerInfo() {
	var ip []byte
	var err error
	var cmd *exec.Cmd
	// 执行单个shell命令时, 直接运行即可
	cmd = exec.Command("/bin/sh", "-c", `dmidecode|grep "System Information" -A9|egrep  "Manufacturer|Product|Serial"|tr -d "\t"`)
	if ip, err = cmd.Output(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(ip))
}

func GetCpuInfo() {
	var ip []byte
	var err error
	var cmd *exec.Cmd
	// 执行单个shell命令时, 直接运行即可
	cmd = exec.Command("/bin/sh", "-c", `dmidecode|grep "Socket Designation: CPU"|wc  -l `)
	if ip, err = cmd.Output(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("CPU的物理数量:", strings.Replace(string(ip),"\n","",-1))

	cmd = exec.Command("/bin/sh", "-c", `cat  /proc/cpuinfo |grep   "processor"|wc -l `)
	if ip, err = cmd.Output(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("CPU的核心数:", strings.Replace(string(ip),"\n","",-1))

	cmd = exec.Command("/bin/sh", "-c", `cat  /proc/cpuinfo |grep "model name"|uniq `)
	if ip, err = cmd.Output(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ArrayLine := strings.Split(string(ip), ":")
	if len(ArrayLine) == 2 {
		fmt.Println("CPU的详细信息:", strings.Replace(ArrayLine[1],"\n","",-1))
	}

}

func GetMemInfo() {
	var ip []byte
	var err error
	var cmd *exec.Cmd
	// 执行单个shell命令时, 直接运行即可
	cmd = exec.Command("/bin/sh", "-c", `cat  /proc/meminfo |grep  MemTotal|tr -s " " `)
	if ip, err = cmd.Output(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ArrayLine := strings.Split(string(ip), ":")
	if len(ArrayLine) == 2 {
		fmt.Println("内存容量:", ArrayLine[1])
	}

}

func GetNetInfo() {
          ips :=  make(map[string]string)

    interfaces, err := net.Interfaces()
    if err != nil {
        fmt.Println(err)
    }

    for _, i := range interfaces {
        byName, err := net.InterfaceByName(i.Name)
        if err != nil {
           fmt.Println(err)
        }
        addresses, err := byName.Addrs()
        for _, v := range addresses {
            ips[byName.Name] = v.String()
            fmt.Println("网卡配置:",byName.Name,v.String())
        }
    }

}


func GetDiskInfo() {
        var ip []byte
        var err error
        var cmd *exec.Cmd
        // 执行单个shell命令时, 直接运行即可
        cmd = exec.Command("/bin/sh", "-c", `cat /proc/scsi/scsi|grep   "Host: scsi"|wc -l `)
        if ip, err = cmd.Output(); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
                fmt.Println("硬盘数量:", strings.Replace(string(ip),"\n","",-1))
        
               cmd = exec.Command("/bin/sh", "-c", `df  -hl`)
        if ip, err = cmd.Output(); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
                fmt.Println("文件系统使用情况:\n", string(ip))


}


func main() {
	GetServerInfo()
	GetCpuInfo()
	GetMemInfo()
        GetDiskInfo()
        GetNetInfo()

}
