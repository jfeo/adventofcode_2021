package submarine

type Submarine struct {
	Depth      int
	Horizontal int
	Aim        int
}

type Command struct {
	Type string
	X    int
}

func (submarine *Submarine) SimplePilot(command Command) {
	switch command.Type {
	case "forward":
		submarine.Horizontal += command.X
	case "up":
		submarine.Depth -= command.X
	case "down":
		submarine.Depth += command.X
	default:
		panic("Invalid command")
	}
}

func (submarine *Submarine) SimplePilotMany(commands []Command) {
	for _, cmd := range commands {
		submarine.SimplePilot(cmd)
	}
}

func (submarine *Submarine) Pilot(command Command) {
	switch command.Type {
	case "forward":
		submarine.Horizontal += command.X
		submarine.Depth += submarine.Aim * command.X
	case "up":
		submarine.Aim -= command.X
	case "down":
		submarine.Aim += command.X
	default:
		panic("Invalid command")
	}
}

func (submarine *Submarine) PilotMany(commands []Command) {
	for _, cmd := range commands {
		submarine.Pilot(cmd)
	}
}
