package db

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

// var (
// 	InvalidConnUrlErr = errors.New("failed to connect to DB, as the connection string is invalid")
// 	ClientCreationErr = errors.New("failed to create new client to connect with db")
// 	ClientInitErr     = errors.New("failed to initialize db client")
// 	ConnectionLeak    = errors.New("unable to disconnect from db, potential connection leak")
// 	connectOnce       sync.Once
// )

// const ConnectionTimeOut = 10 * time.Second

type DBManager interface {
	Database() *sqlx.DB
	//Ping() error
	//Disconnect() error
}

type dbManager struct {
	db *sqlx.DB
}

func (d *dbManager) Database() *sqlx.DB {
	return d.db
}

func NewDBManager(connUrl string) (DBManager, error) {
	log.Debug().
		Str("URL", connUrl).
		Msg("DB Connection URL")

	db, err := sqlx.Open("postgres", connUrl)
	if err != nil {
		log.Error().Err(err).Msg("Unable to open the conection to DB")
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	dbMgr := &dbManager{db: db}

	// Verify connection
	if err := dbMgr.Ping(); err != nil {
		return nil, err
	}

	return dbMgr, nil

}

// // newClient - creates a new Mongo Client to connect to the specified url and initializes the Client
// func newClient(connectionUrl string) (*mongo.Client, error) {
// 	if len(connectionUrl) == 0 {
// 		return nil, InvalidConnUrlErr
// 	}
// 	clientOptions := options.Client().ApplyURI(connectionUrl)
// 	client, err := mongo.NewClient(clientOptions)
// 	if err != nil {
// 		log.Error().Err(err).Msg("Connection Failed to Database")
// 		return nil, ClientCreationErr
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), ConnectionTimeOut)
// 	defer cancel()
// 	connErr := client.Connect(ctx)
// 	if connErr != nil {
// 		log.Error().Err(connErr).Msg("Connection Failed to Database")
// 		return nil, ClientInitErr
// 	}

// 	return client, nil
// }

func (d *dbManager) Ping() error {
	if err := d.db.Ping(); err != nil {
		log.Error().Err(err).Msg("Unable to connect to DB")
		return err
	}
	return nil
}

// Disconnect - Close connection to Database
// func (c *connectionManager) Disconnect() error {
// 	log.Info().Msg("Disconnecting from Database")
// 	if err := c.client.Disconnect(context.Background()); err != nil {
// 		log.Error().Err(err).Msg("unable to disconnect from DB")
// 		return ConnectionLeak
// 	}
// 	log.Info().Msg("Successfully disconnected from DB")
// 	return nil
// }

// func GetDbClient(config Config) *sqlx.DB {
// 	log.Printf("DB DSN: %s", config.DBDsn)

// 	db, err := sqlx.Open("postgres", config.DBDsn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	db.SetConnMaxLifetime(time.Minute * 3)
// 	db.SetMaxOpenConns(10)
// 	db.SetMaxIdleConns(10)
// 	return db
// 	// defer db.Close()
// 	// return db
// }

// type Repo struct {
//     db *sql.DB
// }

// func NewRepo(driverName, connectionString string) (*Repo, error) {
//     db, err := sql.Open(driverName, connectionString)
//     if err != nil {
//         return nil, err
//     }
//     return &Repo{db: db}, nil
// }
