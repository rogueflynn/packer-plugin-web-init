# Packer Plugin web-init

This repository is a basic packer plugin that was made for the purpose of learning how to make a packer-plugin.  
It was generated from [hashicorp/packer-plugin-scaffolding](https://github.com/hashicorp/packer-plugin-scaffolding)   
All this plugin does is create a basic index.html template. Only works on Mac OS.   

- A builder ([builder/web](builder/web))
- A working example ([example](example))

A full guide to creating Packer plugins can be found at [Extending Packer](https://www.packer.io/docs/plugins/creation).

In this repository you will also find a pre-defined GitHub Action configuration for the release workflow
(`.goreleaser.yml` and `.github/workflows/release.yml`). The release workflow configuration makes sure the GitHub
release artifacts are created with the correct binaries and naming conventions.

Please see the [GitHub template repository documentation](https://docs.github.com/en/free-pro-team@latest/github/creating-cloning-and-archiving-repositories/creating-a-repository-from-a-template)
for how to create a new repository from this template on GitHub.

## Installing the plugin from this repository  
Copy prod.pkr.hcl from the example folder to your desktop and run  
```bash
packer init prod.pkr.hcl
packer build prod.pkr.hcl
```

## Building and Installing for developing 
Cloned the repo to your desktop and run this command to build and install  
```bash
make dev
```
## Running  
If everything installed correctly, you can this command to test the plugin out  
```bash
packer build examples/dev.pkr.hcl
```
This should result in a new directory called a web with index.html in the home directory  

## Packer plugin projects

Here's a non exaustive list of Packer plugins that you can checkout:

* [github.com/hashicorp/packer-plugin-docker](https://github.com/hashicorp/packer-plugin-docker)
* [github.com/exoscale/packer-plugin-exoscale](https://github.com/exoscale/packer-plugin-exoscale)
* [github.com/sylviamoss/packer-plugin-comment](https://github.com/sylviamoss/packer-plugin-comment)
* [github.com/hashicorp/packer-plugin-hashicups](https://github.com/hashicorp/packer-plugin-hashicups)

Looking at their code will give you good examples.


# Requirements

-	[packer-plugin-sdk](https://github.com/hashicorp/packer-plugin-sdk) >= v0.1.0
-	[Go](https://golang.org/doc/install) >= 1.16

## Packer Compatibility
This web-init plugin is compatible with Packer >= v1.7.0
