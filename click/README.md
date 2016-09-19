# `click` API

## Problem Being Solved

We need to detect when a user has *clicked* on a node.

*Clicking* is defined as a mouse button-down followed by a mouse button-up on the *same* node.

We also need be hooks into the `life-cycle` of a click event.  These are:

* `click` has be initiated - (i.e. mouse button-down on a node)

* `click` has success concluded (i.e. mouse button-up on the same node)

* `click` has unsuccessfully concluded (i.e. mouse button-up on a *different* node)

## Example Usage

![Sample](example/click.jpg?raw=true "Sample")

See [example usage](click/example)

## API

### click.On(...)

`click.On(cb func(event interface{}))`

* Must be called within a node context (i.e. between `bl.Div()` / `bl.End()` pair)

* `cb` callback function is called when the click has successfully completed

* the type of `event` is `click.Event`.  Use `event.(click.Event)` to cast.

### click.On_WithLifeCycle(...)

`click.On_WithLifeCycle(cb, onDown, onUpAndMiss func(interface{}))`

* Must be called within a node context

* `cb` callback function is called when the click has successfully completed

* `onDown` callback function is called when the click has initiated with the mouse button-down on the node

* `onUpAndMiss` callback function is called when the click has terminated unsuccessfully since the mouse button-up happened on a different node

* the type of `event` is `click.Event`.  Use `event.(click.Event)` to cast.

> Note that on `onUpAndMiss`, `event` will be nil.

We are done!