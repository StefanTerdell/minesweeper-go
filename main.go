package main

func main() {
	init_term()

	size_x := 10
	size_y := 10
	grid := create_grid(size_x, size_y, .1)
	// state := STATE_PLAYING
	pos_x := 0
	pos_y := 0

	print_grid(grid)

Loop:
	for {
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
			reveal(grid, pos_x, pos_y)

			print_grid(grid)
		default:
			break Loop
		}

		clear_brackets()

	}

}
