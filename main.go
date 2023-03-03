package main

import (
	"main/controller"
	"main/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	utils.InitLogger()

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowOriginFunc: func(r *http.Request, origin string) bool {
			utils.Log.Println("CORS origin", origin)
			return true
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))
	r.Use(middleware.Logger)
	r.Route("/kodeyapi", func(r chi.Router) {
		r.Get("/ping", pongCheck)

		r.Route("/v1", func(r chi.Router) {
			// JSON routes
			r.Group(func(r chi.Router) {
				r.Use(utils.JsonRoute)

				r.Route("/code", func(r chi.Router) {

					r.Post("/execute", controller.ExecuteCode)
					r.Post("/edit", controller.FileSubmit)
				})
			})
		})
	})

	utils.Log.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		utils.Log.Fatalln("Failed to start server...", err)
	}
}

func pongCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
	<html>
		<body style="
			min-height:100vh;
			display:flex;
			justify-content:center;
			align-items:center"
		>
			<h2>
				pong
			</h2>
		</body>
	</html>`))
}
