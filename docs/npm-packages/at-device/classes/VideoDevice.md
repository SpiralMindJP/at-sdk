[@spiralmindjp/at-device](../README.md) / [Exports](../modules.md) / VideoDevice

# Class: VideoDevice

VideoDevice は、ビデオデバイスを表します。

## Hierarchy

- [`Device`](Device.md)

  ↳ **`VideoDevice`**

## Table of contents

### Constructors

- [constructor](VideoDevice.md#constructor)

### Properties

- [mediaDeviceInfo](VideoDevice.md#mediadeviceinfo)

### Accessors

- [deviceID](VideoDevice.md#deviceid)
- [direction](VideoDevice.md#direction)
- [groupID](VideoDevice.md#groupid)
- [kind](VideoDevice.md#kind)
- [label](VideoDevice.md#label)

### Methods

- [getMediaStream](VideoDevice.md#getmediastream)
- [setDeviceIDToConstraints](VideoDevice.md#setdeviceidtoconstraints)
- [getAudioConstraints](VideoDevice.md#getaudioconstraints)
- [getAudioInputDevice](VideoDevice.md#getaudioinputdevice)
- [getAudioInputDevices](VideoDevice.md#getaudioinputdevices)
- [getAudioOutputDevice](VideoDevice.md#getaudiooutputdevice)
- [getAudioOutputDevices](VideoDevice.md#getaudiooutputdevices)
- [getDefaultAudioInputDevice](VideoDevice.md#getdefaultaudioinputdevice)
- [getDefaultAudioOutputDevice](VideoDevice.md#getdefaultaudiooutputdevice)
- [getDefaultVideoDevice](VideoDevice.md#getdefaultvideodevice)
- [getDevices](VideoDevice.md#getdevices)
- [getMediaStream](VideoDevice.md#getmediastream-1)
- [getVideoConstraints](VideoDevice.md#getvideoconstraints)
- [getVideoDevice](VideoDevice.md#getvideodevice)
- [getVideoDevices](VideoDevice.md#getvideodevices)

## Constructors

### constructor

• **new VideoDevice**(`info`)

指定された MediaDeviceInfo でインスタンスを初期化します。

#### Parameters

| Name | Type |
| :------ | :------ |
| `info` | `MediaDeviceInfo` |

#### Overrides

[Device](Device.md).[constructor](Device.md#constructor)

## Properties

### mediaDeviceInfo

• `Protected` **mediaDeviceInfo**: `MediaDeviceInfo`

#### Inherited from

[Device](Device.md).[mediaDeviceInfo](Device.md#mediadeviceinfo)

## Accessors

### deviceID

• `get` **deviceID**(): `string`

deviceID プロパティは、デバイスのデバイスIDを返します。

#### Returns

`string`

デバイスのデバイスID。

#### Inherited from

Device.deviceID

___

### direction

• `get` **direction**(): [`DeviceDirection`](../modules.md#devicedirection)

direction プロパティは、デバイスの入出力方向を返します。

#### Returns

[`DeviceDirection`](../modules.md#devicedirection)

デバイスの入出力方向を表す DeviceDirection。

#### Inherited from

Device.direction

___

### groupID

• `get` **groupID**(): `string`

deviceID プロパティは、デバイスのグループIDを返します。

#### Returns

`string`

デバイスのグループID。

#### Inherited from

Device.groupID

___

### kind

• `get` **kind**(): [`DeviceKind`](../modules.md#devicekind)

kind プロパティは、デバイスの種類を返します。

#### Returns

[`DeviceKind`](../modules.md#devicekind)

デバイスの種類を表す DeviceKind。

#### Inherited from

Device.kind

___

### label

• `get` **label**(): `string`

label プロパティは、デバイスのラベルを返します。

#### Returns

`string`

デバイスのラベル。

#### Inherited from

Device.label

## Methods

### getMediaStream

▸ **getMediaStream**(`constraints`): `PromiseResult`<`MediaStream`, `Error`\>

getMediaStream は、指定された MediaStreamConstraints にマッチする MediaStream を返します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `constraints` | `Option`<`MediaStreamConstraints`\> | メディアストリームの制限を指定します。 |

#### Returns

`PromiseResult`<`MediaStream`, `Error`\>

生成された MediaStream を待機する Promise。

#### Overrides

[Device](Device.md).[getMediaStream](Device.md#getmediastream)

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

#### Inherited from

[Device](Device.md).[setDeviceIDToConstraints](Device.md#setdeviceidtoconstraints)

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

#### Inherited from

[Device](Device.md).[getAudioConstraints](Device.md#getaudioconstraints)

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

#### Inherited from

[Device](Device.md).[getAudioInputDevice](Device.md#getaudioinputdevice)

___

### getAudioInputDevices

▸ `Static` **getAudioInputDevices**(): `PromiseResult`<[`AudioDevice`](AudioDevice.md)[], `Error`\>

getAudioInputDevices は、全ての利用可能なオーディオ入力デバイスの一覧を返します。

#### Returns

`PromiseResult`<[`AudioDevice`](AudioDevice.md)[], `Error`\>

オーディオ入力デバイスの一覧を格納した AudioDevice[] を待機する Promise。

#### Inherited from

[Device](Device.md).[getAudioInputDevices](Device.md#getaudioinputdevices)

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

#### Inherited from

[Device](Device.md).[getAudioOutputDevice](Device.md#getaudiooutputdevice)

___

### getAudioOutputDevices

▸ `Static` **getAudioOutputDevices**(): `PromiseResult`<[`AudioDevice`](AudioDevice.md)[], `Error`\>

getAudioOutputDevices は、全ての利用可能なオーディオ出力デバイスの一覧を返します。

#### Returns

`PromiseResult`<[`AudioDevice`](AudioDevice.md)[], `Error`\>

オーディオ出力デバイスの一覧を格納した AudioDevice[] を待機する Promise。

#### Inherited from

[Device](Device.md).[getAudioOutputDevices](Device.md#getaudiooutputdevices)

___

### getDefaultAudioInputDevice

▸ `Static` **getDefaultAudioInputDevice**(): `PromiseOption`<[`AudioDevice`](AudioDevice.md)\>

デフォルトのオーディオ入力デバイスを取得します。

#### Returns

`PromiseOption`<[`AudioDevice`](AudioDevice.md)\>

取得した AudioDevice を待機する Promise。

#### Inherited from

[Device](Device.md).[getDefaultAudioInputDevice](Device.md#getdefaultaudioinputdevice)

___

### getDefaultAudioOutputDevice

▸ `Static` **getDefaultAudioOutputDevice**(): `PromiseOption`<[`AudioDevice`](AudioDevice.md)\>

デフォルトのオーディオ出力デバイスを取得します。

#### Returns

`PromiseOption`<[`AudioDevice`](AudioDevice.md)\>

取得した AudioDevice を待機する Promise。

#### Inherited from

[Device](Device.md).[getDefaultAudioOutputDevice](Device.md#getdefaultaudiooutputdevice)

___

### getDefaultVideoDevice

▸ `Static` **getDefaultVideoDevice**(): `PromiseOption`<[`VideoDevice`](VideoDevice.md)\>

デフォルトのビデオデバイスを取得します。

#### Returns

`PromiseOption`<[`VideoDevice`](VideoDevice.md)\>

取得した VideoDevice を待機する Promise。

#### Inherited from

[Device](Device.md).[getDefaultVideoDevice](Device.md#getdefaultvideodevice)

___

### getDevices

▸ `Static` **getDevices**(): `PromiseResult`<[`Device`](Device.md)[], `Error`\>

getDevices は、全ての利用可能なデバイスの一覧を返します。

#### Returns

`PromiseResult`<[`Device`](Device.md)[], `Error`\>

デバイスの一覧を格納した Device[] を待機する Promise。

#### Inherited from

[Device](Device.md).[getDevices](Device.md#getdevices)

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

#### Inherited from

[Device](Device.md).[getMediaStream](Device.md#getmediastream-1)

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

#### Inherited from

[Device](Device.md).[getVideoConstraints](Device.md#getvideoconstraints)

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

#### Inherited from

[Device](Device.md).[getVideoDevice](Device.md#getvideodevice)

___

### getVideoDevices

▸ `Static` **getVideoDevices**(): `PromiseResult`<[`VideoDevice`](VideoDevice.md)[], `Error`\>

getVideoDevices は、全ての利用可能なビデオデバイスの一覧を返します。

#### Returns

`PromiseResult`<[`VideoDevice`](VideoDevice.md)[], `Error`\>

ビデオデバイスの一覧を格納した VideoDevice[] を待機する Promise。

#### Inherited from

[Device](Device.md).[getVideoDevices](Device.md#getvideodevices)
