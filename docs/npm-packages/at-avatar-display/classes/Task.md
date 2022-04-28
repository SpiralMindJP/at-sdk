[@spiralmindjp/at-avatar-display](../README.md) / [Exports](../modules.md) / Task

# Class: Task

Task class. All tasks are handled synchonously (with await) and in order of submission.

## Table of contents

### Constructors

- [constructor](Task.md#constructor)

### Properties

- [additionalChainedTaskCount](Task.md#additionalchainedtaskcount)
- [id](Task.md#id)
- [name](Task.md#name)
- [startTime](Task.md#starttime)
- [submissionTime](Task.md#submissiontime)
- [taskCancel](Task.md#taskcancel)
- [taskRun](Task.md#taskrun)
- [nextId](Task.md#nextid)

## Constructors

### constructor

• **new Task**(`name`, `taskRun`, `taskCancel`, `additionalChainedTaskCount`)

#### Parameters

| Name | Type |
| :------ | :------ |
| `name` | `string` |
| `taskRun` | [`TaskRunCallbackType`](../interfaces/TaskRunCallbackType.md) |
| `taskCancel` | [`TaskCancelCallbackType`](../interfaces/TaskCancelCallbackType.md) |
| `additionalChainedTaskCount` | `number` |

## Properties

### additionalChainedTaskCount

• **additionalChainedTaskCount**: `number`

A count of additional tasks (chained tasks) that this task will later create as part of its processing.
This allows us to keep track of how many remaining tasks exist (including not-yet-created chained tasks).

___

### id

• **id**: `number`

The task ID

___

### name

• **name**: `string`

The task name

___

### startTime

• **startTime**: ``null`` \| `Date`

The start time of the task

___

### submissionTime

• **submissionTime**: `Date`

The submission time of the task

___

### taskCancel

• **taskCancel**: [`TaskCancelCallbackType`](../interfaces/TaskCancelCallbackType.md)

The function that cancels the task

___

### taskRun

• **taskRun**: [`TaskRunCallbackType`](../interfaces/TaskRunCallbackType.md)

The async function that runs the task

___

### nextId

▪ `Static` **nextId**: `number` = `1`

The next task ID to be assigned
