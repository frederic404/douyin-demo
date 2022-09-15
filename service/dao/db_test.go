package dao

import (
	"encoding/json"
	"fmt"
	"testing"
	"unsafe"
)

func TestInitDataSources(t *testing.T) {
	type main_group_info struct {
		Update_time int64
	}
	// m := make(map[string]int64, 0)
	o := &main_group_info{
		Update_time: 1231234,
	}
	fmt.Println(o)
	j, err := json.Marshal(o)
	fmt.Println(err)
	fmt.Println(j, unsafe.Sizeof(j))
	var val main_group_info
	json.Unmarshal(j, &val)
	fmt.Println(val)
}
