PoC to try out [csi-lib-utils/protosanitizer](https://github.com/kubernetes-csi/csi-lib-utils/tree/master/protosanitizer). 

### Output

```
% go run main.go
I0508 16:59:55.173250   42561 main.go:21] "Incoming NodePublishVolume request" pod-name="" namespace="" uid="" unstripped="secrets:<key:\"password\" value:\"foobar\" > volume_context:<key:\"mount\" value:\"true\" > "
I0508 16:59:55.174179   42561 main.go:22] "Incoming NodePublishVolume request" pod-name="" namespace="" uid="" stripped="{\"secrets\":\"***stripped***\",\"volume_context\":{\"mount\":\"true\"}}"
```