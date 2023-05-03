package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/supertokens/supertokens-golang/recipe/dashboard"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword/tpepmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func main() {
	apiBasePath := "/auth"
	websiteBasePath := "/auth"
	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			// https://try.supertokens.com is for demo purposes. Replace this with the address of your core instance (sign up on supertokens.com), or self host a core.
			ConnectionURI: "http://localhost:3567/",
			// APIKey: <API_KEY(if configured)>,
		},
		AppInfo: supertokens.AppInfo{
			AppName:         "ikarus",
			APIDomain:       "http://localhost:8000",
			WebsiteDomain:   "http://localhost:3000",
			APIBasePath:     &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			thirdpartyemailpassword.Init(&tpepmodels.TypeInput{Providers: []tpmodels.TypeProvider{
				thirdparty.Google(tpmodels.GoogleConfig{
					ClientID:     "1060725074195-kmeum4crr01uirfl2op9kd5acmi9jutn.apps.googleusercontent.com",
					ClientSecret: "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",
				}),
				thirdparty.Github(tpmodels.GithubConfig{
					ClientID:     "467101b197249757c71f",
					ClientSecret: "e97051221f4b6426e8fe8d51486396703012f5bd",
				}),
			}}),
			session.Init(nil), // initializes session features
			dashboard.Init(nil),
		},
	})

	if err != nil {
		panic(err.Error())
	}

	router := mux.NewRouter()
	router.Use(loggingMiddleware)

	// Adding handlers.CORS(options)(supertokens.Middleware(router)))
	fmt.Println("Listening to server: http://localhost:8000")
	http.ListenAndServe(":8000", handlers.CORS(
		handlers.AllowedHeaders(append([]string{"Content-Type"},
			supertokens.GetAllCORSHeaders()...)),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowCredentials(),
	)(supertokens.Middleware(router)))

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		endTime := time.Now()

		log.Printf(
			"%s %s %s %v",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			endTime.Sub(startTime),
		)
	})
}
