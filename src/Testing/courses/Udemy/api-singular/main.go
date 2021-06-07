package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWT struct {
	Token string `json:"token"`
}

type Error struct {
	Message string `json:"message"`
}

var db *sql.DB

func main() {

	pgUrl, err := pq.ParseURL("postgres://itluswin:PJAg3TrcqF-GYFXVZn6DZebPGxRiZQVo@ziggy.db.elephantsql.com:5432/itluswin")

	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(pgUrl)

	db, err = sql.Open("postgres", pgUrl)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/signup", signup).Methods("POST")
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/protected", TokenVerifyMiddleware(protectedEndpoint)).Methods("GET")

	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

}

// error message is repeating so make function
func respondWithErr(w http.ResponseWriter, status int, error Error) {
	//var error Error
	//error.Message = message
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

func responseJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

func signup(w http.ResponseWriter, r *http.Request) {
	var user User
	var error Error

	json.NewDecoder(r.Body).Decode(&user)
	//fmt.Println(user)
	// fmt.Println("--------------")
	// spew.Dump(user)

	//validation
	if user.Email == "" {
		//respond with error
		//error.Message = "Email is missing."
		//send status = bad request to client
		//writerhead accepts only status code
		//w.WriteHeader(http.StatusBadRequest)-for repeated line can create func
		//json.NewEncoder(w).Encode(err)
		error.Message = "Email is missing."
		respondWithErr(w, http.StatusBadRequest, error)
		return

	}

	if user.Password == "" {
		error.Message = "Password is missing."
		respondWithErr(w, http.StatusBadRequest, error)
		return
	}

	/*password getting from user is in plain text it should be be hashed
	so that it shold not be readable in case database is compromised
	for that we use package bcrypt*/

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal(err)
	}
	//hash is slice of byte but table expecting string so need below changes
	user.Password = string(hash)
	// fmt.Println("pass text", user.Password)
	// fmt.Println("hash", hash)
	//fmt.Println("user.Password after hashing", user.Password)

	/*envoke query row on db object which we created and
	in query row we will pass statement which help us insert user record into table in db*/

	stmt := "insert into users (email, password) values($1, $2) RETURNING id;"
	/*scan because we are expecting to return id
	also in query row func says it has to return atleast one row that has scan method
	scan return error*/
	err = db.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)
	//fmt.Println(err)
	if err != nil {
		error.Message = "Server error"
		respondWithErr(w, http.StatusInternalServerError, error)
		return
	}

	//if no error assign to empty pswrd string because we should not reveal password
	user.Password = ""

	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(user)//for repeated line can create func
	responseJSON(w, user)
}

//in below func return string will be token
func GenerateToken(user User) (string, error) {
	var err error
	//secret variable will hold string which we need to sign token
	/*jwt has 3 parts : header, payload and string(here secret)
	secret is string generated by signing the header and payload together
	hence we can not assign any string to secret and we have just created variable handling
	secret string for time being*/
	secret := "secret"

	/*before generating token we need to invoke jwt as below and pass couple of
	attributes such as signingmethod which is algorithm and mapclaims which is
	struct value contains claim we want to use ietf.org site has mentioned which claim syou can pass
	iss is issuer claim taken from it

	newWithClaims returns address to token and it has header and claims and method*/

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})

	//convert secret string to lsice of byte
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}

	//spew.Dump(token)

	return tokenString, nil

}

/*will put generateToken func here as soon as handler login func called we will retieve
user credentials and pass those credentials to generate token func */
func login(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("successfully called login"))
	// fmt.Println(" login Invoked")
	// var user User
	// json.NewDecoder(r.Body).Decode(&user)

	// //return two values token and error
	// token, err := GenerateToken(user)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(token)

	/*below part will verified user credentials if exist then token function will be called
	and resulting token will be sent to user as part of login response
	below variable created w.r.t above models struct */

	var user User
	var jwt JWT
	var error Error

	//decode request body which has user info enter by clien and map it back to our user variable
	json.NewDecoder(r.Body).Decode(&user)
	//spew.Dump(user)

	//server side validation for email and password
	if user.Email == "" {
		error.Message = "Email is missing"
		respondWithErr(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is missing"
		respondWithErr(w, http.StatusBadRequest, error)
		return
	}

	//create password variable to verify if it matches with the system and verified user
	password := user.Password

	//envoke query row to pass email as it is unique identifier to verify user
	//row variable because query row func return atleast one row
	row := db.QueryRow("select * from users where email=$1", user.Email)

	//envoke scan on row which return error
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows { //if no rows matches the query it return ErrNoRows
			error.Message = "The user does not exist"
			respondWithErr(w, http.StatusBadRequest, error)
			return
		}
		log.Fatal(err)
	}

	//compare hashed password provided by user with password in db table

	hashedPassword := user.Password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		error.Message = "Invalid Password"
		respondWithErr(w, http.StatusUnauthorized, error)
		return
	}
	//generate token
	token, err := GenerateToken(user)
	if err != nil {
		log.Fatal(err)
	}
	//response to client with ok status and provide token

	w.WriteHeader(http.StatusOK)
	jwt.Token = token

	//envoke rresponsejson func and pass writer and jwt
	responseJSON(w, jwt)

}

/*below func will validate token we sent from client to server and
give us access to protected endpoint */
func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" ProtectedEndpoint Invoked")
}

/*below func takes argument which is next,
next will be called after token validated and pass protected handle func to
tokenverify func which will return handler func */
func TokenVerifyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	//fmt.Println(" TokenVerifyMiddleware Invoked ")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var errorObject Error
		/*below variable hold the value of auhorized header we send from client to server
		req object has field called header and value of header is map of key value pair and authorization
		value of jwt value which we sent to client are part of this map */
		authHeader := r.Header.Get("Authorization")
		//fmt.Println(authHeader)

		/*authHeader has value string but it has bearer and value of the token and
		we need to extract out the value of token from the string we will use string split
		to split it*/

		bearerToken := strings.Split(authHeader, " ")
		//fmt.Println(bearerToken)

		//after split there will be two values in slice and 2nd value is i.e. index 1 is token we need
		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			//after extracting make sure token is valid
			token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				//validate algorithm
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				//lookout details of token
				//spew.Dump(token)
				return []byte("secret"), nil
			})

			if error != nil {
				errorObject.Message = error.Error()
				respondWithErr(w, http.StatusUnauthorized, errorObject)
				return
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				errorObject.Message = error.Error()
				respondWithErr(w, http.StatusUnauthorized, errorObject)
				return
			}

		} else {
			errorObject.Message = "Invalid token."
			respondWithErr(w, http.StatusUnauthorized, errorObject)
			return
		}
	})
}