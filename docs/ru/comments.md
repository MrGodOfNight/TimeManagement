# Комментирование проекта

## Правила комментирования проекта

1. Комментарии должны быть понятными и развёрнутыми.
2. Комментарии должны быть написаны везде. Чётко комментируйте весь код который вы пишете. Даже если для вас всё понятно-это не значит что это понятно остальным!
3. Комментарии пишутся строго на английском языке.
4. Используйте все необходимые теги и разновидности комментариев.
5. В начале каждого файла должен быть большой комментарий, который описывает что делает этот файл и зачем он нужен.

## Как комментируется проект

Вы можете установить и настроить себе Todo Tree и Better Comments
Мои настройки:
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

### 1. `!` в комментариях используйте всегда, когда вам необходимо о чём-то предупредить в коде
### 2. `?` в комментариях используйте всегда, когда вам необходимо что-то уточнить
### 3. `TODO` в комментариях используйте всегда, когда вы что-то не дописали и хотите, чтоб за вами продолжили
### 4. `FIXME` в комментариях используйте всегда, когда кусок кода, который вы комментируете имеет не критичные недочёты
### 5. `ERROR` в комментариях используйте всегда, когда в данном месте в коде присутствует критическая ошибка, которую необходимо срочно исправить
### 6. `WARNING` в комментариях используйте всегда, когда в данном месте в коде есть недочёты способные привести к ошибкам
### 7. `NOTE` в комментариях используйте всегда, когда просто хотите оставить записку

## Несколько дополнительных, но не менее важных правил

1. Не используйте несколько однострочных комментариев там, где можно использовать один блочный.
2. Если вы используете блочный комментарий, то в начале каждой строчки не добавляйте `*`.
Примеры:

Хорошо:
```js
/*
! we have problems
*/
```
Плохо:
```js
/*
* ! we have problems
*/
```
3. Но вот если вы описываете какую-нибудь функцию/метод/класс и т.д, то необходимо писать блочный комментарий со звёздочками ПЕРЕД САМОЙ ФУНКЦИЕЙ!!! При этом все "специальные" комментарии(TODO, FIXME, ERROR, WARNING, NOTE, ! и ?) должны писаться уже в другом комментарии.
Также хорошенько опишите эту функцию и, если она принимает какие-нибудь параметры, напишите эти параметры через 
```js
/** @params [paramname] [paramdescr] */
```
Примеры:

Хорошо:
```js
/**
* this function outputs string to the console
* @params str string that is displayed in the console
*/
function log(str){}
//NOTE: HELLO!!!
```
Плохо:
```js
function log(str){}
/**
this function outputs string to the console
NOTE: HELLO!!!
*/
```