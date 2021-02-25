#Function to ceate new resource groups
function New-ResourceGroup{
    [CmdletBinding()]

    param (
        # Parameter: Resource Group Name
        [Parameter(Mandatory)]
        [string] $RGName,

        # Parameter: Resource Group Location
        [Parameter(Mandatory)]
        [string] $RGLocation
    )

    $params = @{
        'Name' = $RGName
        'Location' = $RGLocation

    }

    New-AzResourceGroup @params

}