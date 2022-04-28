[@spiralmindjp/at-avatar-display](../README.md) / [Exports](../modules.md) / ATAvatarDisplay

# Class: ATAvatarDisplay

Manager class for the avatar display.
This manager class should be created once at the start of the application and reused
during the lifetime of the webpage. It should not be repeatedly created/destroyed
because this costs time and memory.

## Table of contents

### Constructors

- [constructor](ATAvatarDisplay.md#constructor)

### Properties

- [orderedAsyncTaskQueue](ATAvatarDisplay.md#orderedasynctaskqueue)
- [orderedAsyncTaskQueueCurrentlyRunningTask](ATAvatarDisplay.md#orderedasynctaskqueuecurrentlyrunningtask)
- [orderedAsyncTaskQueueIsBusy](ATAvatarDisplay.md#orderedasynctaskqueueisbusy)
- [taskQueueTimer](ATAvatarDisplay.md#taskqueuetimer)

### Methods

- [avatarDisplayIsReadyCallback](ATAvatarDisplay.md#avatardisplayisreadycallback)
- [destroy](ATAvatarDisplay.md#destroy)
- [dumpCSharpFunctionRegistry](ATAvatarDisplay.md#dumpcsharpfunctionregistry)
- [getCameraPosition](ATAvatarDisplay.md#getcameraposition)
- [getCameraRotation](ATAvatarDisplay.md#getcamerarotation)
- [getDb](ATAvatarDisplay.md#getdb)
- [getDynamicMemorySize](ATAvatarDisplay.md#getdynamicmemorysize)
- [getStaticMemorySize](ATAvatarDisplay.md#getstaticmemorysize)
- [getTimestampString](ATAvatarDisplay.md#gettimestampstring)
- [getTotalMemorySize](ATAvatarDisplay.md#gettotalmemorysize)
- [getTotalStackSize](ATAvatarDisplay.md#gettotalstacksize)
- [isReady](ATAvatarDisplay.md#isready)
- [loadUnity](ATAvatarDisplay.md#loadunity)
- [onPacketReceived](ATAvatarDisplay.md#onpacketreceived)
- [playAnimation](ATAvatarDisplay.md#playanimation)
- [processPendingTasks](ATAvatarDisplay.md#processpendingtasks)
- [remainingTaskCount](ATAvatarDisplay.md#remainingtaskcount)
- [setAvatarCacheSize](ATAvatarDisplay.md#setavatarcachesize)
- [setCameraPosition](ATAvatarDisplay.md#setcameraposition)
- [setCameraRotation](ATAvatarDisplay.md#setcamerarotation)
- [setEnableDebugLoggingForUnity](ATAvatarDisplay.md#setenabledebugloggingforunity)
- [showFPS](ATAvatarDisplay.md#showfps)
- [splitNumber](ATAvatarDisplay.md#splitnumber)
- [startTaskQueue](ATAvatarDisplay.md#starttaskqueue)
- [stopAndClearTaskQueue](ATAvatarDisplay.md#stopandcleartaskqueue)
- [stopAnimation](ATAvatarDisplay.md#stopanimation)
- [submitGenericTask](ATAvatarDisplay.md#submitgenerictask)
- [submitTaskForHideCurrentAvatar](ATAvatarDisplay.md#submittaskforhidecurrentavatar)
- [submitTaskForLoadVRMAvatar](ATAvatarDisplay.md#submittaskforloadvrmavatar)
- [testHideWindow](ATAvatarDisplay.md#testhidewindow)
- [testPassingSplitNumberToCSharp](ATAvatarDisplay.md#testpassingsplitnumbertocsharp)
- [testSplitNumber](ATAvatarDisplay.md#testsplitnumber)
- [unityFree](ATAvatarDisplay.md#unityfree)
- [unityInstance](ATAvatarDisplay.md#unityinstance)
- [unityMalloc](ATAvatarDisplay.md#unitymalloc)

## Constructors

### constructor

• **new ATAvatarDisplay**(`config`)

#### Parameters

| Name | Type |
| :------ | :------ |
| `config` | [`ATAvatarDisplayConfig`](../interfaces/ATAvatarDisplayConfig.md) |

## Properties

### orderedAsyncTaskQueue

• **orderedAsyncTaskQueue**: [`Task`](Task.md)[] = `[]`

(Internal member) This member is currently public but may be changed to private in a future version.

___

### orderedAsyncTaskQueueCurrentlyRunningTask

• **orderedAsyncTaskQueueCurrentlyRunningTask**: ``null`` \| [`Task`](Task.md) = `null`

(Internal member) This member is currently public but may be changed to private in a future version.

___

### orderedAsyncTaskQueueIsBusy

• **orderedAsyncTaskQueueIsBusy**: `boolean` = `false`

(Internal member) This member is currently public but may be changed to private in a future version.

___

### taskQueueTimer

• **taskQueueTimer**: ``null`` \| `number` = `null`

(Internal member) This member is currently public but may be changed to private in a future version.

## Methods

### avatarDisplayIsReadyCallback

▸ **avatarDisplayIsReadyCallback**(): `void`

(Internal function) This function is currently public but may be changed to private in a future version.
Callback invoked by Unity C# code after the avatar display initialization is complete.
This callback in turn calls the user-provided callback (that is passed as a parameter to loadUnity()).

#### Returns

`void`

___

### destroy

▸ **destroy**(): `Promise`<`void`\>

(Experimental) Function to destroy the Unity avatar display application. It is not recommended to use this function
because Unity memory management inside the web browser can be unreliable.

#### Returns

`Promise`<`void`\>

___

### dumpCSharpFunctionRegistry

▸ **dumpCSharpFunctionRegistry**(): `void`

Prints to the console a list of C# functions that were dynamically registered and that are callable from JS code.

#### Returns

`void`

___

### getCameraPosition

▸ **getCameraPosition**(): `Object`

Get the current camera position.

#### Returns

`Object`

| Name | Type |
| :------ | :------ |
| `x` | `number` |
| `y` | `number` |
| `z` | `number` |

___

### getCameraRotation

▸ **getCameraRotation**(): `Object`

Get the current camera rotation in Euler angles.

#### Returns

`Object`

| Name | Type |
| :------ | :------ |
| `x` | `number` |
| `y` | `number` |
| `z` | `number` |

___

### getDb

▸ **getDb**(): `Promise`<`IDBDatabase`\>

#### Returns

`Promise`<`IDBDatabase`\>

___

### getDynamicMemorySize

▸ **getDynamicMemorySize**(): `undefined` \| `number`

Returns the dynamic memory size as seen by Unity

#### Returns

`undefined` \| `number`

___

### getStaticMemorySize

▸ **getStaticMemorySize**(): `undefined` \| `number`

Returns the static memory size as seen by Unity

#### Returns

`undefined` \| `number`

___

### getTimestampString

▸ **getTimestampString**(): `undefined` \| `string`

Get a timestamp string from the Unity avatar application.

#### Returns

`undefined` \| `string`

___

### getTotalMemorySize

▸ **getTotalMemorySize**(): `undefined` \| `number`

Returns the total memory size as seen by Unity

#### Returns

`undefined` \| `number`

___

### getTotalStackSize

▸ **getTotalStackSize**(): `undefined` \| `number`

Returns the total stack size as seen by Unity

#### Returns

`undefined` \| `number`

___

### isReady

▸ **isReady**(): `boolean`

Indicates whether the Unity avatar application is ready to accept commands

#### Returns

`boolean`

___

### loadUnity

▸ **loadUnity**(`canvas`, `unityLoaderUrls`, `userProvidedAvatarDisplayIsReadyCallback`): `void`

Loads the unity loader script from the provided URLs, which in turn will run the Unity avatar display application.
The Unity C# code will call the avatarDisplayIsReadyCallback when the avatar display initialization is finished.

#### Parameters

| Name | Type |
| :------ | :------ |
| `canvas` | `HTMLCanvasElement` |
| `unityLoaderUrls` | [`UnityLoaderUrls`](../interfaces/UnityLoaderUrls.md) |
| `userProvidedAvatarDisplayIsReadyCallback` | [`AvatarDisplayIsReadyCallbackType`](../interfaces/AvatarDisplayIsReadyCallbackType.md) |

#### Returns

`void`

___

### onPacketReceived

▸ **onPacketReceived**(`JSONPacket`): `void`

Apply an incoming face motion packet onto the currently-displayed avatar.

#### Parameters

| Name | Type |
| :------ | :------ |
| `JSONPacket` | `string` |

#### Returns

`void`

___

### playAnimation

▸ **playAnimation**(`animIndex`): `void`

Play the specified animation on the currently-displayed avatar. This executes immediately,
and is not handled as a task, because it is a temporary and uncritical change to world state.

#### Parameters

| Name | Type |
| :------ | :------ |
| `animIndex` | `number` |

#### Returns

`void`

___

### processPendingTasks

▸ **processPendingTasks**(): `Promise`<`void`\>

(Internal function) This function is currently public but may be changed to private in a future version.
Process all pending tasks synchronously (with awit) and in order of submission.
This function is periodically called by a setInterval() timer.

#### Returns

`Promise`<`void`\>

___

### remainingTaskCount

▸ **remainingTaskCount**(): `number`

(Internal function) This function is currently public but may be changed to private in a future version.
Returns the total number of remaining tasks including not-yet created chained tasks (where one tasks creates more tasks).
This can be used to report task progress to the user.

#### Returns

`number`

___

### setAvatarCacheSize

▸ **setAvatarCacheSize**(`size`): `void`

Set the size of the Unity-side avatar cache. Recommended value is 5.
This function can be called inside of the _userProvidedAvatarDisplayIsReadyCallback() function.

#### Parameters

| Name | Type |
| :------ | :------ |
| `size` | `number` |

#### Returns

`void`

___

### setCameraPosition

▸ **setCameraPosition**(`position`): `void`

Set the current camera position.

#### Parameters

| Name | Type |
| :------ | :------ |
| `position` | `Vector3` |

#### Returns

`void`

___

### setCameraRotation

▸ **setCameraRotation**(`rotation`): `void`

Set the current camera rotation in Euler angles.

#### Parameters

| Name | Type |
| :------ | :------ |
| `rotation` | `Vector3` |

#### Returns

`void`

___

### setEnableDebugLoggingForUnity

▸ **setEnableDebugLoggingForUnity**(`onOff`): `void`

Enable or disable console logging of Unity-side debug messages.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `onOff` | `number` | Controls logging of Unity-side debug messages (1=enable, 0=disable) |

#### Returns

`void`

___

### showFPS

▸ **showFPS**(`onOff`): `void`

Show the FPS display in the Unity avatar application.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `onOff` | `number` | Controls visibility of the FPS display (1=visible, 0=not visible) |

#### Returns

`void`

___

### splitNumber

▸ **splitNumber**(`n`): `Object`

Splits a JS Number into two 32-bit ints to be passed to Unity

#### Parameters

| Name | Type |
| :------ | :------ |
| `n` | `number` |

#### Returns

`Object`

| Name | Type |
| :------ | :------ |
| `high32Bits` | `number` |
| `low32Bits` | `number` |

___

### startTaskQueue

▸ **startTaskQueue**(): `void`

Start the setInterval() timer that will periodically execute all tasks in the task queue.

#### Returns

`void`

___

### stopAndClearTaskQueue

▸ **stopAndClearTaskQueue**(): `void`

(Experimental method) Stop the setInterval() timer to prevent further execution of newly-submitted tasks, and cancel all existing tasks in the queue.
Warning: this is an experimental method. Normally user code should not call this, because the avatar display manager
should only be created once and should not be destroyed/re-created. Unpredictable behavior may occur if the task
queue is stopped while an async task is still in the middle of execution.

#### Returns

`void`

___

### stopAnimation

▸ **stopAnimation**(): `void`

Stop any playing animation on the currently-displayed avatar. This executes immediately,
and is not handled as a task, because it is a temporary and uncritical change to world state.

#### Returns

`void`

___

### submitGenericTask

▸ **submitGenericTask**(`task`): `void`

Submit a generic task into the task queue.

#### Parameters

| Name | Type |
| :------ | :------ |
| `task` | [`Task`](Task.md) |

#### Returns

`void`

___

### submitTaskForHideCurrentAvatar

▸ **submitTaskForHideCurrentAvatar**(): `void`

Submit a task for hiding the current avatar. Because avatar loading is handled as a task, avatar hiding
is also handled as a task to ensure synchronous, in-order execution order.
If the browser window is hidden (underneath another window, or mimimized), then the actual avatar hide will not occur,
to reduce unneeded processing when the browser window is hidden. When the browser window again becomes
visible, the last-requested avatar state (shown, or hidden) will be communicated at that time
to the Unity avatar application.

#### Returns

`void`

___

### submitTaskForLoadVRMAvatar

▸ **submitTaskForLoadVRMAvatar**(`deviceId`, `avatarContentId`, `avatarContentKeyInIndexedDb`, `animationContentId`, `animationContentKeyInIndexedDb`): `void`

Submit a task that will load the specified avatar and animation contents into Unity and display the avatar.
If the browser window is hidden (underneath another window, or mimimized), then the loading will not occur,
to reduce unneeded processing when the browser window is hidden. When the browser window again becomes
visible, the last-requested avatar/animation contents will be loaded and displayed.

#### Parameters

| Name | Type |
| :------ | :------ |
| `deviceId` | `number` |
| `avatarContentId` | `number` |
| `avatarContentKeyInIndexedDb` | `string` |
| `animationContentId` | `number` |
| `animationContentKeyInIndexedDb` | `string` |

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

### testPassingSplitNumberToCSharp

▸ **testPassingSplitNumberToCSharp**(): `void`

(Internal function) This function is currently public but may be changed to private in a future version.
Tests passing split numbers to C#

#### Returns

`void`

___

### testSplitNumber

▸ **testSplitNumber**(): `void`

(Internal function) This function is currently public but may be changed to private in a future version.
Tests the operation of splitNumber()

#### Returns

`void`

___

### unityFree

▸ **unityFree**(`ptr`): `void`

(Internal function) This function is currently public but may be changed to private in a future version.
Calls the Unity-side free function to free memory

#### Parameters

| Name | Type |
| :------ | :------ |
| `ptr` | `number` |

#### Returns

`void`

___

### unityInstance

▸ **unityInstance**(): ``null`` \| [`UnityInstance`](../interfaces/UnityInstance.md)

Provides access to the UnityInstance object created by the Unity loader script.

#### Returns

``null`` \| [`UnityInstance`](../interfaces/UnityInstance.md)

___

### unityMalloc

▸ **unityMalloc**(`length`): `number`

(Internal function) This function is currently public but may be changed to private in a future version.
Calls the Unity-side malloc function to allocate memory

#### Parameters

| Name | Type |
| :------ | :------ |
| `length` | `number` |

#### Returns

`number`
