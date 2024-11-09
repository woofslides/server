package main

import (
	"fmt"
)

func main() {
	var presentation Presentation = NewPresentation("/home/acutewoof/Projects/woofslides/server/TEMPORARY.json")
	var show Show = Show{&presentation, &presentation.Slides[0]}
	fmt.Println(*show.Presentation)
	fmt.Println(*show.CurrentSlide)

	// go to slide 2
	show.GoToSlide(2)

	fmt.Println(*show.CurrentSlide)
	fmt.Println(show.ShowCurrentSlide())

	// go forward
	show.NavSlide(1)

	fmt.Println(*show.CurrentSlide)
	fmt.Println(show.ShowCurrentSlide())

	// go backward
	show.NavSlide(-1)

	fmt.Println(*show.CurrentSlide)
	fmt.Println(show.ShowCurrentSlide())
}
