[@spiralmindjp/at-device](../README.md) / [Exports](../modules.md) / Device

# Class: Device

Device クラスは、デバイスを表す抽象クラスです。

## Hierarchy

- **`Device`**

  ↳ [`VideoDevice`](VideoDevice.md)

  ↳ [`AudioDevice`](AudioDevice.md)

## Table of contents

### Constructors

- [constructor](Device.md#constructor)

### Properties

- [mediaDeviceInfo](Device.md#mediadeviceinfo)

### Accessors

- [deviceID](Device.md#deviceid)
- [direction](Device.md#direction)
- [groupID](Device.md#groupid)
- [kind](Device.md#kind)
- [label](Device.md#label)

### Methods

- [getMediaStream](Device.md#getmediastream)
- [setDeviceIDToConstraints](Device.md#setdeviceidtoconstraints)
- [getAudioConstraints](Device.md#getaudioconstraints)
- [getAudioInputDevice](Device.md#getaudioinputdevice)
- [getAudioInputDevices](Device.md#getaudioinputdevices)
- [getAudioOutputDevice](Device.md#getaudiooutputdevice)
- [getAudioOutputDevices](Device.md#getaudiooutputdevices)
- [getDefaultAudioInputDevice](Device.md#getdefaultaudioinputdevice)
- [getDefaultAudioOutputDevice](Device.md#getdefaultaudiooutputdevice)
- [getDefaultVideoDevice](Device.md#getdefaultvideodevice)
- [getDevices](Device.md#getdevices)
- [getMediaStream](Device.md#getmediastream-1)
- [getVideoConstraints](Device.md#getvideoconstraints)
- [getVideoDevice](Device.md#getvideodevice)
- [getVideoDevices](Device.md#getvideodevices)

## Constructors

### constructor

• **new Device**(`mediaDeviceInfo`)

指定された MediaDeviceInfo でインスタンスを初期化します。

#### Parameters

| Name | Type |
| :------ | :------ |
| `mediaDeviceInfo` | `MediaDeviceInfo` |

## Properties

### mediaDeviceInfo

• `Protected` **mediaDeviceInfo**: `MediaDeviceInfo`

## Accessors

### deviceID

• `get` **deviceID**(): `string`

deviceID プロパティは、デバイスのデバイスIDを返します。

#### Returns

`string`

デバイスのデバイスID。

___

### direction

• `get` **direction**(): [`DeviceDirection`](../modules.md#devicedirection)

direction プロパティは、デバイスの入出力方向を返します。

#### Returns

[`DeviceDirection`](../modules.md#devicedirection)

デバイスの入出力方向を表す DeviceDirection。

___

### groupID

• `get` **groupID**(): `string`

deviceID プロパティは、デバイスのグループIDを返します。

#### Returns

`string`

デバイスのグループID。

___

### kind

• `get` **kind**(): [`DeviceKind`](../modules.md#devicekind)

kind プロパティは、デバイスの種類を返します。

#### Returns

[`DeviceKind`](../modules.md#devicekind)

デバイスの種類を表す DeviceKind。

___

### label

• `get` **label**(): `string`

label プロパティは、デバイスのラベルを返します。

#### Returns

`string`

デバイスのラベル。

## Methods

### getMediaStream

▸ `Abstract` **getMediaStream**(`constraints`): `PromiseResult`<`MediaStream`, `Error`\>

#### Parameters

| Name | Type |
| :------ | :------ |
| `constraints` | `Option`<`MediaStreamConstraints`\> |

#### Returns

`PromiseResult`<`MediaStream`, `Error`\>

___

### setDeviceIDToConstraints

▸ `Protected` **setDeviceIDToConstraints**(`constraints`, `exact`): `void`

#### Parameters

| Name | Type |
| :------ | :------ |
| `constraints` | `MediaTrackConstraints` |
| `exact` | `boolean` |

#### Returns

`void`

___

### getAudioConstraints

▸ `Static` `Protected` **getAudioConstraints**(`audio`, `constraints`): `Option`<`MediaTrackConstraints`\>

#### Parameters

| Name | Type |
| :------ | :------ |
| `audio` | `Option`<[`AudioDevice`](AudioDevice.md)\> |
| `constraints` | `Option`<`MediaStreamConstraints`\> |

#### Returns

`Option`<`MediaTrackConstraints`\>

___

### getAudioInputDevice

▸ `Static` **getAudioInputDevice**(`deviceID`): `PromiseResult`<[`AudioDevice`](AudioDevice.md), `Error`\>

指定されたデバイスIDのオーディオ入力デバイスを取得します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `deviceID` | `string` | 取得するオーディオ入力デバイスのデバイスID。 |

#### Returns

`PromiseResult`<[`AudioDevice`](AudioDevice.md), `Error`\>

取得した AudioDevice を待機する Promise。

___

### getAudioInputDevices

▸ `Static` **getAudioInputDevices**(): `PromiseResult`<[`AudioDevice`](AudioDevice.md)[], `Error`\>

getAudioInputDevices は、全ての利用可能なオーディオ入力デバイスの一覧を返します。

#### Returns

`PromiseResult`<[`AudioDevice`](AudioDevice.md)[], `Error`\>

オーディオ入力デバイスの一覧を格納した AudioDevice[] を待機する Promise。

___

### getAudioOutputDevice

▸ `Static` **getAudioOutputDevice**(`deviceID`): `PromiseResult`<[`AudioDevice`](AudioDevice.md), `Error`\>

指定されたデバイスIDのオーディオ出力デバイスを取得します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `deviceID` | `string` | 取得するオーディオ出力デバイスのデバイスID。 |

#### Returns

`PromiseResult`<[`AudioDevice`](AudioDevice.md), `Error`\>

取得した AudioDevice を待機する Promise。

___

### getAudioOutputDevices

▸ `Static` **getAudioOutputDevices**(): `PromiseResult`<[`AudioDevice`](AudioDevice.md)[], `Error`\>

getAudioOutputDevices は、全ての利用可能なオーディオ出力デバイスの一覧を返します。

#### Returns

`PromiseResult`<[`AudioDevice`](AudioDevice.md)[], `Error`\>

オーディオ出力デバイスの一覧を格納した AudioDevice[] を待機する Promise。

___

### getDefaultAudioInputDevice

▸ `Static` **getDefaultAudioInputDevice**(): `PromiseOption`<[`AudioDevice`](AudioDevice.md)\>

デフォルトのオーディオ入力デバイスを取得します。

#### Returns

`PromiseOption`<[`AudioDevice`](AudioDevice.md)\>

取得した AudioDevice を待機する Promise。

___

### getDefaultAudioOutputDevice

▸ `Static` **getDefaultAudioOutputDevice**(): `PromiseOption`<[`AudioDevice`](AudioDevice.md)\>

デフォルトのオーディオ出力デバイスを取得します。

#### Returns

`PromiseOption`<[`AudioDevice`](AudioDevice.md)\>

取得した AudioDevice を待機する Promise。

___

### getDefaultVideoDevice

▸ `Static` **getDefaultVideoDevice**(): `PromiseOption`<[`VideoDevice`](VideoDevice.md)\>

デフォルトのビデオデバイスを取得します。

#### Returns

`PromiseOption`<[`VideoDevice`](VideoDevice.md)\>

取得した VideoDevice を待機する Promise。

___

### getDevices

▸ `Static` **getDevices**(): `PromiseResult`<[`Device`](Device.md)[], `Error`\>

getDevices は、全ての利用可能なデバイスの一覧を返します。

#### Returns

`PromiseResult`<[`Device`](Device.md)[], `Error`\>

デバイスの一覧を格納した Device[] を待機する Promise。

___

### getMediaStream

▸ `Static` **getMediaStream**(`video`, `audio`, `constraints`): `PromiseResult`<`MediaStream`, `Error`\>

getMediaStream は、指定された VideoDevice、AudioDevice、MediaStreamConstraints にマッチする MediaStream を返します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `video` | `Option`<[`VideoDevice`](VideoDevice.md)\> | MediaStream の、ビデオ入力デバイスを指定します。 |
| `audio` | `Option`<[`AudioDevice`](AudioDevice.md)\> | MediaStream の、オーディオ入力デバイスを指定します。 |
| `constraints` | `Option`<`MediaStreamConstraints`\> | メディアストリームの制限を指定します。 |

#### Returns

`PromiseResult`<`MediaStream`, `Error`\>

生成された MediaStream を待機する Promise。

___

### getVideoConstraints

▸ `Static` `Protected` **getVideoConstraints**(`video`, `constraints`): `Option`<`MediaTrackConstraints`\>

#### Parameters

| Name | Type |
| :------ | :------ |
| `video` | `Option`<[`VideoDevice`](VideoDevice.md)\> |
| `constraints` | `Option`<`MediaStreamConstraints`\> |

#### Returns

`Option`<`MediaTrackConstraints`\>

___

### getVideoDevice

▸ `Static` **getVideoDevice**(`deviceID`): `PromiseResult`<[`VideoDevice`](VideoDevice.md), `Error`\>

指定されたデバイスIDのビデオデバイスを取得します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `deviceID` | `string` | 取得するビデオデバイスのデバイスID。 |

#### Returns

`PromiseResult`<[`VideoDevice`](VideoDevice.md), `Error`\>

取得した VideoDevice を待機する Promise。

___

### getVideoDevices

▸ `Static` **getVideoDevices**(): `PromiseResult`<[`VideoDevice`](VideoDevice.md)[], `Error`\>

getVideoDevices は、全ての利用可能なビデオデバイスの一覧を返します。

#### Returns

`PromiseResult`<[`VideoDevice`](VideoDevice.md)[], `Error`\>

ビデオデバイスの一覧を格納した VideoDevice[] を待機する Promise。
