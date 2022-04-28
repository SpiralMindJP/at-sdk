[@spiralmindjp/at-sfu-client](../README.md) / [Exports](../modules.md) / RemoteStream

# Interface: RemoteStream

RemoteStream は、リモートの MediaStream を表します。

## Hierarchy

- `MediaStream`

  ↳ **`RemoteStream`**

## Table of contents

### Properties

- [\_state](RemoteStream.md#_state)
- [active](RemoteStream.md#active)
- [id](RemoteStream.md#id)
- [onaddtrack](RemoteStream.md#onaddtrack)
- [onremovetrack](RemoteStream.md#onremovetrack)

### Methods

- [addEventListener](RemoteStream.md#addeventlistener)
- [addTrack](RemoteStream.md#addtrack)
- [aduio](RemoteStream.md#aduio)
- [clone](RemoteStream.md#clone)
- [deviceID](RemoteStream.md#deviceid)
- [dispatchEvent](RemoteStream.md#dispatchevent)
- [framereate](RemoteStream.md#framereate)
- [getAudioTracks](RemoteStream.md#getaudiotracks)
- [getTrackById](RemoteStream.md#gettrackbyid)
- [getTracks](RemoteStream.md#gettracks)
- [getVideoTracks](RemoteStream.md#getvideotracks)
- [mute](RemoteStream.md#mute)
- [preferFramerate](RemoteStream.md#preferframerate)
- [preferLayer](RemoteStream.md#preferlayer)
- [removeEventListener](RemoteStream.md#removeeventlistener)
- [removeTrack](RemoteStream.md#removetrack)
- [unmute](RemoteStream.md#unmute)
- [video](RemoteStream.md#video)

## Properties

### \_state

• **\_state**: `RemoteStreamState`

___

### active

• `Readonly` **active**: `boolean`

#### Inherited from

MediaStream.active

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

### aduio

▸ **aduio**(): `boolean`

オーディオが有効かどうか。

#### Returns

`boolean`

___

### clone

▸ **clone**(): `MediaStream`

#### Returns

`MediaStream`

#### Inherited from

MediaStream.clone

___

### deviceID

▸ **deviceID**(): `number`

デバイスID。

#### Returns

`number`

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

### framereate

▸ **framereate**(): [`Layer`](../modules.md#layer)

フレームレイトのレイヤー。

#### Returns

[`Layer`](../modules.md#layer)

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

### preferFramerate

▸ **preferFramerate**(`layer`): `void`

preferLayer は、最適なフレームレートを指定します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `layer` | `NoneLayer` | フレームレイトのレイヤー。 |

#### Returns

`void`

___

### preferLayer

▸ **preferLayer**(`layer`): `void`

preferLayer は、最適な解像度レイヤーを指定します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `layer` | `NoneLayer` | 解像度のレイヤー。 |

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

▸ **unmute**(`kind`): `void`

mute は、指定された種類のトラックのミュートを解除します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `kind` | [`TrackKind`](../modules.md#trackkind) | ミュートを解除するトラックの種類。 |

#### Returns

`void`

___

### video

▸ **video**(): `NoneLayer`

ビデオが有効かどうか。

#### Returns

`NoneLayer`
