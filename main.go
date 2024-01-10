package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"Github.com/hashicorp/terraform-json"
	version "github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"

	"github.com/hackebrot/go-repr/repr"
)

func setupEnvironment(terraform_version string) (execPath string, err error) {
	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion(terraform_version)),
	}

	log.Println("Install")
	execPath, err = installer.Install(context.Background())

	return execPath, err
}

func initialize(workingDir string, execPath string) (tf *tfexec.Terraform, err error) {
	tf, err = tfexec.NewTerraform(workingDir, execPath)
	if err != nil {
		//log.Fatalf("error running NewTerraform: %s", err)
		return nil, err
	}

	log.Printf("Init: %s, %s", execPath, workingDir)
	err = tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		// log.Fatalf("error running Init: %s", err)
		return nil, err
	}

	return tf, err
}

func plan(tf *tfexec.Terraform) (plan *tfjson.Plan, err error) {
	tmpDir, err := ioutil.TempDir("", "rover")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tmpDir)

	log.Println("Plan")
	planPath := fmt.Sprintf("%s/%s-%v", tmpDir, "demoplan", time.Now().Unix())

	// Create TF Plan options
	var tfPlanOptions []tfexec.PlanOption
	tfPlanOptions = append(tfPlanOptions, tfexec.Out(planPath))

	diff, err := tf.Plan(context.Background(), tfPlanOptions...)
	if err != nil {
		log.Printf("error running Plan: %s", err)
		return nil, err
	}
	log.Printf("Changes available: %t", diff)

	plan, err = tf.ShowPlanFile(context.Background(), planPath)
	if err != nil {
		log.Printf("error running ShowPlan: %s", err)
		return nil, err
	}
	log.Printf("Terraform version: %s", plan.TerraformVersion)

	return plan, err
}

func main() {
	execPath, err := setupEnvironment("1.1.2")
	if err != nil {
		log.Fatalf("error installing Terraform: %s", err)
	}

	workingDir := "examples/root"
	tf, err := initialize(workingDir, execPath)
	if err != nil {
		log.Fatalf("error running initialize: %s", err)
	}

	log.Printf("Show: %s, %s", execPath, workingDir)
	state, err := tf.Show(context.Background())
	if err != nil {
		log.Fatalf("error running Show: %s", err)
	}
	fmt.Println(state.FormatVersion) // "0.2"

	p, err := plan(tf)
	if err != nil {
		log.Fatalf("error running Plan: %s", err)
	}
	fmt.Printf("Plan: %s", repr.Repr(p.ResourceChanges))
}
