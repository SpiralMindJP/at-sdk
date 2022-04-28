[@spiralmindjp/at-sfu-client](../README.md) / [Exports](../modules.md) / LocalStream

# Class: LocalStream

LocalStream は、ローカルの MediaStream を表します。

## Hierarchy

- `MediaStream`

  ↳ **`LocalStream`**

## Table of contents

### Constructors

- [constructor](LocalStream.md#constructor)

### Properties

- [active](LocalStream.md#active)
- [apiChannel](LocalStream.md#apichannel)
- [audio](LocalStream.md#audio)
- [constraints](LocalStream.md#constraints)
- [id](LocalStream.md#id)
- [onaddtrack](LocalStream.md#onaddtrack)
- [onremovetrack](LocalStream.md#onremovetrack)
- [pc](LocalStream.md#pc)
- [video](LocalStream.md#video)

### Methods

- [addEventListener](LocalStream.md#addeventlistener)
- [addTrack](LocalStream.md#addtrack)
- [clone](LocalStream.md#clone)
- [dispatchEvent](LocalStream.md#dispatchevent)
- [enableLayers](LocalStream.md#enablelayers)
- [getAudioTracks](LocalStream.md#getaudiotracks)
- [getTrackById](LocalStream.md#gettrackbyid)
- [getTracks](LocalStream.md#gettracks)
- [getVideoTracks](LocalStream.md#getvideotracks)
- [mute](LocalStream.md#mute)
- [publish](LocalStream.md#publish)
- [removeEventListener](LocalStream.md#removeeventlistener)
- [removeTrack](LocalStream.md#removetrack)
- [unmute](LocalStream.md#unmute)
- [unpublish](LocalStream.md#unpublish)
- [getMediaStream](LocalStream.md#getmediastream)

## Constructors

### constructor

• **new LocalStream**(`stream`, `constraints`, `video`, `audio`)

指定された MediaStream、Constraints、VideoDevice、AudioDevice で新しいインスタンスを初期化します。

#### Parameters

| Name | Type |
| :------ | :------ |
| `stream` | `MediaStream` |
| `constraints` | [`Constraints`](../interfaces/Constraints.md) |
| `video` | `Option`<`VideoDevice`\> |
| `audio` | `Option`<`AudioDevice`\> |

#### Overrides

MediaStream.constructor

## Properties

### active

• `Readonly` **active**: `boolean`

#### Inherited from

MediaStream.active

___

### apiChannel

• **apiChannel**: `Option`<`RTCDataChannel`\>

API 用の DataChannel オブジェクト。。

___

### audio

• `Readonly` **audio**: `Option`<`AudioDevice`\>

___

### constraints

• `Readonly` **constraints**: [`Constraints`](../interfaces/Constraints.md)

___

### id

• `Readonly` **id**: `string`

#### Inherited from

MediaStream.id

___

### onaddtrack

• **onaddtrack**: ``null`` \| (`this`: `MediaStream`, `ev`: `MediaStreamTrackEvent`) => `any`

#### Inherited from

MediaStream.onaddtrack

___

### onremovetrack

• **onremovetrack**: ``null`` \| (`this`: `MediaStream`, `ev`: `MediaStreamTrackEvent`) => `any`

#### Inherited from

MediaStream.onremovetrack

___

### pc

• **pc**: `Option`<`RTCPeerConnection`\>

WebRTC のコネクションオブジェクト。

___

### video

• `Readonly` **video**: `Option`<`VideoDevice`\>

## Methods

### addEventListener

▸ **addEventListener**<`K`\>(`type`, `listener`, `options?`): `void`

#### Type parameters

| Name | Type |
| :------ | :------ |
| `K` | extends keyof `MediaStreamEventMap` |

#### Parameters

| Name | Type |
| :------ | :------ |
| `type` | `K` |
| `listener` | (`this`: `MediaStream`, `ev`: `MediaStreamEventMap`[`K`]) => `any` |
| `options?` | `boolean` \| `AddEventListenerOptions` |

#### Returns

`void`

#### Inherited from

MediaStream.addEventListener

▸ **addEventListener**(`type`, `listener`, `options?`): `void`

#### Parameters

| Name | Type |
| :------ | :------ |
| `type` | `string` |
| `listener` | `EventListenerOrEventListenerObject` |
| `options?` | `boolean` \| `AddEventListenerOptions` |

#### Returns

`void`

#### Inherited from

MediaStream.addEventListener

___

### addTrack

▸ **addTrack**(`track`): `void`

#### Parameters

| Name | Type |
| :------ | :------ |
| `track` | `MediaStreamTrack` |

#### Returns

`void`

#### Inherited from

MediaStream.addTrack

___

### clone

▸ **clone**(): `MediaStream`

#### Returns

`MediaStream`

#### Inherited from

MediaStream.clone

___

### dispatchEvent

▸ **dispatchEvent**(`event`): `boolean`

Dispatches a synthetic event event to target and returns true if either event's cancelable attribute value is false or its preventDefault() method was not invoked, and false otherwise.

#### Parameters

| Name | Type |
| :------ | :------ |
| `event` | `Event` |

#### Returns

`boolean`

#### Inherited from

MediaStream.dispatchEvent

___

### enableLayers

▸ **enableLayers**(`layers`): `Promise`<`void`\>

enableLayers は、サイマルキャストのレイヤーを有効にします。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `layers` | [`Layer`](../modules.md#layer)[] | 有効にするレイヤー。 |

#### Returns

`Promise`<`void`\>

___

### getAudioTracks

▸ **getAudioTracks**(): `MediaStreamTrack`[]

#### Returns

`MediaStreamTrack`[]

#### Inherited from

MediaStream.getAudioTracks

___

### getTrackById

▸ **getTrackById**(`trackId`): ``null`` \| `MediaStreamTrack`

#### Parameters

| Name | Type |
| :------ | :------ |
| `trackId` | `string` |

#### Returns

``null`` \| `MediaStreamTrack`

#### Inherited from

MediaStream.getTrackById

___

### getTracks

▸ **getTracks**(): `MediaStreamTrack`[]

#### Returns

`MediaStreamTrack`[]

#### Inherited from

MediaStream.getTracks

___

### getVideoTracks

▸ **getVideoTracks**(): `MediaStreamTrack`[]

#### Returns

`MediaStreamTrack`[]

#### Inherited from

MediaStream.getVideoTracks

___

### mute

▸ **mute**(`kind`): `void`

mute は、指定された種類のトラックをミュートにします。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `kind` | [`TrackKind`](../modules.md#trackkind) | ミュートにするトラックの種類。 |

#### Returns

`void`

___

### publish

▸ **publish**(`transport`): `void`

publish は、ストリームを指定された Transport に公開します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `transport` | [`Transport`](Transport.md) | ストリームを公開する Transport。 |

#### Returns

`void`

___

### removeEventListener

▸ **removeEventListener**<`K`\>(`type`, `listener`, `options?`): `void`

#### Type parameters

| Name | Type |
| :------ | :------ |
| `K` | extends keyof `MediaStreamEventMap` |

#### Parameters

| Name | Type |
| :------ | :------ |
| `type` | `K` |
| `listener` | (`this`: `MediaStream`, `ev`: `MediaStreamEventMap`[`K`]) => `any` |
| `options?` | `boolean` \| `EventListenerOptions` |

#### Returns

`void`

#### Inherited from

MediaStream.removeEventListener

▸ **removeEventListener**(`type`, `listener`, `options?`): `void`

#### Parameters

| Name | Type |
| :------ | :------ |
| `type` | `string` |
| `listener` | `EventListenerOrEventListenerObject` |
| `options?` | `boolean` \| `EventListenerOptions` |

#### Returns

`void`

#### Inherited from

MediaStream.removeEventListener

___

### removeTrack

▸ **removeTrack**(`track`): `void`

#### Parameters

| Name | Type |
| :------ | :------ |
| `track` | `MediaStreamTrack` |

#### Returns

`void`

#### Inherited from

MediaStream.removeTrack

___

### unmute

▸ **unmute**(`kind`): `PromiseResult`<`void`, `Error`\>

unmute は、指定された種類のトラックのミュートを解除します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `kind` | [`TrackKind`](../modules.md#trackkind) | ミュートを解除するトラックの種類。 |

#### Returns

`PromiseResult`<`void`, `Error`\>

___

### unpublish

▸ **unpublish**(): `void`

unpublish は、ストリームを非公開にします。

#### Returns

`void`

___

### getMediaStream

▸ `Static` **getMediaStream**(`constraints`, `video`, `audio`): `PromiseResult`<[`LocalStream`](LocalStream.md), `Error`\>

指定された Constraints、VideoDevice、AudioDevice の LocalMedia を取得します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `constraints` | [`Constraints`](../interfaces/Constraints.md) | MediaStream の制限。 |
| `video` | `Option`<`VideoDevice`\> | 使用するビデオ入力デバイス。 |
| `audio` | `Option`<`AudioDevice`\> | 使用するオーディオ入力デバイス。 |

#### Returns

`PromiseResult`<[`LocalStream`](LocalStream.md), `Error`\>

取得した LocalStream を待機する Promise。
