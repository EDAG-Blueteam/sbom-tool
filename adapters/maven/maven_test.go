package maven

import (
	"github.com/kinbiko/jsonassert"
	"os"
	"sbom-tool/structs"
	"testing"
)

func TestMaven_Generate(t *testing.T) {
	type args struct {
		resultInfo structs.ResultInfo
	}

	expected, _ := os.ReadFile("../../test/maven-bom.json")

	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "test example maven",
			args: args{
				resultInfo: structs.ResultInfo{
					Path: "../../example/some/pom.xml",
					Uuid: "unit-test",
					Type: "Maven",
				},
			},
			want: expected,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Maven{}
			got := m.Generate(tt.args.resultInfo)
			gotStr := string(got[:])
			wantStr := string(tt.want[:])
			ja := jsonassert.New(t)
			ja.Assertf(gotStr, wantStr)
		})
	}
}
