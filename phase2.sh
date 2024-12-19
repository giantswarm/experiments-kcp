kubectl ws use :root
kubectl create workspace operators --enter
kubectl create workspace mounts --enter
kubectl create -f config/mounts/resources/apibinding-targets.yaml

kind create cluster --name kind --kubeconfig kind.kubeconfig

kubectl ws use :root:operators:mounts
kubectl create secret generic kind-kubeconfig --from-file=kubeconfig=kind.kubeconfig

# create target cluster:
kubectl create -f config/mounts/resources/example-target-cluster.yaml

# create vcluster target:
kubectl create -f config/mounts/resources/example-target-vcluster.yaml

# get secret string:
secret="$(kubectl get TargetKubeCluster proxy-cluster -o jsonpath='{.status.secretString}')"

# Create a consumer workspace for mounts:
kubectl ws use :root
kubectl create workspace consumer --enter
kubectl create workspace kind-cluster
kubectl create -f config/mounts/resources/apibinding-mounts.yaml

while ! kubectl api-resources | grep -q kubeclusters; do
  sleep 1
done

# !!!!! replace secrets string first in the file bellow :
yq -i ".spec.secretString = \"${secret}\"" config/mounts/resources/example-mount-cluster.yaml
kubectl create -f config/mounts/resources/example-mount-cluster.yaml

# annotate the workspace with mount,  putting the intent that this should be mounted:
kubectl annotate workspace kind-cluster experimental.tenancy.kcp.io/mount='{"spec":{"ref":{"kind":"KubeCluster","name":"proxy-cluster","apiVersion":"mounts.contrib.kcp.io/v1alpha1"}}}'
