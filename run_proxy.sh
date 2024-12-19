docker run --rm -d --network host --name kcp-proxy --user $(id -u):$(id -g) -v $(pwd):/data \
  --workdir=/data --entrypoint="/virtual-workspaces" \
  gsoci.azurecr.io/giantswarm/kcp-mounts:0.0.1 start \
  --kubeconfig=.kcp/admin.kubeconfig \
  --tls-cert-file=.kcp/apiserver.crt \
  --tls-private-key-file=.kcp/apiserver.key \
  --authentication-kubeconfig=.kcp/admin.kubeconfig \
  --virtual-workspaces-proxy-hostname=https://localhost:6444 \
  -v=8
