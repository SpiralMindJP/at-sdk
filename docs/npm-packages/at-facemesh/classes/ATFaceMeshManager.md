[@spiralmindjp/at-facemesh](../README.md) / [Exports](../modules.md) / ATFaceMeshManager

# Class: ATFaceMeshManager

Class to manage access to the facemesh processing.
This manager class should be created once at the start of the application and reused
during the lifetime of the webpage. It should not be repeatedly created/destroyed
because this costs time and memory.

## Table of contents

### Constructors

- [constructor](ATFaceMeshManager.md#constructor)

### Properties

- [cameraStream](ATFaceMeshManager.md#camerastream)
- [faceFrameToWorldFrame](ATFaceMeshManager.md#faceframetoworldframe)
- [faceMeasurementSettings](ATFaceMeshManager.md#facemeasurementsettings)
- [referenceFrameToWorldFrame](ATFaceMeshManager.md#referenceframetoworldframe)
- [renderRequestId](ATFaceMeshManager.md#renderrequestid)
- [restartInProgress](ATFaceMeshManager.md#restartinprogress)
- [startCameraParams](ATFaceMeshManager.md#startcameraparams)
- [workerProps](ATFaceMeshManager.md#workerprops)

### Methods

- [calculateEyeBlinkParameter](ATFaceMeshManager.md#calculateeyeblinkparameter)
- [calculateEyeLookDownUp](ATFaceMeshManager.md#calculateeyelookdownup)
- [calculateEyeLookLeftRight](ATFaceMeshManager.md#calculateeyelookleftright)
- [calculateHeadRotationInReferenceFrame](ATFaceMeshManager.md#calculateheadrotationinreferenceframe)
- [calculateHeadRotationInWorldFrame](ATFaceMeshManager.md#calculateheadrotationinworldframe)
- [calculateMouthOpenParameter](ATFaceMeshManager.md#calculatemouthopenparameter)
- [calculateOpenParameter](ATFaceMeshManager.md#calculateopenparameter)
- [initializeFacemesh](ATFaceMeshManager.md#initializefacemesh)
- [matrixMultiply](ATFaceMeshManager.md#matrixmultiply)
- [matrixToEulerYXZ](ATFaceMeshManager.md#matrixtoeuleryxz)
- [matrixTranspose](ATFaceMeshManager.md#matrixtranspose)
- [pauseCamera](ATFaceMeshManager.md#pausecamera)
- [processRawFaceData](ATFaceMeshManager.md#processrawfacedata)
- [restartCameraAndFaceMesh](ATFaceMeshManager.md#restartcameraandfacemesh)
- [resumeCamera](ATFaceMeshManager.md#resumecamera)
- [setReferenceFrameToCurrent](ATFaceMeshManager.md#setreferenceframetocurrent)
- [stopCameraAndFaceMesh](ATFaceMeshManager.md#stopcameraandfacemesh)
- [testHideWindow](ATFaceMeshManager.md#testhidewindow)
- [vectorAdd](ATFaceMeshManager.md#vectoradd)
- [vectorCross](ATFaceMeshManager.md#vectorcross)
- [vectorDot](ATFaceMeshManager.md#vectordot)
- [vectorLength](ATFaceMeshManager.md#vectorlength)
- [vectorMultiply](ATFaceMeshManager.md#vectormultiply)
- [vectorSubtract](ATFaceMeshManager.md#vectorsubtract)

## Constructors

### constructor

• **new ATFaceMeshManager**(`config`)

Constructor. To begin facemesh processing, call initializeFacemesh().

#### Parameters

| Name | Type |
| :------ | :------ |
| `config` | [`ATFaceMeshManagerConfig`](../interfaces/ATFaceMeshManagerConfig.md) |

## Properties

### cameraStream

• **cameraStream**: ``null`` \| `MediaStream` = `null`

(Internal member) This member is currently public but may become private in a future version.

___

### faceFrameToWorldFrame

• **faceFrameToWorldFrame**: `number`[][]

(Internal member) This member is currently public but may become private in a future version.

___

### faceMeasurementSettings

• **faceMeasurementSettings**: `FaceMeasurementSettingsComplete`

(Internal member) This member is currently public but may become private in a future version.

___

### referenceFrameToWorldFrame

• **referenceFrameToWorldFrame**: `number`[][]

(Internal member) This member is currently public but may become private in a future version.

___

### renderRequestId

• **renderRequestId**: ``null`` \| `number` = `null`

(Internal member) This member is currently public but may become private in a future version.

___

### restartInProgress

• **restartInProgress**: `boolean` = `false`

Read-only property to indicate if the webworker thread is trying to restart in restartCameraAndFaceMesh().
If this value is true for a long time, then maybe the webworker has frozen and cannot be killed.
In this case you can call restartCameraAndFaceMesh(true) to force a restart even though a restart is already
in progress. When the restart finally finishes, restartInProgress will be reset back to false.

___

### startCameraParams

• **startCameraParams**: ``null`` \| `StartCameraParams` = `null`

(Internal member) This member is currently public but may become private in a future version.

___

### workerProps

• **workerProps**: ``null`` \| `WorkerProps` = `null`

(Internal member) This member is currently public but may become private in a future version.

## Methods

### calculateEyeBlinkParameter

▸ **calculateEyeBlinkParameter**(`upperContourVertices`, `lowerContourVertices`): `number`

(Internal function) This function is currently public but may be changed to private in a future version.
Given the upper and lower eye contours, calculate the eye blink parameter between 0 and 1.

#### Parameters

| Name | Type |
| :------ | :------ |
| `upperContourVertices` | `number`[][] |
| `lowerContourVertices` | `number`[][] |

#### Returns

`number`

___

### calculateEyeLookDownUp

▸ **calculateEyeLookDownUp**(`irisVertices`, `eyeUpperVertices`, `eyeLowerVertices`): `number`

(Internal function) This function is currently public but may be changed to private in a future version.
Given the iris position and upper/lower eye contours, calculate the eye down-up look parameter between -1 and 1.

#### Parameters

| Name | Type |
| :------ | :------ |
| `irisVertices` | `number`[][] |
| `eyeUpperVertices` | `number`[][] |
| `eyeLowerVertices` | `number`[][] |

#### Returns

`number`

___

### calculateEyeLookLeftRight

▸ **calculateEyeLookLeftRight**(`irisVertices`, `eyeUpperVertices`, `eyeLowerVertices`, `firstEyeVertexIsAtOperatorsLeftSide`): `number`

(Internal function) This function is currently public but may be changed to private in a future version.
Given the iris position and upper/lower eye contours, calculate the eye left-right look parameter between -1 and 1.

#### Parameters

| Name | Type |
| :------ | :------ |
| `irisVertices` | `number`[][] |
| `eyeUpperVertices` | `number`[][] |
| `eyeLowerVertices` | `number`[][] |
| `firstEyeVertexIsAtOperatorsLeftSide` | `boolean` |

#### Returns

`number`

___

### calculateHeadRotationInReferenceFrame

▸ **calculateHeadRotationInReferenceFrame**(`silhouette`): `number`[][]

(Internal function) This function is currently public but may be changed to private in a future version.
Calculates head rotation (in the reference frame), from the facemesh silhouette vertex list.

#### Parameters

| Name | Type |
| :------ | :------ |
| `silhouette` | `number`[][] |

#### Returns

`number`[][]

___

### calculateHeadRotationInWorldFrame

▸ **calculateHeadRotationInWorldFrame**(`silhouette`): `number`[][]

(Internal function) This function is currently public but may be changed to private in a future version.
Calculates head rotation (in the world frame), from the facemesh silhouette vertex list.

#### Parameters

| Name | Type |
| :------ | :------ |
| `silhouette` | `number`[][] |

#### Returns

`number`[][]

___

### calculateMouthOpenParameter

▸ **calculateMouthOpenParameter**(`upperContourVertices`, `lowerContourVertices`): `number`

(Internal function) This function is currently public but may be changed to private in a future version.
Given the upper and lower mouth contours, calculate the mouth open parameter between 0 and 1.

#### Parameters

| Name | Type |
| :------ | :------ |
| `upperContourVertices` | `number`[][] |
| `lowerContourVertices` | `number`[][] |

#### Returns

`number`

___

### calculateOpenParameter

▸ **calculateOpenParameter**(`measuredAngle`, `minValidAngle`, `maxValidAngle`): `number`

(Internal function) This function is currently public but may be changed to private in a future version.
Given a maximum angle, a minimum angle, and a measured angle, returns a relative value between 0 and 1
expressing the measured angle relative to the minimum and maximum.

#### Parameters

| Name | Type |
| :------ | :------ |
| `measuredAngle` | `number` |
| `minValidAngle` | `number` |
| `maxValidAngle` | `number` |

#### Returns

`number`

___

### initializeFacemesh

▸ **initializeFacemesh**(`htmlVideoElement`, `srcCache`, `getCameraStreamFunction`, `processedFaceDataCallback`, `updateInterval`, `userFaceMeasurementSettingsPartial`, `useWasm`, `processWidth`, `processHeight`): `Promise`<`void`\>

Initialize the facemesh processing.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `htmlVideoElement` | `HTMLVideoElement` | HTML video element used to play the camera stream |
| `srcCache` | `HTMLCanvasElement` | HTML canvas element used to cache the camera stream |
| `getCameraStreamFunction` | () => `PromiseResult`<`LocalStream`, `Error`\> | User-provided function to return the correct camera stream to be used |
| `processedFaceDataCallback` | [`ProcessedFaceDataCallbackType`](../interfaces/ProcessedFaceDataCallbackType.md) | User-provided callback function to handle the processed face data |
| `updateInterval` | ``null`` \| `number` | Interval in millseconds between face processing steps |
| `userFaceMeasurementSettingsPartial` | ``null`` \| [`FaceMeasurementSettingsPartial`](../interfaces/FaceMeasurementSettingsPartial.md) | Set of user-specified parameters controlling eye and mouth calculation. If any setting is null, a default value is used instead. |
| `useWasm` | `boolean` | (Experimental) Flag to control whether to use wasm (true) or WebGL (false) processing. Recommended value false. Currently wasm is not supported, so a value of true still uses WebGL processing. In the future, wasm support may be enabled. |
| `processWidth` | `number` | Pixel width of internal image used to analyze the face data. Recommended value 300. |
| `processHeight` | `number` | Pixel height of internal image used to analyze the face data. Recommended value 300. |

#### Returns

`Promise`<`void`\>

___

### matrixMultiply

▸ **matrixMultiply**(`A`, `B`): `number`[][]

(Internal function) This function is currently public but may be changed to private in a future version.
Multiplies two matrices.

#### Parameters

| Name | Type |
| :------ | :------ |
| `A` | `number`[][] |
| `B` | `number`[][] |

#### Returns

`number`[][]

___

### matrixToEulerYXZ

▸ **matrixToEulerYXZ**(`mat`): `number`[]

(Internal function) This function is currently public but may be changed to private in a future version.
Converts rotation matrix to Euler angles.

#### Parameters

| Name | Type |
| :------ | :------ |
| `mat` | `number`[][] |

#### Returns

`number`[]

___

### matrixTranspose

▸ **matrixTranspose**(`mat`): `number`[][]

(Internal function) This function is currently public but may be changed to private in a future version.
Calculates matrix inverse.

#### Parameters

| Name | Type |
| :------ | :------ |
| `mat` | `number`[][] |

#### Returns

`number`[][]

___

### pauseCamera

▸ **pauseCamera**(): `void`

Pause the facemesh camera stream.

#### Returns

`void`

___

### processRawFaceData

▸ **processRawFaceData**(`predictions`): `Object`

Given the prediction data from the facemesh processing, return the simplified and processed face data parameters.

#### Parameters

| Name | Type |
| :------ | :------ |
| `predictions` | `FaceMeshRawPrediction`[] |

#### Returns

`Object`

| Name | Type |
| :------ | :------ |
| `EyeBlinkLeft` | `number` |
| `EyeBlinkRight` | `number` |
| `EyeLookHorizontalLeft` | `number` |
| `EyeLookHorizontalRight` | `number` |
| `EyeLookVerticalLeft` | `number` |
| `EyeLookVerticalRight` | `number` |
| `FaceAngle` | { `x`: `number` ; `y`: `number` ; `z`: `number`  } |
| `FaceAngle.x` | `number` |
| `FaceAngle.y` | `number` |
| `FaceAngle.z` | `number` |
| `MouthOpenA` | `number` |
| `MouthOpenE` | `number` |
| `MouthOpenI` | `number` |
| `MouthOpenO` | `number` |
| `MouthOpenU` | `number` |

___

### restartCameraAndFaceMesh

▸ **restartCameraAndFaceMesh**(`force?`): `Promise`<`void`\>

Restart the camera and facemesh processing. This should be called occasionally to reset the facemesh memory
usage back to zero.

**`remarks`**

1. This method restartCameraAndFaceMesh() may be slow and take 1-2 minutes to execute.

2. await this method to finish, before calling it many times. If you don't await,
then you may create multiple "dead" webworker threads (view in Chrome task manager).

3. In case this method freezes and never returns due to blocking at await,
then you can pass the parameter force=true to force a restart. but this will probably
leave a dead webworker thread remaining.

4. The camera is in the paused state after this function completes.
User code should then call resumeCamera when needed to start the face tracking.

#### Parameters

| Name | Type | Default value |
| :------ | :------ | :------ |
| `force` | `boolean` | `false` |

#### Returns

`Promise`<`void`\>

___

### resumeCamera

▸ **resumeCamera**(): `void`

Resume (after pause) the facemesh camera stream.

#### Returns

`void`

___

### setReferenceFrameToCurrent

▸ **setReferenceFrameToCurrent**(): `void`

Sets the current head rotation (computed from the web camera) to be the "reference frame",
meaning that the current head rotation will be considered as zero rotation on all axes.
All following rotations are measured relative to the reference frame.

#### Returns

`void`

___

### stopCameraAndFaceMesh

▸ **stopCameraAndFaceMesh**(): `void`

(Internal function) This function is currently public but may be changed to private in a future version.
Completely stop the camera and facemesh processing. This is called as part of restartCameraAndFaceMesh().

#### Returns

`void`

___

### testHideWindow

▸ **testHideWindow**(`onOff`): `void`

(Internal function) This is a test function to simulate the browser window being hidden or shown.
User code does not need to call this function.

#### Parameters

| Name | Type |
| :------ | :------ |
| `onOff` | `boolean` |

#### Returns

`void`

___

### vectorAdd

▸ **vectorAdd**(`a`, `b`): `number`[]

(Internal function) This function is currently public but may be changed to private in a future version.
Adds two vectors.

#### Parameters

| Name | Type |
| :------ | :------ |
| `a` | `number`[] |
| `b` | `number`[] |

#### Returns

`number`[]

___

### vectorCross

▸ **vectorCross**(`a`, `b`): `number`[]

(Internal function) This function is currently public but may be changed to private in a future version.
Calculates vector cross product.

#### Parameters

| Name | Type |
| :------ | :------ |
| `a` | `number`[] |
| `b` | `number`[] |

#### Returns

`number`[]

___

### vectorDot

▸ **vectorDot**(`a`, `b`): `number`

(Internal function) This function is currently public but may be changed to private in a future version.
Calculates vector dot product.

#### Parameters

| Name | Type |
| :------ | :------ |
| `a` | `number`[] |
| `b` | `number`[] |

#### Returns

`number`

___

### vectorLength

▸ **vectorLength**(`a`): `number`

(Internal function) This function is currently public but may be changed to private in a future version.
Calculates the length of a vector.

#### Parameters

| Name | Type |
| :------ | :------ |
| `a` | `number`[] |

#### Returns

`number`

___

### vectorMultiply

▸ **vectorMultiply**(`a`, `b`): `number`[]

(Internal function) This function is currently public but may be changed to private in a future version.
Performs vector*scalar multiplication.

#### Parameters

| Name | Type |
| :------ | :------ |
| `a` | `number`[] |
| `b` | `number` |

#### Returns

`number`[]

___

### vectorSubtract

▸ **vectorSubtract**(`a`, `b`): `number`[]

(Internal function) This function is currently public but may be changed to private in a future version.
Subtracts two vectors.

#### Parameters

| Name | Type |
| :------ | :------ |
| `a` | `number`[] |
| `b` | `number`[] |

#### Returns

`number`[]
