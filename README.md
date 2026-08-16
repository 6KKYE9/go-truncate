# go-truncate

按"显示宽度"截断字符串，而不是按字节或字符数。关键区别：中文、日文这类全角字符在终端里占两格，按字符数截断会算错，结果对不齐还容易把中文砍成乱码。这个工具按每格来量，并支持在末尾加省略号。

## 装

```bash
go build -o truncate ./cmd/truncate
```

## 用

```bash
echo "你好world" | ./truncate -width 6          # "你好w"
echo "hello" | ./truncate -width 3 -ellipsis …  # "he…"
```

参数：
- `-width`：目标显示宽度，默认 20
- `-ellipsis`：截断处补的尾巴，默认 `…`，传空串就不加

## 当库用

```go
import "truncate"

truncate.DisplayWidth("你好")          // 4
truncate.Truncate("你好world", 6, "")  // "你好w"
truncate.Truncate("hello", 3, "…")    // "he…"
```

宽度不够放省略号时，会尽量保留原串前缀，省略号挤不进去就省略掉。

## License

MIT
