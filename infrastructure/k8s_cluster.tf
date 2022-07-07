resource "azurerm_kubernetes_cluster" "cast0r-kubernetes-cluster" {
  name                = "cast0r_kubernetes_cluster-aks0"
  location            = azurerm_resource_group.cast0r-resource-group.location
  resource_group_name = azurerm_resource_group.cast0r-resource-group.name
  dns_prefix          = "cast0raks0"

  default_node_pool {
    name       = "default"
    node_count = 1
    vm_size    = "Standard_D2_v2"
  }

  identity {
    type = "SystemAssigned"
  }

  tags = {
    Environment = "cast0r-test"
  }
}

resource "azurerm_kubernetes_cluster_node_pool" "cast0r-kubernetes-cluster-node-pool" {
  name                  = "internal"
  kubernetes_cluster_id = azurerm_kubernetes_cluster.cast0r-kubernetes-cluster.id
  vm_size               = "Standard_DS2_v2"
  node_count            = 1

  tags = {
    Environment = "cast0r-test"
  }
}

output "client_certificate" {
  value = azurerm_kubernetes_cluster.cast0r-kubernetes-cluster.kube_config.0.client_certificate
}

output "kube_config" {
  value = azurerm_kubernetes_cluster.cast0r-kubernetes-cluster.kube_config_raw
  sensitive = true
}