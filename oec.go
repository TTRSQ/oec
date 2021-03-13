package oec

import "github.com/TTRSQ/imu"

type OrderExecutionChecker interface {
	ReplaceOrders(orders []string)
	ApplyExecutedID(id string) bool
}

type orderExecutionChecker struct {
	executionPool imu.MeetUpper
	executedPool  imu.MeetUpper
	myOrderPool   map[string]bool
}

func NewOrderExecutionChecker() OrderExecutionChecker {
	mp := imu.NewMeetUpper(256)
	optout := imu.NewMeetUpper(256)
	op := map[string]bool{}
	return &orderExecutionChecker{
		executionPool: mp,
		executedPool:  optout,
		myOrderPool:   op,
	}
}

func (oec *orderExecutionChecker) ReplaceOrders(orderIDs []string) {
	op := map[string]bool{}
	for i := range orderIDs {
		op[orderIDs[i]] = true
	}
	oec.myOrderPool = op
}

// ApplyExecutedID return execed or not
func (oec *orderExecutionChecker) ApplyExecutedID(id string) bool {
	meets := oec.executionPool.Apply(id)
	optouted := false
	if meets {
		optouted = oec.executedPool.Apply(id)
	}
	return meets && !optouted
}
