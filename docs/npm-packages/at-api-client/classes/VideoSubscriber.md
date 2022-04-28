[@spiralmindjp/at-api-client](../README.md) / [Exports](../modules.md) / VideoSubscriber

# Class: VideoSubscriber

VideoSubscriber は、ビデオ購読のセッションを表します。

## Hierarchy

- [`SFUSessionBase`](SFUSessionBase.md)

  ↳ **`VideoSubscriber`**

## Table of contents

### Constructors

- [constructor](VideoSubscriber.md#constructor)

### Properties

- [client](VideoSubscriber.md#client)

### Accessors

- [onremovetrack](VideoSubscriber.md#onremovetrack)
- [ontrack](VideoSubscriber.md#ontrack)

### Methods

- [close](VideoSubscriber.md#close)
- [onRemoveTrack](VideoSubscriber.md#onremovetrack-1)
- [onTrack](VideoSubscriber.md#ontrack-1)
- [subscribe](VideoSubscriber.md#subscribe)

## Constructors

### constructor

• **new VideoSubscriber**(`client`)

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

## Accessors

### onremovetrack

• `set` **onremovetrack**(`handler`): `void`

onremovetrack プロパティは、removetrack イベントを処理するための event handler です。

#### Parameters

| Name | Type |
| :------ | :------ |
| `handler` | (`track`: `MediaStreamTrack`, `stream`: `RemoteStream`) => `void` |

#### Returns

`void`

___

### ontrack

• `set` **ontrack**(`handler`): `void`

ontrack プロパティは、track イベントを処理するための event handler です。

#### Parameters

| Name | Type |
| :------ | :------ |
| `handler` | (`track`: `MediaStreamTrack`, `stream`: `RemoteStream`) => `void` |

#### Returns

`void`

## Methods

### close

▸ **close**(): `void`

close は、WebRTC SFU 接続をクローズします。

#### Returns

`void`

#### Inherited from

[SFUSessionBase](SFUSessionBase.md).[close](SFUSessionBase.md#close)

___

### onRemoveTrack

▸ `Protected` **onRemoveTrack**(`track`, `stream`): `void`

onRemoveTrack は、リモートトラックが削除された場合の処理を行います。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `track` | `MediaStreamTrack` | 削除されたリモートの MediaStreamTrack。 |
| `stream` | `RemoteStream` | track を保持する RemoteStream。 |

#### Returns

`void`

___

### onTrack

▸ `Protected` **onTrack**(`track`, `stream`): `void`

onTrack は、リモートトラックが追加された場合の処理を行います。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `track` | `MediaStreamTrack` | 追加されたリモートの MediaStreamTrack。 |
| `stream` | `RemoteStream` | track を保持する RemoteStream。 |

#### Returns

`void`

___

### subscribe

▸ **subscribe**(): `PromiseResult`<`void`, `Error`\>

subscribe は、ビデオの購読を開始します。

#### Returns

`PromiseResult`<`void`, `Error`\>
