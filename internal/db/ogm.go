package db

import (
	"github.com/jhabshoosh/etzer-api/internal/config"
	"github.com/jhabshoosh/etzer-api/internal/models"
	"github.com/mindstand/gogm/v2"
)


func InitNeo4JOGM() *gogm.Gogm {
	env := config.GetEnv()

	config := gogm.Config{
		Host: 				env.DBHost,
	Port: 					env.DBPort,
	    Protocol:			"bolt", //also supports neo4j+s, neo4j+ssc, bolt, bolt+s and bolt+ssc
	    // Specify CA Public Key when using +ssc or +s
	    // CAFileLocation: "my-ca-public.crt",
		Username:			env.DBUser,
		Password:			env.DBPassword,
		PoolSize:			50,
		IndexStrategy:		gogm.ASSERT_INDEX, //other options are ASSERT_INDEX and IGNORE_INDEX
		TargetDbs:			nil,
		// default logger wraps the go "log" package, implement the Logger interface from gogm to use your own logger
		Logger:             gogm.GetDefaultLogger(),
		// define the log level
		LogLevel:           "DEBUG",
		// enable neo4j go driver to log
		EnableDriverLogs:   false,
		// enable gogm to log params in cypher queries. WARNING THIS IS A SECURITY RISK! Only use this when debugging
		EnableLogParams:    false,
		// enable open tracing. Ensure contexts have spans already. GoGM does not make root spans, only child spans
		OpentracingEnabled: false,
		// specify the method gogm will use to generate Load queries
		LoadStrategy: gogm.PATH_LOAD_STRATEGY, // set to SCHEMA_LOAD_STRATEGY for schema-aware queries which may reduce load on the database
	}

	// register all vertices and edges
	// this is so that GoGM doesn't have to do reflect processing of each edge in real time
	// use nil or gogm.DefaultPrimaryKeyStrategy if you only want graph ids
	// we are using the default key strategy since our vertices are using BaseNode
	_gogm, err := gogm.New(&config, gogm.UUIDPrimaryKeyStrategy, &models.Person{})
	if err != nil {
		panic(err)
	}

	return _gogm


}