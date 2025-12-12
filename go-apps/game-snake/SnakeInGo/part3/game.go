package main

import (
    "math/rand"
    "strconv"
    "time"

    "github.com/gdamore/tcell"
)

const (
    RefreshRate = 15
)

type Game struct {
    Screen    tcell.Screen
    snakeBody SnakeBody
    FoodPos   Part
    Score     int
    GameOver  bool
}

// SetContent sets the contents of the given cell location.
// SetContent(x int, y int, primary rune, combining []rune, style Style)
func drawParts(s tcell.Screen,
               snakeParts []Part,
               foodPos Part,
               snakeStyle tcell.Style,
               foodStyle tcell.Style) {
    s.SetContent(foodPos.X, foodPos.Y, '\u25CF', nil, foodStyle)
    for _, part := range snakeParts {
        s.SetContent(part.X, part.Y, ' ', nil, snakeStyle)
    }
}

// Draw a text inside a rectangle on the screen.
func drawText(s tcell.Screen, x1, y1, x2, y2 int, text string) {
    row := y1
    col := x1
    style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
    for _, r := range text {
        s.SetContent(col, row, r, nil, style)
        col++
        if col >= x2 {
            row++
            col = x1
        }
        if row > y2 {
            break
        }
    }
}

// Check if there is a collision between two parts, for example,
// the snake head and the snake body.
func checkCollision(parts []Part, otherPart Part) bool {
    for _, part := range parts {
        if part.X == otherPart.X && part.Y == otherPart.Y {
            return true
        }
    }
    return false
}

// Update the position of the food.
func (g *Game) UpdateFoodPos(width int, height int) {
    g.FoodPos.X = rand.Intn(width)
    g.FoodPos.Y = rand.Intn(height)
    if g.FoodPos.Y == 1 && g.FoodPos.X < 10 {
        g.UpdateFoodPos(width, height)
    }
}

// Run this game.
func (g *Game) Run() {
    defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
    g.Screen.SetStyle(defStyle)
    width, height := g.Screen.Size()
    g.snakeBody.ResetPos(width, height)
    g.UpdateFoodPos(width, height)
    g.GameOver = false
    g.Score = 0
    snakeStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)
    for {
        longerSnake := false
        g.Screen.Clear()

        // Check if the snake hits the food.
        headAsSlice := g.snakeBody.Parts[len(g.snakeBody.Parts)-1:]
        if checkCollision(headAsSlice, g.FoodPos) {
            g.UpdateFoodPos(width, height)
            longerSnake = true
            g.Score++
        }

        // Check if the snake hits its body.
        body := g.snakeBody.Parts[:len(g.snakeBody.Parts)-1]
        head := g.snakeBody.Parts[len(g.snakeBody.Parts)-1]
        if checkCollision(body, head) {
            break
        }

        // Make the next move.
        g.snakeBody.Update(width, height, longerSnake)

        // Draw the snake.
        drawParts(g.Screen, g.snakeBody.Parts, g.FoodPos, snakeStyle, defStyle)

        // Draw the score.
        drawText(g.Screen, 1, 1, 8+len(strconv.Itoa(g.Score)), 1, "Score: "+strconv.Itoa(g.Score))

        // Wait for showing the moved snake.
        time.Sleep((1000 / RefreshRate) * time.Millisecond)

        // Update the screen display.
        g.Screen.Show()
    }

    g.GameOver = true
    drawText(g.Screen, width/2-20, height/2, width/2+20, height/2, "Game Over, Score: "+strconv.Itoa(g.Score)+", Play Again? y/n")
    g.Screen.Show()
}
