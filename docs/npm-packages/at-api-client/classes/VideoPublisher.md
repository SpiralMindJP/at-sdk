[@spiralmindjp/at-api-client](../README.md) / [Exports](../modules.md) / VideoPublisher

# Class: VideoPublisher

VideoPublisher は、ビデオ公開のセッションを表します。

## Hierarchy

- [`SFUSessionBase`](SFUSessionBase.md)

  ↳ **`VideoPublisher`**

## Table of contents

### Constructors

- [constructor](VideoPublisher.md#constructor)

### Properties

- [client](VideoPublisher.md#client)

### Methods

- [close](VideoPublisher.md#close)
- [publish](VideoPublisher.md#publish)

## Constructors

### constructor

• **new VideoPublisher**(`client`)

指定された WebRTC SFU Client で、新しいインスタンスを初期化します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `client` | `Client` | WebRTC SFU Client。 |

#### Overrides

[SFUSessionBase](SFUSessionBase.md).[constructor](SFUSessionBase.md#constructor)

## Properties

### client

• `Protected` `Readonly` **client**: `Client`

#### Inherited from

[SFUSessionBase](SFUSessionBase.md).[client](SFUSessionBase.md#client)

## Methods

### close

▸ **close**(): `void`

close は、WebRTC SFU 接続をクローズします。

#### Returns

`void`

#### Inherited from

[SFUSessionBase](SFUSessionBase.md).[close](SFUSessionBase.md#close)

___

### publish

▸ **publish**(`stream`): `PromiseResult`<`void`, `Error`\>

publish は、指定された LocalStream をセションに公開します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `stream` | `LocalStream` | セッションに公開する LocalStream。 |

#### Returns

`PromiseResult`<`void`, `Error`\>

結果を待機する Promise。
