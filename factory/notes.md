    Summary: The factory method is a creational design pattern that provides an 
interface for creating objects in a superclass, but allows subclasses to alter the 
type of objects that will be created.

    When to use: Use the factory method when you want to delegate object creation
to subclasses, when you need to add new classes without changing existing code, or
when you want to decouple object creation from the rest of your code.  Use the 
Factory Method when you don’t know beforehand the exact types and dependencies of
the objects your code should work with.  Use the Factory Method when you want to 
provide users of your library or framework with a way to extend its internal 
components.  Use the Factory Method when you want to save system resources by 
reusing existing objects instead of rebuilding them each time.

    When not to use: Don't use the factory method when your class doesn't need to 
know how its objects are created, when you only have one concrete product class, 
or when you need to customise object creation for each instance.

    Pros: The factory method allows you to create objects without specifying the 
exact class of object that will be created, it decouples object creation from object 
use, and it simplifies object creation by providing a single point of control.

    Cons: The factory method can be complex to implement, it can result in a large 
number of small classes, and it can make your code more difficult to understand if 
the factory and product classes are not well-designed.

    Complementary techniques: The abstract factory pattern is a related creational
pattern that provides an interface for creating families of related or dependent 
objects.

    Similarities/relationships: The factory method pattern is similar to the template 
method pattern, in that both provide an interface that delegates responsibility to 
subclasses. The factory method pattern is also related to the strategy pattern, in 
that both use composition to vary behavior at runtime.