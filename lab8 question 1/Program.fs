namespace ShapeOOP

// 1. Define the base class with a virtual/abstract method
type Shape() =
    // abstract member serves as the virtual method definition with a default implementation
    abstract member Area : unit -> float
    default this.Area() = 0.0

// 2. Define the Circle derived class overriding Area
type Circle(radius: float) =
    inherit Shape()
    member val Radius = radius

    override this.Area() = System.Math.PI * this.Radius * this.Radius

// 2. Define the Rectangle derived class overriding Area
type Rectangle(width: float, height: float) =
    inherit Shape()
    member val Width = width
    member val Height = height

    override this.Area() = this.Width * this.Height

module Program =
    [<EntryPoint>]
    let main argv =
        // 3. Instantiate objects of each class and store them polymorphically in a list of base type `Shape`
        let shapes: Shape list = [
            Circle(5.0)
            Rectangle(4.0, 6.0)
            Circle(3.0)
            Rectangle(10.5, 2.0)
        ]

        printfn "========================================"
        printfn " Demonstrating OOP Polymorphism in F#   "
        printfn "========================================"

        // Iterate through the shapes and call Area() polymorphically
        for (index, shape) in shapes |> List.indexed do
            let area = shape.Area()
            match shape with
            | :? Circle as c -> 
                printfn "[Shape %d] Circle (Radius: %.1f) -> Area: %.2f" (index + 1) c.Radius area
            | :? Rectangle as r -> 
                printfn "[Shape %d] Rectangle (%.1f x %.1f) -> Area: %.2f" (index + 1) r.Width r.Height area
            | _ -> 
                printfn "[Shape %d] Generic Shape -> Area: %.2f" (index + 1) area

        printfn "========================================"
        0