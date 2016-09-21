# Bellina Common Plugins

Plugins for common behaviors are provided for convenience.

Each package describes usages with example code.

## Common Behaviors made to Plugins

* `anim` allows for animation of a node's property like moving it horizontally or some other direction

* `click` to detect user mouse-clicks on a node

* `double-click` to detect user mouse double-clicks on a node

* `drag-other` when applied to a node, when that node is dragged, the plugin will move *another* node.  This is useful for *title bars* where dragging the title bar actually drags its parent window

* `drag` when applied to a node, makes that node draggable with a mouse

* `focus` allows a node to grab keyboard focus and detect key presses

* `hover` detects when a mouse cursor *hovers* or moves over a node.  This plugin can be used to create *hover-over* effects

* `mouse-drag` detects a mouse-drag gesture over a node.  When a mouse button-press over a node is followed by a mouse move - this is a mouse-drag.  This is different from the `drag` plugin, because this plugin simply events on the gesture, whereas `drag` actually moves the node across the screen.

* ~~`resize-other`~~ is a tech demo - and has not been used in a real application.  This plugin when applied to a node, detects when a node is being resized, but redirects that to *another* node.  Analogous to a `drag-other`.

* ~~`resize`~~ is a tech demo - and has not been used in a real application.  This plugin when applied to a node, allows the node to be resizable.

* `side-resize` when applied to a node and configured for one of the `left`, `top`, `left-top`, `right`, `bottom`, or `right-bottom` will resize its parent when the node is dragged around.  Can be used to implement resizing of a node when grabbed by its side.

* `zindex` when this plugin is applied to a node, it makes all of the node's child nodes clickable.  When a child-node is clicked, it changes its z-index to move to the topmost child.

## Layout Behaviors

* `vert` when applied to a node will arrange its kid nodes vertically with configurable spacing

* `horiz` like `vert` except horizontal

* `docker` allows docking of a node to its parent's borders (e.g. Top-Left docking)

## Suggest others

Use the `issues` capability in Github to suggest new plugins
