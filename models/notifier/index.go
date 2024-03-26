package notifier

import (
	"be101_golang/models/user"
)

type Notifier interface {
	Notify(user user.User, message string)
}
