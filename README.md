This repo contains a set of [Valet](https://github.com/solo-io/valet) workflows (and associated resources) to setup and step through 
a growing set of demos, feature highlights, and deployment architectures. 

## How to Use This Repo

Most of these directories contains a README that can be used for manually executing the workflow. The READMEs are automatically generated by valet and shouldn't be edited directly, to maintain self-documenting workflows. 

Using Valet, each of these can be automatically executed or stepped through. By convention, there are three options:

* `valet ensure -f workflow.yaml`: This runs the entire workflow, as an automated test. This leaves the cluster in the state at the end of the workflow, so `valet teardown -f workflow.yaml` can be used to clean up. 
* `valet ensure -f 1-...yaml`: Each workflow is broken down into a few sub-workflows, which can be individually run execute specific parts of the overall workflow. 
* `valet ensure -f workflow.yaml -s`: The `-s|--step` flag tells Valet to step through the workflow and pause at each step, for the most granular execution. 

Currently this is based on the valet branch `https://github.com/solo-io/valet/tree/updates-for-ref-arch`. 

## Table of Contents

Demos:
* [Petclinic - Migrating to Microservices](demos/extend-monolith/README.md)

Feature Highlights:
* Security
  * TLS
    * Server TLS
      * [Basic](security/tls/server-tls/basic)
      * [SNI](security/tls/server-tls/sni)
    * [Upstream TLS](security/tls/upstream-tls)