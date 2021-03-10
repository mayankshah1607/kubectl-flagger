# kubectl-flagger

A kubectl plugin for manually gating Flagger based Canary deployments.

## Prerequisites

- [flagger-loadtester](https://docs.flagger.app/usage/webhooks#load-testing)
- [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)

## Installation

```bash
# clone this repo
$ git clone https://github.com/mayankshah1607/kubectl-flagger

$ cd kubectl-flagger

# install
$ make
```

Ensure that installation is working correctly:
```bash
$ kubectl flagger --help
```

## Usage

### 1. Manual promotion
```bash
$ kubectl flagger promote [canary name] [canary namespace] [flags]

# promote canary/podinfo in namespace/app
# use `-n` to specify the namespace where flagger-loadtester is installed
$ kubectl flagger promote podinfo app -n tester
```

### 2. Manual rollback
```bash
$ kubectl flagger rollback [canary name] [canary namespace] [flags]

# abort canary/podinfo in namespace/app
# use `-n` to specify the namespace where flagger-loadtester is installed
$ kubectl flagger rollback podinfo app -n tester
```