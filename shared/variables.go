package shared

import "sync"

var ActiveServices sync.WaitGroup
var Devices []Device
