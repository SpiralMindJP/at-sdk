[@spiralmindjp/at-sdk-client](../README.md) / [Exports](../modules.md) / ATSdkConfigItemFaceMesh

# Interface: ATSdkConfigItemFaceMesh

Interface for configuring the facemesh processing.

## Table of contents

### Properties

- [createHTMLElementsCallbackForFaceMesh](ATSdkConfigItemFaceMesh.md#createhtmlelementscallbackforfacemesh)
- [debugFaceMesh](ATSdkConfigItemFaceMesh.md#debugfacemesh)
- [htmlDivId](ATSdkConfigItemFaceMesh.md#htmldivid)
- [processHeight](ATSdkConfigItemFaceMesh.md#processheight)
- [processWidth](ATSdkConfigItemFaceMesh.md#processwidth)
- [processedFaceDataCallback](ATSdkConfigItemFaceMesh.md#processedfacedatacallback)
- [updateInterval](ATSdkConfigItemFaceMesh.md#updateinterval)
- [useWasm](ATSdkConfigItemFaceMesh.md#usewasm)
- [userFaceMeasurementSettingsPartial](ATSdkConfigItemFaceMesh.md#userfacemeasurementsettingspartial)

### Methods

- [getCameraStreamFunction](ATSdkConfigItemFaceMesh.md#getcamerastreamfunction)

## Properties

### createHTMLElementsCallbackForFaceMesh

• **createHTMLElementsCallbackForFaceMesh**: ``null`` \| [`createHtmlElementsCallbackForFaceMeshType`](createHtmlElementsCallbackForFaceMeshType.md)

User-provided callback function to create additional DOM objects required for facemesh processing. If user specifies null, then a default function will be used.

___

### debugFaceMesh

• **debugFaceMesh**: `boolean`

Boolean flag to enable additional console debug messages relating to facemesh processing

___

### htmlDivId

• **htmlDivId**: `string`

HTML ID of the user-provided container div for facemesh-related DOM elements. The AT SDK will create additional DOM objects inside of this div.

___

### processHeight

• **processHeight**: `number`

Pixel height of internal image used to analyze the face data. Recommended value 300.

___

### processWidth

• **processWidth**: `number`

Pixel width of internal image used to analyze the face data. Recommended value 300.

___

### processedFaceDataCallback

• **processedFaceDataCallback**: `ProcessedFaceDataCallbackType`

User-provided callback function to handle the face data from the facemesh processing. Usually, this callback will send the data over the network for reception by another client.

___

### updateInterval

• **updateInterval**: ``null`` \| `number`

Millisecond interval between facemesh processing steps. Recommended value: 100.

___

### useWasm

• **useWasm**: `boolean`

(Experimental) Flag to control whether to use wasm (true) or WebGL (false) processing. Recommended value false. Currently wasm is not supported, so a value of true still uses WebGL processing. In the future, wasm support may be enabled.

___

### userFaceMeasurementSettingsPartial

• **userFaceMeasurementSettingsPartial**: `FaceMeasurementSettingsPartial`

User-provided parameters to control facemesh processing of mouth and eyes. If any parameter is specified as null, then a default value is used instead.

## Methods

### getCameraStreamFunction

▸ **getCameraStreamFunction**(): `PromiseResult`<`LocalStream`, `Error`\>

User-provided function to return the camera stream from the web camera.

#### Returns

`PromiseResult`<`LocalStream`, `Error`\>
