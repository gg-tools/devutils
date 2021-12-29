package timestamp

import (
	"github.com/urfave/cli/v2"
	"time"
)

var Time = &cli.Command{
	Name:    "time",
	Usage:   "covert timestamp to text",
	UsageText: `time [timestamps...]

time 587433600
time 587433600 1597276800
`,
	ArgsUsage: "devutils time [timestamp...]",
	Action: func(c *cli.Context) error {
		if c.NArg() == 0 {
			printTime(time.Now())
			return nil
		}

		for _, arg := range c.Args().Slice() {
			t, err := parseTimestamp(arg)
			if err != nil {
				return err
			}

			printTime(t)
		}

		return nil
	},
}
