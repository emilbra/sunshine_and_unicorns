resource "azurerm_key_vault" "my_kv" { #tfsec:ignore:azure-keyvault-no-purge
  name                = var.my_kv_name
  location            = azurerm_resource_group.my_rg.location
  resource_group_name = azurerm_resource_group.my_rg.name
  sku_name            = "standard"
  tenant_id           = var.my_tenant_id
  network_acls {
    bypass         = "AzureServices"
    default_action = "Deny"
  }
}

