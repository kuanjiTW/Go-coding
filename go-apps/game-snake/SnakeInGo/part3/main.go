package main

import (
    "log"
    "os"

    "github.com/gdamore/tcell"
)

func processKeyEvent(game *Game, event *tcell.EventKey) {
    if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
        // Quit the game when Esc or Ctrl-C is pressed.
        game.Screen.Fini()
        os.Exit(0)
    } else if event.Key() == tcell.KeyUp && game.snakeBody.Yspeed == 0 {
        game.snakeBody.ChangeDir(-1, 0)
    } else if event.Key() == tcell.KeyDown && game.snakeBody.Yspeed == 0 {
        game.snakeBody.ChangeDir(1, 0)
    } else if event.Key() == tcell.KeyLeft && game.snakeBody.Xspeed == 0 {
        game.snakeBody.ChangeDir(0, -1)
    } else if event.Key() == tcell.KeyRight && game.snakeBody.Xspeed == 0 {
        game.snakeBody.ChangeDir(0, 1)
    } else if event.Rune() == 'y' && game.GameOver {
        go game.Run()
    } else if event.Rune() == 'n' && game.GameOver {
        game.Screen.Fini()
        os.Exit(0)
    }
}

func main() {
    // Create a new screen.
    screen, err := tcell.NewScreen()

    if err != nil {
        log.Fatalf("%+v", err)
    }

    // Initialize the screen.
    if err := screen.Init(); err != nil {
        log.Fatalf("%+v", err)
    }

    // Set the default style of the screen.
    defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
    screen.SetStyle(defStyle)

    // Create a new game.
    game := Game { Screen: screen }

    // Run the game.
    go game.Run()
    for {
        // Get the event from the screen.
        switch event := game.Screen.PollEvent().(type) {
        // Resize event
        case *tcell.EventResize: game.Screen.Sync()
        // key event
        case *tcell.EventKey: processKeyEvent(&game, event)
        }
    }
}
