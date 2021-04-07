//
// Show help about our configuration file.
//

package main

import (
	"flag"
	"fmt"

	"github.com/skx/rss2email/configfile"
)

// Structure for our options and state.
type configCmd struct {

	// Configuration file, used for testing
	config *configfile.ConfigFile
}

// Arguments handles argument-flags we might have.
//
// In our case we use this as a hook to setup our configuration-file,
// which allows testing.
func (c *configCmd) Arguments(flags *flag.FlagSet) {
	c.config = configfile.New()
}

// Info is part of the subcommand-API
func (c *configCmd) Info() (string, string) {

	// Get some details of the (new) configuration file.
	if c.config == nil {
		c.config = configfile.New()
	}
	path := c.config.Path()
	exists := c.config.Exists()

	name := "config"
	doc := `Provide documentation for our configuration file.

About
-----

RSS2Email is a simple command-line utility which will fetch items from
remote Atom and RSS feeds and generate emails.  In order to operate it
obviously needs a list of locations to poll.


Config Location
---------------

As of the 2.x series of rss2email releases the configuration file format
and location have changed.  The new configuration file will be read from:

     ` + path

	if !exists {
		doc += `

NOTE: The configuration file does not currently exist!
NOTE: The legacy file will be read if it is present.
NOTE:
NOTE: If nothing is present this application will do nothing useful!`
	}

	doc += `

Configuration File Format
-------------------------

The format of the configuration file is plain-text, and at its simplest
it consists of nothing more than a series of URLs, one per line, like so:

       https://blog.steve.fi/index.rss
       http://floooh.github.io/feed.xml

Entries can be commented out via the '#' character, temporarily:

       https://blog.steve.fi/index.rss
       # http://floooh.github.io/feed.xml

In the future it will be possible to do more, and with that in mind there
is scope for adding options which apply only to specific feeds.  The general
form of this support looks like this:

       https://foo.example.com/
        - key:value
       https://foo.example.com/
        - key2:value2

Here you see that lines prefixed with " -" will be used to specify a key
and value separated with a ":" character.  Configuration-options apply to
the URL above their appearance.

Any option appearing before an URL is a fatal error, and will be reported
as such.

Available Options
------------------

Key           | Purpose
--------------+--------------------------------------------------------------
exclude       | Exclude any item which matches the given regular-expression.
exclude-title | Exclude any item title matching the given regular-expression.
retry         | The maximum number of times to retry a failing HTTP-fetch.
delay         | The amount of time to sleep between retried HTTP-fetches.
`
	return name, doc
}

// Execute is invoked if the user specifies `add` as the subcommand.
func (c *configCmd) Execute(args []string) int {

	fmt.Fprintf(out, "This command only exists to show help, when executed as:")
	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, "rss2email help config")
	fmt.Fprintf(out, "\n")

	// All done, with no errors.
	return 0
}