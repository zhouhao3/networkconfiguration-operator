# API and Resource Definitions

## Port

### Port Spec

### Port Status

#### portID

#### portConfigurationRef

#### deviceRef

#### smartNic


### Port Status

#### state

#### configurationRef

#### vlans

### Port Example

```yaml
metaData:
  name:
  ownerRef:
  finalizers:
spec:
  portID: 0
  portConfigurationRef:
    name: sc1
    kind: SwitchPortConfiguration
    namespace: default
  deviceRef:
    name: switch0
    kind: Switch
    namespace: default
  smrtNic: false

status:
  state:
  configurationRef:
    name:
    namespace:
  vlans:
  - id: 2
    name: 
  - 3
  .....
```

## SwitchPortConfiguration

### SwitchPortConfiguration Spec

### SwitchPortConfiguration Example

```yaml
apiVersion: v1alpha1
metaData:
  name: sc1
  namespace: default
  finalizers:
    - m3m1
spec:
  acl:
    - type: // ipv4, ipv6
      action: // allow, deny
      protocol: // TCP, UDP, ICMP, ALL
      src: // xxx.xxx.xxx.xxx/xx
      srcPortRange: // 22, 22-30
      des: // xxx.xxx.xxx.xxx/xx
      desPortRange: // 22, 22-30
  trunk:
  untaggedVLAN:
  vlans:
    - id:
    - id:
```



## Switch

The fields are:

* **os** --

* **ip** --

* **port** --

* **secret** --

* **restrictedPorts** --

### Switch Example
```yaml
apiVersion: v1alpha1
kind: Switch
metadata:
  name: switch0
  namespace: default
spec:
  os: fos
  ip: 192.168.0.1
  port: 22
  secret:
    name: switch0-secret
    namespace: default
  restrictedPorts:
    - portID: 0
      disabled: false
      vlanRange: 1, 6-10
      trunkDisable: false
```

```yaml
---
apiVersion: v1
kind: Secret
metadata:
  name: switch0-secret
  namespace: default
type: Opaque
data:
  username: YWRtaW4=
  password: cGFzc3dvcmQ=
```