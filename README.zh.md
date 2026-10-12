# randc

基于 crypto/rand 的 Go 安全随机工具。

<!-- TEMPLATE (ZH) BEGIN: BADGES -->

[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/randc/release.yml?branch=main&label=BUILD)](https://github.com/yylego/randc/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/randc)](https://pkg.go.dev/github.com/yylego/randc)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/randc/main.svg)](https://coveralls.io/github/yylego/randc?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.26%2B-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/randc.svg)](https://github.com/yylego/randc/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/randc)](https://goreportcard.com/report/github.com/yylego/randc)
<!-- TEMPLATE (ZH) CLOSE: BADGES -->

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->

## 英文文档

[ENGLISH README](README.md)

<!-- TEMPLATE (ZH) CLOSE: LANGUAGE NAVIGATION -->

## 安装

```bash
go get github.com/yylego/randc
```

## API

| 函数               | 结果                                                   |
| ------------------ | ------------------------------------------------------ |
| `IntN(n)`          | `[0, n)` 内的整数，要求 `n > 0`                        |
| `IntBetween(a, b)` | `[min(a, b), max(a, b))` 内的整数，端点相等时返回该值  |
| `DigitString(n)`   | 恰好 n 位的 ASCII 数字字符串，保留前导零，要求 `n > 0` |

每个函数返回一个值和一个 `error`。底层使用 `crypto/rand.Int` 无偏采样。运行时代码仅依赖 Go 标准库。

## 使用

```go
package main

import (
	"fmt"
	"log"

	"github.com/yylego/randc"
)

func main() {
	code, err := randc.DigitString(6)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(code)) // 6
}
```

`DigitString(6)` 可以生成 `"042317"` 这样的字符串。参数表示位数，不表示取值上限。随机结果可以重复，不保证唯一。有效期、尝试次数限制和验证码存储由业务处理。避免将保密验证码写入日志。

## 检查

```bash
make test
go vet ./...
```

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们完善文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，协助解决性能问题
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：面向用户的更改需要更新文档
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来贡献此项目。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) CLOSE: STANDARD PROJECT FOOTER -->

<!-- TEMPLATE (ZH) BEGIN: GITHUB STARS -->

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/yylego/randc.svg?variant=adaptive)](https://starchart.cc/yylego/randc)
<!-- TEMPLATE (ZH) CLOSE: GITHUB STARS -->
