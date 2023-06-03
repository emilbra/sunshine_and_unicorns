resource "azurerm_key_vault" "my_kv" {
  name                = var.my_kv_name
  location            = azurerm_resource_group.my_rg.location
  resource_group_name = azurerm_resource_group.my_rg.name
}