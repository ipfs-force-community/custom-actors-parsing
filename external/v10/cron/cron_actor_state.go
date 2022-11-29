package cron

import "github.com/filecoin-project/go-state-types/builtin/v10/cron"

type State struct {
	Entries []cron.Entry
}
