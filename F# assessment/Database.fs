namespace StudentManagement

module Database =
    open Microsoft.Data.Sqlite

    let connectionString = "Data Source=university.db"

    // Initialize database table asynchronously
    let initializeDatabaseAsync () =
        async {
            use conn = new SqliteConnection(connectionString)
            do! conn.OpenAsync() |> Async.AwaitTask
            
            use cmd = conn.CreateCommand()
            cmd.CommandText <- """
                CREATE TABLE IF NOT EXISTS Students (
                    Id INTEGER PRIMARY KEY AUTOINCREMENT,
                    Name TEXT NOT NULL,
                    Department TEXT NOT NULL,
                    GPA REAL NOT NULL,
                    Status TEXT NOT NULL
                );
            """
            do! cmd.ExecuteNonQueryAsync() |> Async.AwaitTask |> ignore
        }

    // CREATE: Insert a new student asynchronously
    let insertStudentAsync (student: Student) =
        async {
            use conn = new SqliteConnection(connectionString)
            do! conn.OpenAsync() |> Async.AwaitTask

            use cmd = conn.CreateCommand()
            cmd.CommandText <- """
                INSERT INTO Students (Name, Department, GPA, Status) 
                VALUES (@name, @dept, @gpa, @status);
            """
            cmd.Parameters.AddWithValue("@name", student.Name) |> ignore
            cmd.Parameters.AddWithValue("@dept", Analytics.departmentToString student.Department) |> ignore
            cmd.Parameters.AddWithValue("@gpa", student.GPA) |> ignore
            cmd.Parameters.AddWithValue("@status", Analytics.statusToString student.Status) |> ignore

            do! cmd.ExecuteNonQueryAsync() |> Async.AwaitTask |> ignore
        }

    // READ: Fetch all students asynchronously
    let getAllStudentsAsync () =
        async {
            use conn = new SqliteConnection(connectionString)
            do! conn.OpenAsync() |> Async.AwaitTask

            use cmd = conn.CreateCommand()
            cmd.CommandText <- "SELECT Id, Name, Department, GPA, Status FROM Students;"

            use! reader = cmd.ExecuteReaderAsync() |> Async.AwaitTask
            let mutable students = []
            
            while reader.Read() do
                let s = {
                    Id = reader.GetInt32(0)
                    Name = reader.GetString(1)
                    Department = Analytics.parseDepartment (reader.GetString(2))
                    GPA = reader.GetDouble(3)
                    Status = Analytics.parseStatus (reader.GetString(4))
                }
                students <- s :: students

            return List.rev students
        }

    // UPDATE: Update a student's GPA asynchronously
    let updateGpaAsync (id: int) (newGpa: float) =
        async {
            use conn = new SqliteConnection(connectionString)
            do! conn.OpenAsync() |> Async.AwaitTask

            use cmd = conn.CreateCommand()
            cmd.CommandText <- "UPDATE Students SET GPA = @gpa WHERE Id = @id;"
            cmd.Parameters.AddWithValue("@gpa", newGpa) |> ignore
            cmd.Parameters.AddWithValue("@id", id) |> ignore

            do! cmd.ExecuteNonQueryAsync() |> Async.AwaitTask |> ignore
        }

    // DELETE: Remove a student by ID asynchronously
    let deleteStudentAsync (id: int) =
        async {
            use conn = new SqliteConnection(connectionString)
            do! conn.OpenAsync() |> Async.AwaitTask

            use cmd = conn.CreateCommand()
            cmd.CommandText <- "DELETE FROM Students WHERE Id = @id;"
            cmd.Parameters.AddWithValue("@id", id) |> ignore

            do! cmd.ExecuteNonQueryAsync() |> Async.AwaitTask |> ignore
        }