package main

import (
	"log"
	"net"
	"time"

	"github.com/FrancescoIlario/usplay/cmd/services/bookmaster/server/conf"
	actapi "github.com/FrancescoIlario/usplay/internal/services/bookmaster/activity/api"
	actstore "github.com/FrancescoIlario/usplay/internal/services/bookmaster/activity/storage"
	actmongostore "github.com/FrancescoIlario/usplay/internal/services/bookmaster/activity/storage/mongostore"
	acttypeapi "github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/api"
	acttypestore "github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage"
	acttypemongostore "github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage/mongostore"
	"github.com/FrancescoIlario/usplay/pkg/services/bookmastergrpc"
	"github.com/FrancescoIlario/usplay/pkg/services/ordergrpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	logrus.Debug("Starting server")

	conf, err := conf.GetConfigurationFromEnv()
	if err != nil {
		logrus.Fatalf("unable to parse configuration: %v", err)
	}

	ls, err := net.Listen("tcp", conf.Address)
	if err != nil {
		logrus.Fatalf("failed to listen at %v: %v", conf.Address, err)
	}
	logrus.Debugf("acquired conf.Address %v", conf.Address)

	actrepo, err := buildActivityRepo(&conf)
	if err != nil {
		logrus.Fatalf("error instantiating Activity repository: %v", err)
	}

	acttyperepo, err := buildActivityTypeRepo(&conf)
	if err != nil {
		logrus.Fatalf("error instantiating ActivityType repository: %v", err)
	}

	connOrders, err := grpc.Dial(conf.OrderHost, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("cannot connect to %s: %v", conf.OrderHost, err)
	}
	orderCli := ordergrpc.NewOrderSvcClient(connOrders)

	actServer := actapi.NewActivityServer(actrepo, acttyperepo, orderCli, 1*time.Second)
	grpcServer := grpc.NewServer()
	bookmastergrpc.RegisterActivitySvcServer(grpcServer, actServer)

	actTypeServer := acttypeapi.NewActivityTypeServer(acttyperepo)
	bookmastergrpc.RegisterActivityTypeSvcServer(grpcServer, actTypeServer)

	log.Printf("starting server at %v", conf.Address)
	if err := grpcServer.Serve(ls); err != nil {
		log.Fatalf("Error serving: %v", err)
	}
}

func buildActivityRepo(c *conf.Configuration) (actstore.Repository, error) {
	mc := &actmongostore.Configuration{
		Collection:       c.ActivityMongoCollection,
		ConnectionString: c.ActivityMongoConnstr,
		Database:         c.ActivityMongoDatabase,
		Password:         c.ActivityMongoPassword,
		Username:         c.ActivityMongoUsername,
	}
	return actmongostore.New(mc)
}

func buildActivityTypeRepo(c *conf.Configuration) (acttypestore.Repository, error) {
	mc := &acttypemongostore.Configuration{
		Collection:       c.ActivitytypeMongoCollection,
		ConnectionString: c.ActivitytypeMongoConnstr,
		Database:         c.ActivitytypeMongoDatabase,
		Password:         c.ActivitytypeMongoPassword,
		Username:         c.ActivitytypeMongoUsername,
	}
	return acttypemongostore.New(mc)
}
