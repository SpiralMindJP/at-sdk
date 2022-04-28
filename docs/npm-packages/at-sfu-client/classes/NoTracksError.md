[@spiralmindjp/at-sfu-client](../README.md) / [Exports](../modules.md) / NoTracksError

# Class: NoTracksError

NoTracksError はトラックが存在しないエラーを表します。

## Hierarchy

- `Error`

  ↳ **`NoTracksError`**

## Table of contents

### Constructors

- [constructor](NoTracksError.md#constructor)

### Properties

- [kind](NoTracksError.md#kind)
- [message](NoTracksError.md#message)
- [name](NoTracksError.md#name)
- [stack](NoTracksError.md#stack)

## Constructors

### constructor

• **new NoTracksError**(`kind`)

指定されたトラックの種類でインスタンスを初期化します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `kind` | [`TrackKind`](../modules.md#trackkind) | トラックの種類。 |

#### Overrides

Error.constructor

## Properties

### kind

• `Readonly` **kind**: [`TrackKind`](../modules.md#trackkind)

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
