# Quick Start

## Selecting a target

After you have successfully [installed kubetap](installation.md), find a Service
you'd like to tap. In this example, we use [ArgoCD](https://argoproj.github.io/)
as the target.

```sh
$ kubectl get svc -n argocd
NAME                    TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)             AGE
argocd-dex-server       ClusterIP   10.43.152.9     <none>        5556/TCP,5557/TCP   5d20h
argocd-metrics          ClusterIP   10.43.44.33     <none>        8082/TCP            5d20h
argocd-redis            ClusterIP   10.43.209.196   <none>        6379/TCP            5d20h
argocd-repo-server      ClusterIP   10.43.192.56    <none>        8081/TCP,8084/TCP   5d20h
argocd-server-metrics   ClusterIP   10.43.220.237   <none>        8083/TCP            5d20h
argocd-server           ClusterIP   10.43.118.109   <none>        80/TCP,443/TCP      5d20h
```

## Tapping the service

The proxy container (mitmproxy by default) runs as a sidecar, keeping network
traffic within the cluster. mitmproxy runs in an interactive terminal mode for
real-time packet inspection and modification.

For this example, we target the HTTPS `argocd-server` service:

```sh
$ kubectl tap on -n argocd -p 443 --https argocd-server --port-forward
Establishing port-forward tunnels to service...

Port-Forwards:

  argocd-server - https://127.0.0.1:4000

```

## Connecting to the proxy

With mitmproxy running as a terminal UI in a tmux session, you can interact with it directly.

To attach to the mitmproxy terminal:

```sh
# Find the pod name (usually the target deployment + kubetap suffix)
kubectl get pods -n argocd | grep argocd-server

# Attach to the mitmproxy tmux session
kubectl exec -it <pod-name> -c kubetap -- tmux attach-session -t mitmproxy
```

Once attached, you can:
- View live HTTP/HTTPS traffic flowing through the proxy
- Inspect request/response headers and bodies
- Modify requests before they reach the target service
- Navigate using arrow keys, press `?` for help

To detach from tmux without stopping the proxy, press `Ctrl+B` then `D`.

## Listing active taps

All active taps can be listed using the following command, which can be constrained
to a specific namespace with `-n`:

```sh
$ kubectl tap list
Tapped Namespace/Service:

argocd/argocd-server
```

## Untapping the service

Once we are finished, we can remove the proxy and revert our tap by
turning it off:

```sh
$ kubectl tap off -n argocd argocd-server
Untapped Service "argocd-server"

```
