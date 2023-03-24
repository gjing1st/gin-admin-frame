package rand

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
	"testing"
)

func TestUid(t *testing.T) {
	for i := 0; i < 5; i++ {
		//s := GenerateUUID20()
		//fmt.Println("len=", len(s), "===", s)
		//time.Sleep(time.Nanosecond)
		id := uuid.New()
		fmt.Println(id.String(), "============", len(id.String()), id.Version())
	}
	//fmt.Println(time.Now().UnixNano())

	id := uuid.New()
	ids := strings.ReplaceAll(id.String(), "-", "")
	fmt.Println(ids, "============", len(ids), id.Version())

}
