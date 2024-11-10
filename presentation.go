package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// JSON file structure
type PresentationFile struct {
	Name         string
	Slides       []string
	Presentation *Presentation
}

func (pf PresentationFile) GetSlides() []Slide {
	var slides = make([]Slide, len(pf.Slides))
	for i := 0; i < len(pf.Slides); i++ {
		slides[i] = Slide{i, pf.Slides[i], pf.Presentation}
	}
	return slides
}

func NewPresentationFile(source []byte, associated_presentation *Presentation) PresentationFile {
	pf := PresentationFile{}
	err := json.Unmarshal(source, &pf)
	if err != nil {
		log.Fatal(err)
	}
	pf.Presentation = associated_presentation

	fmt.Println(pf)
	return pf
}

// Slide
type Slide struct {
	Index int
	File  string
	P     *Presentation
}

func (s Slide) IsFirst() bool {
	return s.Index == 0
}

func (s Slide) IsFinal() bool {
	return s.Index == len(s.P.Slides)-1
}

func (s Slide) GetContent() string {
	content, err := os.ReadFile(s.File)

	if err != nil {
		log.Fatal(err)
		return "404"
	}

	return string(content)
}

// Presentation
type Presentation struct {
	File      string
	LoadState bool
	Reference PresentationFile
	Name      string
	Slides    []Slide
}

func NewPresentation(file string) Presentation {
	p := Presentation{}

	// Load the JSON file
	p.File = file
	content, err := os.ReadFile(p.File)

	if err != nil {
		panic(err)
	}

	p.Reference = NewPresentationFile(content, &p)
	p.LoadState = true

	// Use the JSON to load the content into Slides
	p.Slides = p.Reference.GetSlides()

	return p
}

func (p Presentation) GetSlide(i int) string {
	content, err := os.ReadFile(p.Slides[i].File)

	if err != nil {
		log.Fatal(err)
		return "404"
	}

	return string(content)
}

// Show
type Show struct {
	Presentation *Presentation
	CurrentSlide *Slide
	Viewer       *string
	Controller   *string
}

func (s *Show) GoToSlide(i int) *Slide {
	if i < len(s.Presentation.Slides) && i >= 0 {
		s.CurrentSlide = &s.Presentation.Slides[i]
	}
	return s.CurrentSlide
}

func (s *Show) NavSlide(i int) *Slide {
	s.CurrentSlide = s.GoToSlide(s.CurrentSlide.Index + i)
	return s.CurrentSlide
}

func (s Show) ShowCurrentSlide() string {
	return s.CurrentSlide.GetContent()
}
