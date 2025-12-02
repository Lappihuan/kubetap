
<div align="center">
  <img src="img/logo.png" alt="mittens" width="600" />
</div>

# mittens
A kubectl plugin for intercepting HTTP traffic to Kubernetes Services using mitmproxy.

[![Build status][shield-build-status]][build-status]
[![Latest release][shield-latest-release]][latest-release]
[![License][shield-license]][license]

## Usage

```sh
kubectl mittens SERVICE [OPTIONS]
```

**Examples:**
```sh
kubectl mittens my-service -n my-namespace          # Auto-detect port
kubectl mittens my-service -p 8080                  # Explicit port
kubectl mittens my-service -p 443 --https           # HTTPS service
```

**Options:**
- `-n, --namespace STRING`: Target namespace
- `-p, --port INT`: Service port (auto-detected if omitted)
- `--https`: Enable for HTTPS services
- `-i, --image STRING`: Custom proxy image
- `--command-args STRING`: Custom mitmproxy arguments

**What happens:**
1. Deploy mitmproxy sidecar to service pods
2. Redirect traffic through mitmproxy
3. Open interactive mitmproxy TUI
4. Auto-cleanup on exit (Ctrl+C)

## Installation

**Binary:** Download from [Releases](https://github.com/Lappihuan/mittens/releases)

**From source:** `go install github.com/Lappihuan/mittens/cmd/kubectl-mittens@latest`

**With Krew:** `kubectl krew install mittens`

## K9s Integration

Add to `~/.k9s/plugins.yaml`:

```yaml
plugins:
  mittens:
    shortCut: Ctrl-T
    description: "mittens: inject mitmproxy sidecar"
    scopes:
      - services
    command: kubectl
    background: false
    args:
      - mittens
      - $NAME
      - -n
      - $NAMESPACE
```

## GitOps Integration (ArgoCD/Flux)

When using mittens with GitOps tools, Service port modifications may be reconciled back to the desired state. mittens handles this automatically for Flux. For ArgoCD, manual configuration is required.

### Flux

mittens automatically adds `helm.toolkit.fluxcd.io/driftDetection: disabled` annotation to tapped Services, preventing drift detection from rolling back changes. No configuration needed.

### ArgoCD

Add `ignoreDifferences` to your Application to prevent auto-sync from reverting the port changes:

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: my-app
spec:
  ignoreDifferences:
    - group: ""
      kind: Service
      name: my-service
      namespace: default
      jsonPointers:
        - /spec/ports/0/targetPort
```

## License

Apache 2.0. See [LICENSE](LICENSE).

mittens is a fork of [kubetap](https://github.com/soluble-ai/kubetap) by Soluble Inc.

[shield-build-status]: https://github.com/Lappihuan/mittens/workflows/build/badge.svg?branch=master
[shield-latest-release]: https://img.shields.io/github/v/release/Lappihuan/mittens?include_prereleases&label=release&sort=semver
[shield-license]: https://img.shields.io/github/license/Lappihuan/mittens.svg
[license]: https://github.com/Lappihuan/mittens/blob/master/LICENSE
[latest-release]: https://github.com/Lappihuan/mittens/releases
[build-status]: https://github.com/Lappihuan/mittens/actions
