package conn

import "fmt"

func cmdExamine(args commandArgs, c *Conn) {
	m, err := c.User.MailboxByName(args.Arg(0))
	if err != nil {
		fmt.Fprintf(c, "%s NO %s\r\n", args.ID(), err)
		return
	}

	writeMailboxInfo(c, m)
	c.writeResponse(args.ID(), "OK [READ-ONLY] EXAMINE completed")
}
