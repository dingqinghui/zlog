/**
 * @Author: dingQingHui
 * @Description:
 * @File: zlog_test
 * @Version: 1.0.0
 * @Date: 2024/11/14 11:33
 */

package zlog

import "testing"

func TestLog(t *testing.T) {
	log := New()
	log.Info("1111111111111")
}
