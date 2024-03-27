# Preguntas

1. **¿Que son short variables?**
    <br>Se usan para declarar e  inicializar variables dentro de funciones </br>
    ```go
    func main() {
        x := 42
        fmt.Println(x)
    }
    ```
2. **¿Cuando puedes utilizar short variables y cuando no?**
    <br>Se pueden usar dentro funciones y en la redeclaración de una variable. No se pueden usar fuera de funciones y con tipos explicitos sin iniciación.</br>
   **Se puede usar**
   ```go
    func ejemplo() {
        x := 1
        x, y := 2, 3
    }
    ```
   
   **No se puede usar**
   ```go
   
    var globalVar = "Esto se puede hacer"
   //Da error
    globalVar := "Esto no se puede hacer"
   
    ```
3. **¿Que siginica inferencia en tipos de datos?**
    <br>Cuando el compilador en el caso de Go tiene la capacidad de determinar automáticamente el tipo de variable dependiendo el valor con el que se inicializa sin poner el tipo.</br>
     ```go
    nombre := "David" // Se infiere que es string
    edad := 25 // Se infiere que es int
    ```  

4. **¿Puede una constante declararse de manera corta como son “short variables?**
    <br>Las constantes no pueden declararse como short variables porque para eso esta la palabra reservada ``const``</br>
     ```go
    const Pi float64 = 3.14159265358979323846
    const greet = "Hello, World!"
    ```
   
5. **¿Qué es "defer" ?**
   <br>Es una palabra reservada que se utiliza para que una función se ejecute justo antes de que termine la función principal.</br>
    ```go
     func readFile() {
          archivo, err := os.Open("archivo.txt")
            if err != nil {
                  log.Fatal(err)
             }
            defer archivo.Close()
     }
     ```
6. **¿Qué son los pointer?**
    <br>Un pointer es una variable que almacena la dirección en memoria de otra variable. Su principal uso es para la eficiencia en el manejo de datos.</br>
   ```go
       var x int = 10
       var p *int = &x // 'p' es un pointer a 'x'
     ```

7. **¿Qué es struct?**
   <br>Permite agrupar campos bajo un nombre o es como se crean las entidades en Go.</br>
      ```go
        type Persona struct {
            nombre string
            edad int
        }
      ```

8. **¿Qué es un goroutine?**
    <br>Es una función que se ejecuta de manera concurrente y permite realizar tareas en paralelo de una manera mas sencilla; es uno de los conceptos mas fundamentales de go.</br>
    ```go
    
       import (
         "fmt"
         "time"
        )

        func decir(s string) {
           for i := 0; i < 5; i++ {
               time.Sleep(100 * time.Millisecond)
               fmt.Println(s)
              }
      }

        func main() {
        go decir("mundo")
        decir("hola")
      }
      ```