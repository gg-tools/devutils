package hex

import (
	"errors"
	"github.com/urfave/cli/v2"
)

var hexArgs = &struct {
	binary  bool
	octal   bool
	decimal bool
}{}

var hexFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:        "binary",
		Aliases:     []string{"b"},
		Usage:       "Convert to binary system",
		Destination: &hexArgs.binary,
	},
	&cli.BoolFlag{
		Name:        "octal",
		Aliases:     []string{"o"},
		Usage:       "Convert to octal system",
		Destination: &hexArgs.octal,
	},
	&cli.BoolFlag{
		Name:        "decimal",
		Aliases:     []string{"d"},
		Usage:       "Convert to decimal system",
		Destination: &hexArgs.decimal,
	},
}

var Hex = &cli.Command{
	Name:    "hex",
	Usage:   "convert number to hex",
	UsageText: `hex [command] -[flag] [args...]
hex 0b110
hex 06
hex 0x1e
hex -b 1988
hex -o 1988
hex -d 1988
`,
	ArgsUsage: "encode [ids...]",
	Action: func(c *cli.Context) error {
		if c.NArg() == 0 {
			return errors.New("none number is present")
		}

		for _, v := range c.Args().Slice() {
			number, err := parseInt(v)
			if err != nil {
				return err
			}

			if hexArgs.binary {
				printBinary(number)
			} else if hexArgs.octal {
				printOctal(number)
			} else if hexArgs.decimal {
				printDecimal(number)
			}
		}

		return nil
	},
	Flags: hexFlags,
}
