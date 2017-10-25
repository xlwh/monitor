package utils

import (
	"testing"
)

type Input struct {
	apikey string
	cityid string
}

func TestHttpGet(t *testing.T) {
	in := &Input{"", "CN10101010018A"}
	resp, ret := HttpGet("http://apis.baidu.com/heweather/pro/attractions", in, 100)
	if ret != nil {
		t.Errorf("TestHttpGet Fail:%s\n", ret.Error())
		t.Failed()
	} else {
		if resp != "" {
			t.Failed()
		}
	}
}
