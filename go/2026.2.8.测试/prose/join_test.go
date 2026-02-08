package prose

import "testing"

//测试文件的函数名应以Test为开头
func TestJoinWithCommas(t *testing.T) { 

	get := JoinWithCommas([]string{"apple", "orange", "pear"})
	want := "apple, orange, and pear"
	if get != want {
		t.Fatalf("3 items: get %q, want %q", get, want)
	}

	get = JoinWithCommas([]string{"apple", "orange"})
	want = "apple and orange"
	if get != want {
		t.Fatalf("2 items: get %q, want %q", get, want)
	}

}

// 运行测试的命令
// go test lipcoder/prose	 测试
// go test lipcoder/prose -v 展示详细信息
// 