namespace SqliteCrud

open System
open Microsoft.Data.Sqlite

module Program =
    let connectionString = "Data Source=students.db"

    let printAllStudents (conn: SqliteConnection) =
        use cmd = conn.CreateCommand()
        cmd.CommandText <- "SELECT Id, Name, Grade FROM Students;"
        use reader = cmd.ExecuteReader()
        printfn "\n--- Current Students Table ---"
        while reader.Read() do
            let id = reader.GetInt32(0)
            let name = reader.GetString(1)
            let grade = reader.GetString(2)
            printfn "ID: %d | Name: %s | Grade: %s" id name grade
        printfn "------------------------------\n"

    [<EntryPoint>]
    let main argv =
        use connection = new SqliteConnection(connectionString)
        connection.Open()

        // Setup Table
        use createTableCmd = connection.CreateCommand()
        createTableCmd.CommandText <- """
            CREATE TABLE IF NOT EXISTS Students (
                Id INTEGER PRIMARY KEY AUTOINCREMENT,
                Name TEXT NOT NULL,
                Grade TEXT NOT NULL
            );
            DELETE FROM Students; -- Reset for clean run
        """
        createTableCmd.ExecuteNonQuery() |> ignore

        printfn "=== SQLite Relational DB CRUD Operations ==="

        // 1. CREATE (Insert)
        printfn "[CREATE] Inserting initial student records..."
        let insert (name: string) (grade: string) =
            use cmd = connection.CreateCommand()
            cmd.CommandText <- "INSERT INTO Students (Name, Grade) VALUES ($name, $grade);"
            cmd.Parameters.AddWithValue("$name", name) |> ignore
            cmd.Parameters.AddWithValue("$grade", grade) |> ignore
            cmd.ExecuteNonQuery() |> ignore

        insert "Alice Smith" "A"
        insert "Bob Jones" "B"
        printAllStudents connection

        // 2. READ (Select specific record)
        printfn "[READ] Fetching student with ID 1..."
        let readCmd = connection.CreateCommand()
        readCmd.CommandText <- "SELECT Name, Grade FROM Students WHERE Id = 1;"
        use reader = readCmd.ExecuteReader()
        if reader.Read() then
            printfn "Found -> Name: %s, Grade: %s" (reader.GetString(0)) (reader.GetString(1))
        reader.Close()

        // 3. UPDATE
        printfn "\n[UPDATE] Updating Bob Jones's grade to 'A+'..."
        let updateCmd = connection.CreateCommand()
        updateCmd.CommandText <- "UPDATE Students SET Grade = 'A+' WHERE Name = 'Bob Jones';"
        updateCmd.ExecuteNonQuery() |> ignore
        printAllStudents connection

        // 4. DELETE
        printfn "[DELETE] Removing Alice Smith from database..."
        let deleteCmd = connection.CreateCommand()
        deleteCmd.CommandText <- "DELETE FROM Students WHERE Name = 'Alice Smith';"
        deleteCmd.ExecuteNonQuery() |> ignore
        printAllStudents connection

        0