package errorcase

import "errors"

var ErrUnrecognizedTicket = errors.New("unrecognized parking ticket")
var ErrNoPosition = errors.New("no available position")
var ErrCannotParkTwice = errors.New("cannot park twice")
var ErrUnrecognizedStyle = errors.New("unrecognized parking style option")
var ErrLimitInvalid = errors.New("limit is invalid")
var ErrUnrecognizedOptionMenu = errors.New("unrecognized menu option")