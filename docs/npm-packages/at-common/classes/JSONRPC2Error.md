[@spiralmindjp/at-common](../README.md) / [Exports](../modules.md) / JSONRPC2Error

# Class: JSONRPC2Error

JSONRPC2Error クラスは、JSON-RPC2 におけるエラーを表します。

## Hierarchy

- `Error`

  ↳ **`JSONRPC2Error`**

## Table of contents

### Constructors

- [constructor](JSONRPC2Error.md#constructor)

### Properties

- [code](JSONRPC2Error.md#code)
- [message](JSONRPC2Error.md#message)
- [name](JSONRPC2Error.md#name)
- [stack](JSONRPC2Error.md#stack)

### Methods

- [isNotFound](JSONRPC2Error.md#isnotfound)

## Constructors

### constructor

• **new JSONRPC2Error**(`code`, `message`)

JSONRPC2Error クラスのインスタンスを初期化します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `code` | `number` | JSON-RPC2 エラーコード。 |
| `message` | `string` | エラーメッセージ。 |

#### Overrides

Error.constructor

## Properties

### code

• `Readonly` **code**: `number`

___

### message

• `Readonly` **message**: `string`

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

## Methods

### isNotFound

▸ **isNotFound**(): `boolean`

isNotFound は、エラーが NotFound を表すかどうかを返します。

#### Returns

`boolean`

エラーが NotFound を表す場合は `true`、そうでない場合は `false` を返します。
