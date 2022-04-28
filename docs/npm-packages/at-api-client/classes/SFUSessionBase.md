[@spiralmindjp/at-api-client](../README.md) / [Exports](../modules.md) / SFUSessionBase

# Class: SFUSessionBase

SFUSessionBase クラスは、WebRTC SFU セッションの基底となる抽象クラスです。

## Hierarchy

- **`SFUSessionBase`**

  ↳ [`VideoPublisher`](VideoPublisher.md)

  ↳ [`VideoSubscriber`](VideoSubscriber.md)

  ↳ [`OperationSessionBase`](OperationSessionBase.md)

## Table of contents

### Constructors

- [constructor](SFUSessionBase.md#constructor)

### Properties

- [client](SFUSessionBase.md#client)

### Methods

- [close](SFUSessionBase.md#close)

## Constructors

### constructor

• **new SFUSessionBase**(`client`)

指定された WebRTC SFU Client で、新しいインスタンスを初期化します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `client` | `Client` | WebRTC SFU Client。 |

## Properties

### client

• `Protected` `Readonly` **client**: `Client`

## Methods

### close

▸ **close**(): `void`

close は、WebRTC SFU 接続をクローズします。

#### Returns

`void`
