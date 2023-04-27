# No Asteroids

A no asteroids game with simplicity in mind. Right now it's just a prototype

## Features

+ Asteroids scene

## Technologies

It's being developed using:
+ Golang
+ Ebiten
+ Piskel

## Understanding the code

Inside each folder containing code there's a readme that explains better the module but for a overview of the
whole project here we go.

The following modules holds engine code and you won't find anything about game implementation there (mechanics,
resources, etc).

+ animations: possible animations that can be applied to entities
+ assets: things to get images and fonts on the game
+ engine: holds the most important classes being used in game implementation and more about game cycle
+ events: signals, event targets, events, etc.
+ math: transforms, vectors, etc.
+ particles: very simple particle sysmte
+ physics: world, collision tester, etc.
+ render: low level (ebiten code, nothing about opengl) code that helps to draw things on the screen
+ ui: buttons, progress bars, etc.
+ utils: array, lists, text helpers, etc.

The following modules hold game implementation and resources

+ game: holds all the code about game implementation, mechanics, etc.
+ res: game resources; images, fonts, piskel files, etc.