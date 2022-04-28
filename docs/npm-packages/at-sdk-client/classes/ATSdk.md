[@spiralmindjp/at-sdk-client](../README.md) / [Exports](../modules.md) / ATSdk

# Class: ATSdk

## Table of contents

### Constructors

- [constructor](ATSdk.md#constructor)

### Properties

- [apiClientManager](ATSdk.md#apiclientmanager)
- [avatarDisplayManager](ATSdk.md#avatardisplaymanager)
- [faceMeshManager](ATSdk.md#facemeshmanager)
- [sdkConfig](ATSdk.md#sdkconfig)

### Methods

- [convertFaceDataToMotionPacket](ATSdk.md#convertfacedatatomotionpacket)
- [defaultCreateHTMLElementsCallbackForAvatarDisplay](ATSdk.md#defaultcreatehtmlelementscallbackforavatardisplay)
- [defaultCreateHTMLElementsCallbackForFaceMesh](ATSdk.md#defaultcreatehtmlelementscallbackforfacemesh)
- [initializeAvatarDisplay](ATSdk.md#initializeavatardisplay)
- [initializeFaceMesh](ATSdk.md#initializefacemesh)
- [onMotion](ATSdk.md#onmotion)
- [onRoomEvent](ATSdk.md#onroomevent)
- [startAvatarDisplay](ATSdk.md#startavatardisplay)
- [startFacemesh](ATSdk.md#startfacemesh)
- [submitTasksForLoadingContentsAndVRMAvatar](ATSdk.md#submittasksforloadingcontentsandvrmavatar)

## Constructors

### constructor

• **new ATSdk**(`sdkConfig`)

#### Parameters

| Name | Type |
| :------ | :------ |
| `sdkConfig` | [`ATSdkConfigItems`](../interfaces/ATSdkConfigItems.md) |

## Properties

### apiClientManager

• **apiClientManager**: `ApiClientManager`

___

### avatarDisplayManager

• **avatarDisplayManager**: `ATAvatarDisplay`

___

### faceMeshManager

• **faceMeshManager**: ``null`` \| `ATFaceMeshManager`

___

### sdkConfig

• **sdkConfig**: [`ATSdkConfigItems`](../interfaces/ATSdkConfigItems.md)

## Methods

### convertFaceDataToMotionPacket

▸ **convertFaceDataToMotionPacket**(`faceData`): `Motion`

Convert processed face data (from facemesh module) to motion data (to be sent across the network).
Typically, user code will call this function in the processedFaceDataCallback (that was specified
in the config object), and then user code will send the motion data over the network.

#### Parameters

| Name | Type |
| :------ | :------ |
| `faceData` | `ProcessedFaceData` |

#### Returns

`Motion`

___

### defaultCreateHTMLElementsCallbackForAvatarDisplay

▸ **defaultCreateHTMLElementsCallbackForAvatarDisplay**(): [`HtmlElementIdListForAvatarDisplay`](../interfaces/HtmlElementIdListForAvatarDisplay.md)

(Internal function) This function is currently public but may be changed to private in a future version.
This function creates additional DOM elements that are required for the avatar display.
The additional DOM elements are created inside the user-provided container div (that was specified
in the config object). This default function can be overridden with a custom function in the config object,
to allow the user full-control over the additional DOM objects that are created.

#### Returns

[`HtmlElementIdListForAvatarDisplay`](../interfaces/HtmlElementIdListForAvatarDisplay.md)

___

### defaultCreateHTMLElementsCallbackForFaceMesh

▸ **defaultCreateHTMLElementsCallbackForFaceMesh**(): [`HtmlElementIdListForFaceMesh`](../interfaces/HtmlElementIdListForFaceMesh.md)

(Internal function) This function is currently public but may be changed to private in a future version.
This function creates additional DOM elements that are required for facemesh processing.
The additional DOM elements are created inside the user-provided container div (that was specified
in the config object). This default function can be overridden with a custom function in the config object,
to allow the user full-control over the additional DOM objects that are created.

#### Returns

[`HtmlElementIdListForFaceMesh`](../interfaces/HtmlElementIdListForFaceMesh.md)

___

### initializeAvatarDisplay

▸ **initializeAvatarDisplay**(): `void`

Initialize the avatar display. This will create additional DOM elements inside of the user-provided container div
(that was specified in the config object), start the Unity avatar display application and finally call the
user-provided avatarDisplayIsReadyCallback function (that was specified in the config object).

#### Returns

`void`

___

### initializeFaceMesh

▸ **initializeFaceMesh**(): `void`

Initialize the facemesh processing. This will create additional DOM elements inside of the user-provided container div
(that was specified in the config object), start the facemesh webworker thread, start the web camera, then pause
the web camera. This places the facemesh code in "standby" mode, consuming few resources but ready to start quickly.
Later (for example, when the operator joins a room), the user code can quickly unpause the web camera
and the facemesh processing will begin. When a face is recognized, the data will be provided to the
user-provided processedFaceDataCallback function (that was specified in the config object).

#### Returns

`void`

___

### onMotion

▸ **onMotion**(`motion`): `void`

AT SDK function to handle motion events received by the customer client, and to
apply the motion data to the customer client's avatar display as required.
When user code handles motion events, the user code should also call this onMotion
function to allow the AT SDK also to handle the motion event.

#### Parameters

| Name | Type |
| :------ | :------ |
| `motion` | `Motion` |

#### Returns

`void`

___

### onRoomEvent

▸ **onRoomEvent**(`roomEvent`): `Promise`<`void`\>

AT SDK function to handle room events and to load, show, or hide the avatar as required.
When user code handles room events, the user code should also call this onRoomEvent
function to allow the AT SDK also to handle the room event.

#### Parameters

| Name | Type |
| :------ | :------ |
| `roomEvent` | `RoomEvent` |

#### Returns

`Promise`<`void`\>

___

### startAvatarDisplay

▸ **startAvatarDisplay**(`unityLoaderUrls`, `htmlElementIdListForAvatarDisplay`): `void`

(Internal function) This function is currently public but may be changed to private in a future version.
Start the avatar display application. This requires that the extra DOM elements (required for the avatar display)
have already been created.

#### Parameters

| Name | Type |
| :------ | :------ |
| `unityLoaderUrls` | `UnityLoaderUrls` |
| `htmlElementIdListForAvatarDisplay` | [`HtmlElementIdListForAvatarDisplay`](../interfaces/HtmlElementIdListForAvatarDisplay.md) |

#### Returns

`void`

___

### startFacemesh

▸ **startFacemesh**(`htmlElementIdListForFaceMesh`): `Promise`<`void`\>

(Internal function) This function is currently public but may be changed to private in a future version.
Start the facemesh processing. This requires that the extra DOM elements (required for facemesh processing)
have already been created.

#### Parameters

| Name | Type |
| :------ | :------ |
| `htmlElementIdListForFaceMesh` | [`HtmlElementIdListForFaceMesh`](../interfaces/HtmlElementIdListForFaceMesh.md) |

#### Returns

`Promise`<`void`\>

___

### submitTasksForLoadingContentsAndVRMAvatar

▸ **submitTasksForLoadingContentsAndVRMAvatar**(`deviceId`, `avatarContentId`, `avatarContentKeyInIndexedDb`, `animationContentId`, `animationContentKeyInIndexedDb`): `Promise`<`void`\>

(Internal function) This function is currently public but may be changed to private in a future version.
This function submits tasks to do the following: download the specified avatar contents and animation contents
if they are not already downloaded, and load the avatar+animation contents into the Unity avatar display application.
The submitted tasks are then processed synchronously and in order.
User code should call onRoomEvent(), which will then call this function.
User code should not need to call this function directly (unless for testing purposes).

#### Parameters

| Name | Type |
| :------ | :------ |
| `deviceId` | `number` |
| `avatarContentId` | `number` |
| `avatarContentKeyInIndexedDb` | `string` |
| `animationContentId` | `number` |
| `animationContentKeyInIndexedDb` | `string` |

#### Returns

`Promise`<`void`\>
