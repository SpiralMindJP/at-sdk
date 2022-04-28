[@spiralmindjp/at-api-client](../README.md) / [Exports](../modules.md) / CustomerSession

# Class: CustomerSession

CustomerSession は、カスタマーのオペレーションセッションを表します。

## Hierarchy

- [`OperationSessionBase`](OperationSessionBase.md)

  ↳ **`CustomerSession`**

## Table of contents

### Constructors

- [constructor](CustomerSession.md#constructor)

### Properties

- [client](CustomerSession.md#client)
- [joinedRoomState](CustomerSession.md#joinedroomstate)

### Accessors

- [onmotion](CustomerSession.md#onmotion)
- [onplayanimation](CustomerSession.md#onplayanimation)
- [onplaycontent](CustomerSession.md#onplaycontent)
- [onremovetrack](CustomerSession.md#onremovetrack)
- [onstopanimation](CustomerSession.md#onstopanimation)
- [onstopcontent](CustomerSession.md#onstopcontent)
- [ontrack](CustomerSession.md#ontrack)

### Methods

- [close](CustomerSession.md#close)
- [onMotion](CustomerSession.md#onmotion-1)
- [onPlayAnimation](CustomerSession.md#onplayanimation-1)
- [onPlayContent](CustomerSession.md#onplaycontent-1)
- [onRemoveTrack](CustomerSession.md#onremovetrack-1)
- [onStopAnimation](CustomerSession.md#onstopanimation-1)
- [onStopContent](CustomerSession.md#onstopcontent-1)
- [onTrack](CustomerSession.md#ontrack-1)
- [publish](CustomerSession.md#publish)

## Constructors

### constructor

• **new CustomerSession**(`client`, `joinedRoomState`)

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

### onmotion

• `set` **onmotion**(`handler`): `void`

onmotion プロパティは、motion イベントを処理するための event handler です。

#### Parameters

| Name | Type |
| :------ | :------ |
| `handler` | (`motion`: [`Motion`](../interfaces/Motion.md)) => `void` |

#### Returns

`void`

___

### onplayanimation

• `set` **onplayanimation**(`handler`): `void`

onplayanimation プロパティは、playanimation イベントを処理するための event handler です。

#### Parameters

| Name | Type |
| :------ | :------ |
| `handler` | (`index`: `number`) => `void` |

#### Returns

`void`

___

### onplaycontent

• `set` **onplaycontent**(`handler`): `void`

onplaycontent プロパティは、playcontent イベントを処理するための event handler です。

#### Parameters

| Name | Type |
| :------ | :------ |
| `handler` | (`contentID`: `number`) => `void` |

#### Returns

`void`

___

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

### onstopanimation

• `set` **onstopanimation**(`handler`): `void`

onstopanimation プロパティは、stopanimation イベントを処理するための event handler です。

#### Parameters

| Name | Type |
| :------ | :------ |
| `handler` | () => `void` |

#### Returns

`void`

___

### onstopcontent

• `set` **onstopcontent**(`handler`): `void`

onstopcontent プロパティは、stopcontent イベントを処理するための event handler です。

#### Parameters

| Name | Type |
| :------ | :------ |
| `handler` | () => `void` |

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

### onMotion

▸ `Protected` **onMotion**(`motion`): `void`

onMotion は、モーションを受信した場合の処理を行います。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `motion` | [`Motion`](../interfaces/Motion.md) | 受信した Motion。 |

#### Returns

`void`

___

### onPlayAnimation

▸ `Protected` **onPlayAnimation**(`index`): `void`

onPlayAnimation は、アニメーションの再生を受信した場合の処理を行います。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `index` | `number` | 再生するアニメーションのインデックス。 |

#### Returns

`void`

___

### onPlayContent

▸ `Protected` **onPlayContent**(`contentID`): `void`

onPlayContent は、コンテンツの再生を受信した場合の処理を行います。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `contentID` | `number` | 再生するコンテンツのコンテンツID。 |

#### Returns

`void`

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

### onStopAnimation

▸ `Protected` **onStopAnimation**(): `void`

onStopAnimation は、アニメーションの再生停止を受信した場合の処理を行います。

#### Returns

`void`

___

### onStopContent

▸ `Protected` **onStopContent**(): `void`

onStopContent は、コンテンツの再生停止を受信した場合の処理を行います。

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

#### Inherited from

[OperationSessionBase](OperationSessionBase.md).[onTrack](OperationSessionBase.md#ontrack-1)

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
