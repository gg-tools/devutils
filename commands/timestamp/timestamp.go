package timestamp

import (
	"github.com/urfave/cli/v2"
	"time"
)

var Timestamp = &cli.Command{
	Name:    "timestamp",
	Aliases: []string{"ts"},
	Usage:   "covert text to timestamp",
	UsageText: `timestamp [date|datetime|expressions...]

timestamp 
timestamp 1988-08-13
timestamp "1988-08-13 18:06:06"
timestamp +1y -2m 3d 4h 5i 6s
`,
	ArgsUsage: `timestamp [date|datetime|expression...]`,
	Action: func(c *cli.Context) error {
		nArg := c.NArg()
		args := c.Args().Slice()
		if nArg == 0 {
			printTimestamp(time.Now())
			return nil
		} else if nArg == 1 && (len(args[0]) == dateLen || len(args[0]) == timeLen) {
			t, err := parseTime(args[0])
			if err != nil {
				return err
			}
			printTimestamp(t)
			return nil
		} else {
			t, err := str2time(args...)
			if err != nil {
				return err
			}
			printTimestamp(t)
			return nil
		}
	},
}
