[@spiralmindjp/at-sfu-client](../README.md) / [Exports](../modules.md) / Client

# Class: Client

Client は、WebRTC SFU 接続のクライアントを提供します。

## Table of contents

### Constructors

- [constructor](Client.md#constructor)

### Properties

- [config](Client.md#config)
- [signal](Client.md#signal)

### Accessors

- [onactivelayer](Client.md#onactivelayer)
- [onerrornegotiate](Client.md#onerrornegotiate)
- [onmotion](Client.md#onmotion)
- [onoperation](Client.md#onoperation)
- [onspeaker](Client.md#onspeaker)
- [ontrack](Client.md#ontrack)

### Methods

- [close](Client.md#close)
- [join](Client.md#join)
- [leave](Client.md#leave)
- [onActiveLayer](Client.md#onactivelayer-1)
- [onErrorNegotiate](Client.md#onerrornegotiate-1)
- [onMotion](Client.md#onmotion-1)
- [onOperation](Client.md#onoperation-1)
- [onSpeaker](Client.md#onspeaker-1)
- [onTrack](Client.md#ontrack-1)
- [publish](Client.md#publish)
- [sendMotion](Client.md#sendmotion)
- [sendOperation](Client.md#sendoperation)

## Constructors

### constructor

• **new Client**(`signal`, `config?`)

指定されたシグナリングクライアント、構成で、新しいインスタンスを初期化します。

#### Parameters

| Name | Type |
| :------ | :------ |
| `signal` | [`Signal`](Signal.md) |
| `config` | [`Configuration`](../interfaces/Configuration.md) |

## Properties

### config

• `Readonly` **config**: [`Configuration`](../interfaces/Configuration.md)

___

### signal

• `Readonly` **signal**: [`Signal`](Signal.md)

## Accessors

### onactivelayer

• `set` **onactivelayer**(`onactivelayer`): `void`

onactivelayer プロパティは、activelayer イベントを処理するための event handler です。

#### Parameters

| Name | Type |
| :------ | :------ |
| `onactivelayer` | (`layer`: [`ActiveLayer`](../interfaces/ActiveLayer.md)) => `void` |

#### Returns

`void`

___

### onerrornegotiate

• `set` **onerrornegotiate**(`onerrornegotiate`): `void`

onerrornegotiate プロパティは、errornegotiate イベントを処理するための event handler です。

#### Parameters

| Name | Type |
| :------ | :------ |
| `onerrornegotiate` | (`role`: [`Role`](../enums/Role.md), `err`: `Error`, `offer`: `Option`<`RTCSessionDescriptionInit`\>, `answer`: `Option`<`RTCSessionDescriptionInit`\>) => `void` |

#### Returns

`void`

___

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

### onoperation

• `set` **onoperation**(`handler`): `void`

onoperation プロパティは、operation イベントを処理するための event handler です。

#### Parameters

| Name | Type |
| :------ | :------ |
| `handler` | (`operation`: [`Operation`](../interfaces/Operation.md)) => `void` |

#### Returns

`void`

___

### onspeaker

• `set` **onspeaker**(`onspeaker`): `void`

onspeaker プロパティは、speaker イベントを処理するための event handler です。

#### Parameters

| Name | Type |
| :------ | :------ |
| `onspeaker` | (`ev`: `string`[]) => `void` |

#### Returns

`void`

___

### ontrack

• `set` **ontrack**(`ontrack`): `void`

ontrack プロパティは、track イベントを処理するための event handler です。

#### Parameters

| Name | Type |
| :------ | :------ |
| `ontrack` | (`track`: `MediaStreamTrack`, `stream`: [`RemoteStream`](../interfaces/RemoteStream.md)) => `void` |

#### Returns

`void`

## Methods

### close

▸ **close**(): `void`

close は、接続を閉じます。

#### Returns

`void`

___

### join

▸ **join**(): `PromiseResult`<`void`, `Error`\>

join は、WebRTC SFU セッションに接続します。

#### Returns

`PromiseResult`<`void`, `Error`\>

___

### leave

▸ **leave**(`stopStream?`): `Result`<`void`, `Error`\>

leave は、WebRTC SFU セッションから退室します。

#### Parameters

| Name | Type | Default value | Description |
| :------ | :------ | :------ | :------ |
| `stopStream` | `boolean` | `true` | 公開しているローカルストリームを停止します。 |

#### Returns

`Result`<`void`, `Error`\>

___

### onActiveLayer

▸ `Protected` **onActiveLayer**(`layer`): `void`

onActiveLayer は、レイヤーがアクティブになった場合の処理を行います。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `layer` | [`ActiveLayer`](../interfaces/ActiveLayer.md) | アクティブになったレイヤー。 |

#### Returns

`void`

___

### onErrorNegotiate

▸ `Protected` **onErrorNegotiate**(`role`, `err`, `offer`, `answer`): `void`

onErrorNegotiate は、ネゴシエーションにおいて発生したエラーを処理します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `role` | [`Role`](../enums/Role.md) | ロール。 |
| `err` | `Error` | エラー。 |
| `offer` | `Option`<`RTCSessionDescriptionInit`\> | Offer。 |
| `answer` | `Option`<`RTCSessionDescriptionInit`\> | Answer。 |

#### Returns

`void`

___

### onMotion

▸ `Protected` **onMotion**(`motion`): `void`

onMotion は、モーションを受信した場合の処理を行います。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `motion` | [`Motion`](../interfaces/Motion.md) | 受信したモーション。 |

#### Returns

`void`

___

### onOperation

▸ `Protected` **onOperation**(`operation`): `void`

onOperation は、オペレーションを受信した場合の処理を行います。

#### Parameters

| Name | Type |
| :------ | :------ |
| `operation` | [`Operation`](../interfaces/Operation.md) |

#### Returns

`void`

___

### onSpeaker

▸ `Protected` **onSpeaker**(`ev`): `void`

onSpeaker は、サーバーからのスピーカーイベントを処理します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `ev` | `string`[] | イベントデータ。 |

#### Returns

`void`

___

### onTrack

▸ `Protected` **onTrack**(`track`, `stream`): `void`

onTrack は、リモートトラックが追加された場合のイベントを処理します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `track` | `MediaStreamTrack` | 追加されたリモートの MediaStreamTrack。 |
| `stream` | [`RemoteStream`](../interfaces/RemoteStream.md) | track を保持する RemoteStream。 |

#### Returns

`void`

___

### publish

▸ **publish**(`stream`): `Result`<`void`, `Error`\>

publish は、LocalStream をセッションに公開します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `stream` | [`LocalStream`](LocalStream.md) | 公開する LocalStream。 |

#### Returns

`Result`<`void`, `Error`\>

___

### sendMotion

▸ **sendMotion**(`motion`): `Result`<`void`, `Error`\>

sendMotion は、モーションを送信します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `motion` | [`Motion`](../interfaces/Motion.md) | 送信する Motion。 |

#### Returns

`Result`<`void`, `Error`\>

___

### sendOperation

▸ **sendOperation**(`operation`): `Result`<`void`, `Error`\>

sendOperation は、オペレーションを送信します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `operation` | [`Operation`](../interfaces/Operation.md) | 送信する Operation。 |

#### Returns

`Result`<`void`, `Error`\>
