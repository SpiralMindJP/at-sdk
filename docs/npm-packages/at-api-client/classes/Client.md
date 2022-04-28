[@spiralmindjp/at-api-client](../README.md) / [Exports](../modules.md) / Client

# Class: Client

Client は、AT Core API のクライアントを提供します。

## Table of contents

### Constructors

- [constructor](Client.md#constructor)

### Properties

- [config](Client.md#config)

### Methods

- [eventStream](Client.md#eventstream)
- [getAvatar](Client.md#getavatar)
- [getAvatars](Client.md#getavatars)
- [getContent](Client.md#getcontent)
- [getContents](Client.md#getcontents)
- [getDeviceAvatar](Client.md#getdeviceavatar)
- [getDeviceInfo](Client.md#getdeviceinfo)
- [getRoom](Client.md#getroom)
- [getRooms](Client.md#getrooms)
- [joinRoom](Client.md#joinroom)
- [leaveRoom](Client.md#leaveroom)
- [publishVideo](Client.md#publishvideo)
- [reconnectRoom](Client.md#reconnectroom)
- [startCall](Client.md#startcall)
- [subscribeVideo](Client.md#subscribevideo)

## Constructors

### constructor

• **new Client**(`config?`)

指定された構成で、新しいインスタンスを初期化します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `config` | [`ClientConfig`](../interfaces/ClientConfig.md) | Client の構成。 |

## Properties

### config

• `Readonly` **config**: [`ClientConfig`](../interfaces/ClientConfig.md)

## Methods

### eventStream

▸ **eventStream**(`mode`): `PromiseResult`<[`EventStream`](EventStream.md), `Error`\>

eventStream は、指定されたイベント受信モードで EventStream に接続します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `mode` | [`EventStreamMode`](../enums/EventStreamMode.md) | EventStream のイベント受信モード。 |

#### Returns

`PromiseResult`<[`EventStream`](EventStream.md), `Error`\>

EventStream を待機する Promise。

___

### getAvatar

▸ **getAvatar**(`avatarID`): `PromiseResult`<[`Avatar`](../interfaces/Avatar.md), `Error`\>

getAvatar は、指定した avatarID の Avatar を取得します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `avatarID` | `number` | 取得するアバターのアバターID。 |

#### Returns

`PromiseResult`<[`Avatar`](../interfaces/Avatar.md), `Error`\>

取得した Avatar を待機する Promise。

___

### getAvatars

▸ **getAvatars**(): `PromiseResult`<[`Avatar`](../interfaces/Avatar.md)[], `Error`\>

getAvatars は、アバターの一覧を取得します。

#### Returns

`PromiseResult`<[`Avatar`](../interfaces/Avatar.md)[], `Error`\>

取得したアバター一覧の Avatar[] を待機する Promise。

___

### getContent

▸ **getContent**(`contentID`): `PromiseResult`<[`Content`](../interfaces/Content.md), `Error`\>

getContent は、指定した contentID の Content を取得します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `contentID` | `number` | 取得するコンテンツのルームID。 |

#### Returns

`PromiseResult`<[`Content`](../interfaces/Content.md), `Error`\>

取得した Content を待機する Promise。

___

### getContents

▸ **getContents**(): `PromiseResult`<[`Content`](../interfaces/Content.md)[], `Error`\>

getContents は、コンテンツの一覧を取得します。

#### Returns

`PromiseResult`<[`Content`](../interfaces/Content.md)[], `Error`\>

取得したコンテンツ一覧の Content[] を待機する Promise。

___

### getDeviceAvatar

▸ **getDeviceAvatar**(`deviceID`): `PromiseResult`<[`Avatar`](../interfaces/Avatar.md), `Error`\>

getAvatar は、指定した deviceID の Avatar を取得します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `deviceID` | `number` | アバターを取得するデバイスのデバイスID。 |

#### Returns

`PromiseResult`<[`Avatar`](../interfaces/Avatar.md), `Error`\>

取得した Avatar を待機する Promise。

___

### getDeviceInfo

▸ **getDeviceInfo**(): `PromiseResult`<[`DeviceInfo`](../interfaces/DeviceInfo.md), `Error`\>

getDeviceInfo は、デバイス情報を取得します。

#### Returns

`PromiseResult`<[`DeviceInfo`](../interfaces/DeviceInfo.md), `Error`\>

取得した DeviceInfo を待機する Promise。

___

### getRoom

▸ **getRoom**(`roomID`): `PromiseResult`<[`Room`](../interfaces/Room.md), `Error`\>

getRoom は、指定した roomID の Room を取得します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `roomID` | `number` | 取得するルームのルームID。 |

#### Returns

`PromiseResult`<[`Room`](../interfaces/Room.md), `Error`\>

取得した Room を待機する Promise。

___

### getRooms

▸ **getRooms**(): `PromiseResult`<[`Room`](../interfaces/Room.md)[], `Error`\>

getRooms は、ルームの一覧を取得します。

#### Returns

`PromiseResult`<[`Room`](../interfaces/Room.md)[], `Error`\>

取得したルーム一覧の Room[] を待機する Promise。

___

### joinRoom

▸ **joinRoom**(`roomID`, `force?`): `PromiseResult`<[`OperatorSession`](OperatorSession.md), `Error`\>

joinRoom は、ルームにオペレーターデバイスを入室させます。

#### Parameters

| Name | Type | Default value | Description |
| :------ | :------ | :------ | :------ |
| `roomID` | `number` | `undefined` | 入室するルームID。 |
| `force` | `boolean` | `false` | true の場合、既に別のオペレーターデバイスがルームに入室していても、強制的にルームに入室します。 |

#### Returns

`PromiseResult`<[`OperatorSession`](OperatorSession.md), `Error`\>

OperationSession を待機する Promise。

___

### leaveRoom

▸ **leaveRoom**(): `PromiseResult`<[`RoomState`](../interfaces/RoomState.md), `Error`\>

leaveRoom は、オペレーターデバイスをルームから退室させます。

#### Returns

`PromiseResult`<[`RoomState`](../interfaces/RoomState.md), `Error`\>

退室したルームの RoomState を待機する Promise。

___

### publishVideo

▸ **publishVideo**(): `PromiseResult`<[`VideoPublisher`](VideoPublisher.md), `Error`\>

publishVideo は、ビデオ公開者として WebRTC SFU ビデオセッションに接続します。

#### Returns

`PromiseResult`<[`VideoPublisher`](VideoPublisher.md), `Error`\>

VideoPublisher を待機する Promise。

___

### reconnectRoom

▸ **reconnectRoom**(): `PromiseResult`<[`OperatorSession`](OperatorSession.md), `Error`\>

reconnectRoom は、Operator デバイスが入室済の場合にルームの WebRTC SFU セッションに再接続します。

#### Returns

`PromiseResult`<[`OperatorSession`](OperatorSession.md), `Error`\>

OperatorSession を待機する Promise。

___

### startCall

▸ **startCall**(): `PromiseResult`<[`CustomerSession`](CustomerSession.md), `Error`\>

startCall は、Customer デバイスが入室しているルームの WebRTC SFU セッションに接続します。

#### Returns

`PromiseResult`<[`CustomerSession`](CustomerSession.md), `Error`\>

CustomerSession を待機する Promise。

___

### subscribeVideo

▸ **subscribeVideo**(): `PromiseResult`<[`VideoSubscriber`](VideoSubscriber.md), `Error`\>

subscribeVideo は、ビデオ購読者として WebRTC SFU ビデオセッションに接続します。

#### Returns

`PromiseResult`<[`VideoSubscriber`](VideoSubscriber.md), `Error`\>

VideoSubscriber を待機する Promise。
