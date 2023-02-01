# Steadybit Scaffold Extension

This Helm chart adds the Steadybit scaffold extension to your Kubernetes cluster as a deployment.

## Quick Start

### Add Steadybit Helm repository

```
helm repo add steadybit-extension-scaffold https://steadybit.github.io/extension-scaffold
helm repo update
```

### Installing the Chart

To install the chart with the name `steadybit-extension-scaffold`.

```bash
$ helm upgrade steadybit-extension-scaffold \
    --install \
    --wait \
    --timeout 5m0s \
    --create-namespace \
    --namespace steadybit-extension \
    steadybit-extension-scaffold/steadybit-extension-scaffold
```
