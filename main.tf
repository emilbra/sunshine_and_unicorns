terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "3.58.0"
    }
  }
  required_version = "1.4.5"
}

provider "azurerm" {
  features {
  }
}

resource "azurerm_resource_group" "my_rg" {
  name     = var.rg_name
  location = var.rg_location
}