package truncate

import "testing"

func TestTruncate(t *testing.T) {
	// 中文全角算 2 宽
	if w := DisplayWidth("你好"); w != 4 {
		t.Errorf("DisplayWidth(你好)=%d want 4", w)
	}
	if got := Truncate("你好world", 6, ""); got != "你好wo" {
		t.Errorf("Truncate=%q want 你好wo", got)
	}
	if got := Truncate("hello", 3, "…"); got != "he…" {
		t.Errorf("Truncate=%q want he…", got)
	}
	if got := Truncate("hi", 10, "…"); got != "hi" {
		t.Errorf("Truncate noop=%q", got)
	}
	if got := Truncate("你好world", 4, ""); got != "你好" {
		t.Errorf("Truncate=%q want 你好", got)
	}
}
