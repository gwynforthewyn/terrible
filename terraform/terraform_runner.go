package terraform_runner

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-exec/tfexec"
	"log"
	"os"
	"os/exec"
)

type Terraformer struct {
	CliArgs []string
	Context context.Context
}

func (tconf Terraformer) Execute(ctx context.Context, workingDir string) error {
	execPath, err := exec.LookPath("terraform")

	if err != nil {
		log.Fatal(err)
		return err
	}

	tf, err := tfexec.NewTerraform(workingDir, execPath)

	tf.SetStdout(os.Stdout)
	tf.SetStderr(os.Stderr)

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = tf.Apply(ctx)

	if err != nil {
		log.Fatal(err)
		return err
	}

	bar, err := tf.Show(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Errorf("%w", bar)

	return err

}
