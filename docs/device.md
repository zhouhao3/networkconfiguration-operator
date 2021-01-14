# Device Provider Specification


If you want to add support for a new device, you must comply with the following contract:

## Date Types

1. Must define an API type for “device” resources. The type:
    1. Must be implemented as a CustomResourceDefinition
    2. Must be namespace-scoped
    3. Must have the standard Kubernetes “type metadata” and “object metadata”
    4. Must have a `spec` field with the following:
        1. Required fields:
            1. `ip`(string): IP Address of the device
            2. `secret`: The username and password used to connect to the device
    5. Must implement all methods in the `DeviceManager` structure

2. Must define an API type for “configuration” resources. The type:
    1. Must be implemented as a CustomResourceDefinition
    2. Must be namespace-scoped
    3. Must have the standard Kubernetes “type metadata” and “object metadata”
    4. Must match the corresponding `Device`


## RBAC

TODO


