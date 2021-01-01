package main

import (
    "github.com/hashicorp/terraform/plugin"
    "./ldapcrud"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: ldapcrud.Provider,
    })
}