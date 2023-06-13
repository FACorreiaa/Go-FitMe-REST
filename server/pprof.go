package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"net/http/pprof"
)

func InitPprof() {
	//profName := "my_experiment_thing"
	//libProfile = pprof.
	//if libProfile == nil {
	//	libProfile = pprof.NewProfile(profName)
	//}
	router := chi.NewRouter()
	router.Route("/debug/pprof", func(r chi.Router) {
		r.Use(middleware.NoCache)
		r.Get("/", chiWrapper(pprof.Index))
		r.Get("/cmdline", chiWrapper(pprof.Cmdline))
		r.Get("/profile?debug=1", chiWrapper(pprof.Profile))
		r.Post("/symbol", chiWrapper(pprof.Symbol))
		r.Get("/symbol", chiWrapper(pprof.Symbol))
		r.Get("/trace", chiWrapper(pprof.Trace))
		r.Get("/allocs",
			chiWrapper(handlerFunc(pprof.Handler("allocs"))))

		r.Get("/block", chiWrapper(handlerFunc(pprof.Handler("block"))))
		r.Get("/goroutine", chiWrapper(handlerFunc(pprof.Handler("goroutine"))))
		r.Get("/heap", chiWrapper(handlerFunc(pprof.Handler("heap"))))
		r.Get("/mutex", chiWrapper(handlerFunc(pprof.Handler("mutex"))))
		r.Get("/threadcreate", chiWrapper(handlerFunc(pprof.Handler("threadcreate"))))
		//r.Get("/threadcreate", func(w http.ResponseWriter, r *http.Request) {
		//	pprof.Handler("threadcreate")
		//})
		//r.Get("/mutex", func(w http.ResponseWriter, r *http.Request) {
		//	pprof.Handler("threadcreate")
		//})
		//r.Get("/heap", func(w http.ResponseWriter, r *http.Request) {
		//	pprof.Handler("threadcreate")
		//})
		//r.Get("/goroutine", func(w http.ResponseWriter, r *http.Request) {
		//	pprof.Handler("threadcreate")
		//})
	})
	//http.HandleFunc("/functiontrack", func(rw http.ResponseWriter, req *http.Request) {
	//	trackAFunction()
	//})
	go func() {
		log.Println("Running pprof!")

		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()
}

func handlerFunc(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

func chiWrapper(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		fn(w, r.WithContext(ctx))
	}
}
