package main

const (
	GAME_STATE_PLAYING = iota + 1
	GAME_STATE_WON
	GAME_STATE_LOST
	GAME_STATE_QUIT
)

func main() {
	init_term()

	size_x := 10
	size_y := 10
	grid := create_grid(size_x, size_y, .08)

	state := GAME_STATE_PLAYING
	pos_x := 0
	pos_y := 0

	print_grid(grid)

	for state == GAME_STATE_PLAYING {
		print_brackets(pos_x, pos_y)

		switch await_input() {
		case "w":
			if pos_y > 0 {
				pos_y -= 1
			}
		case "a":
			if pos_x > 0 {
				pos_x -= 1
			}
		case "s":
			if pos_y < size_y-1 {
				pos_y += 1
			}
		case "d":
			if pos_x < size_x-1 {
				pos_x += 1
			}
		case " ":
			if reveal(grid, pos_x, pos_y) {
				state = GAME_STATE_LOST
			} else if check_win(grid) {
				state = GAME_STATE_WON
			}

			print_grid(grid)

		case "f":

			cell := &grid[pos_y][pos_x]

			if cell.state != CELL_STATE_VISIBLE {

				if cell.state == CELL_STATE_FLAGGED {
					cell.state = CELL_STATE_INVISIBLE
				} else {
					cell.state = CELL_STATE_FLAGGED

					if check_win(grid) {
						state = GAME_STATE_WON
					}

				}

				print_grid(grid)

			}

		default:
			state = GAME_STATE_QUIT
		}

		clear_brackets()
	}

	if state != GAME_STATE_QUIT {
		if state == GAME_STATE_LOST {
			println("You lost!")
		} else {
			println("You won!")
		}

		await_input()
	}

	cleanup_term()
}
