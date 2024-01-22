package handlers

// import (
// 	"context"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/go-redis/redis/v8"
// 	"github.com/gorilla/mux"
// )

// var redisClient *redis.Client

// func init() {
// 	redisClient = redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379", // Adresse du serveur Redis
// 		Password: "",               // Mot de passe si nécessaire
// 		DB:       0,                // Numéro de base de données par défaut
// 	})
// }

// // Patient struct représente les données du patient
// type Patient struct {
// 	ID       string    `json:"id"`
// 	Name     string    `json:"name"`
// 	DOB      time.Time `json:"dob"`
// 	ExpireAt time.Time `json:"expire_at"`
// }

// // func main() {
// // 	r := mux.NewRouter()

// // 	// Endpoint pour supprimer un patient
// // 	r.HandleFunc("/doctor/patient/{id}", DeletePatientHandler).Methods("DELETE")

// // 	// Handler Lambda pour la fonction Docteur
// // 	lambda.StartHandler(r, common.HandlerConfig{})

// // 	// Cron Job quotidien pour supprimer les comptes expirés de la base de données principale
// // 	go func() {
// // 		for {
// // 			DeleteExpiredAccounts()
// // 			time.Sleep(24 * time.Hour) // Attente d'une journée avant le prochain exécution
// // 		}
// // 	}()

// // 	// Cron Job quotidien pour envoyer des rappels par e-mail depuis Redis
// // 	go func() {
// // 		for {
// // 			SendEmailReminders()
// // 			time.Sleep(24 * time.Hour) // Attente d'une journée avant le prochain exécution
// // 		}
// // 	}()

// // 	select {}
// // }

// // DeletePatientHandler gère la suppression du patient
// func DeletePatientHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	patientID := vars["id"]

// 	// Supprimer le patient de la base de données principale
// 	// (Ajoutez votre logique de suppression ici)

// 	// Stocker l'ID du patient dans Redis avec une expiration d'un mois
// 	expirationTime := time.Now().Add(30 * 24 * time.Hour)
// 	err := redisClient.Set(context.Background(), patientID, true, expirationTime.Sub(time.Now())).Err()
// 	if err != nil {
// 		http.Error(w, "Erreur lors de la sauvegarde dans Redis", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// // DeleteExpiredAccounts supprime les comptes expirés de la base de données principale
// func DeleteExpiredAccounts() {
// 	// Obtenez la liste des clés expirées dans Redis
// 	ctx := context.Background()
// 	keys, err := redisClient.Keys(ctx, "*").Result()
// 	if err != nil {
// 		log.Println("Erreur lors de la récupération des clés Redis expirées:", err)
// 		return
// 	}

// 	// Supprimez les comptes correspondants dans la base de données principale
// 	for _, key := range keys {
// 		// (Ajoutez votre logique de suppression de compte ici)
// 		log.Printf("Suppression du compte expiré: %s\n", key)

// 		// Supprimez la clé de Redis après avoir traité le compte
// 		err := redisClient.Del(ctx, key).Err()
// 		if err != nil {
// 			log.Println("Erreur lors de la suppression de la clé Redis:", err)
// 		}
// 	}
// }

// // SendEmailReminders envoie des rappels par e-mail à partir des données dans Redis
// func SendEmailReminders() {
// 	// Obtenez la liste des clés dans Redis
// 	ctx := context.Background()
// 	keys, err := redisClient.Keys(ctx, "*").Result()
// 	if err != nil {
// 		log.Println("Erreur lors de la récupération des clés Redis:", err)
// 		return
// 	}

// 	// Traitement des rappels par e-mail pour chaque clé
// 	for _, key := range keys {
// 		expireTime, err := redisClient.Get(ctx, key).Result()
// 		if err != nil {
// 			log.Println("Erreur lors de la récupération de la date d'expiration depuis Redis:", err)
// 			continue
// 		}

// 		expireDate, err := time.Parse(time.RFC3339, expireTime)
// 		if err != nil {
// 			log.Println("Erreur lors de la conversion de la date d'expiration:", err)
// 			continue
// 		}

// 		// Logique de traitement des rappels par e-mail ici
// 		// (Envoyez un e-mail en fonction de la logique de rappel spécifiée)

// 		// Supprimez la clé de Redis après avoir traité le rappel
// 		err = redisClient.Del(ctx, key).Err()
// 		if err != nil {
// 			log.Println("Erreur lors de la suppression de la clé Redis:", err)
// 		}
// 	}
// }
