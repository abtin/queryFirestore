# queryFirestore

A utility to run simple queries against google firestore.

To Simplify the command-line parameters, first set the project-id and service account key environment variables and then
call the command.

```bash
   $ export GCP_PROJECTID=<gcp-project-id>
   $ export GCP_JSON_KEY_FILE=<path to service account key json file>
   $ ./queryFirestore -d <document> -f <field> -o "<operator>" -v "<value>"
``` 

Or if you need to print the usage:

```bash
   $ ./queryFirestore
   NAME:
      queryFirestore - A cli to run simple queries against google firestore
   
   USAGE:
      queryFirestore [global options] command [command options] [arguments...]
   
   COMMANDS:
      help, h  Shows a list of commands or help for one command
   
   GLOBAL OPTIONS:
      --projectId value, -p value    Google Project Id [$GCP_PROJECTID]
      --jsonKeyFile value, -j value  The Json Key file [$GCP_JSON_KEY_FILE]
      --document value, -d value     Firestore <Document>
      --field value, -f value        One of '<', '<=', '>', '>=', '=='
      --operator value, -o value     One of '<', '<=', '>', '>=', '=='
      --value value, -v value        <Value> to query for
      --help, -h                     show help (default: false)
   Required flags "document, field, operator, value" not set
```