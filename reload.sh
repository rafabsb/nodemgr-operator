# kubectl delete -f deploy/crds/rafabsb.com_v1alpha1_nodemgr_cr.yaml
# kubectl delete -f deploy/crds/rafabsb.com_nodemgrs_crd.yaml
operator-sdk generate k8s
operator-sdk generate crds
kubectl apply -f deploy/crds/rafabsb.com_nodemgrs_crd.yaml
kubectl apply -f deploy/crds/rafabsb.com_v1alpha1_nodemgr_cr.yaml
operator-sdk up local --namespace=default