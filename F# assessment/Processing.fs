namespace StudentManagement

module Analytics =

    // Pattern matching to safely extract string representations from discriminated unions
    let departmentToString = function
        | ComputerScience -> "ComputerScience"
        | Mathematics     -> "Mathematics"
        | Physics         -> "Physics"
        | Electronics     -> "Electronics"

    let parseDepartment = function
        | "ComputerScience" -> ComputerScience
        | "Mathematics"     -> Mathematics
        | "Physics"         -> Physics
        | _                 -> Electronics

    let statusToString = function
        | Active    -> "Active"
        | Graduated -> "Graduated"
        | Suspended -> "Suspended"

    let parseStatus = function
        | "Active"    -> Active
        | "Graduated" -> Graduated
        | _           -> Suspended

    // Higher-order function: Filter active students by department
    let getActiveStudentsByDept (dept: Department) (students: Student list) =
        students
        |> List.filter (fun s -> s.Department = dept && s.Status = Active)

    // Higher-order function: Map and filter high-achieving honor students
    let getHonorRoll (minGpa: float) (students: Student list) : (string * float) list =
        students
        |> List.filter (fun s -> s.GPA >= minGpa)
        |> List.map (fun s -> (s.Name, s.GPA))

    // Pattern matching for descriptive status evaluation
    let describeStatus = function
        | Active    -> "Enrolled and in good academic standing."
        | Graduated -> "Successfully completed degree requirements."
        | Suspended -> "Enrollment temporarily paused or under review."