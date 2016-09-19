# `click` API

## Problem Being Solved

We need to detect when a user has *clicked* on a node.

*Clicking* is defined as a mouse button-down followed by a mouse button-up on the *same* node.

We also need be hooks into the `life-cycle` of a click event.  These are:

* `click` has be initiated - (i.e. mouse button-down on a node)

* `click` has success concluded (i.e. mouse button-up on the same node)

* `click` has unsuccessfully concluded (i.e. mouse button-up on a *different* node)

## Example Usage

![Sample](click/example/click.jpg?raw=true "Sample")

See [example usage](click/example)

## API

### 
