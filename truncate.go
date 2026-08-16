package truncate

// 一个"显示单元"的宽度：CJK 等宽全角字符算 2，其余算 1。
func runeWidth(r rune) int {
	if r >= 0x1100 && (r <= 0x115F || r == 0x2329 || r == 0x232A ||
		(r >= 0x2E80 && r <= 0xA4CF) ||
		(r >= 0xAC00 && r <= 0xD7A3) ||
		(r >= 0xF900 && r <= 0xFAFF) ||
		(r >= 0xFE30 && r <= 0xFE4F) ||
		(r >= 0xFF00 && r <= 0xFF60) ||
		(r >= 0xFFE0 && r <= 0xFFE6) ||
		(r >= 0x4E00 && r <= 0x9FFF)) {
		return 2
	}
	return 1
}

// DisplayWidth 返回字符串的显示宽度。
func DisplayWidth(s string) int {
	w := 0
	for _, r := range s {
		w += runeWidth(r)
	}
	return w
}

// Truncate 按显示宽度截断到 width，截断处用 ellipsis 补在末尾（默认 "…"，传空串不加）。
// 原串宽度不足则原样返回。
func Truncate(s string, width int, ellipsis string) string {
	if width <= 0 {
		return ""
	}
	if DisplayWidth(s) <= width {
		return s
	}
	ew := DisplayWidth(ellipsis)
	limit := width - ew
	if limit < 0 {
		limit = 0
	}
	var b []rune
	cur := 0
	for _, r := range s {
		w := runeWidth(r)
		if cur+w > limit {
			break
		}
		b = append(b, r)
		cur += w
	}
	return string(b) + ellipsis
}
