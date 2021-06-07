packer {
  required_plugins {
    web-init = {
      version = ">= 0.0.1"
      source  = "github.com/rogueflynn/web-init"
    }
  }
}

source "web-init" "test" {}

build {
  sources = [
    "web-init.test"
  ]
}