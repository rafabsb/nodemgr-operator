# Nodemgr-operator

Its a kubernetes operator to manage nodes metadata.

```shell
operator-sdk new nodemgr-operator --repo github.com/rafabsb/nodemgr-operator
operator-sdk add api --api-version=rafabsb.com/v1alpha1 --kind=Nodemgr
operator-sdk add controller --api-version=rafabsb.com/v1alpha1 --kind=Nodemgr
```
