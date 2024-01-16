package application

import (
	"app/internal/loader"
	"app/internal/ticket/handler"
	"app/internal/ticket/repository"
	"app/internal/ticket/service"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

// ConfigAppDefault represents the configuration of the default application
type ConfigAppDefault struct {
	// serverAddr represents the address of the server
	ServerAddr string
	// dbFile represents the path to the database file
	DbFile string
}

func NewConfigDefault() *ConfigAppDefault {
	return &ConfigAppDefault{
		ServerAddr: os.Getenv("SERVER_ADDR"),
		DbFile:     os.Getenv("DB_FILE"),
	}
}

// NewApplicationDefault creates a new default application
func NewApplicationDefault(cfg *ConfigAppDefault) *ApplicationDefault {
	// default values
	defaultRouter := chi.NewRouter()
	defaultConfig := &ConfigAppDefault{
		ServerAddr: ":8080",
		DbFile:     "",
	}
	if cfg != nil {
		if cfg.ServerAddr != "" {
			defaultConfig.ServerAddr = cfg.ServerAddr
		}
		if cfg.DbFile != "" {
			defaultConfig.DbFile = cfg.DbFile
		}
	}

	return &ApplicationDefault{
		rt:         defaultRouter,
		serverAddr: defaultConfig.ServerAddr,
		dbFile:     defaultConfig.DbFile,
	}
}

// ApplicationDefault represents the default application
type ApplicationDefault struct {
	// router represents the router of the application
	rt *chi.Mux
	// serverAddr represents the address of the server
	serverAddr string
	// dbFile represents the path to the database file
	dbFile string
}

// SetUp sets up the application
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies

	//Loader
	loader := loader.NewLoaderTicketCSV(a.dbFile)

	//DB
	db, err := loader.Load()

	if err != nil {
		return
	}

	//Repository
	rp := repository.NewRepositoryTicketMap(db, len(db))

	// service ...
	sv := service.NewServiceTicketDefault(rp)

	// handler ...
	hd := handler.NewDefaultTicketHandler(sv)

	// routes
	(*a).rt.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("OK"))
	})

	a.rt.Route("/ticket", func(r chi.Router) {
		r.Get("/getByCountry/{dest}", hd.GetTicketsByCountry())
		r.Get("/getAverage/{dest}", hd.GetTicketsAverageByCountry())
	})

	return
}

// Run runs the application
func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.serverAddr, a.rt)
	return
}
