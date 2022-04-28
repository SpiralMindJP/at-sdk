[@spiralmindjp/at-api-client](../README.md) / [Exports](../modules.md) / APICallError

# Class: APICallError

APICallError は、AT Core API の呼び出しエラーを表します。

## Hierarchy

- `Error`

  ↳ **`APICallError`**

## Table of contents

### Constructors

- [constructor](APICallError.md#constructor)

### Properties

- [message](APICallError.md#message)
- [name](APICallError.md#name)
- [stack](APICallError.md#stack)
- [status](APICallError.md#status)

## Constructors

### constructor

• **new APICallError**(`res`)

指定された Response で、新しいインスタンスを初期化します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `res` | `Response` | API の呼び出し応答。 |

#### Overrides

Error.constructor

## Properties

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

___

### status

• `Readonly` **status**: `number`

エラーステータス。
