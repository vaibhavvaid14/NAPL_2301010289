namespace StudentManagement

module Program =
    [<EntryPoint>]
    let main argv =
        async {
            printfn "Initializing Database..."
            do! Database.initializeDatabaseAsync ()

            // Seed sample data
            let student1 = { Id = 0; Name = "Vaibhav Vaid"; Department = ComputerScience; GPA = 3.85; Status = Active }
            let student2 = { Id = 0; Name = "Aadarsh"; Department = Mathematics; GPA = 3.92; Status = Active }
            let student3 = { Id = 0; Name = "Ajit Panja"; Department = Physics; GPA = 3.10; Status = Graduated }

            printfn "Persisting records via async workflows..."
            do! Database.insertStudentAsync student1
            do! Database.insertStudentAsync student2
            do! Database.insertStudentAsync student3

            // Read back all students from database
            let! allStudents = Database.getAllStudentsAsync ()
            printfn "\n--- All Registered Students ---"
            allStudents |> List.iter (fun s -> printfn "%A" s)

            // Process data using higher-order functions (map, filter)
            printfn "\n--- Computer Science Active Students ---"
            let csActive = Analytics.getActiveStudentsByDept ComputerScience allStudents
            csActive |> List.iter (fun s -> printfn "- %s (GPA: %.2f)" s.Name s.GPA)

            printfn "\n--- Honor Roll (GPA >= 3.8) ---"
            let honorRoll = Analytics.getHonorRoll 3.8 allStudents
            honorRoll |> List.iter (fun (name, gpa) -> printfn "* %s scored %.2f" name gpa)

            // Pattern matching demonstration
            printfn "\n--- Status Evaluations ---"
            allStudents |> List.iter (fun s -> 
                let description = Analytics.describeStatus s.Status
                printfn "%s: %s" s.Name description
            )

        } |> Async.RunSynchronously

        0