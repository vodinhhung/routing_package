package api

import (
	"errors"
	"fmt"
	"net/http"
)

type Args struct {
	methodType string
	longitude  float64
	latitude   float64
}

func validate(
	w http.ResponseWriter, r *http.Request, args *Args,
	validationFuncs ...func(w http.ResponseWriter, r *http.Request, args *Args) (err error),
) (err error) {
	for _, validateFunc := range validationFuncs {
		err = validateFunc(w, r, args)
		if err != nil {
			return
		}
	}

	return
}

func validateHttpMethod(w http.ResponseWriter, r *http.Request, args *Args) (err error) {
	if args == nil {
		err = errors.New("no method provided for validation")
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	expectedMethod := args.methodType
	if r.Method != expectedMethod {
		err = errors.New(fmt.Sprint("Method not allowed", http.StatusMethodNotAllowed))
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return err
	}

	return nil
}

func validateLongitudeAndLatitude(w http.ResponseWriter, r *http.Request, args *Args) (err error) {
	if args == nil {
		err = errors.New("no longitude provided for validation")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	longitude := args.longitude
	latitude := args.latitude

	if longitude < -180 || longitude > 180 {
		err = errors.New("longitude must be between -180 and 180")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	if latitude < -90 || latitude > 90 {
		err = errors.New("latitude must be between -90 and 90")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	return nil
}
