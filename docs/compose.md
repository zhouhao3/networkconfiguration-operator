## Configuration object

### Required spec fields

<!-- TODO -->
The status **spec** must have several fields defined:

-
-

### Required status fields

<!-- TODO -->
The status **status** must have several fields defined:

-
-

## Device object

### Required spec fields

<!-- TODO -->
The status **spec** must have several fields defined:

-
-

### Required status fields

The status **status** must have several fields defined:

- os - a string field, operator system of the device, it's different for every brand.
- ip - a string field, the device's ip address.
- secret - a secret's reference, access the device credentials.

## Secrets

The secret's field of the device.

|Field name|Content|
|:-:|:-:|
|username|base64 encoded username|
|password|base64 encoded password|
