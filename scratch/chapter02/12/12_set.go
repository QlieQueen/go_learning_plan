package main

import "fmt"

// map 实现 Set：struct{} vs bool —— bool 版的隐蔽 bug

func main() {
	fmt.Println("=== ❌ bool 版：单值判断的漏判 bug ===")
	done := map[string]bool{}
	done["job-ok"] = true
	done["job-fail"] = false // job-fail 处理过了（key 存在），只是失败了

	for _, id := range []string{"job-ok", "job-fail", "job-new"} {
		if !done[id] { // 单值判断：把"值真假"当"存在性"
			fmt.Println("  视为未处理:", id)
		}
	}
	// 正确结果应只有 job-new 才算未处理；job-fail 却也被当成未处理 → 去重 bug

	fmt.Println("\n=== 救法：bool 版用 _, ok ===")
	for _, id := range []string{"job-ok", "job-fail", "job-new"} {
		if _, ok := done[id]; !ok {
			fmt.Println("  视为未处理:", id)
		}
	}

	fmt.Println("\n=== ✅ struct 版：想看单值判断都不行 ===")
	seen := map[string]struct{}{}
	seen["job-ok"] = struct{}{}
	seen["job-fail"] = struct{}{}
	// if seen["job-ok"] { }  // ✗ 编译错误：non-boolean condition
	for _, id := range []string{"job-ok", "job-fail", "job-new"} {
		if _, ok := seen[id]; !ok {
			fmt.Println("  视为未处理:", id)
		}
	}
}
