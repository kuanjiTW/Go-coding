package main

type Part struct {
    X int
    Y int
}


// A snack is represented as a (FIFO) queue
// storing all its parts.
//   <-- [ tail | ... | ... | ... | head ] <--
// When the snack makes a move, a new head is pushed
// into the queue while the tail is popped out.

// For example, consider the following snake:
//    0\0 1 2 3 4 5 6 (X-axis)
//    1
//    2
//    3 - x x
//    4     x
//    5     x o
//    6
// (Y-axis)
// It is represented as the following slice:
// [(0, 3), (1, 3), (2, 3), (2, 4), (2, 5), (3, 5)].
//
// Suppose the next move of the snake is to the right.
// Then the new snake is represented as:
// [(1, 3), (2, 3), (2, 4), (2, 5), (3, 5), (4, 5)].

type SnakeBody struct {
    Parts  []Part
    Xspeed int
    Yspeed int
}

// Change the direction and speed of the snake.
// There are four directions depending on the signedness of
// vertical and horizontal.
func (sb *SnakeBody) ChangeDir(vertical int, horizontal int) {
    sb.Yspeed = vertical
    sb.Xspeed = horizontal
}

// Update the parts of the snake.
func (sb *SnakeBody) Update(width int, height int, longerSnake bool) {
    // Get the last part, which is the head.
    lastPart := sb.Parts[len(sb.Parts)-1]

    // Add the new part of the next move.
    nextPart := lastPart.GetUpdatedPart(sb, width, height)
    sb.Parts = append(sb.Parts, nextPart)

    // Remove the last part (tail) if the snake does not grow.
    if !longerSnake {
        sb.Parts = sb.Parts[1:]
    }
}

// Reset the parts of the snake.
func (sb *SnakeBody) ResetPos(width int, height int) {
    snakeParts := []Part{
        {
            X: int(width / 2),
            Y: int(height / 2),
        },
        {
            X: int(width/2) + 1,
            Y: int(height / 2),
        },
        {
            X: int(width/2) + 2,
            Y: int(height / 2),
        },
    }
    sb.Parts = snakeParts
    sb.Xspeed = 1
    sb.Yspeed = 0
}

// Compute the new Part of the next move.
// - width: the width of the board
// - height: the height of the board
func (sp *Part) GetUpdatedPart(sb *SnakeBody, width int, height int) Part {
    newPart := *sp
    newPart.X = (newPart.X + sb.Xspeed) % width
    if newPart.X < 0 {
        newPart.X += width
    }
    newPart.Y = (newPart.Y + sb.Yspeed) % height
    if newPart.Y < 0 {
        newPart.Y += height
    }
    return newPart
}
