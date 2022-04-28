[@spiralmindjp/at-sfu-client](../README.md) / [Exports](../modules.md) / Signal

# Class: Signal

Signal クラスは、WebRTC SFU サーバーとのシグナリングを行います。

## Hierarchy

- `JSONRPC2Base`

  ↳ **`Signal`**

## Table of contents

### Properties

- [config](Signal.md#config)
- [socket](Signal.md#socket)

### Accessors

- [onclose](Signal.md#onclose)
- [onerror](Signal.md#onerror)
- [onnegotiate](Signal.md#onnegotiate)
- [onnotify](Signal.md#onnotify)
- [ontrickle](Signal.md#ontrickle)

### Methods

- [answer](Signal.md#answer)
- [call](Signal.md#call)
- [close](Signal.md#close)
- [closeSocket](Signal.md#closesocket)
- [getDeviceInfo](Signal.md#getdeviceinfo)
- [join](Signal.md#join)
- [notify](Signal.md#notify)
- [offer](Signal.md#offer)
- [onClose](Signal.md#onclose-1)
- [onError](Signal.md#onerror-1)
- [onNegotiate](Signal.md#onnegotiate-1)
- [onNotify](Signal.md#onnotify-1)
- [onTrickle](Signal.md#ontrickle-1)
- [trickle](Signal.md#trickle)
- [connect](Signal.md#connect)
- [connectBase](Signal.md#connectbase)

## Properties

### config

• `Protected` **config**: [`SignalConfig`](../interfaces/SignalConfig.md)

___

### socket

• `Protected` **socket**: `WebSocket`

#### Inherited from

JSONRPC2Base.socket

## Accessors

### onclose

• `set` **onclose**(`handler`): `void`

#### Parameters

| Name | Type |
| :------ | :------ |
| `handler` | (`ev`: `Event`) => `void` |

#### Returns

`void`

#### Inherited from

JSONRPC2Base.onclose

___

### onerror

• `set` **onerror**(`handler`): `void`

#### Parameters

| Name | Type |
| :------ | :------ |
| `handler` | (`error`: `Event`) => `void` |

#### Returns

`void`

#### Inherited from

JSONRPC2Base.onerror

___

### onnegotiate

• `set` **onnegotiate**(`onnegotiate`): `void`

onnegotiate プロパティは、negotiate イベントを処理するための event handler です。

#### Parameters

| Name | Type |
| :------ | :------ |
| `onnegotiate` | (`desc`: `RTCSessionDescriptionInit`) => `void` |

#### Returns

`void`

___

### onnotify

• `set` **onnotify**(`handler`): `void`

#### Parameters

| Name | Type |
| :------ | :------ |
| `handler` | (`method`: `string`, `params`: `any`) => `void` |

#### Returns

`void`

#### Inherited from

JSONRPC2Base.onnotify

___

### ontrickle

• `set` **ontrickle**(`ontrickle`): `void`

ontrickle プロパティは、trickle イベントを処理するための event handler です。

#### Parameters

| Name | Type |
| :------ | :------ |
| `ontrickle` | (`trickle`: [`Trickle`](../interfaces/Trickle.md)) => `void` |

#### Returns

`void`

## Methods

### answer

▸ **answer**(`answer`): `void`

answer は、シグナリングサーバーに Answer を通知します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `answer` | `RTCSessionDescriptionInit` | Answer の Session Description。 |

#### Returns

`void`

___

### call

▸ `Protected` **call**<`T`\>(`method`, `params`): `PromiseResult`<`T`, `JSONRPC2Error`\>

#### Type parameters

| Name |
| :------ |
| `T` |

#### Parameters

| Name | Type |
| :------ | :------ |
| `method` | `string` |
| `params` | `any` |

#### Returns

`PromiseResult`<`T`, `JSONRPC2Error`\>

#### Inherited from

JSONRPC2Base.call

___

### close

▸ **close**(): `void`

#### Returns

`void`

#### Inherited from

JSONRPC2Base.close

___

### closeSocket

▸ **closeSocket**(`code?`, `reason?`): `void`

#### Parameters

| Name | Type |
| :------ | :------ |
| `code?` | `number` |
| `reason?` | `string` |

#### Returns

`void`

#### Inherited from

JSONRPC2Base.closeSocket

___

### getDeviceInfo

▸ **getDeviceInfo**(`streamID`): `Promise`<`Result`<[`DeviceInfo`](../interfaces/DeviceInfo.md), `JSONRPC2Error`\>\>

getDeviceInfo は、指定されたストリームIDに対応したデバイスの情報を取得します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `streamID` | `string` | MediaStream の ID。 |

#### Returns

`Promise`<`Result`<[`DeviceInfo`](../interfaces/DeviceInfo.md), `JSONRPC2Error`\>\>

デバイス情報の取得を待機する Promise。

___

### join

▸ **join**(`offer`): `Promise`<`Result`<`RTCSessionDescriptionInit`, `JSONRPC2Error`\>\>

join は、指定された Session Description で接続を開始します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `offer` | `RTCSessionDescriptionInit` | Offer の Session Description。 |

#### Returns

`Promise`<`Result`<`RTCSessionDescriptionInit`, `JSONRPC2Error`\>\>

シグナリングサーバーからの Anser を待機する Promise。

___

### notify

▸ `Protected` **notify**(`method`, `params`): `void`

#### Parameters

| Name | Type |
| :------ | :------ |
| `method` | `string` |
| `params` | `any` |

#### Returns

`void`

#### Inherited from

JSONRPC2Base.notify

___

### offer

▸ **offer**(`offer`): `Promise`<`Result`<`RTCSessionDescriptionInit`, `JSONRPC2Error`\>\>

offer は、シグナリングサーバーに Offer の Session Description を送信します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `offer` | `RTCSessionDescriptionInit` | Offer の Session Description。 |

#### Returns

`Promise`<`Result`<`RTCSessionDescriptionInit`, `JSONRPC2Error`\>\>

シグナリングサーバーからの Anser を待機する Promise。

___

### onClose

▸ `Protected` **onClose**(`e`): `void`

onClose は、WebSocket がクローズした場合の処理を行います。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `e` | `Event` | イベントデータ。 |

#### Returns

`void`

#### Overrides

JSONRPC2Base.onClose

___

### onError

▸ `Protected` **onError**(`e`): `void`

onError は、WebSocket でエラーが発生した場合の処理を行います。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `e` | `Event` | イベントデータ。 |

#### Returns

`void`

#### Overrides

JSONRPC2Base.onError

___

### onNegotiate

▸ `Protected` **onNegotiate**(`description`): `void`

onNotify は、サーバーから offer を受信した場合の処理を行います。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `description` | `RTCSessionDescriptionInit` | Offer の Session Description。 |

#### Returns

`void`

___

### onNotify

▸ `Protected` **onNotify**(`method`, `params`): `void`

onNotify は、リモートから notify を受信した場合の処理を行います。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `method` | `string` | メソッド名。 |
| `params` | `any` | - |

#### Returns

`void`

#### Overrides

JSONRPC2Base.onNotify

___

### onTrickle

▸ `Protected` **onTrickle**(`trickle`): `void`

onNotify は、サーバーから trickle を受信した場合の処理を行います。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `trickle` | [`Trickle`](../interfaces/Trickle.md) | Trickle ICE のメッセージ。 |

#### Returns

`void`

___

### trickle

▸ **trickle**(`trickle`): `void`

trickle は、シグナリングサーバーに Trickle ICE のメッセージを通知します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `trickle` | [`Trickle`](../interfaces/Trickle.md) | Trickle ICE のメッセージ。 |

#### Returns

`void`

___

### connect

▸ `Static` **connect**(`config`): `PromiseResult`<[`Signal`](Signal.md), `Error`\>

connect は、指定された構成でシグナリングサーバーに接続します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `config` | [`SignalConfig`](../interfaces/SignalConfig.md) | シグナリングの構成。 |

#### Returns

`PromiseResult`<[`Signal`](Signal.md), `Error`\>

接続結果を待機する Promise。

___

### connectBase

▸ `Static` `Protected` **connectBase**(`url`): `PromiseResult`<`WebSocket`, `Error`\>

#### Parameters

| Name | Type |
| :------ | :------ |
| `url` | `string` |

#### Returns

`PromiseResult`<`WebSocket`, `Error`\>

#### Inherited from

JSONRPC2Base.connectBase
