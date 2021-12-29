package hashid

import (
	"errors"
	"github.com/urfave/cli/v2"
	"strconv"
)

var args = &struct {
	decode bool
}{}

var flags = []cli.Flag{
	&cli.BoolFlag{
		Name:        "decode",
		Aliases:     []string{"d"},
		Usage:       "Decode HashID",
		Destination: &args.decode,
	},
}

var HashID = &cli.Command{
	Name:  "hashid",
	Usage: "encode ID to HashID",
	UsageText: `hashid [numbers...]
hashid 1988
hashid 1988 1990
hashid -d 17eok0yo
`,
	ArgsUsage: "encode [ids...]",
	Action: func(c *cli.Context) error {
		if c.NArg() == 0 {
			return errors.New("none id present")
		}

		for _, v := range c.Args().Slice() {
			if args.decode {
				id, err := decode(v)
				if err != nil {
					return err
				}

				print(id)
			} else {
				id, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}

				hashed, err := encode(id)
				print(hashed)
			}
		}

		return nil
	},
	Flags: flags,
}
