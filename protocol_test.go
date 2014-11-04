package swf

import (
	"testing"
	"time"
)

func TestDateFormatting(t *testing.T) {
	swftime := Time{time.Now()}

	b, err := swftime.MarshalJSON()

	if err != nil {
		t.Fatal(err)
	}

	testTime := new(Time)
	err = testTime.UnmarshalJSON(b)

	if err != nil {
		t.Fatal(err)
	}

	if testTime.Time.Unix() != swftime.Time.Unix() {
		t.Fatal(swftime, testTime)
	}

}
