terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=2.46.0"
    }
  }
}

provider "azurerm" {
    features {}
}

resource "azurerm_resource_group" "cast0r-resource-group" {
  name     = "cast0r-resource-group"
  location = "West Europe"
}

resource "azurerm_virtual_network" "cast0r-virtual-network" {
  name                = "cast0r-virtual-network"
  resource_group_name = azurerm_resource_group.cast0r-resource-group.name
  location            = azurerm_resource_group.cast0r-resource-group.location
  address_space       = ["10.0.0.0/16"]
}

resource "azurerm_subnet" "cast0r-subnet" {
  name                 = "cast0r-subnet"
  resource_group_name  = azurerm_resource_group.cast0r-resource-group.name
  virtual_network_name = azurerm_virtual_network.cast0r-virtual-network.name
  address_prefixes     = ["10.0.1.0/24"]
}

resource "azurerm_network_interface" "cast0r-network_interface" {
  name = "cast0r-network-interface"
  location            = azurerm_resource_group.cast0r-resource-group.location
  resource_group_name = azurerm_resource_group.cast0r-resource-group.name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = azurerm_subnet.cast0r-subnet.id
    private_ip_address_allocation = "Dynamic"
  }
}

resource "azurerm_public_ip" "cast0r-public-ip" {
  name                = "cast0r-test-public-ip"
  resource_group_name = azurerm_resource_group.cast0r-resource-group.name
  location            = azurerm_resource_group.cast0r-resource-group.location
  allocation_method   = "Dynamic"

  tags = {
    environment = "cast0r-test"
  }
}