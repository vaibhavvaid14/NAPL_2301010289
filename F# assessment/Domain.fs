namespace StudentManagement

type Department = 
    | ComputerScience
    | Mathematics
    | Physics
    | Electronics

type StudentStatus =
    | Active
    | Graduated
    | Suspended

type Student = {
    Id: int
    Name: string
    Department: Department
    GPA: float
    Status: StudentStatus
}