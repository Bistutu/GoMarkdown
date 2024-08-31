# GoMarkdown

GoMarkdown 是一个用于生成 Markdown 文档的 Go 库。它全面支持各种 Markdown 元素，从基本格式到高级功能，如 Mermaid 图表、数学表达式和 GitHub 风格的警报框。

## 功能

- **基本的 Markdown 元素：**
  - 标题（H1-H6）
  - 粗体、斜体、删除线文本
  - 行内代码和代码块
  - 水平分割线
  - 链接和图片（支持自定义大小和缩放）
  - 引用块和脚注
  - 无序和有序列表、任务列表
  - 定义列表和目录

- **高级功能：**
  - Mermaid 图表（时序图、流程图和自定义图表）
  - 使用 LaTeX 语法的数学表达式（行内和块级）
  - GitHub 风格的警报框（提示、警告、危险）
  - 支持嵌入自定义 HTML 内容

## 安装

要安装 GoMarkdown，请使用以下 `go get` 命令：

```bash
go get github.com/Bistutu/GoMarkdown
```

## 使用方法

以下是如何使用 GoMarkdown 创建 Markdown 文档的示例：

```go
package main

import (
	"fmt"
	"os"
	"github.com/Bistutu/GoMarkdown"
)

func main() {
	// 创建一个 MarkdownWriter 实例
	writer, err := gomarkdown.NewMarkdownWriter("example.md", true)
	if err != nil {
		fmt.Printf("创建 MarkdownWriter 失败: %v\n", err)
		return
	}
	defer writer.Close()

	// 生成 Markdown 内容
	writer.H1("GoMarkdown 示例")
	writer.Bold("这是粗体文本")
	writer.Italic("这是斜体文本")
	writer.Code("go", `fmt.Println("Hello, GoMarkdown!")`)
	writer.ImageWithScale("示例图片", "https://example.com/image.png", 50)

	// 更多 Markdown 生成示例...
}
```

## 文档

有关详细的使用方法和所有可用函数，请参阅 [GoDoc 文档](https://pkg.go.dev/github.com/Bistutu/GoMarkdown)。

## 贡献

欢迎贡献！参与贡献的方法：

1. Fork 本仓库。
2. 创建一个新分支 (`git checkout -b feature/YourFeature`)。
3. 进行修改并提交更改 (`git commit -m 'Add YourFeature'`)。
4. 推送到分支 (`git push origin feature/YourFeature`)。
5. 创建一个新的 Pull Request。

贡献时请遵守 [Go 代码风格指南](https://golang.org/doc/effective_go)。

## 许可证

该项目采用 MIT 许可证。详情请参阅 [LICENSE](LICENSE) 文件。

## 鸣谢

- [Go](https://golang.org/) - Go 编程语言
- [Markdown 指南](https://www.markdownguide.org/) - Markdown 语法综合指南
