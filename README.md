# Asset Registry Blockchain
Asset registry blockchain application built using Ignite and Cosmos SDK

---

## Assets
An asset is defined as:
```bash
message Asset {
  
  string name = 1; 
  string description = 2; 
  string owner = 3; 
  uint64 value = 4; 
  uint64 id = 5; 
}
```

---
## Running the application
From the assetregistry folder:
```bash
ignite chain build; ignite chain serve
```

---
## CRUD Operations
You can create new assets, read a list of all assets or details of one asset, update the details of an asset, and delete assets. For example:
```bash
/Users/abi/go/bin/assetregistryd tx assetregistry create-asset "Seiko Lucent" "Seiko dress watch" alice 100 --from alice --chain-id assetregistry

/Users/abi/go/bin/assetregistryd q assetregistry list-asset

/Users/abi/go/bin/assetregistryd q assetregistry show-asset 0

/Users/abi/go/bin/assetregistryd tx assetregistry update-asset 0 "Seiko 5" "Seiko 5 SNZG15 field watch, purchased in Singapore" alice 120 --from alice --chain-id assetregistry

/Users/abi/go/bin/assetregistryd tx assetregistry delete-asset 0 --from alice --chain-id asset
```
