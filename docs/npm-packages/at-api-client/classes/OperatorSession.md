[@spiralmindjp/at-api-client](../README.md) / [Exports](../modules.md) / OperatorSession

# Class: OperatorSession

OperatorSession は、オペレーターのオペレーションセッションを表します。

## Hierarchy

- [`OperationSessionBase`](OperationSessionBase.md)

  ↳ **`OperatorSession`**

## Table of contents

### Constructors

- [constructor](OperatorSession.md#constructor)

### Properties

- [client](OperatorSession.md#client)
- [joinedRoomState](OperatorSession.md#joinedroomstate)

### Accessors

- [onremovetrack](OperatorSession.md#onremovetrack)
- [ontrack](OperatorSession.md#ontrack)

### Methods

- [close](OperatorSession.md#close)
- [onRemoveTrack](OperatorSession.md#onremovetrack-1)
- [onTrack](OperatorSession.md#ontrack-1)
- [playAnimation](OperatorSession.md#playanimation)
- [playContent](OperatorSession.md#playcontent)
- [publish](OperatorSession.md#publish)
- [sendMotion](OperatorSession.md#sendmotion)
- [stopAnimation](OperatorSession.md#stopanimation)
- [stopContent](OperatorSession.md#stopcontent)

## Constructors

### constructor

• **new OperatorSession**(`client`, `joinedRoomState`)

指定された WebRTC SFU Client で、新しいインスタンスを初期化します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `client` | `Client` | WebRTC SFU Client。 |
| `joinedRoomState` | [`RoomState`](../interfaces/RoomState.md) | 入室したルームの状態。 |

#### Overrides

[OperationSessionBase](OperationSessionBase.md).[constructor](OperationSessionBase.md#constructor)

## Properties

### client

• `Protected` `Readonly` **client**: `Client`

#### Inherited from

[OperationSessionBase](OperationSessionBase.md).[client](OperationSessionBase.md#client)

___

### joinedRoomState

• `Readonly` **joinedRoomState**: [`RoomState`](../interfaces/RoomState.md)

#### Inherited from

[OperationSessionBase](OperationSessionBase.md).[joinedRoomState](OperationSessionBase.md#joinedroomstate)

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

#### Inherited from

OperationSessionBase.onremovetrack

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

#### Inherited from

OperationSessionBase.ontrack

## Methods

### close

▸ **close**(): `void`

close は、WebRTC SFU 接続をクローズします。

#### Returns

`void`

#### Inherited from

[OperationSessionBase](OperationSessionBase.md).[close](OperationSessionBase.md#close)

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

#### Inherited from

[OperationSessionBase](OperationSessionBase.md).[onRemoveTrack](OperationSessionBase.md#onremovetrack-1)

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

#### Inherited from

[OperationSessionBase](OperationSessionBase.md).[onTrack](OperationSessionBase.md#ontrack-1)

___

### playAnimation

▸ **playAnimation**(`index`): `void`

playAnimation は、指定されたインデックスのアニメーションを再生します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `index` | `number` | 再生するアニメーションのインデックス。 |

#### Returns

`void`

___

### playContent

▸ **playContent**(`contentID`): `void`

playContent は、指定されたコンテンツIDのコンテンツを再生します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `contentID` | `number` | 再生するコンテンツのコンテンツID。 |

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

#### Inherited from

[OperationSessionBase](OperationSessionBase.md).[publish](OperationSessionBase.md#publish)

___

### sendMotion

▸ **sendMotion**(`motion`): `void`

sendMotion は、モーションを送信します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `motion` | [`Motion`](../interfaces/Motion.md) | 送信する Motion。 |

#### Returns

`void`

___

### stopAnimation

▸ **stopAnimation**(): `void`

stopAnimation は、アニメーションの再生を停止します。

#### Returns

`void`

___

### stopContent

▸ **stopContent**(): `void`

stopContent は、コンテンツの再生を停止します。

#### Returns

`void`
