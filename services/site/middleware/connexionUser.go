package utils

import (
	data "botgvg/internal"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

func CheckUser(w http.ResponseWriter, r *http.Request, database *sql.DB) bool {
	// Récupération des informations du post
	var discordUser data.DiscordUser
	err := json.NewDecoder(r.Body).Decode(&discordUser)
	CheckErr("Erreur de décodage JSON CheckUser", err)

	if discordUser.Id != "" {
		var id string
		stmt, err := database.Prepare("SELECT ID FROM Users WHERE DiscordID = ?")
		CheckErr("db Prepare user_db_uuid : ", err)
		err1 := stmt.QueryRow(discordUser.Id).Scan(&id)
		// fmt.Println("err1 : ", err1)

		if err1 == nil {
			userUUID := uuid.Must(uuid.NewV4())
			new_uuid := userUUID.String()
			dateCookie := time.Now().Add(720 * time.Hour) // 1 mois
			cookie := &http.Cookie{
				Name:    "user_token",
				Value:   new_uuid,
				Expires: dateCookie,
				Domain:  data.SITE_DOMAIN,
				Path:    "/",
				// Secure:  true,
			}

			err := cookie.Valid()
			if err != nil {
				fmt.Printf("invalid cookie: %v\n", err)
			}

			SessionLogger(w, discordUser.Id, cookie, database)
			return true
		}
	}
	return false
}

// Fonction qui crée le cookie et sa correspondance dans la db ainsi que dans la map
func SessionLogger(w http.ResponseWriter, discordId string, cookie *http.Cookie, database *sql.DB) {
	// Convertir en string avec un format spécifique
	layout := "2006-01-02 15:04:05"
	strDate_cookie := cookie.Expires.Format(layout)

	stmt, err := database.Prepare("UPDATE Users SET uuid = ?, DateCookie = ? WHERE DiscordID = ?")
	CheckErr("sessionlogger (SessionLogger): ", err)
	stmt.Exec(cookie.Value, strDate_cookie, discordId)
	// fmt.Println("cookie.Value : ", cookie.Value)
	http.SetCookie(w, cookie)
}

// Fonction qui compare le cookie utilisateur avec la map de gestion des cookies
func CheckToken(c *http.Cookie, database *sql.DB) bool {
	// Récupération des informations cookie de la db
	var id, uuid, dateCookie string
	stmt, err := database.Prepare("SELECT ID, uuid, DateCookie FROM Users WHERE uuid = ?")
	CheckErr("db Prepare user_db_uuid : ", err)
	err1 := stmt.QueryRow(c.Value).Scan(&id, &uuid, &dateCookie)
	// fmt.Println("err1 : ", err1)

	if err1 == nil && dateCookie != "" {
		// Format de date pour convertir en time.Time
		layout := "2006-01-02 15:04:05"
		origine_date_cookie, err := time.Parse(layout, dateCookie)
		if err != nil {
			fmt.Println("Erreur de parsing date cookie:", err)
			return false
		}

		if uuid == c.Value && time.Now().Before(origine_date_cookie) {
			return true
		}
	}

	return false
}

// Fonction de déconnexion (supression du cookie) et suppression de l'uuid dans la db
func Logout(w http.ResponseWriter, r *http.Request, database *sql.DB) {
	c, err := r.Cookie("user_token")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	// CheckErr("1- logout : ", err)
	stmt, err := database.Prepare("UPDATE Users SET uuid = '', DateCookie = '' WHERE uuid = ?")
	CheckErr("2- logout :", err)
	stmt.Exec(c.Value)

	cookie := &http.Cookie{
		Name:    "user_token",
		Value:   "",
		Expires: time.Unix(0, 0),
		MaxAge:  -1,
		Domain:  data.SITE_DOMAIN,
		Path:    "/",
	}
	http.SetCookie(w, cookie)
	// fmt.Println("Logged out successfully")
}
