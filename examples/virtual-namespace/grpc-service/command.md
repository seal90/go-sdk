# config.yaml
```yaml
apiVersion: dapr.io/v1alpha1
kind: Configuration 
metadata:
  name: appconfig
spec:
  nameResolution:
    component: "configurestring"
    version: "v1"
    configuration:
      valueType: "jsonStringValue"
      value: {"hello":[{"ipv4":"127.0.0.1","ipv6":"","port":"7000","extendedInfo":{}},{"ipv4":"127.0.0.1","ipv6":"","port":"7001","extendedInfo":{"virtual-namespace":"gray"}}],"world":[]}
```

# vscode 
```json
["run", "--app-id", "grpc-server", "--app-protocol", "grpc",
                 "--app-port", "50051", "--dapr-grpc-port", "50007", "--config", "/home/seal/dapr_dev/config.yaml"]
```

# cmd
```shell
../dapr/cmd/daprd/debug_bin run --app-id grpc-server          --app-protocol grpc          --app-port 50051          --dapr-grpc-port 50007          --log-level debug           --config ~/dapr_dev/config.yaml
```