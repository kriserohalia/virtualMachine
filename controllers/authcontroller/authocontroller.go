package authcontroller

import (
	"virtualmachine/config"
	"time"
	"gorm.io/gorm"
	"virtualmachine/helper"
	"virtualmachine/models"
	"encoding/json"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4"
)

func Login(w http.ResponseWriter, r * http.Request){

	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response:= map[string]string{"message" :err.Error()}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	var user models.User
	if err := models.DB.Where("username = ?", userInput.Username).First(&user).Error;err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response:=map[string]string{"message":"username or passwordis wrong "}
			helper.ResponseJson(w, http.StatusUnauthorized, response)
			return

		default:
			response := map[string]string{"message":err.Error()}
			helper.ResponseJson(w, http.StatusInternalServerError, response)
			return
		}
	}

	
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response:= map[string]string{"message":"username or password is wrong"}
		helper.ResponseJson(w, http.StatusUnauthorized, response)
		return
	}


	expTime:=time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username:user.Username,
		RegisteredClaims:jwt.RegisteredClaims{
			Issuer : "go-jwt-mux",
			ExpiresAt :jwt.NewNumericDate(expTime),
		},
	}

	
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message":err.Error()}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	
	http.SetCookie(w, &http.Cookie{
		Name:"token",
		Path:"/",
		Value:token,
		HttpOnly:true,
	})

	response:= map[string]string{"message":"Login Succes"}
	helper.ResponseJson(w, http.StatusOK, response)
}

func Register(w http.ResponseWriter, r *http.Request){
	
	var userInput models.User
	decoder:= json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err !=nil {
		response := map[string]string{"message":err.Error()}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	hashPassword, _:= bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)


	if err := models.DB.Create(&userInput).Error; err != nil {
		response:= map[string]string{"message":err.Error()}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message":"success"}
	helper.ResponseJson(w, http.StatusOK, response)

}

func Logout(w http.ResponseWriter, r *http.Request){

	http.SetCookie(w, &http.Cookie{
		Name:"Token",
		Path:"/",
		Value:"",
		HttpOnly:true,
		MaxAge:-1,
	})

	response := map[string]string{"message":"You already logout"}
	helper.ResponseJson(w, http.StatusOK, response)
}