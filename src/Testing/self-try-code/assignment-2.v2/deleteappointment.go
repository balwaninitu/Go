package main

import "errors"

func (d *doctorList) dequeue() (string, error) {
	var patientName string
	if d.head == nil {
		return "", errors.New("empty list")
	}
	patientName = d.head.patientName
	if d.size == 1 {
		d.head = nil
		d.back = nil
	} else {
		d.head = d.head.next
	}
	d.size--
	return patientName, nil
}
