# go_interface_example_use_sql_or_postgres

In this example you can switch database sql or postgres by just changing one line of code 

main.go line 10

if a.InitializeDB(&database.Postgress{}) or if a.InitializeDB(&database.SqlServer{})
