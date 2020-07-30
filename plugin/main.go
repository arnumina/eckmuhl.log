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

package main

import (
	"github.com/arnumina/eckmuhl.core/pkg/command"

	"github.com/arnumina/eckmuhl.log/internal/plugin"
)

var (
	_version string
	_builtAt string
)

// Export AFAIRE.
func Export() command.Command {
	return plugin.New(_version, _builtAt)
}

func main() {
	_ = Export() // avoid linter errors ////////////////////////////////////////////////////////////////////////////////
}

/*
######################################################################################################## @(°_°)@ #######
*/
