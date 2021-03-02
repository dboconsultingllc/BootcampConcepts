# Arm Template Notes and Code

ARM Templates are a way to declare the objects you want, the types, names and properties in a JSON file which can be checked into source control and managed like any other code file. ARM Templates is one way to support the creation of infrastrcuture and services as code in Azure.

**Quick Tips**
* Install or ```az upgrade``` the azure cli

* Use the VS code  [`Azure Resource Manager Tools`](https://marketplace.visualstudio.com/items?itemName=msazurermtools.azurerm-vscode-tools) extension

* Get familiar with [`JSON Schema`](https://json-schema.org/) to become more familiar with the structure

* ARM template [`syntax and structure`](https://docs.microsoft.com/en-us/azure/azure-resource-manager/templates/template-syntax)

* Azure CLI Commnands [`reference`](https://docs.microsoft.com/en-us/cli/azure/reference-index?view=azure-cli-latest)

* [`Visualize`](http://armviz.io/designer) all of the resources in your ARM template. 

* Many more tips to come!

## High level description of an arm template 
* `$schema` Version of the template
* `contentVersion` Required element for template
* `parameters` Customize input to make them more dynamic/reusable
* `functions` Define reusable functions that can be reused within the template
* `variables` Constants or globals to use within the template
* `resources` Resources to create within Azure
* `outputs` Values that will be available or shown when a resource is created

```
{
    "$schema": "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#",
    "contentVersion": "1.0.0.0",
    "parameters": {},
    "functions": [],
    "variables": {},
    "resources": [],
    "outputs": {}
}
```
## ArmViz example using sample ARM template 
| ARM Template | Diagram |
| --- | ----------- |
| [`Link to sample json`](https://github.com/dboconsultingllc/BootcampConcepts/blob/mainBranch/ARM/ARM-NewVM.json) | ![armViz](https://github.com/dboconsultingllc/BootcampConcepts/blob/mainBranch/images/armVizSample.PNG)
|



