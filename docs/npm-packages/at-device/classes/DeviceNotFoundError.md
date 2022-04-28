[@spiralmindjp/at-device](../README.md) / [Exports](../modules.md) / DeviceNotFoundError

# Class: DeviceNotFoundError

DeviceNotFoundError は、デバイスが存在しないことを表すエラーです。

## Hierarchy

- `Error`

  ↳ **`DeviceNotFoundError`**

## Table of contents

### Constructors

- [constructor](DeviceNotFoundError.md#constructor)

### Properties

- [deviceID](DeviceNotFoundError.md#deviceid)
- [message](DeviceNotFoundError.md#message)
- [name](DeviceNotFoundError.md#name)
- [stack](DeviceNotFoundError.md#stack)

## Constructors

### constructor

• **new DeviceNotFoundError**(`deviceID`)

指定されたデバイスIDでインスタンスを初期化します。

#### Parameters

| Name | Type |
| :------ | :------ |
| `deviceID` | `string` |

#### Overrides

Error.constructor

## Properties

### deviceID

• `Readonly` **deviceID**: `string`

___

### message

• **message**: `string`

#### Inherited from

Error.message

___

### name

• **name**: `string`

#### Inherited from

Error.name

___

### stack

• `Optional` **stack**: `string`

#### Inherited from

Error.stack
