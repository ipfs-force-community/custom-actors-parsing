package cron

import "github.com/filecoin-project/go-state-types/builtin/v10/cron"

type ConstructorParams struct {
	Entries []cron.Entry
}
