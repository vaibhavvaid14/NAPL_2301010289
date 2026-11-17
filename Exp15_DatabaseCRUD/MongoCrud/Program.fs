namespace MongoCrud

open System
open MongoDB.Bson
open MongoDB.Driver

// Define Student Document mapping model
type StudentDoc = {
    [<BsonId>]
    [<BsonRepresentation(BsonType.ObjectId)>]
    Id : string
    Name : string
    Grade : string
}

module Program =
    let printAllStudents (collection: IMongoCollection<StudentDoc>) =
        printfn "\n--- Current MongoDB Students Collection ---"
        let students = collection.Find(Builders.Filter.Empty).ToList()
        for s in students do
            printfn "ID: %s | Name: %s | Grade: %s" s.Id s.Name s.Grade
        printfn "-------------------------------------------\n"

    [<EntryPoint>]
    let main argv =
        // Connect to local MongoDB instance
        let connectionString = "mongodb://localhost:27017"
        let client = MongoClient(connectionString)
        let database = client.GetDatabase("SchoolDB")
        let collection = database.GetCollection<StudentDoc>("Students")

        // Clean collection for demo
        collection.DeleteMany(Builders.Filter.Empty) |> ignore

        printfn "=== MongoDB NoSQL CRUD Operations ==="

        // 1. CREATE (Insert)
        printfn "[CREATE] Inserting student documents..."
        let student1 = { Id = ObjectId.GenerateNewId().ToString(); Name = "Charlie Brown"; Grade = "B+" }
        let student2 = { Id = ObjectId.GenerateNewId().ToString(); Name = "Diana Prince"; Grade = "A" }
        
        collection.InsertMany([student1; student2])
        printAllStudents collection

        // 2. READ (Find specific document)
        printfn "[READ] Querying student where Name is 'Diana Prince'..."
        let filter = Builders<StudentDoc>.Filter.Eq(fun s -> s.Name, "Diana Prince")
        let found = collection.Find(filter).FirstOrDefault()
        if found <> null then
            printfn "Found Document -> ID: %s, Grade: %s" found.Id found.Grade

        // 3. UPDATE
        printfn "\n[UPDATE] Updating Charlie Brown's grade to 'A'..."
        let updateFilter = Builders<StudentDoc>.Filter.Eq(fun s -> s.Name, "Charlie Brown")
        let updateDef = Builders<StudentDoc>.Update.Set(fun s -> s.Grade, "A")
        collection.UpdateOne(updateFilter, updateDef) |> ignore
        printAllStudents collection

        // 4. DELETE
        printfn "[DELETE] Removing Diana Prince document..."
        let deleteFilter = Builders<StudentDoc>.Filter.Eq(fun s -> s.Name, "Diana Prince")
        collection.DeleteOne(deleteFilter) |> ignore
        printAllStudents collection

        0