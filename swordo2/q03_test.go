package swordo2

import (
	"fmt"
	"testing"
)

// 数组中的重复数字

func TestQ03(t *testing.T) {

	q03UseSwap()
}

func q03UseSwap() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}

	for i, num := range nums {

		for i != num {
			if num == nums[num] {
				fmt.Printf("final:%d", num)
				return
			}
			nums[i], nums[num] = nums[num], nums[i]
			num = nums[i]
		}

	}
}


[2, 3, 1, 0, 2, 5, 3]
[2, 3, 1, 0, 2, 5, 3]

func q03UseHash() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}

	arr := make([]int, len(nums))

	for _, num := range nums {
		if arr[num] > 0 {
			print(num)
			break
		}
		arr[num]++
	}
}



curl -X 'POST' \
  'https://opsdev.kadev-cicd.statusfeishu.cn/open/api/alert/send_custom_notify' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  --header 'X-OPS-Auth: kju#!yv6%T1i^4Q9REwHUNY#Eke4cFLA' \
  -d '{
  "content": "{\n  \"header\": {\n    \"title\": {\n      \"tag\": \"plain_text\",\n      \"content\": \"标题\"\n    }\n  },\n  \"elements\": [\n    {\n      \"tag\": \"div\",\n      \"fields\": [\n        {\n          \"is_short\": true,\n          \"text\": {\n            \"tag\": \"lark_md\",\n            \"content\": \"内容自定义通知 [newapi]\"\n          }\n        }\n      ]\n    }\n    \n  ]\n}",
  "notify_channel": "lark",
  "receiver_id": "jinshun.0731@bytedance.com",
  "receiver_type": "email"
}'