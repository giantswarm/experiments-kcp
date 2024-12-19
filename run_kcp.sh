docker run --rm -d --network host --name kcp-server --user $(id -u):$(id -g) \
  -v $(pwd):/data --workdir=/data \
  gsoci.azurecr.io/giantswarm/kcp-mounts:0.0.1 start \
  --mapping-file=./contrib/mounts-vw/assets/path-mapping.yaml \
  --feature-gates=WorkspaceMounts=true
