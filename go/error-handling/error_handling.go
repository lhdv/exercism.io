package erratum

// Use will call functions and deal with errors
func Use(o ResourceOpener, input string) (err error) {

	var res Resource

	res, err = o()

	if err != nil {

		// If it's a TransientError keep trying to open resource
		if _, ok := err.(TransientError); ok {
			return Use(o, input)
		}

		return err
	}

	defer func() {

		if r := recover(); r != nil {
			if e, ok := r.(FrobError); ok {
				res.Defrob(e.defrobTag)
				err = e
			}

			err = r.(error)
		}
		res.Close()
	}()

	res.Frob(input)

	return err
}
