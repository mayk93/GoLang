/*
protect is basically a try except - If g may cause an error, call it
through protect, so execution won't break if g encounters a panic.
*/

func protect(g func()) {
	defer func() {
		log.Println("done")  // Println executes normally even if there is a panic
		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
		}
	}()
	log.Println("start")
	g()
}

// My version of protect:

// Feed this function a function that may cause errors
func protect(to_run func() int) int {
	defer func() {
		stack_trace := recover()
        if stack_trace != nil {
			fmt.Println("Run time panic: ", stack_trace)
		}
	}()

	result := to_run()

    return result
}
