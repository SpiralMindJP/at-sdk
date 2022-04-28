[@spiralmindjp/at-sdk-client](../README.md) / [Exports](../modules.md) / ATSdkConfigItemAvatarDisplay

# Interface: ATSdkConfigItemAvatarDisplay

Interface for configuring the avatar display.

## Table of contents

### Properties

- [avatarDisplayIsReadyCallback](ATSdkConfigItemAvatarDisplay.md#avatardisplayisreadycallback)
- [createHTMLElementsCallbackForAvatarDisplay](ATSdkConfigItemAvatarDisplay.md#createhtmlelementscallbackforavatardisplay)
- [debugAvatarDisplay](ATSdkConfigItemAvatarDisplay.md#debugavatardisplay)
- [htmlDivId](ATSdkConfigItemAvatarDisplay.md#htmldivid)
- [unityIframeMode](ATSdkConfigItemAvatarDisplay.md#unityiframemode)
- [unityIframeUrlForAvatarDisplayModule](ATSdkConfigItemAvatarDisplay.md#unityiframeurlforavatardisplaymodule)
- [unityLoaderUrls](ATSdkConfigItemAvatarDisplay.md#unityloaderurls)

### Methods

- [taskProgressCallback](ATSdkConfigItemAvatarDisplay.md#taskprogresscallback)

## Properties

### avatarDisplayIsReadyCallback

• **avatarDisplayIsReadyCallback**: ``null`` \| [`VoidFunctionCallbackType`](VoidFunctionCallbackType.md)

User-provided callback function that is invoked when the avatar display application is ready to accept API commands, such as loading VRM avatar data for display.

___

### createHTMLElementsCallbackForAvatarDisplay

• **createHTMLElementsCallbackForAvatarDisplay**: ``null`` \| [`createHtmlElementsCallbackForAvatarDisplayType`](createHtmlElementsCallbackForAvatarDisplayType.md)

User-provided callback function to create additional DOM objects required for the avatar display. If user specifies null, then a default function will be used.

___

### debugAvatarDisplay

• **debugAvatarDisplay**: `boolean`

Boolean flag to enable additional console debug messages relating to avatar display

___

### htmlDivId

• **htmlDivId**: `string`

HTML ID of the user-provided container div for avatar-related DOM elements. The AT SDK will create additional DOM objects inside of this div.

___

### unityIframeMode

• **unityIframeMode**: `boolean`

Enable experimental Unity-in-iframe mode. Recommended value is false.

___

### unityIframeUrlForAvatarDisplayModule

• **unityIframeUrlForAvatarDisplayModule**: `string`

URL pointing to the at-avatar-display module folder on the web server. This folder is required when unityIframeMode is set to true.

___

### unityLoaderUrls

• **unityLoaderUrls**: `UnityLoaderUrls`

URLs pointing to the Unity data files required for the avatar display

## Methods

### taskProgressCallback

▸ **taskProgressCallback**(`id`, `name`, `elapsedTime`, `isFinished`, `remainingTaskCount`): `void`

User-provided callback function that is automatically invoked by the AT SDK when tasks are being processed (tasks such as downloading data, or loading VRM data into Unity). This allows user code to track the progress of long-running tasks.

#### Parameters

| Name | Type |
| :------ | :------ |
| `id` | `number` |
| `name` | `string` |
| `elapsedTime` | `number` |
| `isFinished` | `boolean` |
| `remainingTaskCount` | `number` |

#### Returns

`void`
