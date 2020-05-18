package extratime

import (
	"context"
	"os/exec"
	"testing"
)

func Test_generate(t *testing.T) {
	ctx := context.Background()
	if err := exec.CommandContext(ctx, "go", "generate", ".").Run(); err != nil {
		t.Fatal(err)
	}
	_out, err := exec.CommandContext(ctx,
		"git", "diff", "extratime_gen.go", "extratime_gen_test.go").Output()
	if err != nil {
		t.Fatal(err)
	}
	out := string(_out)
	if out != "" {
		t.Fatal("diff exists\n", out)
	}
}
