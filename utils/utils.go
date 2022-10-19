/************************************************************
 * Author:        jackey
 * Date:        2022/10/19
 * Description: 工具类
 * Version:    V1.0.0
 **********************************************************/

package utils

import (
	"bytes"
	"gin/logs"
	"os/exec"
)

// ExecNoRes 执行shell指令，无返回值
func ExecNoRes(command string) error {
	in := bytes.NewBuffer(nil)
	cmd := exec.Command("sh")
	cmd.Stdin = in
	in.WriteString(command)
	if err := cmd.Run(); err != nil {
		logs.Error("ExecNoRes", command)
		return err
	}
	return nil
}
