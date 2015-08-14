// Package presentation implments a simple CLI based "PowerPoint"
package presentation

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

const (
	// BulletStandard is the const to represent a normal bullet
	BulletStandard = 1 << iota

	// BulletSubStandard is the const to represent an indented sub-bullet
	BulletSubStandard

	// SlideNormal is the type of slide you expect to see with bullets
	SlideNormal

	// SlideTitle is a simple slide with just a title
	SlideTitle
)

// Presentation is the total collection of slides
type Presentation struct {
	Title  string
	Slides []*Slide
}

// AddSlide adds a single slide to this presentation
func (p *Presentation) AddSlide(slide *Slide) {
	p.Slides = append(p.Slides, slide)
}

// AddSlides adds a slice of slides to this presentation
func (p *Presentation) AddSlides(slides []*Slide) {
	for _, slide := range slides {
		p.Slides = append(p.Slides, slide)
	}
}

// Present will present this presentation
func (p *Presentation) Present() {
	for _, slide := range p.Slides {
		clearTerminal()
		slide.Present()
	}
	clearTerminal()
}

// NewPresentation creates a new presentation, but only sets up the title
func NewPresentation() *Presentation {
	return &Presentation{}
}

// Slide is a single slide in a presentation
type Slide struct {
	Type    int
	Title   string
	Bullets []*Bullet
}

// AddBullet adds a single bullet to this slide
func (s *Slide) AddBullet(bullet *Bullet) {
	s.Bullets = append(s.Bullets, bullet)
}

// AddBullets adds a slice of bullets to this slide
func (s *Slide) AddBullets(bullets []*Bullet) {
	for _, bullet := range bullets {
		s.Bullets = append(s.Bullets, bullet)
	}
}

// Present presents this slide to the console
func (s *Slide) Present() {
	fmt.Println()

	if s.Type == SlideNormal {
		lineBreaker()
		fmt.Printf("\t\t\t%s\n", s.Title)
		lineBreaker()

		if len(s.Bullets) != 0 {
			fmt.Println()
			for pos, bullet := range s.Bullets {
				switch bullet.Type {
				case BulletStandard:
					fmt.Printf("\t\t\t(%d/%d) %s\n", pos+1, len(s.Bullets), bullet.Content)
				case BulletSubStandard:
					fmt.Printf("\t\t\t\t(%d/%d)- %s\n", pos+1, len(s.Bullets), bullet.Content)
				}
				bufio.NewReader(os.Stdin).ReadBytes('\n')
			}
			lineBreaker()
		}
	} else if s.Type == SlideTitle {
		lineBreaker()
		lineBreaker()
		fmt.Println()
		fmt.Println()
		fmt.Println()

		fmt.Printf("\t\t")
		for pos := 0; pos < 30-(len(s.Title)/2); pos++ {
			fmt.Printf(" ")
		}
		fmt.Printf("%s\n", s.Title)

		fmt.Println()
		fmt.Println()
		fmt.Println()
		lineBreaker()
		lineBreaker()
	}
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// NewSlide creates a new slide, but only sets up the title
func NewSlide(title string, t int) *Slide {
	return &Slide{Title: title, Type: t}
}

// Bullet is a single bullet point in a slide
type Bullet struct {
	Type    int
	Content string
}

// NewBullet creates a new bullet for a slide
func NewBullet(t int, c string) *Bullet {
	return &Bullet{Type: t, Content: c}
}

func lineBreaker() {
	fmt.Println("\t\t------------------------------------------------------------")
}

func clearTerminal() {
	clearCmd := exec.Command("clear")
	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
}
