package main

import (
	"github.com/deliryo-io/VerificationStatusUpdater/config"
	"github.com/deliryo-io/VerificationStatusUpdater/internal/handlers/upload_status_api"
	"github.com/kabukky/httpscerts"
	"log"
	"net/http"
)

const (
	update_verification_status = "/verification_status_uploader/upload_status"
)

func main() {
	log.Println("start application")
	//ctx := context.Background()
	//DbConnection := database_tools.NewConnection(ctx)

	err := httpscerts.Check("cert.pem", "key.pem")
	// Если креды недоступны, то генерируем новые.
	if err != nil {
		err = httpscerts.Generate("cert.pem", "key.pem", "127.0.0.1:8081")
		if err != nil {
			log.Fatal("Ошибка: Не можем сгенерировать https сертификат.")
		}
	}
	//giris

	go http.HandleFunc(update_verification_status, upload_status_api.Upload_status_handler)

	//err = http.ListenAndServeTLS(mediator_https.Get_rest_api_creds(), "./cert.pem", "./key.pem", nil)
	err = http.ListenAndServe(config.Verification_listener_creds.Host+config.Verification_listener_creds.Port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Println("server is listening!")
}
