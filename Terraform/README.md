# Terraform Notes and Code

Terraform is an open-source infrastructure as code software tool created by HashiCorp using a declarative configuration language known as HashiCorp Configuration Language, or optionally JSON.

**Quick Tips**
* Install the [`cli`](https://www.terraform.io/downloads.html)

* Use the VS code  [`Azure Resource Manager Tools`](https://marketplace.visualstudio.com/items?itemName=HashiCorp.terraform) extension

* Get familiar with [`JSON Schema`](https://json-schema.org/) to become more familiar with the structure

* Link to Terraform [`providers`](https://registry.terraform.io/browse/providers)

* Pre-defined gitignore file for [`terraform`](https://github.com/github/gitignore/blob/master/Terraform.gitignore)

## Concepts and examples

| Terraform concepts | Reference | Example
| --- | ----------- |---------|
| Interpolation can be used for referencing values from resources to be used in other files/resources. | [Reference](https://www.terraform.io/docs/configuration-0-11/interpolation.html)| ![Example](https://github.com/dboconsultingllc/BootcampConcepts/blob/mainBranch/images/tf-InterpolationExample.png)
| The terraform `plan` command is used to prepare to execute your scripts | [Reference](https://www.terraform.io/docs/cli/commands/plan.html)| N/A
| The terraform `apply` command is used to execute your scripts based on the plan | [Reference](https://www.terraform.io/docs/cli/commands/apply.html)| N/A
| Take a guess on what the terraform `destroy` command does. :blush:  | [Reference](https://www.terraform.io/docs/cli/commands/destroy.html)| N/A

