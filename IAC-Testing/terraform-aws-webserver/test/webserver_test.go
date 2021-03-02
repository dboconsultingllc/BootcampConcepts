package test

import (
	"fmt"
	"time"
	"testing"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
	//Terratest library above
	//NOTE: Run go get -t -v to get all of the libraries to run the tests
)

//NOTE: All function names must be written in "Title Case"
func TestTerraformWebserverExample(t *testing.T){
	//Values to pass into the terraform CLI
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		//Path where the example terraform code is located	
		TerraformDir: "../examples/webserver",

		//Variables to pass to the terraform code using -var options
		Vars: map[string]interface{}{
			"region": "us-west-2",
			"serverName": "testwebserver2",
			"serverSize": "t3.micro",	
		},
	})
	//Function within terratest library to apply and execute
	terraform.InitAndApply(t, terraformOptions)

	//If tet bombs out, then destroy
	defer terraform.Destroy(t, terraformOptions)

	//public ip variable in GoLang reading from the terraform outputs file
	publicIP := terraform.Output(t, terraformOptions, "public_ip")
	
	//the fmt library is for printing output. A string replace function to replace with the public ip
	url := fmt.Sprintf("http://%s:8080", publicIP)

	//http helper library is in the terratest library
	//t for test, url variable, nil is for tls, 200 http status, "body of html", 30 retries, wait 5 seconds between tests 
	http_helper.HttpGetWithRetry(t, url, nil, 200, "I made a Terraform Module!", 30, 5*time.Second)
}