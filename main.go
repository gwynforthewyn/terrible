package main

import (
	"context"
	ansible_runner "terrible.playtechnique.io/ansible"
	terraform_runner "terrible.playtechnique.io/terraform"
)

type AnsibleConfig ansible_runner.Config

type TerribleConfig struct {
	terraform_runner.Terraformer
	AnsibleConfig
}

func main() {
	terraformArgs := []string{"version"}
	workingDir := "/Users/gwyn/Developer/azure/terraform/default_infra"

	ctx := context.Background()

	var foo = TerribleConfig{
		Terraformer: terraform_runner.Terraformer{CliArgs: terraformArgs},
	}

	err := foo.Terraformer.Execute(ctx, workingDir)

	if err != nil {
		panic(err)
	}
}
