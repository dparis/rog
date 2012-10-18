package main

type App interface {
    // required 
    Update()
    Render()
    // provided
    Width() int
    Height() int
    Title() string
}

type AppController struct {
    width, height int
    window_title string
}


func (self *AppController) Width() int {
    return self.width
}

func (self *AppController) Height() int {
    return self.height
}

func (self *AppController) Title() string {
    return self.window_title
}

