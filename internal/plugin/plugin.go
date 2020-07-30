/*
#######
##                __               __   __      __
##       ___ ____/ /__ __ _  __ __/ /  / /     / /__  ___ _
##      / -_) __/  '_//  ' \/ // / _ \/ / _   / / _ \/ _ `/
##      \__/\__/_/\_\/_/_/_/\_,_/_//_/_/ (_) /_/\___/\_, /
##                                                  /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package plugin

import (
	"time"

	"github.com/arnumina/eckmuhl.core/pkg/command"
	"github.com/arnumina/failure"

	"github.com/arnumina/eckmuhl.log/internal/log"
)

type (
	plugin struct {
		version string
		builtAt time.Time
	}
)

// New AFAIRE.
func New(version, builtAt string) command.Command {
	return &plugin{
		version: version,
		builtAt: command.UnixToTime(builtAt),
	}
}

func (p *plugin) Name() string {
	return "log"
}

func (p *plugin) Description() string {
	return "print the log file in real time"
}

func (p *plugin) Version() string {
	return p.version
}

func (p *plugin) BuiltAt() time.Time {
	return p.builtAt
}

func (p *plugin) Run(args []string) error {
	cf := command.NewCmdFlag(p)
	pFileName := cf.String("file", "", "the log file to be printed")

	if err := cf.Parse(args); err != nil {
		return err
	}

	if *pFileName == "" {
		return failure.New(nil).
			Msg("the log file name has not been provided") /////////////////////////////////////////////////////////////
	}

	return log.TailFile(*pFileName)
}

/*
######################################################################################################## @(°_°)@ #######
*/
