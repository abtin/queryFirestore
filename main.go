package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/urfave/cli/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"os"
)

func main() {
	app := &cli.App{
		Name: "queryFirestore - A cli to run simple queries against google firestore",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "projectId",
				Aliases: []string{"p"},
				Usage:   "Google Project Id",
				EnvVars: []string{"GCP_PROJECTID"},
			},
			&cli.StringFlag{
				Name:    "jsonKeyFile",
				Aliases: []string{"j"},
				Usage:   "The Json Key file",
				EnvVars: []string{"GCP_JSON_KEY_FILE"},
			},
			&cli.StringFlag{
				Name:     "document",
				Aliases:  []string{"d"},
				Usage:    "Firestore <Document>",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "field",
				Aliases:  []string{"f"},
				Usage:    "One of '<', '<=', '>', '>=', '=='",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "operator",
				Aliases:  []string{"o"},
				Usage:    "One of '<', '<=', '>', '>=', '=='",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "value",
				Aliases:  []string{"v"},
				Usage:    "<Value> to query for",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			ctx := context.Background()
			client, err := firestore.NewClient(ctx, c.String("projectId"),
				option.WithCredentialsFile(c.String("jsonKeyFile")))
			if err != nil {
				fail(err)
			}
			fmt.Println("Query Result:")
			iter := client.Collection(c.String("document")).Where(
				c.String("field"),
				c.String("operator"),
				c.String("value")).Documents(ctx)
			for {
				doc, err := iter.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					fail(err)
				}
				fmt.Printf("%#+v\n", doc.Data())
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fail(err)
	}

}

func fail(err error) {
	fmt.Println(err)
	os.Exit(1)
}
