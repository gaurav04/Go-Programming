package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//  return an array of all primes smaller than or equal to n. If number is prime mark it as false else true in Array. EX num=15
//[false false false false true false true false true true true false true false true true]

func sieve(num int) []bool {

	prime := make([]bool, num+1)
	for i := 0; i < num+1; i++ {
		prime[i] = false
	}

	for i := 2; i*i <= num; i++ {
		if prime[i] == false {
			for j := i * 2; j <= num; j += i {
				prime[j] = true
			}
		}
	}
	return prime
}

func left(num int, list []bool) bool {
	count := 0
	temp1 := 0
	for temp := num; temp != 0; temp = temp / 10 {
		count++
		temp1 = temp % 10
		if temp1 == 0 {
			return false
		}
	}

	for i := count; i > 0; i-- {
		mod := math.Pow(10, float64(i))
		if list[num%int(mod)] {
			return false
		}
	}

	return true
}

func right(num int, list []bool) bool {
	for num > 0 {
		if list[num] {
			return false
		}
		num = num / 10
	}
	return true
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/{number}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		prime := vars["number"]

		i, _ := strconv.Atoi(prime)

		list := sieve(i)

		var leftprime = left(i, list)

		var righprimte = right(i, list)

		if righprimte && leftprime {
			fmt.Fprintf(w, "The number is Two-sided Prime: %s\n", "True")
		} else {
			fmt.Fprintf(w, "The number is  Two-sided Prime: %s\n", "False")
		}
	})

	http.ListenAndServe(":8080", r)
}
