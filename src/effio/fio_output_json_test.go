package effio

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestLoadFioJsonData(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name      string
		args      args
		wantFdata FioJsonData
	}{
		{
			"client_stats",
			args{
				"../../suites/maybe/datera-trivial/output.json",
			},
			FioJsonData{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if gotFdata := LoadFioJsonData(tt.args.filename); !reflect.DeepEqual(gotFdata, tt.wantFdata) {
			// 	t.Errorf("LoadFioJsonData() = %v, want %v", gotFdata, tt.wantFdata)
			// }
			spew.Dump(LoadFioJsonData(tt.args.filename).GetJobs())
		})
	}
}
