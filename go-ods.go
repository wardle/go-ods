package main

import (
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/wardle/go-ods/store"
	"github.com/wardle/go-ods/trud"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

//go:generate protoc -I=trud --go_out=plugins=grpc:trud trud/nhspd.proto

var doImport = flag.String("import", "", "import data files from zip file specified")
var database = flag.String("db", "", "filename of database to open or create (e.g. ./ods.db)")
var runRPCServer = flag.Bool("rpcserver", false, "Run a RPC server")
var port = flag.Int("port", 8080, "Port to run server")
var readOnly = flag.Bool("ro", false, "Run in read-only mode")

func main() {
	flag.Parse()
	if len(os.Args) < 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *database == "" {
		fmt.Fprint(os.Stderr, "error: missing mandatory database file\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
	var db *store.Database
	var err error
	// if we are importing, open database in read-write mode.
	if *doImport != "" {
		db, err = store.OpenDatabase(*database, false)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: could not open database file '%s'", *database)
			os.Exit(1)
		}
		fmt.Printf("Running import from '%s'", *doImport)
		trud.Import(*doImport, 200, func(ft trud.ODSFileType, result []proto.Message) {
			err := db.Put(ft.GetName(), ft.GetPropertyKey(), result)
			if err != nil {
				panic(err)
			}
		})
	}

	// if we haven't already opened database, now open in read-only mode, if requested
	if db == nil {
		db, err = store.OpenDatabase(*database, *readOnly)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: could not open database file '%s'", *database)
			os.Exit(1)
		}
	}

	if *runRPCServer {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
		if err != nil {
			panic(err)
		}
		var opts []grpc.ServerOption
		server := grpc.NewServer(opts...)
		trud.RegisterTrudServer(server, &rpcServer{db: db})
		fmt.Printf("Starting RPC server on port %d\n", *port)
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to serve : %v", err)
		}
	}

}

type rpcServer struct {
	db *store.Database
}

func (rs *rpcServer) GetPostcode(ctx context.Context, postcode *trud.Postcode) (*trud.NHSPD, error) {
	r := new(trud.NHSPD)
	if err := rs.db.Get(trud.ODSGridAll.GetName(), postcode.Postcode, r); err != nil {
		return nil, err
	}
	return r, nil
}
