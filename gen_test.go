package extratime

import (
	"context"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

func ci() bool {
	ci, _ := strconv.ParseBool(os.Getenv("CI"))
	return ci
}

func Test_generate(t *testing.T) {
	if !ci() {
		t.Skip()
	}
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
