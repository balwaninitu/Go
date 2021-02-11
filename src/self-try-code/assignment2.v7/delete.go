package main

type appointmentList []appointment

func (a *appointmentList) popLast() (appointment, bool) {
	if len(*a) == 0 {
		return appointment{}, false

	}
	i := len(*a) - 1
	x := (*a)[i]
	*a = (*a)[:i]
	return x, true

}
