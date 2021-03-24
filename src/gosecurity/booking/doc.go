package booking

/*

package booking includes data structure and handlers which get invoke from main func
and help user to book appointment.

Data structure struct at the top of each file
indicats the information particular to that file.

Each file has different handlers to help user book appointment
When the user redirected to booking package

Handlers:

         Handler is an interface that has a method called ServeHTTP,
		 which takes a value of type ResponseWriter and another type
		 of request as parameters.

	1. Index:

	   Name of handlers :

	   #IndexA
	   #IndexD
	   #IndexS
	   #IndexP

	     Index handler with redirect replies to the request with a
		 redirect to url, which may be a path relative to the request path.
		 The provided code is in 303 range and is StatusSeeOther.


	2. Prefix+Index:

	   Name of handlers :

	   #AppointmentIndex
	   #DoctorIndex
	   #ScheduleIndex
	   #PatientIndex

	    Above handlers supports the POST method, once invoked it fetch data available
		in database. Same data from html template get parse and executed at browser.
		Data fetching from database is done by invoking query which executes a query
		that returns rows, typically a SELECT.The args are for any placeholder parameters
		in the query.Since query returns pointer to rows, slice can be created from struct
		and can loop over the elements by rows.Next method.


	3. Show:


	   Name of handlers :

	   #showAppointments
	   #showDoctor
	   #showSchedule
	   #showPatient

	    showhandlers implemented by POST method and it fetch selected data form database.
		It has QueryRow method which executes a query that is expected to return at most one row.
		QueryRow always returns a non-nil value. Errors are deferred until Row's Scan method is called.
		If the query selects no rows, the *Row's Scan will return ErrNoRows. Otherwise, the *Row's Scan
		scans the first selected row and discards the rest. Data parsed through html template and get executed.

	4. CreateAppointments :

	     Above handle is only in apppointment.go file it executes html template and get user input.


    5. CreateAppointmentProcess :

	   Handler only in appointment.go file which when get invoked by user can able to
	   book appointment.It includes database function Exec which executes a create query without returning
	   any rows. The args are for ID to create parameters in the query.


	6. UpdateDoctor:

	   Handler is available in docotr.go file, when invoked can update doctor details such as ID and name.
	   ID can be updated only when it is not scheduled for any appointment.It takes QueryRow method which
	   executes a query that is expected to return at most one row.QueryRow always returns a non-nil value.
	   Errors are deferred until Row's Scan method is called.If the query selects no rows, the *Row's Scan
	   will return ErrNoRows. Otherwise, the *Row's Scan scans the first selected row and discards the rest.
	   Data parsed through html template and get executed.Process same like show handler only difference is to
	   select particular ID which need to update.


	6. UpdateDoctorProcess:


	   Handler is available in docotr.go file, when invoked executes parsed html template and display
	   updated details through database function Exec which executes a update query without returning
	   any rows. The args are for ID to update parameters in the query.


	7. Delete:

	   Name of handlers :

	   #deleteAppointments
	   #deleteDoctor
	   #deleteSchedule
	   #deletePatient

	   Handler when get invoked will use database func Exec which executes a delete query without returning
	   any rows. The args are for ID to delete parameters in the query.



















*/
