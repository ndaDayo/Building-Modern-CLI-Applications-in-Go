package command

import "github.com/ndaDayo/Audio-Metadata-CLI/internal/interfaces"

type Parser struct {
	commands []interfaces.Command
}

func NewParser(commands []interfaces.Command) *Parser {
	return &Parser{commands: commands}
}

func (p *Parser) Parse(args []string) error {
	if len(args) < 1 {
		help()
		return nil
	}
}
