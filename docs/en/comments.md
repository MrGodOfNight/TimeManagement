# Commenting on the project

## Project commenting rules

1. Comments must be clear and detailed.
2. Comments must be written everywhere. Clearly comment on all the code you write. Even if everything is clear to you, it does not mean that it is understandable to others!
3. Comments must be written strictly in English.
4. use all the necessary tags and varieties of comments.
5. At the beginning of each file there should be a big comment describing what the file does and why it is needed.

## How a project is commented

You can install and customize your own Todo Tree and Better Comments
My settings:
```json
"better-comments.multilineComments": true,
  "better-comments.tags": [
  
    {
      "tag": "!",
      "color": "#FF2D00",
      "strikethrough": false,
      "underline": false,
      "backgroundColor": "transparent",
      "bold": false,
      "italic": false
    },
    {
      "tag": "?",
      "color": "#3498DB",
      "strikethrough": false,
      "underline": false,
      "backgroundColor": "transparent",
      "bold": false,
      "italic": false
    },
    {
      "tag": "todo",
      "color": "#FF8C00",
      "strikethrough": false,
      "underline": false,
      "backgroundColor": "transparent",
      "bold": false,
      "italic": false
    },
    {
      "tag": "error",
      "color": "#B30600",
      "strikethrough": false,
      "underline": false,
      "backgroundColor": "transparent",
      "bold": false,
      "italic": false
    },
    {
      "tag": "fixme",
      "color": "#04D175",
      "strikethrough": false,
      "underline": false,
      "backgroundColor": "transparent",
      "bold": false,
      "italic": false
    },
    {
      "tag": "warning",
      "color": "#F0E807",
      "strikethrough": false,
      "underline": false,
      "backgroundColor": "transparent",
      "bold": false,
      "italic": false
    },
    {
      "tag": "note",
      "color": "#49018C",
      "strikethrough": false,
      "underline": false,
      "backgroundColor": "transparent",
      "bold": false,
      "italic": false
    }
  ],
  "todo-tree.highlights.enabled": true,
  "todo-tree.regex.enableMultiLine": true,
  "todo-tree.highlights.customHighlight": {
    "TODO": {
      "type": "text",
      "icon": "pin",
      "foreground": "#FF8C00",
    },
    "ERROR": {
      "type": "text",
      "icon": "stop",
      "foreground": "#B30600",
    },
    "FIXME": {
      "type": "text",
      "icon": "shield",
      "foreground": "#04D175",
    },
    "WARNING": {
      "type": "text",
      "icon": "alert",
      "foreground": "#F0E807",
    },
    "NOTE": {
      "type": "text",
      "icon": "checklist",
      "foreground": "#49018C",
    }
  },
  "todo-tree.general.tags": [ 
    "TODO", 
    "FIXME", 
    "ERROR",
    "WARNING",
    "NOTE"
  ]
```

### 1. Use `!` in comments whenever you need a warning in your code
### 2. Use `?` in the comments whenever you need to clarify something
### 3. Use `TODO` in the comments whenever you haven't finished something and you want someone to follow up on it
### 4. Use `FIXME` in comments whenever there's a piece of code you're commenting on that is not critically flawed
### 5. Use `ERROR` in the commentary whenever there is a critical error at any point in the code that needs to be corrected right away
### 6. Use `WARNING` in comments whenever there is something in the comment that may lead to an error
### 7. Use `NOTE` in the comments whenever you just want to leave a note

### A few extra, but equally important, rules

1. Don't use multiple single-line comments where you can use a single block comment.
2. If you use a block comment, do not add `*` at the beginning of each line.
Examples:

Good:
```js
/*
! we have problems
*/
```

Bad:
```js
/*
* ! we have problems
*/
```
3. But if you describe any function/method/class etc., it is necessary to write a block comment with asterisks BEFORE the FUNCTION itself!!! All "special" comments (TODO, FIXME, ERROR, WARNING, NOTE, ! and ?) should be written in another comment.
Also describe the function well, and if it takes any parameters, write those parameters in
```js
/** @params [paramname] [paramdescr] */
```
Examples:

Good:
```js
/**
* this function outputs string to the console
* @params str string that is displayed in the console
*/
function log(str){}
//NOTE: HELLO!!!
```
Bad:
```js
function log(str){}
/**
this function outputs string to the console
NOTE: HELLO!!!
*/
```