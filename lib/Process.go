package lib

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

func ShutDownWin() error {

	arg := []string{"-s", "-t", "1"}
	cmd := exec.Command("shutdown", arg...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}

func ProcessExists(processName string) bool {
	// 使用 tasklist 命令查询系统中正在运行的进程
	cmd := exec.Command("tasklist")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return false
	}

	// 解析进程列表，并查找指定的进程
	processes := strings.Split(out.String(), "\n")
	for _, p := range processes {
		if strings.Contains(p, processName) {
			return true
		}
	}
	return false
}

func OpenProcessByShortcut(shortcut string) error {
	var temArg []string
	// 启动目标程序
	temArg = append(temArg, "-Command")
	temArg = append(temArg, "Start-Process")
	temArg = append(temArg, shortcut)

	cmd := exec.Command("PowerShell.exe", temArg...)
	// 启动时隐藏powershell窗口,没有这句会闪一下powershell窗口
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	err := cmd.Run()

	if nil != err {
		return err
	}

	return nil
}

func OpenProcessByName(processName string) error {
	// 使用 exec.Command 函数打开进程
	cmd := exec.Command(processName)

	// 执行命令并等待完成
	err := cmd.Run()
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err != nil {
		return err
	}

	return nil
}

func KillProcessByName(processName string) error {
	// 构造 taskkill 命令
	cmd := exec.Command("taskkill", "/IM", processName, "/F")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	// 执行命令并等待完成
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
