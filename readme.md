## Presentation
Presentation is a small and simple go package that gives to oppertunity
to make a small and simple "powerpoint" kinda presentation that will
be presented in the CLI.

## Examples

Example Usage:
```
package main

import "github.com/maxtors/presentation"

func main() {
	pres := presentation.NewPresentation()

	first := presentation.NewSlide("My Fancy Presentation", presentation.SlideTitle)
	pres.AddSlide(first)

    second := presentation.NewSlide("Some Content", presentation.SlideNormal)
    second.AddBullet(presentation.NewBullet(presentation.BulletStandard, "First"))
    second.AddBullet(presentation.NewBullet(presentation.BulletStandard, "Second"))
    second.AddBullet(presentation.NewBullet(presentation.BulletStandard, "Third"))
    pres.AddSlide(second)

    third := presentation.NewSlide("Thank You", presentation.SlideTitle)
	pres.AddSlide(third)

    pres.Present()
}
```

The Second slide will look as following:
```
------------------------------------------------------------
	Some Content
------------------------------------------------------------

	(1/3) First

	(2/3) Second

	(3/3) Third

------------------------------------------------------------
```
