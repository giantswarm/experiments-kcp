# create operators namespace where platforms operators will create objects. This could be many of them.
# for this example we will use only one.
kubectl ws use :root
# create provider workspaces
kubectl create workspace providers --enter
kubectl create workspace mounts --enter

for file in $(find config/mounts/resources -name 'apiresource*' -o -name 'apiexport*'); do
  kubectl apply -f ${file}
done

echo "NOW RUN THE PROXY"
