output "keyvault_uri" {
  value = azurerm_key_vault.my_kv.vault_uri
}
output "keyvault_name" {
  value = azurerm_key_vault.my_kv.name
}

output "keyvault_location" {
  value = azurerm_key_vault.my_kv.location
}

output "rg_name" {
  value = azurerm_resource_group.my_rg.name
}