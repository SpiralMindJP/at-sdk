[@spiralmindjp/at-avatar-display](../README.md) / [Exports](../modules.md) / UnityModule

# Interface: UnityModule

UnityModule object created by the Unity loader script

## Table of contents

### Properties

- [HEAPU8](UnityModule.md#heapu8)
- [asmLibraryArg](UnityModule.md#asmlibraryarg)

### Methods

- [UTF8ToString](UnityModule.md#utf8tostring)
- [\_free](UnityModule.md#_free)
- [\_malloc](UnityModule.md#_malloc)

## Properties

### HEAPU8

• **HEAPU8**: `ArrayBufferView`

___

### asmLibraryArg

• **asmLibraryArg**: `UnityAsmLibrary`

## Methods

### UTF8ToString

▸ **UTF8ToString**(`ptr`): `string`

#### Parameters

| Name | Type |
| :------ | :------ |
| `ptr` | `number` |

#### Returns

`string`

___

### \_free

▸ **_free**(`ptr`): `void`

#### Parameters

| Name | Type |
| :------ | :------ |
| `ptr` | `number` |

#### Returns

`void`

___

### \_malloc

▸ **_malloc**(`size`): `number`

#### Parameters

| Name | Type |
| :------ | :------ |
| `size` | `number` |

#### Returns

`number`
