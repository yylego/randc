# randc

Go utilities with secure randomness based on crypto/rand.

<!-- TEMPLATE (EN) BEGIN: BADGES -->

[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/randc/release.yml?branch=main&label=BUILD)](https://github.com/yylego/randc/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/randc)](https://pkg.go.dev/github.com/yylego/randc)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/randc/main.svg)](https://coveralls.io/github/yylego/randc?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.26%2B-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/randc.svg)](https://github.com/yylego/randc/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/randc)](https://goreportcard.com/report/github.com/yylego/randc)
<!-- TEMPLATE (EN) CLOSE: BADGES -->

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)

<!-- TEMPLATE (EN) CLOSE: LANGUAGE NAVIGATION -->

## Setup

```bash
go get github.com/yylego/randc
```

## API

| Function           | Result                                                                |
| ------------------ | --------------------------------------------------------------------- |
| `IntN(n)`          | An int in `[0, n)`; requires `n > 0`                                  |
| `IntBetween(a, b)` | An int in `[min(a, b), max(a, b))`; matching bounds return that value |
| `DigitString(n)`   | A string of n ASCII digits, retaining leading zeros; requires `n > 0` |

Each function returns a value and an `error`. Results use `crypto/rand.Int` with unbiased sampling. The package has no runtime dependencies outside Go itself.

## Usage

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

`DigitString(6)` can produce strings such as `"042317"`. The argument specifies a digit count, not a maximum. Repeated outputs remain possible; random generation does not guarantee uniqueness. Applications manage expiration, attempt limits, and code storage. Avoid logging secret codes.

## Checks

```bash
make test
go vet ./...
```

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) CLOSE: STANDARD PROJECT FOOTER -->

<!-- TEMPLATE (EN) BEGIN: GITHUB STARS -->

## GitHub Stars

[![Stargazers](https://starchart.cc/yylego/randc.svg?variant=adaptive)](https://starchart.cc/yylego/randc)
<!-- TEMPLATE (EN) CLOSE: GITHUB STARS -->
