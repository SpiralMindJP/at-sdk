[@spiralmindjp/at-api-client](../README.md) / [Exports](../modules.md) / OperationSessionBase

# Class: OperationSessionBase

OperationSessionBase は、オペレーションのセッションを表します。

## Hierarchy

- [`SFUSessionBase`](SFUSessionBase.md)

  ↳ **`OperationSessionBase`**

  ↳↳ [`OperatorSession`](OperatorSession.md)

  ↳↳ [`CustomerSession`](CustomerSession.md)

## Table of contents

### Constructors

- [constructor](OperationSessionBase.md#constructor)

### Properties

- [client](OperationSessionBase.md#client)
- [joinedRoomState](OperationSessionBase.md#joinedroomstate)

### Accessors

- [onremovetrack](OperationSessionBase.md#onremovetrack)
- [ontrack](OperationSessionBase.md#ontrack)

### Methods

- [close](OperationSessionBase.md#close)
- [onRemoveTrack](OperationSessionBase.md#onremovetrack-1)
- [onTrack](OperationSessionBase.md#ontrack-1)
- [publish](OperationSessionBase.md#publish)

## Constructors

### constructor

• **new OperationSessionBase**(`client`, `joinedRoomState`)

指定された WebRTC SFU Client で、新しいインスタンスを初期化します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `client` | `Client` | WebRTC SFU Client。 |
| `joinedRoomState` | [`RoomState`](../interfaces/RoomState.md) | 入室したルームの状態。 |

#### Overrides

[SFUSessionBase](SFUSessionBase.md).[constructor](SFUSessionBase.md#constructor)

## Properties

### client

• `Protected` `Readonly` **client**: `Client`

#### Inherited from

[SFUSessionBase](SFUSessionBase.md).[client](SFUSessionBase.md#client)

___

### joinedRoomState

• `Readonly` **joinedRoomState**: [`RoomState`](../interfaces/RoomState.md)

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
